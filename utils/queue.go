package utils

import "shopping-cart-api/models"

type Queue struct {
  elements []models.CartItem
}

func (q *Queue) Enqueue(item models.CartItem) {
  q.elements = append(q.elements, item)
}

func (q *Queue) Dequeue() (models.CartItem, bool) {
  if len(q.elements) == 0 {
    return models.CartItem{}, false
  }
  item := q.elements[0]
  q.elements = q.elements[1:]
  return item, true
}

func (q *Queue) IsEmpty() bool {
  return len(q.elements) == 0
}
