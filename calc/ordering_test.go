package calc

import "testing"

func TestOrderList(t *testing.T) {

	unsortedList := make([]int, 0)
	expectedResult := make([]int, 0)

	unsortedList = append(unsortedList, 6, 5, 4, 3, 2, 1)
	expectedResult = append(expectedResult, 1, 2, 3, 4, 5, 6)

	sortedList := OrderList(unsortedList)

	for i := 0; i < len(unsortedList); i++ {
		if unsortedList[i] != expectedResult[i] {
			t.Errorf("Expected result: %d - Result: %d", expectedResult, sortedList)
		}
	}
}
