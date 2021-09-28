package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LettersCountMap map[int32]int
type SortedKeysSlice []int32

func getLetterCount(s string) LettersCountMap {
	// Функция возвращает map следующего вида:
	// ключ - символ строки в формате rune
	// значение - количество повторений символа.

	counter := make(LettersCountMap)

	for _, value := range s {
		counter[value] += 1
	}

	return counter
}

func (f SortedKeysSlice) Len() int {
	// Реализация интерфейса для сортировки SortedKeysSlice.
	return len(f)
}

func (f SortedKeysSlice) Less(i, j int) bool {
	// Реализация интерфейса для сортировки SortedKeysSlice.
	return f[i] < f[j]
}

func (f SortedKeysSlice) Swap(i, j int) {
	// Реализация интерфейса для сортировки SortedKeysSlice.
	f[i], f[j] = f[j], f[i]
}

func getSortedKeys(counter LettersCountMap) SortedKeysSlice {
	// Функция возвращает slice с отсортированными ключами LettersCountMap.

	keys := make(SortedKeysSlice, 0, len(counter))
	for k := range counter {
		keys = append(keys, k)
	}
	sort.Sort(keys)

	return keys
}

func CreateAnswer(counter LettersCountMap, keys SortedKeysSlice) string {
	// Функция создает строку ответа в формате: <Символ><Число вхождений>...

	var finalString strings.Builder

	for _, k := range keys {
		finalString.WriteString(string(k))
		finalString.WriteString(strconv.Itoa(counter[k]))
	}
	return finalString.String()
}

func ShrinkString(s string) string {
	// Функция получает строку в качестве параметра и возвращает преобразованную
	// строку с использованными символами и их количеством.
	// Например: aabbccc -> a2b2c3

	count := getLetterCount(s)
	keys := getSortedKeys(count)

	return CreateAnswer(count, keys)
}

func main() {

	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	fmt.Println(ShrinkString(input))
}
