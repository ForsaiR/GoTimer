package main

import (
	"awesomeProject1/restapiserver/controller"
	"awesomeProject1/restapiserver/service"
)

func main() {
	service.StartCounter()
	controller.QueryProcessor()
}
