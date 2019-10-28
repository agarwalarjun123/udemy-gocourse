package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	a := newDeck()
	if len(a) != 52 {
		t.Errorf("Expected deck length of 52 but got")
	}
}
func TestIO(t *testing.T) {
	d := newDeck()
	d.saveDeck()
	_, err := os.Open("objectFile.txt")
	if err != nil {
		t.Errorf("File not exists or does not have the right permission")
	}
	_, err = readDeck("objectFile.txt")
	if err != nil {
		t.Errorf("error in reading from file")
	}
}
