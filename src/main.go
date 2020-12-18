package main

func main() {
	var dLL = newDoubleLinkedList(1)
	println(dLL.firstnode.value)
}

type doubleLinkedList struct {
	firstnode *dllNode
}

func newDoubleLinkedList(value int) doubleLinkedList {
	return doubleLinkedList{firstnode: &dllNode{value: value, next: nil, prev: nil}}
}

func (dll doubleLinkedList) AddValue(value int) {
	if dll.firstnode != nil {
		dll.firstnode.addNext(value)
	} else {
		dll.firstnode = &dllNode{value: value, next: nil, prev: nil}
	}
}

func (dll doubleLinkedList) invert() {
	if dll.firstnode == nil {

	} else {

	}
}

type dllNode struct {
	value int
	next  *dllNode
	prev  *dllNode
}

func (node dllNode) addNext(value int) {
	node.next = &dllNode{value: value, next: nil, prev: nil}
}

func (node dllNode) invert() {

}
