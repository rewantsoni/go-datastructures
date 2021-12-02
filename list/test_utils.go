package list

type testMultiply struct {
	Val int
}

func (m testMultiply) Apply(element int) int {
	return element * m.Val
}

type testAdd struct {
	Val int
}

func (a testAdd) Apply(e int) int {
	return e + a.Val
}