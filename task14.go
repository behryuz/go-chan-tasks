package main

import (
	"fmt"
	"sync"
	"time"
)

// 14. Реализуй простой семафор на буферизированном канале chan struct{} с ёмкостью 2,
// чтобы одновременно "работали" не больше 2 горутин из 6 запущенных:
// перед началом работы горутина отправляет struct{}{} в канал-семафор (если он полон — ждёт),
// после работы (time.Sleep(500*time.Millisecond) как имитация работы)
// освобождает место, читая один элемент из семафора.
// Выведи для каждой горутины момент начала и конца работы (time.Since от старта программы)
// и убедись по логам, что одновременно работают не больше 2.

func task14() {
	start := time.Now()

	semaphore := make(chan struct{}, 2)

	var wg sync.WaitGroup

	for i := 1; i <= 6; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			semaphore <- struct{}{}

			fmt.Printf("Горутина %d: начало работы: %v\n",
				id, time.Since(start))

			time.Sleep(500 * time.Millisecond)

			fmt.Printf("Горутина %d: конец работы: %v\n",
				id, time.Since(start))

			<-semaphore
		}(i)
	}

	wg.Wait()
}
