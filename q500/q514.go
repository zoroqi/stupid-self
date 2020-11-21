package q500

func FindRotateSteps(ring string, key string) int {
	return 10
}

func FindRotateSteps1(ring string, key string) int {
	bs := []byte(ring)
	length := len(bs)
	rotate := func(start int, count int, lr bool) int {
		if lr {
			// 右转, 0->max
			return (start + count) % length
		} else {
			// 左转 max->0
			return (start - count + length) % length
		}
	}

	minRotate := func(b byte, start int) (int, index int) {
		if bs[start] == b {
			return 0, start
		}
		l := length / 2
		for i := 1; i <= l; i++ {
			if b == bs[rotate(start, i, true)] {
				return i, rotate(start, i, true)
			} else if b == bs[rotate(start, i, false)] {
				return i, rotate(start, i, false)
			}
		}
		return 0, start
	}

	if len(key) == 1 {
		c, _ := minRotate(key[0], 0)
		return c + 1
	}

	count := len(key)
	index := 0
	c := 0
	for _, v := range []byte(key) {
		c, index = minRotate(v, index)
		count += c
	}

	return count
}
