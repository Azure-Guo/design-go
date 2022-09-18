package main

import (
	"design-go/router"
)

func main() {
	engine := router.InitRouter()
	engine.Run(":8000")
}
