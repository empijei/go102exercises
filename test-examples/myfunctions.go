package myfunc

func theAnswer() int {
	i := 0
	for ; i != 42; i++ {
	}
	return i
}

func myBar(s string) string {
	return s
}

func hazard() int {
	a := 0
	go func() { a++ }()
	return a
}
