// Start the execution of the tic tac toe game
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"board"
)

type player struct {
	wins  int
	piece string
}

// initialize players
var playerX = player{0, board.Xplayer}
var playerO = player{0, board.Oplayer}

// initialize board
func init() {
	board.ResetBoard()
}

func main() {
	// determine who moves first
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	moveFirstDeterminer := r1.Intn(2)
	// begin game
	startGame(moveFirstDeterminer)

}

func startGame(moveFirstDetermin int) {
	var playerOrder [2]player
	winner := false
	roundWinnerIndex := -1
	// set order for players
	if moveFirstDetermin == 0 {
		playerOrder[0] = playerX
		playerOrder[1] = playerO
	} else {
		playerOrder[0] = playerO
		playerOrder[1] = playerX
	}
	// start game loop, infinite loop
	for {
		// keep track of how many moves
		moveCount := 0

		// singe game instance
		for {
			// first player move
			winner = getPlayerMove(playerOrder[0])
			moveCount++
			if winner {
				roundWinnerIndex = 0
				break
			}
			if moveCount == 9 {
				fmt.Println("Game has ended in a tie!")
				moveCount = 0
				break
			}
			// second player move
			winner = getPlayerMove(playerOrder[1])
			moveCount++
			if winner {
				roundWinnerIndex = 1
				break
			}

		}
		// increment player wins
		if winner {
			playerOrder[roundWinnerIndex].wins++
			fmt.Printf("Player %s won! Current win count: %d\n", playerOrder[roundWinnerIndex].piece, playerOrder[roundWinnerIndex].wins)
		}
		board.PrintState()
		// reset state to new game
		board.ResetBoard()
		winner = false
		// swap player move order to reflect winnner and lose
		switch roundWinnerIndex {
		// -1 represents a tie
		case 0, -1:
			placeholder := playerOrder[1]
			playerOrder[1] = playerOrder[0]
			playerOrder[0] = placeholder

		case 1:
			placeholder := playerOrder[0]
			playerOrder[0] = playerOrder[1]
			playerOrder[1] = placeholder

		}

	}
}

// given a player, get their move from stdin and check if they won or not
func getPlayerMove(p player) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Player %s's move\n", p.piece)
		board.PrintNextMove()
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		win, err := board.PlayerMove(p.piece, text)
		if err != nil {
			fmt.Println(err)
		} else {
			return win
		}
	}

}
