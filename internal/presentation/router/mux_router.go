package router

import (
	"github.com/gorilla/mux"
	"todoapp/internal/presentation/handler/task"
)

type Router struct {
	Router       *mux.Router
	taskUseCases task.UseCases
}

func NewRouter(taskUseCases task.UseCases) Router {
	r := mux.NewRouter()
	mr := Router{Router: r, taskUseCases: taskUseCases}
	mr.configRoutes()
	return mr
}

func (r *Router) configRoutes() {
	taskHandler := task.NewHandler(r.taskUseCases)
	r.Router.HandleFunc("/task", taskHandler.Create).Methods("POST")
	r.Router.HandleFunc("/task", taskHandler.FindAll).Methods("GET")
}
