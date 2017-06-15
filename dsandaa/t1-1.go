package dsandaa

import (
	"math/rand"
	"time"
)

//新建一个1~n的n长度的int数组，然后随机打乱顺序输出
func getrandint(n int) []int {
	is := []int{}
	for i := 1; i <= n; i++ {
		is = append(is, i)
	}
	//打乱
	for j := 0; j <= n-1; j++ {
		r1 := rand.New(rand.NewSource(time.Now().UnixNano() + int64(j)))
		r2 := rand.New(rand.NewSource(time.Now().AddDate(1, 0, 0).UnixNano() + int64(j)))
		a := r1.Intn(n)
		b := r2.Intn(n)
		is[a], is[b] = is[b], is[a]
	}
	return is
}

//冒泡排序法
func bubbleSort(a []int) []int {
	for j := 0; j < len(a)-1; j++ {
		for i := 0; i < len(a)-j-1; i++ {
			if a[i] < a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a
}

//用冒泡排序法进行从大到小排序，然后取出第k个值和消耗时间
func getK(n, k int) (int, float64) {
	begin := time.Now()
	list := getrandint(n)
	r := bubbleSort(list)
	if len(r) > k {
		return r[k], time.Since(begin).Seconds()
	}
	return 0, 0
}
