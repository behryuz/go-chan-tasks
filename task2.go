package main

import "fmt"

// 2. Напиши функцию square(n int, ch chan int), которая кладёт в канал n*n.
// Запусти её как горутину для числа 7,
// в main получи результат из канала и выведи "7 в квадрате = 49".

func square(n int, ch chan int) {
	ch <- n * n
}

func task2() {
	ch := make(chan int)

	go square(7, ch)

	fmt.Printf("7 в квадрате = %v\n", <-ch)
}
