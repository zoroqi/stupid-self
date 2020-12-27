package other

import (
	"math/rand"
	"time"
)

// 猴子排序
// 验证 -> 交换两个元素 -> 验证 -> 交换 ....
func MonkeySort(arr []int) (a []int, loopCount int) {
	a = arr
	if isOrder(arr) {
		return
	}
	length := len(arr)
	rand.Seed(time.Now().UnixNano())
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	for !isOrder(arr) {
		loopCount++
		r1 := abs(rand.Int() % length)
		r2 := abs(rand.Int() % length)
		arr[r1], arr[r2] = arr[r2], arr[r1]
	}
	return
}

func isOrder(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

// *   如果最后一个值小于第一个值，则交换这两个数
// *   如果当前集合元素数量大于等于3：
//		1.  使用臭皮匠排序法排序前2/3的元素
//		2.  使用臭皮匠排序法排序后2/3的元素
//		3.  再次使用臭皮匠排序法排序前2/3的元素
func StoogeSort(arr []int) (a []int, loopCount int) {
	loopCount = stoogeSort(arr, 0, len(arr)-1)
	a = arr
	return
}

func stoogeSort(arr []int, l, r int) int {
	if arr[l] > arr[r] {
		arr[l], arr[r] = arr[r], arr[l]
	}
	n := 1
	if (r - l + 1) >= 3 {
		sl := (r - l + 1) / 3
		n += stoogeSort(arr, l, r-sl)
		n += stoogeSort(arr, l+sl, r)
		n += stoogeSort(arr, l, r-sl)
	}
	return n
}

/**
在[珠排序](https://zh.wikipedia.org/wiki/%E7%8F%A0%E6%8E%92%E5%BA%8F)中，一行（row）表示一个数字。如果一行里有2颗珠子，该行代表数字2；如果一行里有4颗珠子，该行代表数字4。当给定一个数组，数组里有多少个数字，就要有多少行；数组里最大的数字是几，就要准备多少根杆子。

准备就绪后，释放珠子，珠子按重力下落，就完成了排序。

珠排序可以类比于珠子在平行的竖直杆上滑动，就像算盘一样。然而，每一竖直杆都有珠子数目的限制。因此，初始化就相当于在竖直的杆上悬挂珠子，在第一步中，排列就被显示为n=5行的珠子在m=4列队竖直杆上。每一行右边的数字意味着该行在问题中被表示的数；第1，2行表示正整数3（因为它们都有3个珠子）而顶层的一行表示正整数2（因为它只含有2个珠子）。

如果我们要允许珠子掉落，那么每行表示已排序的整数。第1行表示在集合中最大的数，而第n行表示最小的数。如果按照前面提到的规则（行包含一系列在竖直杆1到k的珠子，并且让k+1到m竖直杆都空），那么它会出现这种情况。

允许珠子掉落的行为在物理意义上就是允许珠子从高的行掉落至低的行。如果被行a表示的值小于被行a+1表示的值，那么一些珠子就会从a+1掉落至a；因为行a不包含足够的珠子防止珠从a+1行掉落，所以这一定会发生。

用机械设备实现的珠排序类似于计数排序；每一杆上的数字与那些在所有数中等于或大于该数字的数量相当。

理论上**无论是电子还是实物上的实现，珠排序都能在O(n)时间内完成.** 仅仅是理论上

时间复杂度

* O(1)：即所有珠子都同时移动，但这种算法只是概念上的，无法在计算机中实现。 **最快的**
* O(\sqrt{n})：在真实的物理世界中用引力实现，所需时间正比于珠子最大高度的平方根，而最大高度正比于n。
* O(n)：一次移动一行珠子，可以用模拟和数字的硬件实现。
* O(S)，S是所有输入数据的和：一次移动一个珠子，能在软件中实现。

*/
func BeadSort(arr []uint) (a []uint, loopCount int) {
	var findMax func([]uint, uint) uint
	// 这里只是不想用循环来实现查找最大值
	findMax = func(ar []uint, before uint) uint {
		loopCount++
		if len(ar) == 0 {
			return before
		}
		if ar[0] > before {
			return findMax(ar[1:], ar[0])
		}
		return findMax(ar[1:], before)
	}
	length := len(arr)
	max := findMax(arr, 0)
	// 构建算盘空间
	abacus := make([][]bool, length)
	for i, v := range arr {
		abacus[i] = make([]bool, max)
		loopCount++
		for j := uint(0); j < v; j++ {
			loopCount++
			abacus[i][j] = true
		}
	}

	// 算盘滑落
	// 重新赋值
	//for i := uint(0); i < max; i++ {
	//	loopCount++
	//	beadCount := 0
	//	for j := 0; j < length; j++ {
	//		loopCount++
	//		if abacus[j][i] {
	//			beadCount++
	//		}
	//	}
	//	nilCount := length - beadCount
	//	for j := 0; j < length; j++ {
	//		loopCount++
	//		if j < nilCount {
	//			abacus[j][i] = false
	//		} else {
	//			abacus[j][i] = true
	//		}
	//	}
	//}
	// 交换式
	for i := uint(0); i < max; i++ {
		l, r := 0, length-1
		for l < r {
			loopCount++
			lt := length
			for l < r {
				loopCount++
				if abacus[l][i] {
					lt = l
					break
				}
				l++
			}
			rf := 0
			for l < r {
				loopCount++
				if !abacus[r][i] {
					rf = r
					break
				}
				r--
			}
			if lt < rf {
				abacus[lt][i] = false
				abacus[rf][i] = true
			}
		}
	}
	a = arr
	count := func(a []bool) uint {
		for i, v := range a {
			loopCount++
			if !v {
				return uint(i)
			}
		}
		return uint(len(a))
	}
	for i, v := range abacus {
		loopCount++
		a[i] = count(v)
	}

	return
}
