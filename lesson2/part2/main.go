package main

import "fmt"

// Ввод: 2 числа через пробел
// Вывод: произведение суммы двух чисел
// В случае если происходит какая-то ошибка, то выводим ее и завершаем программу.

func main() {
	var a, b int
	_, err := fmt.Scanf("%d %d\n", &a, &b) // 123 228
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a * b)
}
