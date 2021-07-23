package strategy

import (
	"testing"
)

func TestAppend(t *testing.T) {
	listobj := ListInstance()
	var (
		testData = []interface{}{
			1, "a", []int{1, 3},
		}
		expected = true
	)
	for i := 0; i < len(testData); i++ {
		actual := listobj.Append(testData[i])
		if actual != expected {
			t.Errorf("append is wrong")
		}
	}
}
