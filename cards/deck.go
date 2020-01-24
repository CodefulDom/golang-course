package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, errorObject := ioutil.ReadFile(filename)
	if errorObject != nil {
		fmt.Println("Error:", errorObject)
		os.Exit(1)
	}
	stringSlice := strings.Split(string(byteSlice), ",")
	return deck(stringSlice)
}

func (d deck) cardRandomize() {
	for index := range d {
		source := rand.NewSource(time.Now().UnixNano())
		randomGeneratedNumber := rand.New(source)
		newPosition := randomGeneratedNumber.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}

/*
leftover card values
, "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"
*/
