package collection

import "computer_club/analyze/data"

type Queue struct {
	items []*data.Client
}

func (q *Queue) Enqueue(item *data.Client) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() *data.Client {
	if len(q.items) == 0 {
		return nil
	}

	dequeuedItem := q.items[0]
	q.items = q.items[1:]
	return dequeuedItem
}

func (q *Queue) Length() int {
	return len(q.items)
}
