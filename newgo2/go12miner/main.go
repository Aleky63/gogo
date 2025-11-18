package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {

	var coal int
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

	coalTransferPoint := miner.MinerPool(minerContext, 3)
	mailTransferPoint := postman.PostmanPool(postmanContext, 3)

	isCoalClosed := false
	isMailClosed := false

	for !isCoalClosed || isMailClosed {

		select {
		case c, ok := <-coalTransferPoint:
			if !ok {
				isCoalClosed = true
				continue

			}

			coal += c

		case m, ok := <-mailTransferPoint:
			if !ok {
				isMailClosed = true
				continue

			}
			mails = append(mails, m)
		}

	}
	red := color.New(color.FgHiRed).SprintFunc()
	fmt.Println(red("СУММАРНО ДОБЫТЫЙ УГОЛЬ:", coal))
	fmt.Println(red("СУММАРНОЕ КОЛИЧЕСТВО ПОЛУЧЕННЫХ ПИСЕМ:", len(mails)))

}
