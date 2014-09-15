package bimaru

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/zigloo/godlx/dlx"
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

type shippos struct {
	r,c,dr,dc,size int
}

type fullships struct {
	ships map[int]*shippos
}

func getShipPos(r,c,dr,dc,size int) *shippos {
	var s *shippos

	s = new(shippos)
	s.r = r
	s.c = c
	s.dr = dr
	s.dc = dc
	s.size = size

	return s
}

func getFullShips() *fullships {
	var fs *fullships

	fs = new(fullships)
	fs.ships = make(map[int]*shippos)

	return fs
}

func (fs *fullships) addShip(r,c,dr,dc,size int) {
	fs.ships[len(fs.ships)] = getShipPos(r,c,dr,dc,size)
}

func (fs *fullships) print() {
	var sp *shippos

	fmt.Println("Full ships")
	fmt.Println("==========")
	for p:= range fs.ships {
		sp = fs.ships[p]
		fmt.Println(sp.size,"[",sp.r,",",sp.c,"] => (",sp.dr,",",sp.dc,")")
	}
}

func (fs *fullships) hasShip(r,c int) bool {
	for p:= range fs.ships {
		if fs.ships[p].r == r && fs.ships[p].c == c {
				return true
		}
	}
	return false
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

func (b *Bimaru) SolveBimaru() {
	var root *dlx.Column
	//var d dlx.RowData
	var name dlx.RowName
	var dataSize, nship int
	var s_h_size, s_v_size, size, max_pass int
	var f *Bimaru
	var ok, valid bool
	var numberRows uint64
	var sh *Shuffle
	var fs *fullships

	fs = getFullShips()

	sh = GetShuffle()

	f = b.initialize(b.h_size, b.v_size)

	//f. Print()

	nship = len(f.Ship)

	dataSize = nship + f.v_size * f.h_size

	root = dlx.GetRoot()

	// add water around ship elements
	// and check grid validity
	valid = true
	for ri, rv:= range f.Grid {
		for ci, cv:= range rv {
			if cv != f.Empty() && cv != f.Water() {
				// for all ship element, add water on diagonals
				switch cv {
					case f.Up() : {
						if ! f.addBoxWater(ri-1,ci-1,ri+2,ci-1) {
							valid = false
						}
						if ! f.addWater(ri-1,ci) {
							valid = false
						}
						if ! f.addBoxWater(ri-1,ci+1,ri+2,ci+1) {
							valid = false
						}
						if ! f.addUnknow(ri+1,ci,1,0,fs) {
							valid = false
						}
					}
					case f.Down() : {
						if ! f.addBoxWater(ri-2,ci-1,ri+1,ci-1) {
							valid = false
						}
						if ! f.addWater(ri+1,ci) {
							valid = false
						}
						if ! f.addBoxWater(ri-2,ci+1,ri+1,ci+1) {
							valid = false
						}
						if ! f.addUnknow(ri-1,ci,-1,0,fs) {
							valid = false
						}
					}
					case f.Left() : {
						if ! f.addBoxWater(ri+1,ci-1,ri+1,ci+2) {
							valid = false
						}
						if ! f.addWater(ri,ci-1) {
							valid = false
						}
						if ! f.addBoxWater(ri-1,ci-1,ri-1,ci+2) {
							valid = false
						}
						if ! f.addUnknow(ri,ci+1,0,1,fs) {
							valid = false
						}
					}
					case f.Right() : {
						if ! f.addBoxWater(ri-1,ci-2,ri-1,ci+1) {
							valid = false
						}
						if ! f.addWater(ri,ci+1) {
							valid = false
						}
						if ! f.addBoxWater(ri+1,ci-2,ri+1,ci+1) {
							valid = false
						}
						if ! f.addUnknow(ri,ci-1,0,-1,fs) {
							valid = false
						}
					}
					case f.One() : {
						if ! f.addBoxWater(ri-1,ci-1,ri-1,ci+1) {
							valid = false
						}
						if ! f.addWater(ri,ci-1) {
							valid = false
						}
						if ! f.addWater(ri,ci+1) {
							valid = false
						}
						if ! f.addBoxWater(ri+1,ci-1,ri+1,ci+1) {
							valid = false
						}
						// we found a full one ship
						fs.addShip(ri,ci,0,0,1)
					}
					case f.Center() : {
						if ! f.addWater(ri-1,ci-1) {
							valid = false
						}
						if ! f.addWater(ri-1,ci+1) {
							valid = false
						}
						if ! f.addWater(ri+1,ci-1) {
							valid = false
						}
						if ! f.addWater(ri+1,ci+1) {
							valid = false
						}
						// check if other Center are around
						// on the left
						if ci-1 >= 0 {
							if f.Grid[ri][ci-1] == f.Center() {
								if ! f.addUnknow(ri,ci-2,0,-1,fs) {
									valid = false
								}
							}
						} else {
							valid = false
						}
						// on the right
						if ci+1 < f.h_size {
							if f.Grid[ri][ci+1] == f.Center() {
								if ! f.addUnknow(ri,ci+2,0,1,fs) {
									valid = false
								}
							}
						} else {
							valid = false
						}
						// upsters
						if ri-1 >= 0 {
							if f.Grid[ri-1][ci] == f.Center() {
								if ! f.addUnknow(ri-2,ci,-1,0,fs) {
									valid = false
								}
							}
						} else {
							valid = false
						}
						// downsters
						if ri+1 < f.v_size {
							if f.Grid[ri+1][ci] == f.Center() {
								if ! f.addUnknow(ri+2,ci,1,0,fs) {
									valid = false
								}
							}
						} else {
							valid = false
						}
					}
				}
			}
		}
	}

	name = make(dlx.RowName, dataSize)

	numberRows = 0
	// add known full ships
	for p:= range fs.ships {
		found := false
		for si, ss:= range f.Ship {
			if int(ss) == fs.ships[p].size {
				found = true

				name[si] = strconv.Itoa(int(ss))

				f.Ship[si] = 0

				s_v_size = (fs.ships[p].size - 1) * fs.ships[p].dr + 1
				s_h_size = (fs.ships[p].size - 1) * fs.ships[p].dc + 1

				d := make(dlx.RowData,dataSize)

				d[si] = 1

				for ri:= fs.ships[p].r; ri < fs.ships[p].r + s_v_size; ri++ {
					if f.Constraint[1][ri] > 0 {
						f.Constraint[1][ri] -= 1
					} else {
						valid = false
					}
					for ci:= fs.ships[p].c; ci < fs.ships[p].c + s_h_size; ci++ {
						if f.Constraint[0][ci] > 0 {
							f.Constraint[0][ci] -= 1
						} else {
							valid = false
						}

						d[nship + ri * f.h_size + ci] = 1

						// add water around ship (below and on the right)

						// below the ship
						if ri + s_v_size < f.v_size {
							for sci:= 0; sci <= s_h_size; sci++ {
								if ci + sci < f.h_size {
									if f.canHaveWater(ri + s_v_size, ci + sci) {
										d[nship + (ri + s_v_size) * f.h_size + ci + sci] = 1
									}
								}
							}
						}

						// on the right of the ship
						if ci + s_h_size < f.h_size {
							for sri:= 0; sri < s_v_size; sri++ {
								if ri + sri < f.v_size {
									if f.canHaveWater(ri + sri, ci + s_h_size) {
										d[nship + (ri + sri) * f.h_size + ci + s_h_size] = 1
									}
								}
							}
						}

						numberRows++
					//	root.PrintRowData(d)
						sh.AddRow(&d)
						//root.AddRow(d)
					}
				}
				break
			}
		}
		if found == false {
			valid = false
		}
	}

	if ! valid {
		panic("Invalid grid.")
	}

	f.Print()

	fs.print()

	// add constraints
	for si, ship := range f.Ship {
		// skip full ship previously found
		if ship > 0 {
			s_h_size = 1
			s_v_size = int(ship)
			max_pass = 2
			if s_v_size * s_h_size == 1 {
				max_pass = 1
			}
			numberbyship := 0
			for pass:= 0; pass < max_pass; pass++ {
				// add ship on grid
				for ri := 0; ri <= f.v_size - s_v_size; ri++ {
					for ci:= 0; ci <= f.h_size - s_h_size; ci++ {
						// skip full ship
						if ! fs.hasShip(ri,ci) {
							// ship horizontal
							d := make(dlx.RowData,dataSize)
							ok = true // if constraints are satisfied
							// add ship
							d[si] = 1
							// the ship is placed in grid if constraints (v and h) are satisfied
							for sri:= 0; ok && sri < s_v_size; sri++ {
								if  f.Constraint[1][ri + sri] >= byte(s_h_size)  {
									for sci:= 0; ok && sci < s_h_size; sci++ {
										if f.Constraint[0][ci + sci] >= byte(s_v_size) {
											value:= f.Grid[ri+sri][ci+sci]
											// check ship head
											if sri == 0 && sci == 0 {
												if ! ( value == f.Empty() ||
												   ( max_pass == 2 &&
												      ( value == f.Left() || value == f.Up() || value == f.unknown() ) ||
												     max_pass == 1 && value == f.One() ) ) {
												//	fmt.Println(s_v_size,s_h_size,ri+1,ci+1,"refused: head",value)
													ok = false
												}
											} else if ( sri == s_v_size - 1 ) && ( sci == s_h_size - 1 ){// check ship tail
												if ! ( value == f.Empty() ||
												   ( max_pass == 2 &&
												      ( value == f.Right() || value == f.Down() || value == f.unknown() ) ||
												     max_pass == 1 && value == f.One() ) ) {
												//	fmt.Println(s_v_size,s_h_size,ri+1,ci+1,"refused: tail",value)
													ok = false
												}
											} else { // check ship center
												if ! ( value == f.Empty() ||
												   ( max_pass == 2 && 
												      ( value == f.Center() || value == f.unknown() ) ||
												     max_pass == 1 && value == f.One() ) ) {
												//	fmt.Println(s_v_size,s_h_size,ri+1,ci+1,"refused: center",value)
													ok = false
												}
											}
											if ok {
												d[nship + (ri + sri) * f.h_size + ci + sci] = 1
											}
										} else {
											ok = false
										}
									}
								} else {
									ok = false
								}
							}

							// add water around ship (below and on the right)

							// below the ship
							if ri + s_v_size < b.v_size {
								for sci:= 0; ok && sci <= s_h_size; sci++ {
									if ci + sci < f.h_size {
										if f.canHaveWater(ri + s_v_size, ci + sci) {
											d[nship + (ri + s_v_size) * f.h_size + ci + sci] = 1
										} else {
											ok = false
										}
									}
								}
							}

							// on the right of the ship
							if ci + s_h_size < f.h_size {
								for sri:= 0; ok && sri < s_v_size; sri++ {
									if ri + sri < f.v_size {
										if f.canHaveWater(ri + sri, ci + s_h_size) {
											d[nship + (ri + sri) * f.h_size + ci + s_h_size] = 1
										} else {
											ok = false
										}
									}
								}
							}

							if ok {
								numberRows++
								numberbyship++
							//	root.PrintRowData(d)
								sh.AddRow(&d)
								//root.AddRow(d)
							}
						}
					}
				}
				// swap ship size (h->v and v->h)
				size = s_h_size
				s_h_size = s_v_size
				s_v_size = size
			}
			fmt.Println("For ship",ship,numberbyship)
		}
	}

	//sh.Permut(numberRows)
	//sh.Permut(10000)
	//sh.AddToRoot(numberRows,root)

	sh.AddToRoot(10000,root)
	//sh.Print()

	fs.print()

	// add names
	for p := range name {
		if p < nship {
			if f.Ship[p] > 0 {
				name[p] = strconv.Itoa(int(f.Ship[p]))
			}
		} else {
			name[p] = "(" + strconv.Itoa((p - nship) / f.h_size) + "," + strconv.Itoa((p - nship) % f.h_size) + ")"
		}
	}

	root.SetColumnName(name)

	root.PrintColumn()

	fmt.Println("Number rows",numberRows)

	//root.Solve(uint64(nship),dlx.First,2,dlx.GetPrint())
	root.Solve(uint64(nship),dlx.First,2,GetPrint())
}
