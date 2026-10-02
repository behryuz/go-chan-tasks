package main

import (
	"fmt"
	"sync"
	"time"
)

// 15. Мини-проект "Проверка сайтов": дан срез []string{"site1.ru", "site2.ru", "site3.ru", "site4.ru", "site5.ru"}.
// Напиши функцию checkSite(name string, results chan string),
// которая после паузы time.Sleep(время в зависимости от длины имени сайта,
// например time.Duration(len(name))*100*time.Millisecond)
// отправляет в канал results строку "<name> - OK".
// Запусти проверку всех сайтов одновременно горутинами на общий буферизированный канал results с ёмкостью len(сайтов),
// дождись всех через sync.WaitGroup и закрой канал в отдельной горутине после wg.Wait().
// В main через range собери и выведи результаты по мере готовности,
// а в конце — отдельно исходный порядок сайтов.
// Комментарием объясни, почему порядок результатов из канала может не совпадать с исходным порядком среза.

func checkSite(name string, results chan string) {
	time.Sleep(time.Duration(len(name)) * 100 * time.Millisecond)

	results <- name + " - OK"
}

func task15() {
	sites := []string{
		"google.com",
		"facebook.com",
		"somon.tj",
		"apple.com",
		"go.dev",
	}

	results := make(chan string, len(sites))

	var wg sync.WaitGroup

	for _, site := range sites {
		wg.Add(1)

		go func(name string) {
			defer wg.Done()
			checkSite(name, results)
		}(site)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("Результаты проверки:")
	for result := range results {
		fmt.Println(result) // сайт с меньшими символами в названии выполняется быстрее из-за логики sleep
	}

	fmt.Println("\nИсходный порядок сайтов:")
	for _, site := range sites {
		fmt.Println(site)
	}
}
