package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []int64{1, 2, 3, 4, 5, 6, 7}
	cs := []Content{{1}, {2}, {3}, {7}, {8}}
	cc := filterValidTargetIDs(a, cs)
	fmt.Println("cc = ", cc)

	s := []string{"茅台", "五粮液", "泸州老窖"}

	b := []int64{3, 17}

	fmt.Println(arrayContains(a, b))

	fmt.Println(strings.Join(s, ","))
}

func filterValidTargetIDs(targetIds []int64, excontents []Content) (validTargetIDs []int64) {
	if len(targetIds) == len(excontents) {
		return
	}
	m := make(map[int64]int)
	for _, content := range excontents {
		m[content.ID]++
	}

	for _, targetId := range targetIds {
		if _, ok := m[targetId]; !ok {

			validTargetIDs = append(validTargetIDs, targetId)
		}
	}
	return
}

type Content struct {
	ID int64
}

func arrayContains(arr []int64, subArr []int64) bool {
	times := 0
	for _, v := range arr {
		for _, v2 := range subArr {
			if v == v2 {
				times++
			}
		}
	}
	return times == len(subArr)
}
