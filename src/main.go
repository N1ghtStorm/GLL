package main

func main() {
	var dLL = newDoubleLinkedList(17)
	dLL.AddValue(2)
	dLL.AddValue(33)
	dLL.AddValue(14)
	dLL.AddValue(58)
	dLL.AddValue(6)
	dLL.AddValue(27)
	dLL.AddValue(8)
	dLL.AddValue(88)

	dLL.PrintValues()
	println("")
	println(dLL.Count)
	println(dLL.count())
	dLL.invert()
	dLL.PrintValues()
	println("=======================")
	println(dLL.getValueByIndex(2))
}

type doubleLinkedList struct {
	firstnode *dllNode
	Count     int
}

func newDoubleLinkedList(value int) doubleLinkedList {
	return doubleLinkedList{firstnode: &dllNode{value: value, next: nil, prev: nil}, Count: 1}
}

func (dll *doubleLinkedList) getValueByIndex(index int) int {
	if dll.firstnode != nil {
		var current = 0
		return dll.firstnode.getValueByIndex(&index, &current)
	}

	panic("list is empty")
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

func (dll *doubleLinkedList) count() int {
	var count = 0
	if dll.firstnode != nil {
		count = 1
		dll.firstnode.iterCount(&count)
	}
	return count
}

func (dll *doubleLinkedList) Sort() {
	// var count = dll.count()
}

// func (dll *doubleLinkedList) Divide() (doubleLinkedList, doubleLinkedList) {
// 	//var count = dll.count()
// 	//var firstList = doubleLinkedList{firstnode: dll.firstnode}
// 	return error
// }

type dllNode struct {
	value int
	next  *dllNode
	prev  *dllNode
}

func (node *dllNode) getValueByIndex(index *int, current *int) int {
	if *index == *current {
		return node.value
	}

	if node.next != nil {
		*current++
		return node.next.getValueByIndex(index, current)
	}

	panic("index is out of boundaries of list")
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
		println("")
	}
}

func (node *dllNode) iterCount(count *int) {
	if node.next != nil {
		*count++
		node.next.iterCount(count)
	}
}
