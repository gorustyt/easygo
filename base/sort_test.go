package base

import (
	"fmt"
	"testing"
)

var sortArr = []int{1, 2, 3, 4, 5, 7, 99}

func TestBubbleSort(t *testing.T) {
	var arr1 = &SortInt{99, 1, 2, 3, 7, 4, 5}
	orgin := fmt.Sprintf("%v", arr1)
	BubbleSort(arr1)
	for i := 0; i < arr1.Len(); i++ {
		if (*arr1)[i] != sortArr[i] {
			t.Errorf("sort value not equal:%v:%v\norgin:%v\n,after:%v\n,expect:%v\n",
				i, arr[i], *arr1, orgin, arr1,
			)
		}
	}
}

func TestSelectSort(t *testing.T) {
	var arr1 = &SortInt{99, 1, 2, 3, 7, 4, 5}
	orgin := fmt.Sprintf("%v", arr1)
	BubbleSort(arr1)
	for i := 0; i < arr1.Len(); i++ {
		if (*arr1)[i] != sortArr[i] {
			t.Errorf("sort value not equal:%v:%v\norgin:%v\n,after:%v\n,expect:%v\n",
				i, arr[i], *arr1, orgin, arr1,
			)
		}
	}
}

func TestInsertSort(t *testing.T) {
	var arr1 = &SortInt{99, 1, 2, 3, 7, 4, 5}
	orgin := fmt.Sprintf("%v", arr1)
	InsertSort(arr1)
	for i := 0; i < arr1.Len(); i++ {
		if (*arr1)[i] != sortArr[i] {
			t.Errorf("sort value not equal:%v:%v\norgin:%v\n,after:%v\n,expect:%v\n",
				i, arr[i], *arr1, orgin, arr1,
			)
		}
	}
}

func TestMergeSort1(t *testing.T) {
	var arr = []int{99, 1, 2, 3, 7, 4, 5}
	origin := fmt.Sprintf("%v", arr)
	MergeSort1(arr)
	for i := 0; i < len(arr); i++ {
		if (arr)[i] != sortArr[i] {
			t.Errorf("sort value not equal:%v:%v\norgin:%v\n,after:%v\n,expect:%v\n",
				i, arr[i], arr, origin, arr,
			)
		}
	}
}

func TestQuickSort(t *testing.T) {
	var arr1 = &SortInt{99, 1, 2, 3, 7, 4, 5}
	orgin := fmt.Sprintf("%v", arr1)
	InsertSort(arr1)
	for i := 0; i < arr1.Len(); i++ {
		if (*arr1)[i] != sortArr[i] {
			t.Errorf("sort value not equal:%v:%v\norgin:%v\n,after:%v\n,expect:%v\n",
				i, arr[i], *arr1, orgin, arr1,
			)
		}
	}
}
