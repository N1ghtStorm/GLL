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
	print("")
	println("=======================")
	println(dLL.getValueByIndex(2))
	println(dLL.getNodeByIndex(0).value)
	println("=======================")
	_, dLL2 := dLL.Divide()
	dLL.PrintValues()
	println("")
	dLL2.PrintValues()
	// var dLL2 = newDoubleLinkedList(17)
	// println(dLL2.getNodeByIndex(3).value)

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

func (dll *doubleLinkedList) getNodeByIndex(index int) *dllNode {
	if dll.firstnode != nil {
		var current = 0
		return dll.firstnode.getNodeByIndex(&index, &current)
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
	//first, second = dll.Divide()
	panic("unimplemented")
}

func (dll *doubleLinkedList) SortDivide() (*doubleLinkedList, *doubleLinkedList) {
	if dll.count() > 1 {
		return dll.Divide()
	}

	return dll, nil
}

func (dll *doubleLinkedList) Divide() (*doubleLinkedList, *doubleLinkedList) {
	var firstLinkedList = dll
	var halfCount = dll.count() / 2
	var secondLinkedListFirstNode = dll.getNodeByIndex(halfCount)
	secondLinkedListFirstNode.prev.next = nil
	secondLinkedListFirstNode.prev = nil
	var secondLinkedList = doubleLinkedList{firstnode: secondLinkedListFirstNode, Count: 0}
	return firstLinkedList, &secondLinkedList
}

func (dll *doubleLinkedList) Merge(dll2 *doubleLinkedList) *doubleLinkedList {
	if dll2 != nil {
		var dllLast = dll.takeLast()
		dllLast.next = dll2.firstnode
		dll2.firstnode.prev = dllLast
	}

	return dll
}

func (dll *doubleLinkedList) takeLast() *dllNode {
	if dll.firstnode == nil {
		panic("cant take last element - list is empty")
	}

	return dll.firstnode.takeLast()
}

//============================================================================================================
//============================================================================================================
//============================================================================================================

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

func (node *dllNode) getNodeByIndex(index *int, current *int) *dllNode {
	if *index == *current {
		return node
	}

	if node.next != nil {
		*current++
		return node.next.getNodeByIndex(index, current)
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

func (node *dllNode) takeLast() *dllNode {
	if node.next == nil {
		return node
	}

	return node.next.takeLast()
}
