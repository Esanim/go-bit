package gobit

// IsOdd checks if a number is odd.
func IsOdd(x int) bool {
	return x&1 == 1
}

// IsEven checks if a number is even.
func IsEven(x int) bool {
	return !IsOdd(x)
}

// TheSameSigns checks if numbers a and b have the same sign.
func TheSameSigns(a, b int) bool {
	return a^b >= 0
}

// IsPowerOfTwo checks if a number is a power of two.
func IsPowerOfTwo(x int) bool {
	return x != 0 && x&(x-1) == 0
}

// LargestPowerOf2LessThanN returns the largest power of 2 that is less than a given number n.
func LargestPowerOf2LessThanN(x int) int {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++
	return x >> 1
}

// CountOnes counts the number of 1s in a bit. Uses x & (x-1) method to do so [x & (x-1) flips the rightmost 1 only].
func CountOnes(x int) uint {
	var count uint
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

// CheckIfTheBitIsSet checks if the bit k of the number n is set (1) or not (0).
func CheckIfTheBitIsSet(x int, k uint) bool {
	return (x & (1 << k)) == 1<<k
}

// GetFullSetBitOfLengthN gets the full set bit of length n-1.
func GetFullSetBitOfLengthN(x int) int {
	return (1 << uint(x)) - 1
}

// GetRightMostOneBinary returns the rightmost 1 in binary representation of x.
// Example: If x = 11010 it returns 10
func GetRightMostOneBinary(x int) int {
	return x ^ (x & (x - 1))
}

// IndexFound checks if a stringIndex is -1.
func IndexFound(stringIndex int) (bool, int) {
	return ^stringIndex != 0, stringIndex
}

// GetMaxUInt returns the largest signed integer.
func GetMaxUInt() uint {
	return ^uint(0)
}

// GetMaxInt returns the largest unsigned integer.
func GetMaxInt() int {
	return int(GetMaxUInt() >> 1)
}

// GetMinInt returns the smallest unsigned integer.
func GetMinInt() int {
	return -GetMaxInt() - 1
}

// Other Examples:

// Total calculates the total number of grains on the chessboard.
// Story: Place a grain of salt on first square of the chessboard, 2 pieces on the second, 4 of the third, 8 on the fourth etc..
// What is total amount of grains of salt on the whole chessboard ?
// 2^1 + 2^2 + ... + 2^n = 2(2^n - 1) // induction
func Total(squares uint) uint64 {
	return 2*(1<<(squares-1)-1) + 1
}
