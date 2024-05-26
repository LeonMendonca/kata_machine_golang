package heapmethods

import (
  "math"
)

func ToParent(idx int) int {
  //math.Round expects floating number, therefore used typecasting
  return int(math.Round(float64((idx-1)/2)))
}

func ToLeft(idx int) int {
  return (2*idx)+1
}

func ToRight(idx int) int {
  return (2*idx)+2
}
