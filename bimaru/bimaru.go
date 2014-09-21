package bimaru

import (
	"fmt"
	"strconv"
	"github.com/zigloo/godlx/dlx"
)

func (b *Bimaru) SolveBimaru() {
	var root *dlx.Column
	var name dlx.RowName
	var dataSize, nship, orient int
	var s_h_size, s_v_size, size, max_pass int
	var f *Bimaru
	var ok, valid bool
	var numberRows uint64
	var sh *Shuffle
	var fs *fullships
	var c *constraint
	var grid *constraints_grid

	fs = getFullShips()

	sh = GetShuffle()

	f = b.initialize(b.h_size, b.v_size)

	max_constraint:= 0
	for _,cv:= range f.Constraint {
		for _,v:= range cv {
			if int(v) > max_constraint {
				max_constraint = int(v)
			}
		}
	}

	c = getConstraint(max_constraint)

	c.print()

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

	nship = len(f.Ship)

	constraints_size:= f.constraints_size() - fs.size()

	delta:= nship + constraints_size

	dataSize = delta + f.v_size * f.h_size + 1

	// orient to solve 2 ship on bottom indetermination (see print.go)
	orient = dataSize - 1

	fmt.Println("nship",nship,"constarints",constraints_size,"delta",delta)
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
				if fs.ships[p].dc == 1 {
					d[orient] = 1
				}

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

						d[delta + ri * f.h_size + ci] = 1

						// add water around ship (below and on the right)

						// below the ship
						if ri + s_v_size < f.v_size {
							for sci:= 0; sci <= s_h_size; sci++ {
								if ci + sci < f.h_size {
									if f.canHaveWater(ri + s_v_size, ci + sci) {
										d[delta + (ri + s_v_size) * f.h_size + ci + sci] = 1
									}
								}
							}
						}

						// on the right of the ship
						if ci + s_h_size < f.h_size {
							for sri:= 0; sri < s_v_size; sri++ {
								if ri + sri < f.v_size {
									if f.canHaveWater(ri + sri, ci + s_h_size) {
										d[delta + (ri + sri) * f.h_size + ci + s_h_size] = 1
									}
								}
							}
						}

						//grid = b.constraints(c,s_v_size,s_h_size,ri,ci)

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
												d[delta + (ri + sri) * f.h_size + ci + sci] = 1
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
											d[delta + (ri + s_v_size) * f.h_size + ci + sci] = 1
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
											d[delta + (ri + sri) * f.h_size + ci + s_h_size] = 1
										} else {
											ok = false
										}
									}
								}
							}

							if ok {
								if pass == 1 {
									d[orient] = 1
								}

								grid = f.constraints(c,s_v_size,s_h_size,ri,ci)

								numberRows += uint64(len(*grid))
								numberbyship += len(*grid)
							//	root.PrintRowData(d)
								//sh.AddRow(&d)
								sh.extend(nship,&d,grid)
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

	sh.AddToRoot(root)
	//sh.Print()

	fs.print()

	// add names
	for p := range name {
		if p < nship {
			if f.Ship[p] > 0 {
				name[p] = strconv.Itoa(int(f.Ship[p]))
			}
		} else if p == orient {
			name[p] = "h"
		} else if p >= delta {
			name[p] = "(" + strconv.Itoa((p - delta) / f.h_size) + "," + strconv.Itoa((p - delta) % f.h_size) + ")"
		} else {
			//name[p] = strconv.Itoa(int(p-nship))
			name[p] = ""
		}
	}

	root.SetColumnName(name)

	root.PrintColumn()

	fmt.Println("Number rows",numberRows)

	root.Solve(uint64(nship),dlx.First,1,GetPrint())
	//root.Solve(uint64(delta),dlx.First,2,GetPrint())
}
