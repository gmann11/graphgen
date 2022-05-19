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
