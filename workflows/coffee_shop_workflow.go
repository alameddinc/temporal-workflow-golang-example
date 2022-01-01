package workflows

import (
	"github.com/alameddinc/temporal-workflow-golang-example/activities"
	"github.com/alameddinc/temporal-workflow-golang-example/signals"
	"go.temporal.io/sdk/workflow"
	"log"
	"time"
)

func CoffeeShopWorkflow(ctx workflow.Context, customerName string) error {
	// We set our activity options with ActivtiyOptions
	// if we want to use childworkflow and if we want to set custom settings for that
	// we should use ChildWorkflowOptions like that.
	ao := workflow.ActivityOptions{
		StartToCloseTimeout:    50 * time.Second,
		ScheduleToCloseTimeout: 100 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	// the workflow is preparing coffee with PrepareCoffee activity
	workflow.ExecuteActivity(ctx, activities.PrepareCoffee, nil).Get(ctx, nil)
	workflow.Sleep(ctx, 5*time.Second)
	log.Println("Coffee is ready to serve")
	// Customer paid bill
	if status := signals.ReciveSignal(ctx, signals.PAYMENT_SIGNAL); !status {
		log.Println("Payment couldn't be completed! ")
		// We sent customerName to WriteAsDept activity for It can write an dept to him
		workflow.ExecuteActivity(ctx, activities.WriteAsDept, customerName).Get(ctx, nil)
	}
	// Customer took coffee
	workflow.ExecuteActivity(ctx, activities.GiveCoffee, customerName).Get(ctx, nil)
	return nil
}
