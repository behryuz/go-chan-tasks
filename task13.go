package main

import (
	"fmt"
)

// 13. Сделай простой двухэтапный конвейер (pipeline) на каналах:
// функция generator(nums ...int) chan int создаёт канал,
// в отдельной горутине кладёт туда переданные числа и закрывает канал;
// функция square(in chan int) chan int создаёт новый канал, в горутине читает числа из in через range,
// кладёт в новый канал их квадраты и тоже закрывает его.
// В main вызови generator(1, 2, 3, 4, 5),
// передай результат в square и через range выведи все квадраты.

func generator(nums ...int) chan int {
	ch := make(chan int)
	go func() {
		for _, n := range nums {
			ch <- n
		}
		close(ch)
	}()

	return ch
}

func squareChan(in chan int) chan int {
	ch := make(chan int)
	go func() {
		for n := range in {
			ch <- n * n
		}
		close(ch)
	}()

	return ch
}

func task13() {
	in := generator(1, 2, 3, 4, 5)
	ch := squareChan(in)

	for v := range ch {
		fmt.Println(v)
	}

}
