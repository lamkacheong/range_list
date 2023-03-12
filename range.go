package range_list

import "fmt"

type Range [2]int

func (r Range) Left() int {
	return r[0]
}

func (r Range) Right() int {
	return r[1]
}

func (r Range) intersect(r2 Range) bool {
	return r.Left() >= r2.Left() && r.Left() <= r2.Right() ||
		r2.Left() >= r.Left() && r2.Left() <= r.Right()
}

func (r Range) String() string {
	return fmt.Sprintf("[%d, %d)", r[0], r[1])
}
