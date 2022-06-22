package anonim

func GetAn() func(int) int {
	return func(x int) int {
		return 2 * x
	}
}
