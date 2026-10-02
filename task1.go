package main

import "fmt"

// 1. Заведи небуферизированный канал ch := make(chan string).
// В горутине отправь в него строку "Привет из горутины!" (ch <- "..."),
// а в main получи значение (v := <-ch) и выведи его.

func task1() {
	ch := make(chan string)

	go func() {
		ch <- "Привет из горутины!"
	}()

	fmt.Println(<-ch)
}
