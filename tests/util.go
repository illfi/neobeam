package main

import (
	"fmt"
	"ill.fi/neobeam/interp"
	"testing"
)

func ShouldBe(name, src, check string) string {
	return fmt.Sprintf("%s failed, is %q, should be %q", name, src, check)
}

func EShouldBe(t *testing.T, name, src, check string) {
	t.Errorf(ShouldBe(name, src, check))
}

func AllUnitsAre(t interp.UnitType, w *interp.World) bool {
	flag := true
	for _, u := range w.FlatUnits() {
		flag = flag && u.Type == t
	}
	return flag
}
