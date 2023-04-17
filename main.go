package main

import "homework/homework_2/pkg/router"

func main() {
	r := router.Router()
	r.Run(":8001")
}