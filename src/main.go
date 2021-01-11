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
	dLL.Sort()
	//dLL.PrintValues()
	//dLL.PrintValues()
	// println("")
	// println(dLL.Count)
	// println(dLL.count())
	// dLL.invert()
	// dLL.PrintValues()
	// print("")
	// println("=======================")
	// println(dLL.getValueByIndex(2))
	// println(dLL.getNodeByIndex(0).value)
	// println("=======================")
	// _, dLL2 := dLL.Divide()
	// dLL.PrintValues()
	// println("")
	// dLL2.PrintValues()
	// dLL2.Sort()
	// dLL2.PrintValues()
	// // var dLL2 = newDoubleLinkedList(17)
	// println(dLL2.getNodeByIndex(3).value)

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
	targetNodePtr.next = last
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
	// 	if A is None or A.next is None:
	// 	return A
	// leftHalf, rightHalf = splitTheList(A)
	// left = mergeSortLinkedList(leftHalf)
	// right = mergeSortLinkedList(rightHalf)
	// return mergeTheLists(left, right)
	if dll.firstnode == nil || dll.firstnode.next == nil {
		return
	}

	var leftHalf, rightHalf = dll.Divide()

	leftHalf.Sort()
	rightHalf.Sort()
	leftHalf.Merge(rightHalf)
}

func MergeSorted(dll1 *doubleLinkedList, dll2 *doubleLinkedList) *doubleLinkedList {

	// 	if list1.data < list2.data:
	// 	temp = list1
	// 	temp.next = self.mergeSorted(list1.next, list2)
	//   else:
	// 	temp = list2
	// 	temp.next = self.mergeSorted(list1, list2.next)
	//   return temp
	if dll2 == nil {
		return dll1
	}

	if dll1 == nil {
		return dll2
	}

	//var temp *doubleLinkedList

	//if dll1.

	panic("unimplemented")
}

func (dll *doubleLinkedList) OutPutDivide() (*doubleLinkedList, *doubleLinkedList) {
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

func (left *doubleLinkedList) Merge(right *doubleLinkedList) *doubleLinkedList {
	// if dll2 != nil {
	// 	var dllLast = dll.takeLast()
	// 	dllLast.next = dll2.firstnode
	// 	dll2.firstnode.prev = dllLast
	// }

	var newList = doubleLinkedList{}

	for left.count() > 0 || right.count() > 0 {
		var leftCount = left.count()
		var rightCount = right.count()
		if leftCount > 0 && rightCount > 0 {
			if left.firstnode.value > right.firstnode.value {
				newList.Add(right.firstnode)
				right.removeFirstNode()
			} else {
				newList.Add(left.firstnode)
				left.removeFirstNode()
			}
		} else if leftCount > 0 {
			newList.Add(left.firstnode)
			left.removeFirstNode()
		} else if rightCount > 0 {
			newList.Add(right.firstnode)
			right.removeFirstNode()
		}
	}
	return &newList
}

func (dll *doubleLinkedList) removeFirstNode() {
	if dll.firstnode.next != nil {
		dll.firstnode = dll.firstnode.next
		dll.firstnode.prev = nil
	} else {
		dll.firstnode = nil
	}

}

func (dll *doubleLinkedList) Add(node *dllNode) {
	if dll.firstnode != nil {
		dll.firstnode.addNodeNext(node)
	} else {
		dll.firstnode = node
	}
}

// Replases first node from left to right
func (left *doubleLinkedList) replaceFirstNode(right *doubleLinkedList) {
	var replacedNode = left.firstnode
	replacedNode.next = right.firstnode
	left.firstnode = left.firstnode.next
	left.firstnode.prev = nil
	right.firstnode.next.prev = replacedNode
	right.firstnode = replacedNode
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
	} else {
		node.next.addNodeNext(newNode)
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
