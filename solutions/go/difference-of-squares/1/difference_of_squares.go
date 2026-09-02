package differenceofsquares

func SquareOfSum(n int) int {
	s := (n * (n + 1)) >> 1
    return s * s 
}

func SumOfSquares(n int) int {
    return (n * (n + 1) * (n << 1 + 1)) / 6 
}

func Difference(n int) int {
	if n <= 0 {
        return 0
    }

    return SquareOfSum(n) - SumOfSquares(n)
}
