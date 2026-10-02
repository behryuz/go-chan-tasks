package main

import "fmt"

// 11. Заведи буферизированный канал ch := make(chan int, 10).
// Запусти горутину-производителя, которая кладёт в канал квадраты чисел от 1 до 10
// и закрывает канал по завершении.
// В main через range посчитай сумму всех полученных чисел и выведи её.

func task11() {
	ch := make(chan int, 10)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i * i
		}
		close(ch)
	}()

	sum := 0
	for v := range ch {
		sum += v
	}

	fmt.Println("sum:", sum)
}
