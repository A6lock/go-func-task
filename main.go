package main

import (
	"bufio"
	"fmt"
	"os"
)

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

//const testString string = "Hello, its my page: http://localhost123.com and another page http://localhost123234.com See you //case"

func isProtocol(text string) bool {
	fmt.Println("condition is protocol", text)

	return text == "http://"
}

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

		strByteSlice = append(strByteSlice, byte(char))

		if string(char) == "/" && strByteSlice[index-1] == '/' {
			isLink = isProtocol(string(strByteSlice[index-6 : index+1]))
		}
	}

	return string(strByteSlice)
}

func main() {
	inputString := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите строку")
	inputString.Scan()

	res := maskingLinks(inputString.Text())

	fmt.Println(res)
}
