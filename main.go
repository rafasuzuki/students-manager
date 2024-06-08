package main

import (
	"github.com/rafasuzuki/students-manager/configs"
	"github.com/rafasuzuki/students-manager/router"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	router.Initialize()
}
