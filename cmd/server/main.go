package main

import "application/internal/handler"

func main() {
	router := handler.InitRouter()
	if err := router.Run(":8080"); err != nil {
		return
	}
}
