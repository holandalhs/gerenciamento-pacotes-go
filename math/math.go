package math

func Sum(a, b int) int {
	c := mult(2)
	return (a + b) + c
}

func Sub(a, b int) int {
	return a - b
}

func mult(a int) int { //função privada
	return a * 100
}
