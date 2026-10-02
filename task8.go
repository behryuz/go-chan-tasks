package main

import "fmt"

// 8. Напиши функцию generateNumbers(n int, ch chan int),
// которая в цикле кладёт в канал числа от 1 до n, а после цикла закрывает канал через close(ch).
// Запусти её как горутину, а в main прочитай числа через range ch (без ручного comma-ok) и выведи их все.

func generateNumbers(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)
}

func task8() {
	ch := make(chan int)

	go generateNumbers(10, ch)

	for v := range ch {
		fmt.Println(v)
	}
}
