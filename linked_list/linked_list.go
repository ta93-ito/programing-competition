package main

import "fmt"

type Cell struct {
	item int
	next *Cell
}

type List struct {
	top *Cell
}

func main() {
	list := newList()
	cell := newCell(1, new(Cell))
	cell2 := newCell(2, new(Cell))

	list.insertCellInIndex(1, cell.item)
	result := list.insertCellInIndex(2, cell2.item)

	if !result {
		panic("error occured")
	}
	first, _ := list.getItemByIndex(1)
	second, _ := list.getItemByIndex(2)
	fmt.Println(first)
	fmt.Println(second)
	list.printList()
}

func newCell(x int, cp *Cell) *Cell {
	return &Cell{item: x, next: cp}
}

func newList() *List {
	return &List{top: new(Cell)}
}

func (cp *Cell) getCellByIndex(id int) *Cell {
	for i := 0; cp != nil; i++ {
		if i == id {
			return cp
		}
		cp = cp.next
	}
	return nil
}

func (lst *List) getItemByIndex(id int) (*Cell, bool) {
	cp := lst.top.getCellByIndex(id)
	if cp == nil {
		return nil, false
	} else {
		return cp, true
	}
}

func (list *List) insertCellInIndex(id, item int) bool {
	before, result := list.getItemByIndex(id - 1)
	if !result {
		return false
	}
	next := before.next
	cell := &Cell{item: item, next: next}
	before.next = cell
	return true
}

func (list *List) deleteByIndex(id int) bool {
	before, result := list.getItemByIndex(id - 1)
	if !result {
		return false
	}
	cell := before.next
	next := cell.next
	before.next = next
	return true
}

func (list *List) printList() {
	cp := list.top.next
	for i := 1; cp != nil; cp = cp.next {
		fmt.Println(i, cp.item)
	}
}

func (list *List) isEmpty() bool {
	return list.top.next == nil
}
