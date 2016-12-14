package main

import (
	"fmt"

	"github.com/kpovoc/lib/playingcards"
)

func main() {

	var player1Hand [7]playingcards.Card
	deck := playingcards.NewDeck()

	for i := 0; i < 7; i++ {
		player1Hand[i] = deck.DrawCard()
	}

	for i := 0; i < 7; i++ {
		fmt.Println(player1Hand[i].GetName())
	}

	something(player1Hand)

}

func something(pcopy [7]playingcards.Card) {
	sortHandHighToLow(pcopy[:])
	is, weight, pair := checkForOnePair(pcopy[:])
	if is == true {
		fmt.Println("Hand (One Pair): ", is, weight, "\nAt cards ", pair[0], pair[1])
	}
}

/***************************
 * Hand Checking Functions *
 **************************/
func checkForOnePair(pcopy []playingcards.Card) (is bool, highWeight int, pairSlice []int) {
	var pair [2]int
	for i := 0; i < 6; i++ {
		if pcopy[i].GetWeight() == pcopy[i+1].GetWeight() {
			is = true
			highWeight = pcopy[i].GetWeight()
			pair[0] = i
			pair[1] = i + 1
			pairSlice = pair[:]
			return
		}
	}
	return
}

/**************************
 * Hand Sorting Functions *
 *************************/
// quicksort Hand from high to low (Ace High)
func sortHandHighToLow(pcopy []playingcards.Card) {
	lo := 0
	hi := len(pcopy) - 1
	quicksort(pcopy, lo, hi)

	fmt.Println("\nSortedHand")
	for i := 0; i < 7; i++ {
		fmt.Println(pcopy[i].GetName())
	}
}

// Nico Lomuto quicksort partition scheme
func quicksort(pcopy []playingcards.Card, lo int, hi int) {
	if lo < hi {
		p := partition(pcopy, lo, hi)
		quicksort(pcopy, lo, p-1)
		quicksort(pcopy, p+1, hi)
	}
}

func partition(pcopy []playingcards.Card, lo int, hi int) int {
	pivot := pcopy[hi].GetWeight()
	i := lo
	for j := lo; j < hi; j++ {
		if pcopy[j].GetWeight() >= pivot {
			temp := pcopy[i]
			pcopy[i] = pcopy[j]
			pcopy[j] = temp
			i++
		}
	}
	temp := pcopy[i]
	pcopy[i] = pcopy[hi]
	pcopy[hi] = temp
	return i
}
