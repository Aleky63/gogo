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

type PaymentModuleWithSlice struct {
	s []PaymentInfo
}

func (m *PaymentModuleWithSlice) AddInfo(info PaymentInfo) {
	m.s = append(m.s, info)
}

func (m *PaymentModuleWithSlice) GetInfo(id int) PaymentInfo {
	for _, info := range m.s {
		if info.ID == id {
			return info
		}
	}
	return PaymentInfo{}
}

// ------------------------------

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

	pSlice := PaymentModuleWithSlice{}
	pMap := PaymentModuleWithMap{
		m: make(map[int]PaymentInfo),
	}

	pSlice.AddInfo(info1)
	pMap.AddInfo(info1)

	pp.Println("pSlice", pSlice)
	pp.Println("pMap", pMap)

	i1 := pSlice.GetInfo(10)
	i2 := pMap.GetInfo((10))
	fmt.Println("i1", i1)
	fmt.Println("i2", i2)
}
