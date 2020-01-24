package main

import "testing"

import "os"

func TestNewDeck(t *testing.T)  {
  d := newDeck()

  if len(d) != 16 {
    t.Errorf("Expected deck length of 16 but got %v", len(d))
  }
}
func TestFirstCard(t *testing.T)  {
  d := newDeck()
  if d[0] != "Ace of Spades" {
    t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
  }
}
func TestLastCard(t *testing.T)  {
  d := newDeck()
  if d[len(d) - 1] != "Four of Clubs" {
    t.Errorf("Expected first card of Four of Clubs, but got %v", d[0])
  }
}

func TestSaveToDeckAndNewDeckTestFile(t *testing.T)  {
  os.Remove("_decktesting")

  deck := newDeck()
  deck.saveToFile("_decktesting")

  loadedDeck := newDeckFromFile("_decktesting")

  if len(loadedDeck) != 16 {
    t.Errorf("Expected deck length of 16 but got %v", len(deck))
  }

  os.Remove("_decktesting")
}