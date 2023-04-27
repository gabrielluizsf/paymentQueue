package payment

import "container/list"

type Queue struct {
  queue *list.List
}

func NewQueue() *Queue {
  return &Queue{
    queue: list.New(),
  }
}

func (pq *Queue) AddPayment(payment Payment) {
  pq.queue.PushBack(payment)
}

func (pq *Queue) ProcessPayment() any {
  if pq.queue.Len() == 0 {
    return nil
  }
  payment := pq.queue.Front().Value.(Payment)
  pq.queue.Remove(pq.queue.Front())
  return payment
}
