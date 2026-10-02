package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 12. Запусти две горутины-производителя на один и тот же небуферизированный канал результатов:
// первая отправляет в канал строки "A1".."A5", вторая — "B1".."B5"
// (с небольшой паузой time.Sleep между отправками, чтобы значения перемежались).
// Третья горутина через sync.WaitGroup дожидается завершения обеих и закрывает канал.
// В main через range выведи все 10 строк по мере прихода и посчитай их количество.

func task12() {
	ch := make(chan string)
	wg := &sync.WaitGroup{}

	wg.Go(func() {
		for i := 1; i <= 5; i++ {
			ch <- "A" + strconv.Itoa(i)
			time.Sleep(50 * time.Millisecond)
		}
	})

	wg.Go(func() {
		for i := 1; i <= 5; i++ {
			ch <- "B" + strconv.Itoa(i)
			time.Sleep(50 * time.Millisecond)
		}
	})

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

}
