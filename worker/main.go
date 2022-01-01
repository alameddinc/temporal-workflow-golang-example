package main

import (
	"github.com/alameddinc/temporal-workflow-golang-example/activities"
	"github.com/alameddinc/temporal-workflow-golang-example/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	log.Println("Worker Starting...")
	opt := client.Options{
		HostPort: client.DefaultHostPort,
	}
	c, err := client.NewClient(opt)
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	w := worker.New(c, "worker-group-1", worker.Options{})
	w.RegisterWorkflow(workflows.CoffeeShopWorkflow)
	w.RegisterActivity(activities.PrepareCoffee)
	w.RegisterActivity(activities.WriteAsDept)
	w.RegisterActivity(activities.GiveCoffee)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}

}
