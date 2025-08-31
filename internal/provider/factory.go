package provider

import "fmt"

import "gatewayAdapterReInvent/provider"

func NewProvider(name string) (provider.PaymentProvider, error) {
	switch name {
	case "A":
		return provider.NewProviderA(), nil
	case "B":
		return provider.NewProviderB(), nil
	default:
		return nil, fmt.Errorf("unknown Provider: %s", name)
	}
}
