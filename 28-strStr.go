package main

func strStr(haystack string, needle string) int {
	if haystack == "" && needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}

	hs := []byte(haystack)
	nes := []byte(needle)
	nesLen := len(nes)
	for i := 0; i < len(hs); i++ {
		if nesLen+i > len(hs) {
			return -1
		}
		cs := hs[i : nesLen+i]
		flag := true
		for j := 0; j < nesLen; j++ {
			if cs[j] != nes[j] {
				flag = false
				break
			}
		}
		if flag == true {
			return i
		}
	}

	return -1
}
