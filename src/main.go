package main

func main() {
	var dLL = newDoubleLinkedList(17)
	dLL.AddValue(2)
	dLL.AddValue(3)
	dLL.AddValue(4)
	dLL.AddValue(5)
	dLL.AddValue(6)
	dLL.AddValue(7)
	dLL.AddValue(8)

	dLL.PrintValues()
	println("")
	println("================")
	println("================")
	println(dLL.Count)
	println("================")
	println("================")
	dLL.invert()
	dLL.PrintValues()
}

type doubleLinkedList struct {
	firstnode *dllNode
	Count     int
}

func newDoubleLinkedList(value int) doubleLinkedList {
	return doubleLinkedList{firstnode: &dllNode{value: value, next: nil, prev: nil}, Count: 1}
}

func (dll *doubleLinkedList) PrintValues() {
	dll.firstnode.printNode()
}

func (dll *doubleLinkedList) AddValue(value int) {
	if dll.firstnode != nil {
		dll.Count = dll.Count + 1
		dll.firstnode.addNext(value)
	} else {
		dll.Count = 1
		dll.firstnode = &dllNode{value: value, next: nil, prev: nil}
	}
}

func (dll *doubleLinkedList) invert() {
	if dll.firstnode != nil && dll.firstnode.next != nil {
		dll.firstnode.invert(dll)
	} else {
	}
}

type dllNode struct {
	value int
	next  *dllNode
	prev  *dllNode
}

func (node *dllNode) addNext(value int) {
	if node.next == nil {
		node.next = &dllNode{value: value, next: nil, prev: node}
	} else {
		node.next.addNext(value)
	}
}

func (node *dllNode) printNode() {
	print(node.value)
	if node.next != nil {
		print("->")
		node.next.printNode()
	} else {
		print("->nil")
	}
}

func (node *dllNode) invert(dll *doubleLinkedList) {
	if node.next != nil {
		node.next, node.prev = node.prev, node.next
		node.prev.invert(dll)
	} else {
		node.next = node.prev
		dll.firstnode = node
	}
}
