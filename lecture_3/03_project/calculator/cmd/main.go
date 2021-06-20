package main

import (
	"log"

	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/cmd/bootstrap"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/cmd/config"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/tasks"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	rabbitMqChannel := bootstrap.RabbitMq()
	db := bootstrap.Sqlite()

	signalHandler := bootstrap.SignalHandler()
	engine := bootstrap.CalculatorEngine(rabbitMqChannel, db)

	log.Println("Bootstrap finished. Engine is starting")

	tasks.RunTasks(signalHandler, engine)

	log.Println("Service finished gracefully")
}
