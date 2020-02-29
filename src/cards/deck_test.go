package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected dek length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spade" {
		t.Errorf("Expected first card of deck is Ace of Spade, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of deck is Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSavetoFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected dek length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
