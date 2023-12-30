package api

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todoapp/internal/infrastructure/database"
	"todoapp/internal/infrastructure/repository"
	"todoapp/internal/presentation/handler/task"
	"todoapp/internal/presentation/router"
	"todoapp/internal/usecase"
)

type Api struct {
	server     *http.Server
	httpRouter router.Router
}

func NewApi() *Api {
	db := database.MysqlConnection("root:@tcp(localhost:3306)/tasks?parseTime=true")
	repo := repository.NewTaskRepositoryMySQL(db)
	findAllUseCase := usecase.NewFindAllUseCase(repo)
	createTaskUseCase := usecase.NewCreateTaskUseCase(repo)
	taskUseCases := task.NewUseCases(createTaskUseCase, findAllUseCase)
	r := router.NewRouter(taskUseCases)
	return &Api{httpRouter: r}
}

func (a *Api) Start() {
	a.setupServer()
	a.listenAndServe()
	a.waitForShutdownSignal()
}

func (a *Api) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

func (a *Api) listenAndServe() {
	go func() {
		log.Println("Server up and running")
		if err := a.server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}

func (a *Api) setupServer() {
	a.server = &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.httpRouter.Router,
	}
}

func (a *Api) waitForShutdownSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	err := a.server.Shutdown(ctx)
	if err != nil {
		return
	}

	log.Println("shutting down")
	os.Exit(0)
}
