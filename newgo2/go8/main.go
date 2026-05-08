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

func foo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("endEnd_foo", n)
			return
		default:
			fmt.Println(red("maksim", n))
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
			fmt.Println(green("kat"))
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)

	go foo(parentContext, 1)
	go foo(parentContext, 2)
	go foo(parentContext, 3)
	go foo(parentContext, 4)
	go boo(childContext)

	time.Sleep(1 * time.Second)
	childCancel()
	time.Sleep(3 * time.Second)
	parentCancel()

	time.Sleep(1 * time.Second)
	fmt.Println("END😍END😍END😍END😍END")
}
