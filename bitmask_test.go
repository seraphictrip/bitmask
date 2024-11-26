package bm

import (
	"fmt"
	"strconv"
	"testing"
)

// BitMask
var ConstantsTests = []struct {
	symbol             string
	constant, expected Bits
}{
	{"0", 0, 0},
	{"1", 1, 1 << 0},
	{"F0", F0, 1},
	{"F0", F0, 0x1},
	{"F1", F1, 2},
	{"F1", F1, 1 << 1},
	{"F2", F2, 4},
	{"F2", F2, 1 << 2},
	{"F3", F3, 8},
	{"F3", F3, 1 << 3},
	{"F4", F4, 16},
	{"F4", F4, 1 << 4},
	{"F5", F5, 32},
	{"F5", F5, 1 << 5},
	{"F6", F6, 64},
	{"F6", F6, 1 << 6},
	{"F7", F7, 128},
	{"F7", F7, 1 << 7},
}

func TestConstants(t *testing.T) {
	for i, e := range ConstantsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if e.constant != e.expected {
				t.Fatalf("%v: %08b != %08b", e.symbol, e.constant, e.expected)
			}
		})
	}
}

// Set = b | flag
var SetTests = []struct {
	b, flag, expected Bits
}{
	// 00000000
	{},
	// 00000001
	{0, F0, F0},
	// 00000010
	{0, F1, F1},
	// 00000100
	{0, F2, F2},
	// 00001000
	{0, F3, F3},
	// 00010000
	{0, F4, F4},
	// 00100000
	{0, F5, F5},
	// 01000000
	{0, F6, F6},
	// 10000000
	{0, F7, F7},
	// 00000001 (set 1 to 1/noop)
	{1, F0, F0},
	// 00000011 (2+1)
	{1, F1, 3},
	{1, F1, 1 | F1},
	// 00000101 (4+1)
	{1, F2, 5},
	{1, F2, 1 | F2},
	// 00001001 (8+1)
	{1, F3, 9},
	// 00010001 (16+1)
	{1, F4, 17},
	// 00100001 (32+1)
	{1, F5, 33},
	// 01000001 (64+1)
	{1, F6, 65},
	// 10000001 (128+1)
	{1, F7, 129},
	{128, F0, 129},
	// 10000010 (128+2)
	{128, F1, 130},
	// 10000100 (128+4)
	{128, F2, 132},
	// 10001000 (128+8)
	{128, F3, 136},
	// 10010000 (128+16)
	{128, F4, 144},
	// 10100000 (128+32)
	{128, F5, 160},
	// 11000000 (128+64)
	{128, F6, 192},
	// 10000000 (128)
	{128, F7, 128},
}

func TestSet(t *testing.T) {
	for i, e := range SetTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := Set(e.b, e.flag)
			if actual != e.expected {
				t.Fatalf("Set(%08b, %08b) = %08b, want %08b", e.b, e.flag, actual, e.expected)
			}
			fmt.Printf("%08b\n", actual)
			bitwiseOr := e.b | e.flag
			if bitwiseOr != e.expected {
				t.Fatalf("(%08b | %08b) = %08b, want %08b", e.b, e.flag, actual, e.expected)
			}
		})
	}
}

// Clear, bitclear, and not &^
var ClearTests = []struct {
	b, flag, expected Bits
}{
	{},
	// 11111111
	{255, 0, 255},
	// 11111110
	{255, F0, 254},
	// 11111101
	{255, F1, 253},
	// 11111011
	{255, F2, 251},
	// 11110111
	{255, F3, 247},
	// 11101111
	{255, F4, 239},
	// 11011111
	{255, F5, 255 - 32},
	// 10111111
	{255, F6, 255 - 64},
	// 01111111
	{255, F7, 255 - 128},
}

func TestClear(t *testing.T) {
	for i, e := range ClearTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := Clear(e.b, e.flag)
			if actual != e.expected {
				t.Fatalf("Clear(%08b, %08b) = %08b, want %08b", e.b, e.flag, actual, e.expected)
			}
			fmt.Printf("%08b\n", actual)
			bitclear := e.b &^ e.flag
			if bitclear != e.expected {
				t.Fatalf("(%08b &^ %08b) = %08b, want %08b", e.b, e.flag, actual, e.expected)
			}
		})
	}
}

var ToggleTests = []struct {
	b, flag, expected Bits
}{
	{},
	// 00000000
	{255, 255, 0},
	// 11111111
	{0, 255, 255},
	// 00000001
	{0, F0, 1},
	// 00000010
	{0, F1, 2},
	// 00000100
	{0, F2, 4},
	// 00001000
	{0, F3, 8},
	// 00010000
	{0, F4, 16},
	// 00100000
	{0, F5, 32},
	// 01000000
	{0, F6, 64},
	// 10000000
	{0, F7, 128},
	{255, 0, 255},
	// 11111110
	{255, F0, 254},
	// 11111101
	{255, F1, 253},
	// 11111011
	{255, F2, 251},
	// 11110111
	{255, F3, 247},
	// 11101111
	{255, F4, 239},
	// 11011111
	{255, F5, 255 - 32},
	// 10111111
	{255, F6, 255 - 64},
	// 01111111
	{255, F7, 255 - 128},
}

func TestToggle(t *testing.T) {
	for i, e := range ToggleTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := Toggle(e.b, e.flag)
			if actual != e.expected {
				t.Fatalf("Toggle(%08b, %08b) = %08b, want %08b", e.b, e.flag, actual, e.expected)
			}
			fmt.Printf("%08b\n", actual)
			not := e.b ^ e.flag
			if not != e.expected {
				t.Fatalf("(%08b ^ %08b) = %08b, want %08b", e.b, e.flag, actual, e.expected)
			}
		})
	}
}

var HasTests = []struct {
	b, flag  Bits
	expected bool
}{
	{},
	// no bitmask Has 0 mask
	// 0 set is not subset of any set
	{0, 0, false},
	{255, 0, false},
	// every set has self
	{F0, F0, true},
	{F1, F1, true},
	{F2, F2, true},
	{F3, F3, true},
	{F4, F4, true},
	{F5, F5, true},
	{F6, F6, true},
	{F7, F7, true},
	// 255 is superset
	{255, F0, true},
	{255, F1, true},
	{255, F2, true},
	{255, F3, true},
	{255, F4, true},
	{255, F5, true},
	{255, F6, true},
	{255, F7, true},
}

func TestHas(t *testing.T) {
	for i, e := range HasTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := Has(e.b, e.flag)
			if result != e.expected {
				t.Fatalf("Has(%08b, %08b) = %v, %v", e.b, e.flag, result, e.expected)
			}
		})
	}
}

var ArithmeticShiftTests = []struct {
	input    int8
	shift    int8
	expected int8
}{
	{-127, 1, 0},
	{127, 1, 0},
}

func TestArithmeticShift(t *testing.T) {
	for i, e := range ArithmeticShiftTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if e.shift >= 0 {
				result := e.input >> e.shift
				fmt.Printf("%08b; %v\n", result, result)
			}
		})
	}
}
