package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	sumMap := buildWordSumMap()
	findWordWithValueOfThreeNineteen(sumMap)
	totalWordsWithOddSum(sumMap)
	for i, arg := range args {
		res := fmt.Sprintf(`#%v Arg has a sum of %v`, i+1, letterSum(arg))
		fmt.Println(res)
	}
}

func letterSum(value string) int {
	total := 0
	baseChar := int('a') - 1
	for _, letter := range value {
		total += (int(letter) - baseChar)
	}

	return total
}

func getWordList() []string {
	resp, err := http.Get("https://raw.githubusercontent.com/dolph/dictionary/master/enable1.txt")

	if err != nil {
		log.Fatal(err)
	}

	res, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	words := strings.Split(string(res[:]), "\n")

	return words
}

func buildWordSumMap() map[int][]string {
	wordList := getWordList()
	sumMap := make(map[int][]string)

	for _, word := range wordList {
		value := letterSum(word)
		if val, ok := sumMap[value]; ok {
			val = append(val, word)
		} else {
			sumMap[value] = []string{word}
		}
	}

	return sumMap
}

func findWordWithValueOfThreeNineteen(sumMap map[int][]string) {
	threeNineteen := 319
	if val, ok := sumMap[threeNineteen]; ok {
		for _, word := range val {
			res := fmt.Sprintf(`Found %q which is has a sum of 319`, word)
			fmt.Println(res)
		}
	}
}

func totalWordsWithOddSum(sumMap map[int][]string) {
	total := 0
	for key, value := range sumMap {
		if key%2 != 0 {
			total += len(value)
		}
	}
	res := fmt.Sprintf(`There are a total of %x words with an odd value`, total)
	fmt.Println(res)
}
