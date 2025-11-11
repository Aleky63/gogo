package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(("FOO end 😒"))
			return
		default:
			fmt.Println(("FOO"))
		}

		time.Sleep(200 * time.Millisecond)
	}
}
func boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(("BOO end 😒"))
			return
		default:
			fmt.Println(("BOO"))
		}

		time.Sleep(200 * time.Millisecond)
	}
}
func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())

	childContext, childCancel := context.WithCancel(parentContext)
	go foo(parentContext)
	go boo(childContext)
	time.Sleep(1 * time.Second)
	childCancel()
	time.Sleep(1 * time.Second)
	parentCancel()
	time.Sleep(4 * time.Second)
	fmt.Println("End---------------")
}
