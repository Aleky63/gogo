package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
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


// --------------------------------------

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


// ----------------------------------






func main() {

	pSlice := PaymentModuleWithSlice{}
	pMap := PaymentModuleWithMap{
		m: make(map[int]PaymentInfo),
	}
	//  добавляем информацию в слайс и мапу
	interations := 20000

	before := time.Now()

	for i := range interations {
		pMap.AddInfo(PaymentInfo{

			ID:          i,
			Description: strconv.Itoa(i),
		})

	}
	mapAddTime := time.Since(before)

	before = time.Now()

	for i := range interations {
		pSlice.AddInfo(PaymentInfo{

			ID:          i,
			Description: strconv.Itoa(i),
		})

	}
	sliceAddTime := time.Since(before)
	// --------------------
	//  ищем информацию

	before = time.Now()

	for i := range interations {

		info := pMap.GetInfo(i)
		_ = info

	}
	mapGetTime := time.Since(before)

	before = time.Now()

	for i := range interations {

		info := pSlice.GetInfo(i)
		_ = info

	}
	sliceGetTime := time.Since(before)

	proba := "---PROBA---"
	blue := color.New(color.FgHiBlue).SprintFunc()
	red := color.New(color.FgHiRed).SprintFunc()
	fmt.Println(proba)

	i1s := pSlice.GetInfo((10))
	i2m := pMap.GetInfo((10))
	fmt.Println(i1s)
	fmt.Println(i2m)


fmt.Println(blue("map add-- ", mapAddTime))
fmt.Println(blue("map get-- ", mapGetTime)) 

	fmt.Println(red("slice add-- ", sliceAddTime))
	fmt.Println(red("slice get-- ", sliceGetTime))
	
}
