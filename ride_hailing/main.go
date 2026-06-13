
package main

import "fmt"

type FareCalculator interface {
	CalculateFare(distance float64) float64
}

type EconomyRide struct {
}

func (ecoRide EconomyRide) CalculateFare(distance float64) float64 {
	return distance * 200.0
}

type PremiumRide struct {
	baseFare float64
}

func (premRide PremiumRide) CalculateFare(distance float64) float64 {
	return premRide.baseFare + distance * 400.0
}

func PrintReceipt(fc FareCalculator, kms float64) {
	fare := fc.CalculateFare(kms)
	fmt.Printf("Print receipt is %.2f\n", fare)
}

func main() {
	economyRide := EconomyRide{}
	premiumRide := PremiumRide{baseFare: 15.5}

	PrintReceipt(economyRide, 10)
	PrintReceipt(premiumRide, 10)
}