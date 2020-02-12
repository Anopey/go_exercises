package diffsquares

//SquareOfSum calculates the square of the sum of all the first n natural numbers
func SquareOfSum(n int) int {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

//SumOfSquares takes the sum of the squares of all of the first n natural numbers
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

//Difference takes the difference of SquareOfSum and SumOfSquares of the input number.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
