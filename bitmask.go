package bm

// see: https://yourbasic.org/golang/bitmask-flag-set-clear/
// this is code described there and tests in bitmask_test.go to re-enforce understanding
type Bits uint8

const (
	F0 Bits = 1 << iota
	F1
	F2
	F3
	F4
	F5
	F6
	F7
)

// Set bits of b found in flag to 1
// Set(00000000, 10000001) > 10000001
func Set(b, flag Bits) Bits { return b | flag }

// Clear bits of b in flag
// Clear(11111111, 10000001) => 01111110
func Clear(b, flag Bits) Bits { return b &^ flag }

// Toggle bits of b in flag
// Toggle(11110000, 10001000) => 01111000
func Toggle(b, flag Bits) Bits { return b ^ flag }

// Check if b contains flag
// Nothing contains 0 (empty bitmask)
// 11111111 contains all non-empty flags
func Has(b, flag Bits) bool { return b&flag != 0 }

/*
| Operation 	| Result 	| Desciption 	|
| 0011 & 0101	|	0001	|	Bitwise AND	|
| 0011 | 0101	|	0111	|	Bitwise OR	|
| 0011 ^ 0101	|   0110	| 	Bitwise XOR |
| ^0101			| 1010		| Bitwise NOT (same as 1111 ^ 0101) |
| 0011 &^ 0101	| 0010		| Bitclear (AND NOT) |
| 00110101<<2	| 11010100	|	Left shift	|
| 00110101<<100	| 00000000	|	No upper limit on shift count |
| 00110101>>2	| 00001101	| Right shift	|
* The binary numbers in the examples are for explanation only. Integer literals in Go must be specified in octal, decimal or hexadecimal.
* The bitwise operators take both signed and unsigned integers as input. The right-hand side of a shift operator, however, must be an unsigned integer.
* Shift operators implement arithmetic shifts if the left operand is a signed integer and logical shifts if it is an unsigned integer.
*/
