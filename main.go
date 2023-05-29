package main

import (
	"container/heap"
	"fmt"
	"os"
)

type Food struct {
	name         string
	expiryInDays int
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func priorityQueueFood() {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for {
		var option string
		fmt.Println("Enter i to insert, p to know element that gets expired first, l to list elements and q to quit")
		fmt.Scanf("%s", &option)

		switch option {
		case "i":
			var name string
			fmt.Println("Enter the food name")
			fmt.Scanf("%s", &name)

			var expiryInDays int
			fmt.Println("Enter the days in which food gets expired")
			fmt.Scanf("%d", &expiryInDays)
			item := &Item{
				value:    name,
				priority: expiryInDays,
			}
			heap.Push(&pq, item)

		case "p":
			item := heap.Pop(&pq).(*Item)
			fmt.Println("Element that gets expired first in given set of data is ", item.value)

		case "l":
			for _, value := range pq {
				fmt.Printf("data is as follows: %+v", value)
				fmt.Println()
			}

		default:
			os.Exit(0)
		}
	}
}

func main() {
	fmt.Println("Entering the main functionnn!!!")
	priorityQueueFood()
}
