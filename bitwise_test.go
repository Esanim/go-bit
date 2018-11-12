package gobit

import (
	"testing"
)

func TestIsOdd(t *testing.T) {
	x := 3
	var expected = true
	if actual := IsOdd(x); actual != expected {
		t.Errorf("IsOdd expected %t, Actual %t", expected, actual)
	}
	x = 4
	expected = false
	if actual := IsOdd(x); actual != expected {
		t.Errorf("IsOdd expected %t, Actual %t", expected, actual)
	}
}

func TestIsEven(t *testing.T) {
	x := 4
	var expected = true
	if actual := IsEven(x); actual != expected {
		t.Errorf("IsEven expected %t, Actual %t", expected, actual)
	}
}

func TestTheSameSigns(t *testing.T) {
	a := 2
	b := -2
	var expected = false
	if actual := TheSameSigns(a, b); actual != expected {
		t.Errorf("TheSameSigns expected %t, Actual %t", expected, actual)
	}

	b = 3
	expected = true
	if actual := TheSameSigns(a, b); actual != expected {
		t.Errorf("TheSameSigns expected %t, Actual %t", expected, actual)
	}
}

func TesIsPowerOfTwo(t *testing.T) {
	x := 64
	var expected = true
	if actual := IsPowerOfTwo(x); actual != expected {
		t.Errorf("IsPowerOfTwo expected %t, Actual %t", expected, actual)
	}

	x = 13
	expected = false
	if actual := IsPowerOfTwo(x); actual != expected {
		t.Errorf("IsPowerOfTwo expected %t, Actual %t", expected, actual)
	}
}

func TestLargestPowerOf2LessThanN(t *testing.T) {
	x := 20
	expected := 16
	if actual := LargestPowerOf2LessThanN(x); actual != expected {
		t.Errorf("LargestPowerOf2LessThanN expected %d, Actual %d", expected, actual)
	}

	x = 4
	expected = 2
	if actual := LargestPowerOf2LessThanN(x); actual != expected {
		t.Errorf("LargestPowerOf2LessThanN expected %d, Actual %d", expected, actual)
	}

	x = 82
	expected = 64
	if actual := LargestPowerOf2LessThanN(x); actual != expected {
		t.Errorf("LargestPowerOf2LessThanN expected %d, Actual %d", expected, actual)
	}
}

func TestCountOnes(t *testing.T) {
	x := 37 // 100101
	var expected uint = 3
	if actual := CountOnes(x); actual != expected {
		t.Errorf("CountOnes expected %d, Actual %d", expected, actual)
	}

	x = 1
	expected = 1
	if actual := CountOnes(x); actual != expected {
		t.Errorf("CountOnes expected %d, Actual %d", expected, actual)
	}

	x = 9 // 1001
	expected = 2
	if actual := CountOnes(x); actual != expected {
		t.Errorf("CountOnes expected %d, Actual %d", expected, actual)
	}
}

func TestCheckIfTheBitIsSet(t *testing.T) {
	var k uint = 1
	x := 3 // 10
	var expected = true
	if actual := CheckIfTheBitIsSet(x, k); actual != expected {
		t.Errorf("CheckIfTheBitIsSet expected %t, Actual %t", expected, actual)
	}

	x = 13
	expected = false
	if actual := CheckIfTheBitIsSet(x, k); actual != expected {
		t.Errorf("CheckIfTheBitIsSet expected %t, Actual %t", expected, actual)
	}
}

func TestGetFullSetBitOfLengthN(t *testing.T) {
	x := 4
	var expected = 15 // 1111
	if actual := GetFullSetBitOfLengthN(x); actual != expected {
		t.Errorf("GetFullSetBitOfLengthN expected %d, Actual %d", expected, actual)
	}

	x = 3
	expected = 7 //111
	if actual := GetFullSetBitOfLengthN(x); actual != expected {
		t.Errorf("GetFullSetBitOfLengthN expected %d, Actual %d", expected, actual)
	}
}

func TestGetRightMostOneBinary(t *testing.T) {
	x := 13
	var expected = 1 // from 1101 get 1
	if actual := GetRightMostOneBinary(x); actual != expected {
		t.Errorf("GetRightMostOneBinary expected %d, Actual %d", expected, actual)
	}

	x = 10
	expected = 2 // from 1010 get 10
	if actual := GetRightMostOneBinary(x); actual != expected {
		t.Errorf("GetRightMostOneBinary expected %d, Actual %d", expected, actual)
	}
}

func TestIndexFound(t *testing.T) {
	x := -1
	var expected = false
	if actual, _ := IndexFound(x); actual != expected {
		t.Errorf("IndexFound expected %t, Actual %t", expected, actual)
	}

	x = 0
	expected = true
	if actual, _ := IndexFound(x); actual != expected {
		t.Errorf("IndexFound expected %t, Actual %t", expected, actual)
	}
}

func TestTotal(t *testing.T) {
	var expected uint64 = 18446744073709551615
	if actual := Total(64); actual != expected {
		t.Errorf("Total() expected %d, Actual %d", expected, actual)
	}
}
