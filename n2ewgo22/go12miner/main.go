package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fatih/color"
)

func main() {
	var coal atomic.Int64
	mtx := sync.Mutex{}
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	initTime := time.Now()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("🔔 Время работы шахтёров истекло — останавливаю добычу.")
		minerCancel()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("👌 Время работы почтальонов истекло — останавливаю доставку.")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 5)

	mailTransferPoint := postman.PostmanPool(postmanContext, 5)

	wg := &sync.WaitGroup{}

	wg.Go(func() {
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	})

	wg.Go(func() {
		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	})
	wg.Wait()

	red := color.New(color.FgHiRed).SprintFunc()
	green := color.New(color.FgHiGreen).SprintFunc()
	fmt.Println(red("😊😊__СУММАРНО ДОБЫТЫЙ УГОЛЬ:", coal.Load()))

	mtx.Lock()

	fmt.Println(red("😊😊__СУММАРНОЕ КОЛИЧЕСТВО ПОЛУЧЕННЫХ ПИСЕМ:", len(mails)))

	mtx.Unlock()

	fmt.Println(green(" ❤️ __ ЗАТРАЧЕННОЕ ВРЕМЯ: ", time.Since(initTime)))
	fmt.Println("------------")
}
