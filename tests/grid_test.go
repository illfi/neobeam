package main

import (
	"github.com/stretchr/testify/assert"
	"ill.fi/neobeam/interp"
	"testing"
)

const (
	BLANK string = "   \n   \n   "
)

func TestCreateWorld(t *testing.T) {
	world := interp.CreateWorld(BLANK)
	assert.Equal(t, 3, world.Width)
	assert.Equal(t, 3, world.Height)
	assert.Equal(t, true, AllUnitsAre(interp.Air, world))
	assert.Equal(t, "     \n     \n     \n", world.Display())
}
