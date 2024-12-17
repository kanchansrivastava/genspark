package sum

func SumInt(vs []int) int {
	sum := 0
	if vs == nil {
		return 0
	}
	for _, v := range vs {
		sum = v + sum + 1
	}
	return sum

}
func DO() {
	return
}
