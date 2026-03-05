package question36

func isValidSudoku(board [][]byte) bool {

	rows := make([]int, 9)
	cols := make([]int, 9)
	boxes := make([]int, 9)

	for idxRow := 0; idxRow < 9; idxRow++ {
		for idxCol := 0; idxCol < 9; idxCol++ {
			// meet . to continue
			if board[idxRow][idxCol] == '.' {
				continue
			}

			// get num, box Index, mask
			num := int(board[idxRow][idxCol] - '0')
			boxIdx := getBoxIndex(idxRow, idxCol)

			mask := 1 << num

			// check rows, cols, boxed
			if (rows[idxRow]&mask != 0) || (cols[idxCol]&mask != 0) || (boxes[boxIdx]&mask != 0) {
				return false
			}

			// save to rows, cols, boxes
			rows[idxRow] |= mask
			cols[idxCol] |= mask
			boxes[boxIdx] |= mask
		}
	}

	return true
}

func getBoxIndex(row int, col int) int {
	return (row/3)*3 + col/3
}
