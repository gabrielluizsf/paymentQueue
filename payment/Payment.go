package payment

import "time"

type Payment struct {
  CustomerName string
  Value        float64
  Date         time.Time
}