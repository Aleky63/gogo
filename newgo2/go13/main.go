package main

import (
	"context"
	"fmt"

	"go13/miner"
	"go13/postman"

	"github.com/fatih/color"
)

func main() {
	var coal int
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	coalTransferPoint := miner.MinerPool(minerContext, 2)

	mailTransferPoint := postman.PostmanPool(postmanContext, 2)

	isCoalClosed := false
	isMailClosed := false

	for !isCoalClosed || !isMailClosed {
		select {
		case c, ok := <-coalTransferPoint:

			if !ok {
				isCoalClose = true
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
	red := color.New(color.FgRed).SprintFunc()
	fmt.Println(red("total coal produced:", coal))
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Println(yellow("total number of letters received:", len(mails)))
}
