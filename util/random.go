package util

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

var dictionary []string

func init() {
	dictionary, _ = readDictionary()
	fmt.Printf("dictionary loaded with %v words\n", len(dictionary))
}

func readDictionary() ([]string, error) {
	file, err := os.Open("words")
	if err != nil {
		return nil, fmt.Errorf("couldn't find dictionary file %w", err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), "\n"), nil
}

func RandomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}
