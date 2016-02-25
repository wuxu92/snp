package test
import (
  "testing"
  "reflect"
)

func TestSliceRemove(t *testing.T) {
  var slice = []int{1, 2, 3, 4, 5}
  var result = []int{1,3,4,5}
  p := append(slice[:1], slice[2:]...)
  if reflect.DeepEqual(p, result) {
    t.Log("slice and plus right ", p)
  } else {
    t.Error("can do slice and plus operate ", p)
  }
}