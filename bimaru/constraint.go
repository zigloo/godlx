package bimaru

import (
	"fmt"
	"github.com/zigloo/godlx/dlx"
)

type constraint struct {
	// dim 1: total number of bits
	// dim 2: a value with dim 1 bits
	// dim 3: bit representation of dim 2 value
	//        dim 3 [0] contains number of bits in value
	d [][][]byte
}

type constraints_grid [][]byte

func getConstraint(max int) *constraint {
	var c *constraint

	c = new(constraint)
	// max is the number of bits
	c.d = make([][][]byte,max + 1)

	for n:= range c.d {
		// dim 2 index contains all values formed with n bits (size 2^n)
		// index 0 [k] contains the number of values with k bits
		c.d[n] = make([][]byte,1 << uint(n))
		// iterate on values of n bits
		for d:= range c.d[n] {
			// dim 3 contains the bit representation of value
			// index 0 contains number of bits in value
			c.d[n][d] = make([]byte,n + 1)
			for p:= 0; p < n; p++ {
				bit := 1 << uint(p)
				if d & bit != 0 {
					// index 0 contains number ones
					c.d[n][d][0]++
					// bits representation
					// higher bit on lower index
					c.d[n][d][n - p] = 1
				}
			}
		}
		// Then number of bits is the combinaison of k ones over n bits
		// Pascal rows
		// 1
		// 2 1
		// 3 3 1
		// ...
		for i:=1; i < len(c.d[n]); i++ {
			c.d[n][0][c.d[n][i][0]] += 1
		}
	}

	return c 
}

func (c *constraint) print() {

	fmt.Println("Constraint")
	fmt.Println("==========")

	fmt.Println("Number_bits value : [number_ones] bit_representation (higher_bit_first)")
	fmt.Println("-----------------------------------------------------------------------")

	for ni,n:= range c.d {
		for vi,v:= range n {
			fmt.Print(ni," ",vi," : ")
			for di,d:= range v { 
				if di == 0 {
					fmt.Print("[",d,"]")
				} else {
					fmt.Print(d)
				}
				fmt.Print(" ")
			}
			fmt.Println()
		}
	}
}

func (b *Bimaru) constraints_size() int {
	// the number of bits we need is the sum of all constraints
	bit_size := 0
	for _,v :=range b.Constraint[0] {
		bit_size += int(v)
	}
	for _,v :=range b.Constraint[1] {
		bit_size += int(v)
	}

	return bit_size

}

func (b *Bimaru) constraints(co *constraint, s_v_size, s_h_size, r, c int) *constraints_grid {
	var bit_size, co_size int
	var grid constraints_grid


	// the number of bits we need is the sum of all constraints
	bit_size = b.constraints_size()

	// the number of constraints (row) is the product of the number of way to write ship sizes
	// on each constraint  
	co_size = 1
	for ri:= r; ri < r + s_v_size ; ri++ {
		// each row is s_h_size size
		co_size *= int(co.d[b.Constraint[1][ri]][0][s_h_size])
	}
	for ci:=c; ci < c + s_h_size; ci++ {
		// each row is s_v_size size
		co_size *= int(co.d[b.Constraint[0][ci]][0][s_v_size])
	}


	// the grid of constraints
	grid = make(constraints_grid,co_size)
	for ri:= range grid {
		grid[ri] = make([]byte,bit_size)
	}

	// index will contains the constraints indexes used to fill the grid
	index:= make([]int,b.h_size + b.v_size)
	pos:= 0
	for ci,cv:= range b.Constraint[0] {
		index[ci] = pos
		pos += int(cv)
	}
	for ri,rv:=range b.Constraint[1] {
		index[b.h_size + ri] = pos
		pos += int(rv)
	}

	// insert constraints in grid
	// for rows
	// repeat: number of times to repeat a row line
	repeat:=co_size
	for ri:= r; ri < r + s_v_size ; ri++ {
		// blocrepeat: number of times to repeat all lines
		blocrepeat:= co_size / repeat
		blocsize:= repeat

		grid_v_delta:= int(co.d[b.Constraint[1][ri]][0][s_h_size])
		repeat = repeat / grid_v_delta
		for bloc:= 0; bloc < blocrepeat; bloc++ {
			pos:= 0
			for ti,t:= range co.d[b.Constraint[1][ri]] {
				if int(t[0]) == s_h_size {
					for bi,bv:= range co.d[b.Constraint[1][ri]][ti] {
						if bi != 0 {
							for l:= 0; l < repeat; l++ {
								grid[bloc*blocsize + pos*repeat + l][index[b.h_size + ri] + bi - 1] = bv
							}
						}
					}
					pos++
				}
			}
		}
	}

	// for columns
	for ci:= c; ci < c + s_h_size ; ci++ {
		// blocrepeat: number of times to repeat all lines
		blocrepeat:= co_size / repeat
		blocsize:= repeat

		grid_v_delta:= int(co.d[b.Constraint[0][ci]][0][s_v_size])
		repeat = repeat / grid_v_delta
		for bloc:= 0; bloc < blocrepeat; bloc++ {
			pos:= 0
			for ti,t:= range co.d[b.Constraint[0][ci]] {
				if int(t[0]) == s_v_size {
					for bi,bv:= range co.d[b.Constraint[0][ci]][ti] {
						if bi != 0 {
							for l:= 0; l < repeat; l++ {
								grid[bloc*blocsize + pos*repeat + l][index[ci] + bi - 1] = bv
							}
						}
					}
					pos++
				}
			}
		}
	}

/*
	// print the grid to check result
	for _,rv:= range grid {
		pos:= 0
		for ci,cv:= range rv {
			if ci == index[pos] {
				fmt.Print("|")
				for {
					if pos+1 < len(index) {
						pos++
						if index[pos] != ci {
							break
						}
					} else {
						break
					}
				}
			}
			fmt.Print(cv)
		}
		fmt.Println("|")
	}
*/
	return &grid
}

func (s *Shuffle) extend(delta int, d *dlx.RowData, g *constraints_grid) {
	for _,rv:= range *g {
		completed_data:= make(dlx.RowData,len(*d))
		for di,dv:= range *d {
			completed_data[di] = dv
		}
		for ci,cv:= range rv {
			completed_data[delta + ci] = cv
		}

		s.AddRow(&completed_data)
	}
}
