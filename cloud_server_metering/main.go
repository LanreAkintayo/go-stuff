package main

import "fmt"

type UsageLog struct {
	tenant  string
	service string
	cost    float64
}

func AggregateBilling(logs []UsageLog) map[string]map[string]float64 {
	aggregate := make(map[string]map[string]float64)
	for _, log := range logs {
		if _, tenantExist := aggregate[log.tenant]; !tenantExist {
			aggregate[log.tenant] = make(map[string]float64)
		}
		aggregate[log.tenant][log.service] += log.cost

	}
	return aggregate
}

func main() {
	usageLogs := []UsageLog{
		{tenant: "Paystack", service: "Database", cost: 120.0},
		{tenant: "Mtn", service: "Storage", cost: 300.0},
		{tenant: "Paystack", service: "Storage", cost: 80.0},
		{tenant: "Paystack", service: "Database", cost: 50.0},
		{tenant: "Mtn", service: "Storage", cost: 150.0},
	}

	aggregateBillings := AggregateBilling(usageLogs)

	fmt.Println(aggregateBillings)
}
