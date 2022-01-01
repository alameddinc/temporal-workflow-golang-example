package starters

import (
	"context"
	"github.com/alameddinc/temporal-workflow-golang-example/workflows"
	"go.temporal.io/sdk/client"
)

func StartWorkflowFunc(workflowID string, customerName string) {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		panic(err)
	}
	defer c.Close()
	opt := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "worker-group-1",
	}
	ctx := context.Background()
	if _, err := c.ExecuteWorkflow(ctx, opt, workflows.CoffeeShopWorkflow, customerName); err != nil {
		panic(err)
	}
}
