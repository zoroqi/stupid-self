package q800

import (
	"encoding/json"
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"os"
	"testing"
)

func TestNumBusesToDestination(t *testing.T) {
	stupid_self.AssertEqual(t, numBusesToDestinationPlanA([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6), 2)
	stupid_self.AssertEqual(t, numBusesToDestinationPlanA([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 7), 1)
	stupid_self.AssertEqual(t, numBusesToDestinationPlanA([][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12), -1)
	stupid_self.AssertEqual(t, numBusesToDestinationPlanA([][]int{{1, 7}, {3, 5}}, 5, 5), 0)
	bigData, _ := os.ReadFile("_815_test.txt")
	routes := [][]int{}
	json.Unmarshal(bigData, &routes)
	m := map[int]bool{}
	for _, v := range routes[0] {
		m[v] = true
	}
	for _, v := range routes[1] {
		m[v] = true
	}
	stupid_self.AssertEqual(t, numBusesToDestinationPlanA(routes, 0, 100000), -1)
}
