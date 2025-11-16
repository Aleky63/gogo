package postman

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

func Postman(ctx context.Context, wg *sync.WaitGroup, transferPoint <-chan string, n int, mail string) {
	defer wg.Done()
	red := color.New(color.FgHiRed).SprintFunc()
	for {

		select {
		case <-ctx.Done():

			fmt.Println(red("Я почтальон номер:", n, "Мой рабочий день закончен!"))
			return
		default:
			fmt.Println("Я почтальон номер:", n, "Взял письмо!")
			time.Sleep(1 * time.Second)

			fmt.Println("Я почтальон номер:", n, "Донес письмо!")
			time.Sleep(1 * time.Second)

			transferPoint <- mail

			fmt.Println(red("Я почтальон номер:"), n, "Передал письмо!")
		}

	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, mailTransferPoint, i, postmanToMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}
func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "Семейный привет",
		2: "Приглашение",
		3: "Информация",
	}
	mail, ok := ptm[postmanNumber]

	if !ok {
		return "Лотерея"
	}
	return mail
}
