package main

func isValidFail(s string) bool {

	det := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	srn := []rune(s)
	times := len(srn) / 2

	for i := 0; i < times; i++ {
		first := srn[i]
		last := srn[len(srn)-1-i]
		if det[first] != last {
			return false
		}

	}
	return true

}

func isValid(s string) bool {
	det := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	fl := []rune{'(', '{', '['}
	// ll := []rune{')', '}', ']'}

	inArray := func(s rune, arr []rune) bool {
		for _, v := range arr {
			if v == s {
				return true
			}
		}
		return false
	}
	stack := []rune{}

	srn := []rune(s)
	for i := 0; i < len(srn); i++ {
		if inArray(srn[i], fl) {
			stack = append(stack, srn[i])
		} else {
			if len(stack) == 0 {
				stack = append(stack, srn[i])
				continue
			}

			//fmt.Printf("stack befor %s", string(stack))
			stl := stack[len(stack)-1]
			//fmt.Printf("stack last %s", string(stl))
			//fmt.Printf("det srn %s", string(det[srn[i]]))
			if det[stl] == srn[i] {
				stack = stack[:len(stack)-1]
				//fmt.Printf("stack pop %s", string(stack))
			} else {
				stack = append(stack, srn[i])
			}
		}
	}
	// fmt.Printf("stack %s", string(stack))
	return (len(stack) == 0)
}
