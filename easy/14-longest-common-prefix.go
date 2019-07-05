package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) < 2 {
		return strs[0]
	}
	resn := []rune{}
	frsn := []rune(strs[0])
	for i := 0; i < len(frsn); i++ {
		t := frsn[i]
		f := true
		for j := 1; j < len(strs); j++ {
			sbn := []rune(strs[j])
			if len(sbn) <= i || sbn[i] != t {
				f = false
				break
			}
		}
		if f {
			resn = append(resn, t)
		} else {
			break
		}
	}
	return string(resn)
}
