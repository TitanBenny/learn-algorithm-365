package sort

func bubble_sort(a []int) []int {
	for j := 0; j < len(a)-1; j++ {
		for i := 0; i < len(a)-j-1; i++ {
			if a[i] < a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a
}
