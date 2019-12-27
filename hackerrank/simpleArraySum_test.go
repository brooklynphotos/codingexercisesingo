package hackerrank

import (
	"reflect"
	"testing"
)

func TestSimpleArraySum(t *testing.T) {
	res := simpleArraySum([]int32{1, 2, 3})
	expected := 6
	if reflect.DeepEqual(expected, res) {
		t.Errorf("%d != %d", res, expected)
	}
}
