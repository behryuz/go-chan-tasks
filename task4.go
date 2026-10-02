package main

import "fmt"

// 4. Заведи буферизированный канал ch := make(chan int, 3)
// и без единой горутины положи в него три числа подряд (ch <- 1, ch <- 2, ch <- 3),
// затем прочитай и выведи все три. Выведи также len(ch) и cap(ch) после заполнения и после чтения.

func task4() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("len:", len(ch), ", cap:", cap(ch))

	fmt.Println("value:", <-ch)
	fmt.Println("value:", <-ch)
	fmt.Println("value:", <-ch)

	fmt.Println("len:", len(ch), ", cap:", cap(ch))
}
