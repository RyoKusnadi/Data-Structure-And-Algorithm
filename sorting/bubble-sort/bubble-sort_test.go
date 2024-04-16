package bubble_sort

import (
	"reflect"
	"testing"
)

func Test_bubble_sort(t *testing.T) {
	randomArray := []int{64, 25, 12, 22, 11}
	bubbleSort(randomArray, true)
	if !reflect.DeepEqual(randomArray, []int{11, 12, 22, 25, 64}) {
		t.Error("Wrong Sorted randomArray: ", randomArray)
	}

	randomArray2 := []int{64, 25, 12, 22, 11}
	bubbleSort(randomArray2, false)
	if !reflect.DeepEqual(randomArray2, []int{64, 25, 22, 12, 11}) {
		t.Error("Wrong Sorted randomArray2: ", randomArray2)
	}
}
