package main

import (
	"errors"
	"fmt"
)

type Product struct {
	barCode string
	value   float64
}

type POS struct {
	availableProducts []Product
}

func NewPOS() *POS {
	return &POS{
		[]Product{
			{"12345", 7.25},
			{"489", 10.0},
			{"555", 10.0},
		},
	}
}

func (p POS) Scan(barCode string) (string, error) {
	if len(barCode) == 0 {
		return "", errors.New("empty barcode")
	}

	if scannedProduct, err := p.findProduct(barCode); err != nil {
		return "", errors.New("barcode not found")
	} else {
		return p.formatValue(scannedProduct.value), nil
	}
}

func (p POS) findProduct(barCode string) (*Product, error) {
	for _, product := range p.availableProducts {
		if product.barCode == barCode {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

func (p POS) formatValue(value float64) string {
	return fmt.Sprintf("$%.2f", value)
}

func (p POS) Total(barCodes []string) string {
	sumOfProductValues := 0.0
	for _, barCode := range barCodes {
		if scannedProduct, err := p.findProduct(barCode); err == nil {
			sumOfProductValues += scannedProduct.value
		}
	}
	return p.formatValue(sumOfProductValues)
}
