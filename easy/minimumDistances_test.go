package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumDistance(t *testing.T) {
	a := []int{7, 1, 3, 4, 1, 7}
	result := MinimumDistance(a)

	assert.Equal(t, 3, result, "Result must be 3")
}
