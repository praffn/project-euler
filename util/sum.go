package util

func Sum(v []int) (t int) {
	for _, n := range v {
		t += n
	}
	return t
}
