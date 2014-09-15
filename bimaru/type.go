package bimaru

import (
	"fmt"
	"strconv"
	"strings"
)

type Bimaru struct {
	// v_size (rows) x h_size (columns)
	Grid [][]byte
	// 0 : h_size (column constraints)
	// 1 : v_size (line constraints)
	Constraint [][]byte
	Ship []byte
	v_size int
	h_size int
}

// Empty
func (b *Bimaru) Empty() byte {
	return 0
}

// One
func (b *Bimaru) One() byte {
	return 1
}

// Down
func (b *Bimaru) Down() byte {
	return 2
}

// Up
func (b *Bimaru) Up() byte {
	return 3
}

// Left
func (b *Bimaru) Left() byte {
	return 4
}

// Right
func (b *Bimaru) Right() byte {
	return 5
}

// Center
func (b *Bimaru) Center() byte {
	return 6
}

// Water
func (b *Bimaru) Water() byte {
	return 7
}

// unknow to complete initial Bimaru
func (b *Bimaru) unknown() byte {
	return 8
}

func GetBimaru(h_size, v_size int) *Bimaru {
	var f Bimaru

	f.h_size = h_size
	f.v_size = v_size

	// full cleared Bimaru grid
	f.Grid = make([][]byte,v_size)
	for ri := range f.Grid {
		f.Grid[ri] = make([]byte,h_size)
	}

	f.Constraint = make([][]byte,2)
	f.Constraint[0] = make([]byte,h_size)
	f.Constraint[1] = make([]byte,v_size)

	return &f
}

func (b *Bimaru) initialize(h_size, v_size int) *Bimaru {
	var f Bimaru

	f = *GetBimaru(h_size,v_size)

	// initialize from Bimaru pointer
	for ri, r := range f.Grid {
		for  ci := range r {
			if ri < len(b.Grid) {
				if ci < len(f.Grid[ri]) {
					v := b.Grid[ri][ci]
					//TODO: check more constraints (no D on bottom, ...)
					switch v {
						case b.Empty(),
						     b.One(),
						     b.Down(),
						     b.Up(),
						     b.Left(),
						     b.Right(),
						     b.Center(),
						     b.Water(),
						     b.unknown() : f.Grid[ri][ci] = v
						default :
							panic("Invalid value (" + strconv.Itoa(int(v)) + ") at (" + strconv.Itoa(ri) + "," + strconv.Itoa(ci) + ")")
					}
				}
			}
		}
	}

	//TODO: check constraints
	f.Ship = make([]byte,len(b.Ship))

	for ri,r := range b.Ship {
		f.Ship[ri] = r
	}

	for ri,r := range f.Constraint {
		if ri < len(b.Constraint) {
			for ci := range r {
				if ci < len(b.Constraint[ri]) {
					v := b.Constraint[ri][ci]
					// TODO: get v constraint from ships
					if ri == 0 {
						if int(v) <= h_size {
							f.Constraint[ri][ci] = v
						}
					} else {
						if int(v) <= v_size {
							f.Constraint[ri][ci] = v
						}
					}
				}
			}
		}
	}

	return &f
}

func (b *Bimaru) Print() {
	var f *Bimaru

	f = b.initialize(b.h_size, b.v_size)

	for ri, r := range f.Grid {
		fmt.Print(f.Constraint[1][ri])
		fmt.Print(" ")
		for _, v := range r {
			switch v {
				case f.Empty(): fmt.Print("-")
				case f.One(): fmt.Print("O")
				case f.Down(): fmt.Print("v")
				case f.Up(): fmt.Print("^")
				case f.Left(): fmt.Print("<")
				case f.Right(): fmt.Print(">")
				case f.Center(): fmt.Print("*")
				case f.Water(): fmt.Print("W")
				case f.unknown(): fmt.Print("?")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Print("  ")
	for _,r := range f.Constraint[0] {
		fmt.Print(r)
		fmt.Print(" ")
	}
	fmt.Println()

	//TODO: sort by size and print a count of ships by size
	fmt.Println()
	for _,r := range f.Ship {
		if r > 0 {
			fmt.Print(r)
			fmt.Print(" ")
			fmt.Print(strings.Repeat("*",int(r)))
			fmt.Println()
		}
	}
}

func (b *Bimaru) canHaveWater(r, c int) bool {
	var cv byte

	if r >= 0 && r < int(len(b.Grid)) && c >= 0 && c < int(len(b.Grid[0])) {
		cv = b.Grid[r][c]
		if cv == b.Empty() || cv == b.Water() {
			return true
		} else {
			return false
		}
	}

	return true
}

func (b *Bimaru) addWater(r, c int) bool {
	var cv byte

	if r >= 0 && r < int(len(b.Grid)) && c >= 0 && c < int(len(b.Grid[0])) {
		cv = b.Grid[r][c]
		if cv == b.Empty() || cv == b.Water() {
			b.Grid[r][c] = b.Water()
			return true
		} else {
			return false
		}
	}

	return true
}

func (b *Bimaru) addBoxWater(rf,cf,rt,ct int) bool {
	var valid bool

	valid = true
	for ri:= rf; ri <= rt; ri++ {
		for ci:= cf; ci <=ct; ci++ {
			if ! b.addWater(ri,ci) {
				valid =false
			}
		}
	}

	return valid
}

func (b *Bimaru) addUnknow(rf,cf,dr,dc int,fs *fullships) bool {
	var valid bool
	var r,c int
	var end byte

	if dr == 0 {
		if dc > 0 {
			end = b.Right()
		} else {
			end = b.Left()
		}
	} else if dr > 0 {
		end = b.Down()
	} else {
		end = b.Up()
	}

	valid = true
	r = rf
	c = cf
	for {
		if r >= 0 && r < b.v_size && c >= 0 && c < b.h_size {
			if b.Grid[r][c] == b.Empty() {
				b.Grid[r][c] = b.unknown()
				// add water on diagonals
				if ! b.addWater(r+1,c-1) {
					valid = false
				}
				if ! b.addWater(r+1,c+1) {
					valid = false
				}
				if ! b.addWater(r-1,c-1) {
					valid = false
				}
				if ! b.addWater(r-1,c+1) {
					valid = false
				}
				break
			} else if b.Grid[r][c] == b.Center() {
				r += dr
				c += dc
			} else if b.Grid[r][c] == end {
				if (end == b.Down() || end == b.Right() ) && b.Grid[rf-dr][cf-dc] != b.Center() {
					fs.addShip(rf-dr,cf-dc,dr,dc,(r-rf+2*dr)*dr+(c-cf+2*dc)*dc)
				}
				break
			} else if b.Grid[r][c] == b.unknown() {
				break
			} else {
				valid = false
				break
			}
		} else {
			break
		}
	}

	return valid
}
