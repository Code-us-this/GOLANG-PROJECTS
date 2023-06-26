// Package operations is used for arithmetic operations of calculator project.
package operations

// Add returns the sum of two numbers.
func Add(numA float32, numB float32) float32 {
	return numA + numB
}

// Sub returns the difference between two numbers.
func Sub(numA float32, numB float32) float32 {
	return numA - numB
}

// Mul returns multiplication of two numbers.
func Mul(numA float32, numB float32) float32 {
	return numA * numB
}

// Div returns division of two numbers.
func Div(numA float32, numB float32) (int, int) {
	return int(numA) / int(numB), int(numA) % int(numB)
}
