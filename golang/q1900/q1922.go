package q1900

const countGoodNubmersMod = 1e9 + 7

func countGoodNumbers(n int64) int {
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % countGoodNubmersMod
			}
			x = x * x % countGoodNubmersMod
		}
		return res
	}

	if n%2 == 0 {
		return pow(20, int(n/2))
	} else {
		return (pow(20, int(n/2)) * 5) % countGoodNubmersMod
	}
	return 0
}
