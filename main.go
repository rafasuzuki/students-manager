package main

import (
	"github.com/rafasuzuki/students-manager/configs"
	"github.com/rafasuzuki/students-manager/initializer"
	"github.com/rafasuzuki/students-manager/router"
)

func main() {
	initializer.LoadEnvVariables()
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	router.Initialize()
}
