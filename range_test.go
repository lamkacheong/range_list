package range_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeIntersect(t *testing.T) {
	r1 := Range{5, 10}
	r2 := Range{9, 12}
	assert.Equal(t, true, r1.intersect(r2))
	assert.Equal(t, true, r2.intersect(r1))

	r3 := Range{5, 10}
	r4 := Range{12, 14}
	assert.Equal(t, false, r3.intersect(r4))
}
