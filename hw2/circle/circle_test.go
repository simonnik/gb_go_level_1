package circle_test

import (
	"github.com/simonnik/gb_go_level_1/hw2/circle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircle(t *testing.T) {
	parameters := []struct {
		r, expectedDiameter, expectedLength float64
	}{
		{123.4, 246.8, 775.345066905961},
		{1, 2, 6.283185307179586},
		{0, 0, 0},
	}

	for i := range parameters {
		d := circle.Diameter(parameters[i].r)
		l := circle.Length(d)

		assert.Equal(t, parameters[i].expectedDiameter, d)
		assert.Equal(t, parameters[i].expectedLength, l)
	}
}
