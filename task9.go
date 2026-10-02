package main

import "fmt"

// 9. Попробуй закрыть один и тот же канал дважды подряд (close(ch); close(ch))
// и отдельно — отправить значение в уже закрытый канал (close(ch); ch <- 1).
// Запусти оба случая по очереди (закомментировав один, пока проверяешь другой),
// посмотри на текст паники в терминале и комментарием рядом с каждым случаем
// кратко запиши, что говорит паника.

func task9() {
	ch := make(chan int)

	go generateNumbers(10, ch)

	for v := range ch {
		fmt.Println(v)
	}

	close(ch) // panic: close of closed channel - закрытие уже закрытого канала
	ch <- 11  // panic: send on closed channel - отправка в закрытый канал
}
