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

func testCreateNodes(elements ...int) (*node, *node) {
	var first, prev, curr *node
	for _, element := range elements {
		curr = newNode(element)

		if first == nil {
			first = curr
		}

		if prev != nil {
			curr.prev = prev
			prev.next = curr
		}

		prev = curr
	}
	return first, curr
}
