package main

func similarity(x [1000]int, y [1000]int) (int) {
	var sim int = 0
	var c int
	for i:=0; i<1000; i++ {
		c = 0
		for j:=0; j<1000; j++ {
			if x[i] == y[j] {
				c++
			}
		}
		sim += c*x[i]
	}

	return sim
}