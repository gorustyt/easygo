package base

import "testing"

var arr = []int{1, 2, 3, 4, 5, 6}

func TestBinarySearch(t *testing.T) {
	for index, v := range arr {
		i, ok := BinarySearch(arr, v)
		if !ok {
			t.Errorf("error:not found:%v:%v", index, v)
		}
		if i != index {
			t.Errorf("error:not equal,expect,%v:%v,real:%v", index, v, i)
		}
	}
	_, ok := BinarySearch(arr, 7)
	if ok {
		t.Errorf("error:%vexpect not found", 7)
	}
}
