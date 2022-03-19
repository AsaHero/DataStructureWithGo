package LinkedList

import "errors"

type node struct {
	data interface{}
	next *node
}

func (n *node) setData(data interface{}) *node {
	n.data = data
	return n
}

type linkedlist struct {
	head *node
}

func (l *linkedlist) Add(data ...interface{}) {
	for i := 0; i < len(data); i++ {
		if l.head != nil {
			current_node := l.head
			for ; current_node.next != nil; current_node = current_node.next {
			}
			current_node.next = new(node).setData(data[i])
		} else {
			l.head = new(node).setData(data[i])
		}
	}
}

func (l *linkedlist) AddAt(index int, data interface{}) error {
	if l.Len() < index+1 {
		return errors.New("error: Required index is more than actual size of list")
	}
	if index <= 0 {
		new_node := new(node)
		new_node.data = data
		new_node.next = l.head.next
		l.head = new_node
		return nil
	}

	current_node := l.head
	for i := 0; i < index-1; i++ {
		current_node = current_node.next

	}
	new_node := new(node)
	new_node.data = data
	new_node.next = current_node.next
	current_node.next = new_node
	return nil
}

func (l *linkedlist) Len() int {
	if l.head != nil {
		count := 1
		current_node := l.head
		for ; current_node.next != nil; current_node = current_node.next {
			count++
		}
		return count
	}
	return 0
}

func (l *linkedlist) At(index int) (interface{}, error) {
	if l.Len() < index+1 {
		return nil, errors.New("error: Required index is more than actual size of list")
	}
	current_node := l.head
	for i := 0; i < index; i++ {
		current_node = current_node.next
	}
	return current_node.data, nil
}

func (l *linkedlist) DeleteAt(index int) error {
	if l.Len() < index+1 {
		return errors.New("error: Required index is more than actual size of list")
	}

	if index <= 0 {
		l.head = l.head.next
		return nil
	}

	current_node := l.head
	for i := 0; i < index-1; i++ {
		current_node = current_node.next
	}
	current_node.next = current_node.next.next
	return nil
}

func (l *linkedlist) Delete(data interface{}) error {
	if l.head == nil {
		return errors.New("error: List is empty")
	}

	if l.head.data == data && l.head.next == nil {
		l.head = nil
	} else if l.head.data == data && l.head.next != nil {
		l.head = l.head.next
	} else {
		current_node := l.head
		for ; current_node.next != nil; current_node = current_node.next {
			if current_node.next.data == data {
				current_node.next = current_node.next.next
				return nil
			}
		}
	}
	return errors.New("Not Found !")
}

func New() linkedlist {
	return linkedlist{}
}
