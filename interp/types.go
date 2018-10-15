package interp

import (
	"strings"
)

/////////////////
// string shit //
/////////////////

type Rope struct {
	Source  []rune
	Current *rune
}

func NewRope(source string) Rope {
	r := []rune(source)
	return Rope{
		Source:  r,
		Current: &r[0],
	}
}

func (r *Rope) Consume(pred func(rune) bool) []rune {
	buf := []rune("")
	sc := r.Source
	w := &sc[0]
	for w != nil && pred(*w) && len(sc) > 0 {
		buf = append(buf, *w)
		sc = append(sc[:0], sc[1:]...) // deletes head
		if len(sc) == 0 {
			w = nil
		} else {
			w = &sc[0]
		}
	}
	r.Reset(sc)
	return buf
}

func (r *Rope) Reset(n []rune) {
	if len(n) == 0 {
		r.Source = n
		r.Current = nil
	} else {
		r.Source = n
		r.Current = &r.Source[0]
	}
}

type UnitState int

const (
	Benign UnitState = iota
	Neutral
	Activated
	// TODO: think of more states
)

type UnitType int

const (
	Unrecognized UnitType = iota
	Air                   // WS
	Caster                // *
	Emitter               // > ^ < v
	Wall                  // ~ I
	Router                // + | \ /
	Reference             // digits
)

///////////////
// important //
///////////////

func supplementType(r string, t UnitType) {
	for _, c := range []rune(r) {
		TypeMap[c] = t
	}
}

func AddSupplementTypes() {
	// routers
	supplementType("<>^v", Router)
	// references
	supplementType("1234567890", Reference)
}

// root 1 char types
var TypeMap = map[rune]UnitType{
	'*': Caster,
	'~': Wall,
	'I': Wall,
	' ': Air,
}

type World struct {
	Units         [][]Unit
	Width, Height int
}

type Unit struct {
	Type                  UnitType
	State                 UnitState
	Left, Right, Up, Down *Unit
	Row, Col              int
}

func charFromType(t UnitType) *rune {
	for k, v := range TypeMap {
		if v == t {
			return &k
		}
	}
	return nil
}

func (u Unit) Display() *rune {
	return charFromType(u.Type)
}

func (w *World) Interweave() {
	for ri, r := range (*w).Units {
		for ci, c := range r {
			c.Up = w.At(ri, ci-1)
			c.Down = w.At(ri, ci+1)
			c.Right = w.At(ri+1, ci)
			c.Left = w.At(ri-1, ci)
		}
	}
}

func (w *World) At(x, y int) *Unit {
	nx := x
	ny := y
	if x < 0 {
		// wrap around by diff
		nx = w.Width - x
	}
	if y < 0 {
		ny = w.Height - y
	}

	return w.PAt(nx, ny)
}

// only does positive wrapping...
func (w *World) PAt(x, y int) *Unit {
	return w.AtUnsafe(x%w.Width, y%w.Height)
}

// doesnt wrap
func (w *World) AtUnsafe(x, y int) *Unit {
	for ri, r := range (*w).Units {
		for ci, c := range r {
			if ri == x && ci == y {
				return &c
			}
		}
	}
	return nil
}

func (w *World) FlatUnits() []Unit {
	units := make([]Unit, 0)
	for i := range w.Units {
		for x := range w.Units[i] {
			units = append(units, w.Units[i][x])
		}
	}
	return units
}

func CreateWorld(s string) *World {
	// split by newline, then split by character?
	lines := strings.Split(s, "\n")
	world := &World{}
	world.Units = make([][]Unit, len(lines))
	for i := range world.Units {
		world.Units[i] = make([]Unit, len(strings.Split(lines[i], "")))
	}
	for li, l := range lines {
		for ci, c := range strings.Split(l, "") {
			t := ParseType(c)
			world.Units[li][ci] = Unit{
				Type:  t,
				State: DefaultUnitState(t),
				Row:   ci,
				Col:   li,
				Up:    nil, Down: nil, Left: nil, Right: nil,
			}
		}
	}
	world.Height = len(lines)
	world.Width = len(strings.Split(lines[0], "")) // TODO: remove dodgy code (or comment)
	return world
}

func (w *World) Display() string {
	buf := ""
	for ri := range w.Units {
		for ci := range w.Units[ri] {
			d := w.Units[ri][ci].Display()
			if d == nil {
				buf += "?"
			} else {
				buf += string(*d)
			}
			if ci != (len(w.Units[ri]) - 1) {
				buf += " "
			}
		}
		buf += "\n"
	}
	return buf
}

func DefaultUnitState(t UnitType) UnitState {
	switch t {
	case Unrecognized:
	case Air:
	case Wall:
	case Reference:
	default:
		return Benign
	case Router:
	case Emitter:
	case Caster:
		return Neutral
	}
	return Benign
}

func ParseType(t string) UnitType {
	// disregard all further runes, first rune is always the one that matters
	// parsing will do the rest
	// could potentially cause bugs with runes, look into later
	return TypeMap[[]rune(t)[0]] // will return type or unrecognized if zero / not found
}
