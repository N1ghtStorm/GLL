package main

func main() {
	var dLL = newDoubleLinkedList(17)
	dLL.AddValue(88)
	dLL.AddValue(2)
	dLL.AddValue(33)
	dLL.AddValue(14)
	dLL.AddValue(58)
	dLL.AddValue(42)
	dLL.AddValue(6)
	dLL.AddValue(27)
	dLL.AddValue(8)

	dLL.PrintValues()

	dLL.createCycleAtBack(5)
	var value, aaa = dLL.getCycleInfo()
	println(value)
	println(aaa)
}

type doubleLinkedList struct {
	firstnode *dllNode
	Count     int
}

func newDoubleLinkedList(value int) doubleLinkedList {
	return doubleLinkedList{firstnode: &dllNode{value: value, next: nil, prev: nil}, Count: 1}
}

func (dll *doubleLinkedList) createCycleAtBack(cycleTargetNodeIndex int) {
	var last = dll.firstnode.takeLast()
	var targetNodePtr = dll.getNodeByIndex(cycleTargetNodeIndex)
	last.next = targetNodePtr
}

func (dll *doubleLinkedList) getCycleInfo() (cycleStartNumber int, isCycled bool) {
	var turtle = dll.firstnode
	var rabbit = dll.firstnode

	for rabbit != nil && rabbit.next != nil {
		turtle = turtle.next
		rabbit = rabbit.next.next
		if turtle == rabbit {
			rabbit = dll.firstnode
			break
		}
	}

	if rabbit == nil || rabbit.next == nil {
		return 0, false
	}

	for rabbit != turtle {
		rabbit = rabbit.next
		turtle = turtle.next
	}

	return rabbit.value, true
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
	if dll.firstnode != nil {
		dll.firstnode.printNode()
	} else {
		println("nil list")
	}
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

func (dll *doubleLinkedList) Add(node *dllNode) {
	if dll.firstnode != nil {
		node.next = nil
		node.prev = nil
		dll.firstnode.addNodeNext(node)
	} else {
		node.next = nil
		node.prev = nil
		dll.firstnode = node
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

func (unsorted *doubleLinkedList) Sort() *doubleLinkedList {
	if unsorted.count() <= 1 {
		return unsorted
	}

	var leftHalf, rightHalf = unsorted.Divide()

	var left = leftHalf.Sort()
	var right = rightHalf.Sort()
	var mergedList *doubleLinkedList

	mergedList = Merge(left, right)
	return mergedList
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

func Merge(left *doubleLinkedList, right *doubleLinkedList) *doubleLinkedList {
	var newList = &doubleLinkedList{}

	for left.count() > 0 || right.count() > 0 {
		var leftCount = left.count()
		var rightCount = right.count()
		if leftCount > 0 && rightCount > 0 {
			if left.firstnode.value <= right.firstnode.value {
				node := left.firstnode
				left.removeFirstNode()
				newList.Add(node)
			} else {
				node := right.firstnode
				right.removeFirstNode()
				newList.Add(node)
			}
		} else if leftCount > 0 {
			node := left.firstnode
			left.removeFirstNode()
			newList.Add(node)
		} else if rightCount > 0 {
			node := right.firstnode
			right.removeFirstNode()
			newList.Add(node)
		}
	}
	return newList
}

func (dll *doubleLinkedList) removeFirstNode() {
	if dll.firstnode == nil {
		return
	}
	if dll.firstnode.next != nil {
		dll.firstnode = dll.firstnode.next
		dll.firstnode.prev = nil
	} else {
		dll.firstnode = nil
	}
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

func (node *dllNode) addNodeNext(newNode *dllNode) {
	if node.next == nil {
		node.next = newNode
		newNode.prev = node
	} else {
		node.next.addNodeNext(newNode)
	}
}

func (node *dllNode) printNode() {
	if node != nil {
		print(node.value)
	} else {
		println("nil node")
		return
	}

	if node.next != nil {
		print("->")
		node.next.printNode()
	} else {
		println("->nil")
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
