package main

import (
	"fmt"
	"gatewayAdapterReInvent/provider"
	"log"
)

func main() {
	providerA := provider.NewProviderA()
	resultA, errA := providerA.ProcessPayment(1000, "USD")
	if errA != nil {
		log.Printf("Provider A Error %s", &errA)
	} else {
		fmt.Printf("Provider A Result %vv\n", *&resultA)
	}

	providerB := provider.NewProviderA()
	resultB, errB := providerB.ProcessPayment(1000, "USD")
	if errB != nil {
		log.Printf("Provider A Error %s", &errB)
	} else {
		fmt.Printf("Provider A Result %vv\n", *&resultB)
	}
}
