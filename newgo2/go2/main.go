package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
	// "github.com/k0kubun/pp"
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

// ---------------------------------

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

func main() {

	pSlice := PaymentModuleWithSlice{}
	pMap := PaymentModuleWithMap{
		m: make(map[int]PaymentInfo),
	}
	//  добавляем информацию в слайс и мапу
	interations := 10000

	before := time.Now()

	for i := 0; i < interations; i++ {
		pSlice.AddInfo(PaymentInfo{

			ID:          i,
			Description: strconv.Itoa(i),
		})

	}
	sliceAddTime := time.Since(before)

	// ____________

	before = time.Now()

	for i := 0; i < interations; i++ {
		pMap.AddInfo(PaymentInfo{

			ID:          i,
			Description: strconv.Itoa(i),
		})

	}
	mapAddTime := time.Since(before)
	// --------------------
	//  ищем информацию

	before = time.Now()

	for i := 0; i < interations; i++ {

		info := pSlice.GetInfo(i)
		_ = info

	}
	sliceGetTime := time.Since(before)

	before = time.Now()

	for i := 0; i < interations; i++ {

		info := pMap.GetInfo(i)
		_ = info

	}
	mapGetTime := time.Since(before)

	proba := "PROBA"
	blue := color.New(color.FgHiBlue).SprintfFunc()
	fmt.Println(blue("%v", proba))

	// pp.Println(pSlice)
	// pp.Println(pMap)

	i1s := pSlice.GetInfo((10))
	i2m := pMap.GetInfo((10))
	fmt.Println(i1s)
	fmt.Println(i2m)

	fmt.Println("slice add", sliceAddTime)
	fmt.Println("map add", mapAddTime)

	fmt.Println("slice get", sliceGetTime)
	fmt.Println("map get", mapGetTime)
}
