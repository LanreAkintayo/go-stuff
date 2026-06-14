package main

import "fmt"

type Product struct {
	price          float64
	quantityBought int
}

type SaleMetric struct {
	totalRevenue float64
	totalSold    int
}

type Sale struct {
	vendor   string
	category string
	product  Product
}

func ProcessSales(allSales []Sale) map[string]map[string]SaleMetric {
	saleMetrics := make(map[string]map[string]SaleMetric)

	for _, sale := range allSales {
		if _, vendorExists := saleMetrics[sale.vendor]; !vendorExists {
			saleMetrics[sale.vendor] = make(map[string]SaleMetric)
		}

		currentSaleMetric := saleMetrics[sale.vendor][sale.category]
		currentProduct := sale.product

		saleMetrics[sale.vendor][sale.category] = SaleMetric{
			totalRevenue: currentSaleMetric.totalRevenue + (currentProduct.price * float64(currentProduct.quantityBought)),
			totalSold:    currentSaleMetric.totalSold + currentProduct.quantityBought,
		}
	}

	return saleMetrics
}

func main() {
	allSales := []Sale{
		{vendor: "AlabaExpress", category: "Electronics", product: Product{price: 50_000.0, quantityBought: 2}},
		{vendor: "GbagadaStores", category: "Fashion", product: Product{price: 15_000.0, quantityBought: 1}},
		{vendor: "AlabaExpress", category: "Electronics", product: Product{price: 45_000.0, quantityBought: 1}},
		{vendor: "GbagadaStores", category: "Fashion", product: Product{price: 12_000.0, quantityBought: 3}},
		{vendor: "AlabaExpress", category: "Fashion", product: Product{price: 2_000.0, quantityBought: 5}},
	}

	processedSales := ProcessSales(allSales)

	fmt.Println(processedSales)
}
