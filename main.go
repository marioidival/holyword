package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Book struct {
	Abbrev   string                         `json:"abbrev"`
	Book     string                         `json:"book"`
	Chapters []map[string]map[string]string `json:"chapters"`
}

func ChooseBook(books []Book) (book Book) {
	book = books[rand.Intn(len(books))]
	return
}

func ChooseChapter(chapters []map[string]map[string]string) (chapter map[string]map[string]string) {
	chapter = chapters[rand.Intn(len(chapters))]
	return
}

func ChooseVerse(verses map[string]string) (key, verse string) {
	var keys []string
	for key, _ := range verses {
		keys = append(keys, key)
	}
	key = keys[rand.Intn(len(keys))]
	verse = verses[key]
	return
}

func main() {
	file, err := ioutil.ReadFile("src/pt_nvi.json")
	if err != nil {
		fmt.Printf("could not load file: %v\n", err)
		os.Exit(1)
	}

	var bible []Book
	json.Unmarshal(file, &bible)

	rand.Seed(time.Now().Unix())

	book := ChooseBook(bible)
	chapter := ChooseChapter(book.Chapters)

	var keys []string
	for key, _ := range chapter {
		keys = append(keys, key)
	}

	verseNum, verse := ChooseVerse(chapter[keys[0]])

	fmt.Printf("Livro: %s, Capitulo: %s, Versiculo: %s \n%s\n", book.Book, keys[0], verseNum, verse)
}
