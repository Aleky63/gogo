package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type PaymentInfo struct {
	ID          int
	Description string
	Usd         int
	Cancelled   bool
}

type PaymentModuleWithMap struct {
	m map[int]PaymentInfo
}

func (m *PaymentModuleWithMap) AddInfo(info PaymentInfo) {
	m.m[info.ID] = info
}

func (m *PaymentModuleWithMap) GetInfo(id int) PaymentInfo {
	info, ok := m.m[id]
	if !ok {
		return PaymentInfo{}
	}
	return info
}

// ------------------------

func main() {
	info1 := PaymentInfo{
		ID:          10,
		Description: "Kasha",
		Usd:         5,
		Cancelled:   false,
	}

	// pSlice := PaymentModuleWithSlice{}
	pMap := PaymentModuleWithMap{
		m: make(map[int]PaymentInfo),
	}

	pMap.AddInfo(info1)

	pp.Println("pMap", pMap)

	i2 := pMap.GetInfo((10))

	fmt.Println("i2", i2)
}
