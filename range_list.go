package range_list

import (
	"math"
	"strings"
)

type RangeList struct {
	Members []Range
}

func (rl *RangeList) String() string {
	s := make([]string, len(rl.Members))
	for i, v := range rl.Members {
		s[i] = v.String()
	}
	return strings.Join(s, " ")
}

// FirstNotLeft RangeList中第一个不在rangeToAdd左边的
// 也就是待插入的地方
// 全部都在左边，就返回n
func (rl *RangeList) FirstNotLeft(rangeToAdd Range) int {
	for i, _ := range rl.Members {
		if rl.Members[i].Left() >= rangeToAdd.Left() {
			return i
		}
	}
	return len(rl.Members)
}

// Add
// 1。找到和rangeToAdd相交的intersectRanges
// 2。如果没有，找第一个在rangeToAdd右边的地方插入
// 3。如果有，rangeToAdd和这些intersectRanges合并
func (rl *RangeList) Add(rangeToAdd Range) {
	if !(rangeToAdd.Left() < rangeToAdd.Right()) {
		return
	}
	min := rangeToAdd.Left()
	max := rangeToAdd.Right()

	intersectRanges := make([]int, 0)
	for i, v := range rl.Members {
		if v.intersect(rangeToAdd) {
			if v.Left() < min {
				min = v.Left()
			}
			if v.Right() > max {
				max = v.Right()
			}
			intersectRanges = append(intersectRanges, i)
		}
	}
	if len(intersectRanges) > 0 {
		// 和所有交集合并后，在第一个交集后插入
		s := intersectRanges[0]
		t := intersectRanges[len(intersectRanges)-1]
		rl.Members = append(rl.Members[:s], rl.Members[t+1:]...)
		rl.Members = SliceInsert(rl.Members, s, Range{min, max})
	} else {
		// 没有交集，在第一个不在rangeToAdd左边处插入
		placeToInsert := rl.FirstNotLeft(rangeToAdd)
		rl.Members = SliceInsert(rl.Members, placeToInsert, rangeToAdd)
	}
}

func (rl *RangeList) Remove(rangeToRemove Range) {
	//输入不正确的处理
	if !(rangeToRemove.Left() < rangeToRemove.Right()) {
		return
	}

	//找出和rangeToAdd相交的intersectRanges
	min := math.MaxInt
	max := -math.MaxInt
	intersectRanges := make([]int, 0)
	for i, v := range rl.Members {
		if v.intersect(rangeToRemove) {
			if v.Left() < min {
				min = v.Left()
			}
			if v.Right() > max {
				max = v.Right()
			}
			intersectRanges = append(intersectRanges, i)
		}
	}

	if len(intersectRanges) == 0 {
		//如果没有相交，无事发生
		return
	} else {
		//s,t以内的都删除，处理首尾的切割
		s := intersectRanges[0]
		t := intersectRanges[len(intersectRanges)-1]
		rl.Members = append(rl.Members[:s], rl.Members[t+1:]...)
		if max > rangeToRemove.Right() {
			rl.Members = SliceInsert(rl.Members, s, Range{rangeToRemove.Right(), max})
		}
		if min < rangeToRemove.Left() {
			rl.Members = SliceInsert(rl.Members, s, Range{min, rangeToRemove.Left()})
		}
	}
}
