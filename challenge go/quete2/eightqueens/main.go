package main

func main() {
	EightQueens()
}

func EightQueens() {
	var board [8]int
	placeQueen(&board, 0)
}

func placeQueen(board *[8]int, row int) {
	if row == 8 {
		printBoard(board)
		return
	}
	for i := 0; i < 8; i++ {
		if checkPosition(board, row, i) {
			board[row] = i
			placeQueen(board, row+1)
		}
	}
}

func checkPosition(board *[8]int, row int, column int) bool {
	for i := 0; i < row; i++ {
		if board[i] == column || (i-row) == (board[i]-column) || (i-row) == (column-board[i]) {
			return false
		}
	}
	return true
}

func printBoard(board *[8]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i] == j {
				print("Q ")
			} else {
				print(". ")
			}
		}
		println()
	}
	println()
}
