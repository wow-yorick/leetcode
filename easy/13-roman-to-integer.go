package main

func romanToInt(s string) int {
	romp := map[string]int{}
	romp["I"] = 1
	romp["V"] = 5
	romp["X"] = 10
	romp["L"] = 50
	romp["C"] = 100
	romp["D"] = 500
	romp["M"] = 1000
	romp["IV"] = 4
	romp["IX"] = 9
	romp["XL"] = 40
	romp["XC"] = 90
	romp["CD"] = 400
	romp["CM"] = 900

	res := 0
	chars := []rune(s)
	i := 0
	for i < len(chars) {
		if integer, ok := romp[string(chars[i:i+2])]; ok {
			res += integer
			i += 2
		} else {
			res += romp[string(chars[i])]
			i++
		}
	}
	return res
}
