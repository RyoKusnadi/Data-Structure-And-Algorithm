package merge_sort

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	arr := []int{12, 11, 13, 5, 6, 7}
	sorted := mergeSort(arr)

	if !reflect.DeepEqual(sorted, []int{5, 6, 7, 11, 12, 13}) {
		t.Error("Wrong Sorted randomArray: ", sorted)
	}
}
