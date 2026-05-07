package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	red   = color.New(color.FgHiRed).SprintFunc()
	green = color.New(color.FgHiGreen).SprintFunc()
)

func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("endEnd_foo")
			return
		default:
			fmt.Println(red("maksim"))
		}

		time.Sleep(1 * time.Second)
	}
}

func boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("endEnd_boo")
			return
		default:
			fmt.Println("kat")
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)

	go foo(parentContext)
	go boo(childContext)

	time.Sleep(3 * time.Second)
	childCancel()
	time.Sleep(6 * time.Second)
	parentCancel()

	time.Sleep(1 * time.Second)
	fmt.Println(green("END😍END😍END😍END😍END"))
}
