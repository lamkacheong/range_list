package range_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	rl := RangeList{
		Members: []Range{Range{1, 10}, Range{20, 30}},
	}
	assert.Equal(t, "[1, 10) [20, 30)", rl.String(), "Print range list should work")
}

func TestAddMiddle(t *testing.T) {
	var rl RangeList
	rl = RangeList{
		Members: []Range{Range{10, 20}, Range{30, 40}},
	}
	rl.Add(Range{20, 30})
	assert.Equal(t, rl.Members[0], Range{10, 40})

	rl = RangeList{
		Members: []Range{Range{10, 20}, Range{30, 40}},
	}
	rl.Add(Range{21, 29})
	assert.Equal(t, rl.Members[0], Range{10, 20})
	assert.Equal(t, rl.Members[1], Range{21, 29})
	assert.Equal(t, rl.Members[2], Range{30, 40})

	rl = RangeList{
		Members: []Range{Range{10, 20}, Range{30, 40}},
	}
	rl.Add(Range{21, 50})
	assert.Equal(t, rl.Members[0], Range{10, 20})
	assert.Equal(t, rl.Members[1], Range{21, 50})

	rl = RangeList{
		Members: []Range{Range{10, 20}, Range{30, 40}},
	}
	rl.Add(Range{0, 100})
	assert.Equal(t, 1, len(rl.Members))
	assert.Equal(t, rl.Members[0], Range{0, 100})

}

func TestAddOver(t *testing.T) {
	var rl RangeList
	rl = RangeList{
		Members: []Range{Range{10, 20}, Range{30, 40},
			Range{50, 60},
		},
	}
	rl.Add(Range{25, 55})
	assert.Equal(t, rl.Members[0], Range{10, 20})
	assert.Equal(t, rl.Members[1], Range{25, 60})

	rl = RangeList{
		Members: []Range{Range{10, 20}, Range{30, 40},
			Range{50, 60},
		},
	}
	rl.Add(Range{25, 35})
	assert.Equal(t, rl.Members[0], Range{10, 20})
	assert.Equal(t, rl.Members[1], Range{25, 40})
	assert.Equal(t, rl.Members[2], Range{50, 60})
}

func TestRemove(t *testing.T) {

	var rl RangeList
	rl = RangeList{
		Members: []Range{Range{10, 20}, Range{30, 40},
			Range{50, 60},
		},
	}

	rl.Remove(Range{31, 39})
	assert.Equal(t, rl.Members[0], Range{10, 20})
	assert.Equal(t, rl.Members[1], Range{30, 31})
	assert.Equal(t, rl.Members[2], Range{39, 40})
	assert.Equal(t, rl.Members[3], Range{50, 60})

	rl = RangeList{
		Members: []Range{Range{10, 20}, Range{30, 40},
			Range{50, 60},
		},
	}

	rl.Remove(Range{0, 100})
	assert.Equal(t, len(rl.Members), 0)

	rl = RangeList{
		Members: []Range{Range{10, 20}, Range{30, 40},
			Range{50, 60},
		},
	}
	rl.Remove(Range{29, 55})
	assert.Equal(t, rl.Members[0], Range{10, 20})
	assert.Equal(t, rl.Members[1], Range{55, 60})
}
