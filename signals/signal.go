package signals

import (
	"context"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
	"log"
)

const (
	PAYMENT_SIGNAL = "payment_signal"
)

func SendPaymentSignal(workflowID string, paymentStatus bool) (err error) {
	temporalClient, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
		return
	}

	err = temporalClient.SignalWorkflow(context.Background(), workflowID, "", PAYMENT_SIGNAL, paymentStatus)
	if err != nil {
		log.Fatalln("Error signaling client", err)
		return
	}

	return nil
}

func ReciveSignal(ctx workflow.Context, signalName string) (paymentStatus bool) {
	workflow.GetSignalChannel(ctx, signalName).Receive(ctx, &paymentStatus)
	return
}
