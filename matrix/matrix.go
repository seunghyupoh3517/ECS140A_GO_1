package matrix

// If needed, you may define helper functions here.

// AreAdjacent returns true iff a and b are adjacent in lst.
func AreAdjacent(lst []int, a, b int) bool {
	if lst == nil {
		return false
	}

	for i := 0; i < len(lst) - 1; i++ {
		if lst[i] == a && lst[i + 1] == b {
			return true
		}
		if lst[i] == b && lst[i + 1] == a {
			return true
		}
	}
	return false
}

// Transpose returns the transpose of the 2D matrix mat.
func Transpose(mat [][]int) [][]int {
	if len(mat) == 0 {
		return mat
	}

	rowNum := len(mat[0])
	colNum := len(mat)

	t_mat := make([][]int, rowNum)
	for i := range t_mat {
		t_mat[i] = make([]int, colNum)
	}

	for i := range t_mat {
		for j := range t_mat[0] {
			t_mat[i][j] = mat[j][i]
		}
	}
	return t_mat
}

// AreNeighbors returns true iff a and b are Manhattan neighbors in the 2D
// matrix mat.
func AreNeighbors(mat [][]int, a, b int) bool {
	if len(mat) == 0{
		return false
	}

	rowNum := len(mat)
	colNum := len(mat[0])

	for i := 0; i < rowNum; i++ {
		if i == rowNum - 1 {
			for j := 0; j < colNum - 1; j++ {
					if mat[i][j] == a && mat[i][j + 1] == b {
						return true
					}
					if mat[i][j] == b && mat[i][j + 1] == a {
						return true
					}
			}
			continue
		}

		for j := 0; j < colNum - 1; j++ {
			if mat[i][j] == a && mat[i + 1][j] == b {
				return true
			}
			if mat[i][j] == b && mat[i + 1][j] == a {
				return true
			}
			if mat[i][j] == a && mat[i][j + 1] == b {
				return true
			}
			if mat[i][j] == b && mat[i][j + 1] == a {
				return true
			}
		}

		if mat[i][colNum - 1] == a && mat[i + 1][colNum - 1] == b {
			return true
		}
		if mat[i][colNum - 1] == b && mat[i + 1][colNum - 1] == a {
			return true
		}
	}

	return false
}
