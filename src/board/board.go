package board

import (
	"errors"
	"fmt"
)

type moveSet struct {
	X, Y int
}

const Xplayer = "X"
const Oplayer = "O"

var moveMapping = map[string]moveSet{
	"1": moveSet{0, 0},
	"2": moveSet{0, 1},
	"3": moveSet{0, 2},
	"4": moveSet{1, 0},
	"5": moveSet{1, 1},
	"6": moveSet{1, 2},
	"7": moveSet{2, 0},
	"8": moveSet{2, 1},
	"9": moveSet{2, 2},
}

var boardstatus = [3][3]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
}

// ResetBoard change the state of the board to the start state
func ResetBoard() {
	boardstatus = [3][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
}

// PrintState  print the current moves on the board
func PrintState() {
	fmt.Println("State of board")
	fmt.Printf(" ___ ___ ___ \n")
	fmt.Printf("| %s | %s | %s |\n", boardstatus[0][0], boardstatus[0][1], boardstatus[0][2])
	fmt.Printf(" ___ ___ ___ \n")
	fmt.Printf("| %s | %s | %s |\n", boardstatus[1][0], boardstatus[1][1], boardstatus[1][2])
	fmt.Printf(" ___ ___ ___ \n")
	fmt.Printf("| %s | %s | %s |\n", boardstatus[2][0], boardstatus[2][1], boardstatus[2][2])
}

// PrintNextMove print the prompt to provide the next move
func PrintNextMove() {
	fmt.Println("Enter the number of the square to place your piece in. For example 1. ")
	fmt.Println(" ___  ___  ___ ")
	fmt.Printf(" | %s | %s | %s |\n", boardstatus[0][0], boardstatus[0][1], boardstatus[0][2])
	fmt.Printf("  ___  ___  ___ \n")
	fmt.Printf(" | %s | %s | %s |\n", boardstatus[1][0], boardstatus[1][1], boardstatus[1][2])
	fmt.Printf("  ___  ___  ___ \n")
	fmt.Printf(" | %s | %s | %s |\n", boardstatus[2][0], boardstatus[2][1], boardstatus[2][2])
}

// PlayerMove submit a move of the player
func PlayerMove(player string, move string) (winner bool, err error) {
	moveX, moveY, err := parseMove(move)
	if err != nil {
		return false, err
	}
	boardstatus[moveX][moveY] = player
	winner = checkIfWinningMove(moveX, moveY)
	return winner, nil
}

// Tie check to see if the players have reached a tie
func Tie() bool {
	for r := 0; r <= 2; {
		for c := 0; c <= 2; {
			if boardstatus[r][c] != Xplayer && boardstatus[r][c] != Oplayer {
				return false
			}
			c++
		}
		r++
	}
	return true
}

// checkIfWinningMove given the x and y cords for the players move, check if a line if three is present on the board
// based on the x and y
func checkIfWinningMove(x int, y int) bool {
	// check rows
	for r, c := 0, 0; r <= 2; {
		if boardstatus[r][c] == boardstatus[r][c+1] && boardstatus[r][c] == boardstatus[r][c+2] && boardstatus[r][c+1] == boardstatus[r][c+2] {
			return true
		}
		r++
	}
	// check colums
	for r, c := 0, 0; c <= 2; {
		if boardstatus[r][c] == boardstatus[r+1][c] && boardstatus[r][c] == boardstatus[r+2][c] && boardstatus[r+1][c] == boardstatus[r+2][c] {
			return true
		}
		c++
	}

	// check diagnols
	var centerBoardX = 1
	var centerBoardY = 1
	// left to right diagnol
	if boardstatus[centerBoardX-1][centerBoardY-1] == boardstatus[centerBoardX][centerBoardY] && boardstatus[centerBoardX][centerBoardY] == boardstatus[centerBoardX+1][centerBoardY+1] && boardstatus[centerBoardX-1][centerBoardY-1] == boardstatus[centerBoardX+1][centerBoardY+1] {
		return true
	}
	// right to left diagnol
	if boardstatus[centerBoardX-1][centerBoardY+1] == boardstatus[centerBoardX][centerBoardY] && boardstatus[centerBoardX][centerBoardY] == boardstatus[centerBoardX+1][centerBoardY-1] && boardstatus[centerBoardX-1][centerBoardY+1] == boardstatus[centerBoardX+1][centerBoardY-1] {
		return true
	}
	return false
}

// parseMove convert the string representation move to an x and y cord.
// This function will also make sure that the input is in the correct format
func parseMove(move string) (x int, y int, err error) {
	// get move set from map
	m, ok := moveMapping[move]
	if !ok {
		return 0, 0, errors.New("Illegal move " + move + " move not in the range 1-9")
	}

	// make sure board does not already have a move here
	if boardstatus[m.X][m.Y] == Xplayer || boardstatus[m.X][m.Y] == Oplayer {
		return 0, 0, errors.New("Illegal move " + move + " postion has already been played on board")
	}
	return m.X, m.Y, nil
}
