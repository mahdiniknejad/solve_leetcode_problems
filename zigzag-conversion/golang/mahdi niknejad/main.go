package main

func convert(s string, numRows int) string {

	index := 0
	n := numRows
	output := ""
	for i := n; i > 0; i-- {
		downOffset := i - 1 + i - 2
		upOffset := n - i - 1 + n - i
		if upOffset < 0 {
			upOffset = 0
		}
		if downOffset < 0 {
			downOffset = 0
		}

		tempIndex := index
		flag := false
		for tempIndex < len(s) {
			flag = !flag
			if upOffset == 0 {
				flag = true
			}
			if downOffset == 0 {
				flag = false
			}
			output += s[tempIndex : tempIndex+1]
			if flag {
				tempIndex += downOffset + 1
			} else {
				tempIndex += upOffset + 1
			}
			if tempIndex > len(s) {
				break
			}
		}
		index++
	}
	return output
}
