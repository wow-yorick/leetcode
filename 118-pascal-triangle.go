package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	ret := generate(numRows - 1)
	baseRow := ret[len(ret)-1]

	tempRow := []int{1}
	for i := 0; i+1 < len(baseRow); i++ {
		tempRow = append(tempRow, baseRow[i]+baseRow[i+1])
	}
	tempRow = append(tempRow, 1)

	ret = append(ret, tempRow)
	return ret
}
