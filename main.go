package main

import "fmt"

/*
Задача: необходимо написать функцию, которая будет принимать на вход строку (сообщение)
и маскировать там все ссылки, заменяя их на звездочки.

Правила выполнения:
- Не использовать стандартную библиотечную функцию или сторонние пакеты (strings.replace и тп)
Нужно решить задачу, только манипулируя байтами напрямую.
- Не используйте конкатенацию строк (оператор +)
- Маскировать только ссылки, начинающиеся с http://
- Не проверять наличие прописных/строчных букв

Подсказка: создайте новый байтовый срез в качестве буфера из заданного строкового аргумента.
Затем манипулируйте им во время вашей программы.
И затем распечатайте этот буфер.

Пример выполнения:
input: Hello, its my page: http://localhost123.com See you
Output: Hello, its my page: http://**************** See you

Решение залить на гитхаб, создать отдельную ветку и прислать пулл реквест.
*/

const testString string = "Hello, its my page: http://localhost123.com and another page http://localhost123234.com See you"

func maskingLinks(str string) string {
	isLink := false

	strByteSlice := make([]byte, 0, len(str))

	for index, char := range str {
		if string(char) == " " {
			isLink = false
		}

		if isLink {
			strByteSlice = append(strByteSlice, byte('*'))

			continue
		}

		if string(char) == "/" && string(str[index-1]) == "/" {
			isLink = true
		}

		strByteSlice = append(strByteSlice, byte(char))
	}

	return string(strByteSlice)
}

func main() {
	res := maskingLinks(testString)

	fmt.Println(res)
}
