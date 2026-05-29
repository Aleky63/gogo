package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postman(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- string,
	n int,
	mail string,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():

			fmt.Println("I'm a postman number:", n, "the working day is over")
			return
		default:
			fmt.Println("I'm a postman number:", n, "took the letter")
			time.Sleep(1 * time.Second)

			fmt.Println("I'm a postman number:", n, "delivered the letter to the post office", mail)
			transferPoint <- mail
			fmt.Println("I'm a postman number:", n, "handed over the letter", mail)
		}
	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}
	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "family greetings",
		2: "tax notice",
		3: "Hello from Trump",
	}

	mail, ok := ptm[postmanNumber]
	if !ok {
		return "lotto"
	}
	return mail
}
