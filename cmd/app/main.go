package main

import (
	"todoapp/internal/presentation/api"
)

func main() {
	a := api.NewApi()
	a.Start()
}
