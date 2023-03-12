package range_list

// SliceInsert 帮助方法，在index地方插入value到a
func SliceInsert(a []Range, index int, value Range) []Range {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
