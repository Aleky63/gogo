package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"go13/miner"
	"go13/postman"

	"github.com/fatih/color"
)

func main() {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	var coal atomic.Int64

	mtx := sync.Mutex{}
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(2 * time.Second)

		minerCancel()
	}()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("<<<--------------Working day postman the  is over----------")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 5)

	mailTransferPoint := postman.PostmanPool(postmanContext, 5)

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

	fmt.Println(red("total coal produced:", coal.Load()))
	mtx.Lock()
	fmt.Println(yellow("total number of letters received:", len(mails)))
	mtx.Unlock()
}
