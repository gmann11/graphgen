package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

var dictionary []string

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	dictionary, _ = readDictionary()
	log.Printf("dictionary loaded with %v words", len(dictionary))
}

func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func readDictionary() ([]string, error) {
	file, err := os.Open("words")
	if err != nil {
		return nil, fmt.Errorf("couldn't find dictionary file %w", err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), "\n"), nil
}

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}
