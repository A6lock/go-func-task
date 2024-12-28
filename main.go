package main

import (
	"bufio"
	"fmt"
	"os"
)

const prefix = "http://"
const prefixLength = len(prefix)

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

func maskingLinks(inputStr string) string {
	isLink := false

	strByteSlice := []byte(inputStr)

	for i := 0; i < len(strByteSlice)-prefixLength; i++ {
		if string(strByteSlice[i]) == " " {
			isLink = false
		}

		if isLink {
			strByteSlice[i] = '*'

			continue
		}

		if string(inputStr[i:i+prefixLength]) == prefix {
			isLink = true

			i += prefixLength - 1
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
