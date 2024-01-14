package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	EmptyCell = " "
	PlayerX   = "X"
	PlayerO   = "O"
	BoardSize = 8
)

type Board [BoardSize][BoardSize]string

func initializeBoard() Board {
	board := Board{}
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			board[i][j] = EmptyCell
		}
	}
	board[3][3] = PlayerX
	board[3][4] = PlayerO
	board[4][3] = PlayerO
	board[4][4] = PlayerX
	return board
}

func printBoard(board Board) {
	fmt.Println("  0 1 2 3 4 5 6 7")
	for i := 0; i < BoardSize; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < BoardSize; j++ {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	board := initializeBoard()
	player := PlayerX
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Current player: %s\n", player)
		printBoard(board)
		fmt.Print("Enter row and column (e.g., '2 3'): ")
		scanner.Scan()
		input := scanner.Text()
		coords := strings.Split(input, " ")

		if len(coords) != 2 {
			fmt.Println("Invalid input. Please enter row and column.")
			continue
		}

		// Parse row and column
		row, col := 0, 0
		fmt.Sscanf(coords[0], "%d", &row)
		fmt.Sscanf(coords[1], "%d", &col)

		// Check if the move is valid and update the board
		if isValidMove(board, row, col, player) {
			board = makeMove(board, row, col, player)
			player = togglePlayer(player)
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
}

func isValidMove(board Board, row, col int, player string) bool {
	// Check if the cell is empty
	if board[row][col] != EmptyCell {
		return false
	}
	// Check if the move flips any opponent pieces
	// Implement your Reversi move validation logic here
	return true
}

func makeMove(board Board, row, col int, player string) Board {
	// Update the cell with the player's piece
	board[row][col] = player

	// Define the eight possible directions to check for opponent's pieces to flip
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// Iterate through each direction to check for opponent's pieces to flip
	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		x, y := row+dx, col+dy
		oppPiecesToFlip := []int{}

		// Check if there are opponent's pieces in the current direction
		for x >= 0 && x < BoardSize && y >= 0 && y < BoardSize && board[x][y] != EmptyCell && board[x][y] != player {
			oppPiecesToFlip = append(oppPiecesToFlip, x*BoardSize+y)
			x += dx
			y += dy
		}

		// If there are opponent's pieces to flip in this direction, flip them
		if len(oppPiecesToFlip) > 0 && x >= 0 && x < BoardSize && y >= 0 && y < BoardSize && board[x][y] == player {
			for _, pos := range oppPiecesToFlip {
				r, c := pos/BoardSize, pos%BoardSize
				board[r][c] = player
			}
		}
	}

	return board
}

func togglePlayer(player string) string {
	if player == PlayerX {
		return PlayerO
	} else {
		return PlayerX
	}
}
