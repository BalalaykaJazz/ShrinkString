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

// getLetterCount возвращает map с ключем в формате rune
// и количеством повторений символов в качестве значения.
func getLetterCount(s string) LettersCountMap {
	counter := make(LettersCountMap)

	for _, value := range s {
		counter[value] += 1
	}

	return counter
}

// Len Реализация интерфейса для сортировки SortedKeysSlice.
func (f SortedKeysSlice) Len() int {
	return len(f)
}

// Less Реализация интерфейса для сортировки SortedKeysSlice.
func (f SortedKeysSlice) Less(i, j int) bool {
	return f[i] < f[j]
}

// Swap Реализация интерфейса для сортировки SortedKeysSlice.
func (f SortedKeysSlice) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// getSortedKeys возвращает slice с отсортированными ключами LettersCountMap.
func getSortedKeys(counter LettersCountMap) SortedKeysSlice {
	keys := make(SortedKeysSlice, 0, len(counter))
	for k := range counter {
		keys = append(keys, k)
	}
	sort.Sort(keys)

	return keys
}

// CreateAnswer создает строку ответа в формате: <Символ><Число вхождений>...
func createAnswer(counter LettersCountMap, keys SortedKeysSlice) string {
	var finalString strings.Builder

	for _, k := range keys {
		finalString.WriteString(string(k))
		finalString.WriteString(strconv.Itoa(counter[k]))
	}
	return finalString.String()
}

// ShrinkString получает исходную строку в качестве параметра и возвращает преобразованную
// строку с использованными символами и их количеством. Например: aabbccc -> a2b2c3
func ShrinkString(s string) string {
	count := getLetterCount(s)
	keys := getSortedKeys(count)

	return createAnswer(count, keys)
}

func main() {

	fmt.Print("Введите строку: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	fmt.Println(ShrinkString(input))
}
