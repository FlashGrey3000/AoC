package main

func X_MAS(i, j int, mat []string) (bool, int) {
	cnt := 0
	rows := len(mat)
	cols := len(mat[0]) // technically rows and columns are same so we can just take one variable dimension
	
	if i-1>=0 && j-1>=0 && j+1<cols && i+1<rows && (string(mat[i-1][j-1])+string(mat[i][j])+string(mat[i+1][j+1]) == "MAS" || string(mat[i-1][j-1])+string(mat[i][j])+string(mat[i+1][j+1]) == "SAM") {
		if i-1 >= 0 && j+1 < cols && i+1 < rows && j-1 >= 0 && (string(mat[i-1][j+1])+string(mat[i][j])+string(mat[i+1][j-1]) == "MAS" || string(mat[i-1][j+1])+string(mat[i][j])+string(mat[i+1][j-1]) == "SAM") {
			cnt++
		}
	}

	return cnt > 0, cnt
}