package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("~(OwO)~")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Data processing
	processedData := strings.Split(string(data), "\n")

	player1Deck := []int{}
	player2Deck := []int{}
	player := 0
	for _, line := range processedData {
		if strings.Contains(line, "Player") {
			player++
		}
		number, err := strconv.Atoi(string(line))
		if err != nil {
			continue
		}
		if player == 1 {
			player1Deck = append(player1Deck, number)
		} else {
			player2Deck = append(player2Deck, number)
		}

	}
	fmt.Println(player1Deck, player2Deck)
	player1DeckOG := deepCopy(player1Deck)
	player2DeckOG := deepCopy(player2Deck)

	// Part 1

	for len(player1Deck) != 0 && len(player2Deck) != 0 {

		player1Card := player1Deck[0]
		player2Card := player2Deck[0]
		player1Deck = player1Deck[1:]
		player2Deck = player2Deck[1:]
		if player1Card > player2Card {
			player1Deck = append(player1Deck, player1Card)
			player1Deck = append(player1Deck, player2Card)
		} else {
			player2Deck = append(player2Deck, player2Card)
			player2Deck = append(player2Deck, player1Card)
		}
	}

	fmt.Println(player1Deck, player2Deck)

	winningDeck := []int{}
	if len(player1Deck) > len(player2Deck) {
		winningDeck = player1Deck
	} else {
		winningDeck = player2Deck
	}
	fmt.Println(calculateWinner(winningDeck))

	// Part 2
	player1Deck, player2Deck = player1DeckOG, player2DeckOG
	fmt.Println("Part 2 decks", player1Deck, player2Deck)
	winner, score := recursiveCombat(player1Deck, player2Deck)
	fmt.Println(winner, score)
}

func recursiveCombat(deck1, deck2 []int) (int, int) {

	roundsPlayed := map[string]bool{}
	winningDeck := []int{}
	round := 0
	var gameWinner int
	for gameWinner == 0 {
		fmt.Println("player1:", deck1)
		fmt.Println("player2:", deck2)
		fmt.Println()
		// Check if game has been played before
		deckstring := ""
		for _, card := range deck1 {
			deckstring += fmt.Sprint(card) + " "
		}
		deckstring += " | "
		for _, card := range deck2 {
			deckstring += fmt.Sprint(card) + " "
		}
		_, playedBefore := roundsPlayed[deckstring]
		if playedBefore {
			// fmt.Println("Round played before")
			gameWinner = 1
			break
		} else {
			roundsPlayed[deckstring] = true
		}
		// Check for normal winners
		if len(deck2) == 0 {
			gameWinner = 1
			winningDeck = deck1
			break
		}
		if len(deck1) == 0 {
			gameWinner = 2
			winningDeck = deck2
			break
		}

		round++
		// fmt.Println("-------")
		// fmt.Println("round", round)

		var handWinner int

		card1 := deck1[0]
		card2 := deck2[0]
		deck1 = deck1[1:]
		deck2 = deck2[1:]

		// fmt.Println("Player 1 deck", deck1)
		// fmt.Println("Player 2 deck", deck2)
		// fmt.Println("Player 1 plays", card1)
		// fmt.Println("Player 2 plays", card2)
		// Check for recursive Combat
		if card1 <= len(deck1) && card2 <= len(deck2) {
			deck1Copy := deepCopy(deck1[0:card1])
			deck2Copy := deepCopy(deck2[0:card2])
			handWinner, _ = recursiveCombat(deck1Copy, deck2Copy)
		} else {
			if card1 > card2 {
				handWinner = 1
			} else if card2 > card1 {
				handWinner = 2
			}
		}

		if handWinner == 1 {
			deck1 = append(deck1, card1)
			deck1 = append(deck1, card2)
			// fmt.Println("deck 1 added", card1, card2)
		} else if handWinner == 2 {
			deck2 = append(deck2, card2)
			deck2 = append(deck2, card1)
			// fmt.Println("deck 2 added", card2, card1)
		} else {
			panic("Noone won")
		}
	}
	//  fmt.Println("Winning Deck", len(winningDeck), winningDeck)
	return gameWinner, calculateWinner(winningDeck)
}

func calculateWinner(winningDeck []int) int {
	sum := 0
	for idx, num := range winningDeck {
		sum += num * (len(winningDeck) - idx)
	}
	return sum
}

func deepCopy(deck []int) []int {
	output := []int{}
	for _, num := range deck {
		output = append(output, num)
	}
	return output
}
