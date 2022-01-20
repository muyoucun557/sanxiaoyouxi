package main

import "fmt"

func main() {
	boards := make([][]rune, 0)
	boards = append(boards, []rune{'0', 'R', '0', '0'})
	boards = append(boards, []rune{'B', '0', 'W', 'Y'})
	boards = append(boards, []rune{'0', 'B', '0', '0'})

	fmt.Println(F(boards))
}

func F(boards [][]rune) (int, int) {
	rows := len(boards)
	cols := len(boards[0])

	rowsScopes := make([][][]int, 0) // 存储每一行被W分割的区间
	rowsNums := make([][]int, 0)     // 存储行每个区间能被消除的方块的个数

	// 统计行的数据
	for i := 0; i < rows; i++ {
		scopes := make([][]int, 0)
		nums := make([]int, 0)
		// 用开区间
		start := -1
		num := 0 // 区间里的可消除的个数
		for j := 0; j < cols; j++ {
			if boards[i][j] == 'W' {
				scopes = append(scopes, []int{start, j})
				nums = append(nums, num)
				num = 0
				start = j
			} else {
				if boards[i][j] != '0' {
					num++
				}
			}
		}
		scopes = append(scopes, []int{start, cols})
		nums = append(nums, num)

		rowsScopes = append(rowsScopes, scopes)
		rowsNums = append(rowsNums, nums)
	}

	colsScopes := make([][][]int, 0) // 存储每一列被W分割的区间
	colsNums := make([][]int, 0)     // 存储列每个区间能被消除的方块的个数
	// 统计列的数据
	for j := 0; j < cols; j++ {
		scopes := make([][]int, 0)
		nums := make([]int, 0)
		// 用开区间
		start := -1
		num := 0 // 区间里的可消除的个数
		for i := 0; i < rows; i++ {
			if boards[i][j] == 'W' {
				scopes = append(scopes, []int{start, i})
				nums = append(nums, num)
				num = 0
				start = j
			} else {
				if boards[i][j] != '0' {
					num++
				}
			}
		}
		scopes = append(scopes, []int{start, rows})
		nums = append(nums, num)

		colsScopes = append(colsScopes, scopes)
		colsNums = append(colsNums, nums)
	}

	// 最佳位置的初始数据
	bestRow := -1
	bestCol := -1
	max := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if boards[i][j] == '0' {
				// 找到当前行能消除的方块数量
				m := 0
				rowNum := 0
				for ; m < len(rowsScopes[i]); m++ {
					if rowsScopes[i][m][0] < j && j < rowsScopes[i][m][1] {
						break
					}
				}
				if m == len(rowsScopes[i]) {
					rowNum = 0
				} else {
					rowNum = rowsNums[i][m]
				}

				// 找到当前列能消除的方块数量
				m = 0
				colNum := 0
				for ; m < len(colsScopes[i]); m++ {
					if colsScopes[i][m][0] < j && j < colsScopes[i][m][1] {
						break
					}
				}
				if m == len(colsScopes[i]) {
					colNum = 0
				} else {
					colNum = colsNums[i][m]
				}

				if rowNum+colNum > max {
					bestRow, bestCol, max = i, j, rowNum+colNum
				}
			}
		}
	}
	return bestRow, bestCol
}
