package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	firstItem *ListItem
	lastItem  *ListItem
	length    int
}

func NewList() List {
	return new(list)
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *ListItem {
	return l.firstItem
}

func (l list) Back() *ListItem {
	return l.lastItem
}

func (l *list) insertAfter(item *ListItem, newItem *ListItem) {
	newItem.Prev = item

	if item.Next == nil {
		newItem.Next = nil
		l.lastItem = newItem
	} else {
		newItem.Next = item.Next
		item.Next.Prev = newItem
	}

	item.Next = newItem
}

func (l *list) insertBefore(item *ListItem, newItem *ListItem) {
	newItem.Next = item

	if item.Prev == nil {
		newItem.Prev = nil
		l.firstItem = newItem
	} else {
		newItem.Prev = item.Prev
		item.Prev.Next = newItem
	}

	item.Prev = newItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.firstItem == nil {
		l.firstItem = newItem
		l.lastItem = newItem
	} else {
		l.insertBefore(l.firstItem, newItem)
	}

	l.length++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	var newItem *ListItem

	if l.lastItem == nil {
		newItem = l.PushFront(v)
	} else {
		newItem = &ListItem{Value: v}
		l.insertAfter(l.lastItem, newItem)
	}

	l.length++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.firstItem = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.lastItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.firstItem == i {
		return
	}

	if l.lastItem == i {
		l.lastItem = i.Prev
	}

	i.Prev.Next = i.Next
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	tmpFirstItem := l.firstItem
	l.firstItem = i

	i.Prev = nil
	i.Next = tmpFirstItem

	tmpFirstItem.Prev = i
}
