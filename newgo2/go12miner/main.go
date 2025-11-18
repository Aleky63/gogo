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

	go func() {
		time.Sleep(3 * time.Second)
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 5)
	mailTransferPoint := postman.PostmanPool(postmanContext, 8)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()
	wg.Wait()
	// isCoalClosed := false
	// isMailClosed := false

	// for !isCoalClosed || !isMailClosed {

	// 	select {
	// 	case c, ok := <-coalTransferPoint:
	// 		if !ok {
	// 			isCoalClosed = true
	// 			continue

	// 		}

	// 		coal += c

	// 	case m, ok := <-mailTransferPoint:
	// 		if !ok {
	// 			isMailClosed = true
	// 			continue

	// 		}
	// 		mails = append(mails, m)
	// 	}

	// }
	red := color.New(color.FgHiRed).SprintFunc()
	fmt.Println(red("СУММАРНО ДОБЫТЫЙ УГОЛЬ:", coal.Load()))

	mtx.Lock()
	fmt.Println(red("СУММАРНОЕ КОЛИЧЕСТВО ПОЛУЧЕННЫХ ПИСЕМ:", len(mails)))
	mtx.Unlock()
}
