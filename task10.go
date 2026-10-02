package main

import "fmt"

// 10. Дан срез студентов []string{"Аня", "Борис", "Вика", "Данил", "Ева"}.
// Напиши функцию, которая в отдельной горутине отправляет имена в канал одно за другим,
// а после всех имён закрывает канал. В main через range собери имена в новый срез
// и в конце выведи, сколько всего имён получено и сам итоговый срез.

func task10() {
	students := []string{"Аня", "Борис", "Вика", "Данил", "Ева"}
	ch := make(chan string)

	go func() {
		for _, s := range students {
			ch <- s
		}
		close(ch)
	}()

	var newStudents []string

	for s := range ch {
		newStudents = append(newStudents, s)
	}

	fmt.Println(len(newStudents), newStudents)
}
