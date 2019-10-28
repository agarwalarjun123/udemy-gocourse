package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(n int) (deck, deck) {

	if len(d) < n {
		return d, deck{}
	}

	return d[:n], d[n:]

}

func newDeck() deck {

	d := deck{}
	cardSuit := []string{"Diamond", "Spades", "Hearts", "Club"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardType := range cardSuit {
		for _, cardNo := range cardValue {
			d = append(d, cardNo+" of "+cardType)
		}
	}
	return d
}
func (d deck) deckToString() string {
	return strings.Join(d, ",")
}

func (d deck) saveDeck() error {

	err := ioutil.WriteFile("objectFile.txt", []byte(d.deckToString()), 0777)

	if err != nil {
		return err
	}
	return nil
}

func readDeck(filePath string) (deck, error) {
	bs, err := ioutil.ReadFile(filePath)
	if err == nil {
		return deck(strings.Split(string(bs), ",")), err
	}
	return deck{}, err

}
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		index := r.Intn(52)
		d[index], d[i] = d[i], d[index]
	}
}
