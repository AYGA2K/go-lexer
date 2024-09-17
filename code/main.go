func test(x, y int) int {
	if x == 10 {
		var z = y * 2
		for z != 0 {
			z = z - 1
		}
		return z + x
	} else if y != 5 {
		return y / x
	} else {
		return 0
	}
}
