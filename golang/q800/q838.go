package q800

func pushDominoesPlanA(dominoes string) string {
	if len(dominoes) <= 1 {
		return dominoes
	}
	current := []rune(dominoes)
	next := []rune(dominoes)
	l := len(current)
Outer:
	for {
		current, next = next, current
		for i := 0; i < l; i++ {
			r := current[i]
			if r == '.' {
				if i == 0 {
					right := current[i+1]
					if right == 'L' {
						next[i] = right
					} else {
						next[i] = '.'
					}
				} else if i == l-1 {
					left := current[i-1]
					if left == 'R' {
						next[i] = left
					} else {
						next[i] = '.'
					}
				} else {
					left := current[i-1]
					right := current[i+1]
					if left == 'R' && right == 'L' {
						next[i] = '.'
					} else if left == right {
						next[i] = left
					} else if left == '.' && right == 'L' {
						next[i] = right
					} else if right == '.' && left == 'R' {
						next[i] = left
					} else {
						next[i] = '.'
					}
				}
			} else {
				next[i] = r
			}
		}
		for i, r := range next {
			if current[i] != r {
				continue Outer
			}
		}
		break
	}
	return string(next)
}

func pushDominoesPlanB(dominoes string) string {
	if len(dominoes) <= 1 {
		return dominoes
	}
	strs := []rune(dominoes)
	offset := 0
	l := len(strs)
	for offset < l {
		leftI := -2
		rightI := -2
		for i := offset; i < l; i++ {
			if leftI == -2 && strs[i] == '.' {
				leftI = i - 1
			} else if leftI != -2 && strs[i] != '.' {
				rightI = i
				break
			}
		}

		if leftI == -2 && rightI == -2 {
			return string(strs)
		}

		left := '.'
		right := '.'
		if leftI >= 0 {
			left = strs[leftI]
		} else {
			leftI = -1
		}
		if rightI <= l-1 && rightI >= 0 {
			right = strs[rightI]
		} else {
			rightI = l
		}

		if left == 'R' && right == 'L' {
			for i, j := leftI+1, rightI-1; i < j; i, j = i+1, j-1 {
				strs[i] = 'R'
				strs[j] = 'L'
			}
		} else if left == right {
			for i := leftI + 1; i < rightI; i++ {
				strs[i] = left
			}
		} else if left == '.' && right == 'L' {
			for i := leftI + 1; i < rightI; i++ {
				strs[i] = 'L'
			}
		} else if right == '.' && left == 'R' {
			for i := leftI + 1; i < rightI; i++ {
				strs[i] = 'R'
			}
		}
		offset = rightI + 1
	}
	return string(strs)
}
