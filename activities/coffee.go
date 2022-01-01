package activities

import (
	"context"
	"log"
)

func PrepareCoffee(ctx context.Context) error {
	log.Println("Coffee is preparing...")
	return nil
}

func GiveCoffee(ctx context.Context, customerName string) error {
	log.Printf("%s adlı müşteriye kahve teslim edildi.", customerName)
	return nil
}

func WriteAsDept(ctx context.Context, customerName string) error {
	log.Printf("%s borrowed a coffee", customerName)
	return nil
}
