package bimaru

import (
	"fmt"
	"math"
	"strings"
	"strconv"
	"github.com/zigloo/godlx/dlx"
)

func CreateSolutionGrid(sol *dlx.Solution, k uint64) *Bimaru {
	var c, r *dlx.Column
	var i, row, col, val, rmin, cmin, rmax, cmax, v_size, h_size, dr, dc uint64
	var n string
	var fs *fullships
	var b *Bimaru
	var horizontal bool

	fs = getFullShips()

	v_size = 0
	h_size = 0
	for i = 0; i < k; i++ {
		r = sol.GetRow(i)
		c = r
		row = 0
		col = 0
		rmin = math.MaxUint64
		cmin = math.MaxUint64
		rmax = 0
		cmax = 0
		horizontal = false
		for {
			n = c.GetCenterName()
			if len(n) > 0 {
				if strings.HasPrefix(n,"(") {
					coord := strings.Split(strings.TrimSuffix(strings.TrimPrefix(n,"("),")"),",")
					row,_ = strconv.ParseUint(coord[0],10,64)
					col,_ = strconv.ParseUint(coord[1],10,64)
					if row < rmin {
						rmin = row
					} else if row > rmax {
						rmax = row
					}
					if col < cmin {
						cmin = col
					} else if col > cmax {
						cmax = col
					}
				} else if strings.HasPrefix(n,"h") {
					horizontal = true
				} else {
					val,_ = strconv.ParseUint(n,10,64)
				}
			}
			c = c.GetNext()
			if c == r {
				break
			}
		}
		if rmax > v_size {
			v_size = rmax
		}
		if cmax > h_size {
			h_size = cmax
		}

		// last data column solves the 2 ship bottom indetermination
		// solution data
		// ...
		// ... - - - - ...
		// ... - 1 1 - ...
		// ... - 1 1 - ...
		// where is the 2 ship
		// here
		// ... - x - - ...
		// ... - x - - ...
		// or there
		// ... - x x - ...
		// ... - - - - ...
		// Actually it is there
		// ... - x x - ...
		// ... - - - - ...
		if val == 1 {
			dr = 0
			dc = 0
		} else if horizontal {
			dr = 0
			dc = 1
		} else {
			dr = 1
			dc = 0
		}
		fs.addShip(int(rmin),int(cmin),int(dr),int(dc),int(val))
	}

	h_size++
	v_size++

	b = GetBimaru(int(h_size),int(v_size))

	for _,s:= range fs.ships {
		s_v_size := (s.size - 1) * s.dr + 1
		s_h_size := (s.size - 1) * s.dc + 1
		for ri:= 0; ri < s_v_size; ri++ {
			for ci:= 0; ci < s_h_size; ci++ {
				if ri + s.r < int(v_size) && ci + s.c < int(h_size) {
					element := b.unknown()
					if s.size == 1 {
						element = b.One()
					} else if ri == 0 && ci == 0 {
						if s.dr == 1 {
							element = b.Up()
						} else {
							element = b.Left()
						}
					} else if ri == (s_v_size - 1) && ci == (s_h_size -1) {
						if s.dr == 1 {
							element = b.Down()
						} else {
							element = b.Right()
						}
					} else {
						element = b.Center()
					}
					b.Grid[ri + s.r][ci + s.c] = element
				}
			}
		}
	}

	return b
}

func PrintSolution() (dlx.GetNumberDLXSolutions, dlx.PrintDLXSolution) {
	var numberSolution uint64

	numberSolution = 0

	return func() uint64 {
			return numberSolution
		}, func(sol *dlx.Solution, k uint64) {
		var b *Bimaru

		numberSolution++

		fmt.Print("--- #")
		fmt.Print(numberSolution)
		fmt.Println(" ---")

		b = CreateSolutionGrid(sol,k)

		b.Print(false)
	}
}

func GetPrint() dlx.PrintDLX {
	var ps dlx.PrintDLX

	ps.PrintHeader = dlx.PrintHeaderDefault
	ps.GetNumberSolutions, ps.PrintSolution = PrintSolution()
	ps.PrintFooter = dlx.PrintFooterDefault

	return ps
}
