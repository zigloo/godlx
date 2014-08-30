package sudoku

import (
	"fmt"
	"strconv"
	"math"
	"github.com/zigloo/godlx/dlx"
)

type Sudoku [][]byte

func (s *Sudoku) full(n int) *Sudoku {
	var f Sudoku
	var bs int // block size

	bs = int(math.Sqrt(float64(n)))

	// check that n is a square
	if bs * bs != n {
		panic("n must be a square.")
	}

	// full cleared Sudoku grid
	f = make(Sudoku,n)
	for ri := range f {
		f[ri] = make([]byte,n)
	}

	// initialize from Sudoku pointer
	for ri, r := range f {
		for  ci := range r {
			if ri < len(*s) {
				if ci < len((*s)[ri]) {
					v := (*s)[ri][ci]
					if v >= 0 && int(v) <= n {
						f[ri][ci] = v
					} else {
						panic("Invalid value (" + strconv.Itoa(int(v)) + ") at (" + strconv.Itoa(ri) + "," + strconv.Itoa(ci) + ")")
					}
				}
			}
		}
	}

	return &f
}

func (s *Sudoku) Print(n int) {
	var f *Sudoku

	f = s.full(n)

	for _, r := range *f {
		for  ci, v := range r {
			if v != 0 {
				fmt.Print(v)
			} else {
				fmt.Print("-")
			}
			// Sanitize output (no trailing space for Examples testing)
			if ci < n - 1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Solve a Sudoku using the DLX algorithm
func (s * Sudoku) SolveSudoku(n int, maxsolutions uint64, ps dlx.PrintDLX) {
	var dsudoku *Sudoku
	var root *dlx.Column
	var d dlx.RowData
	var name dlx.RowName
	var bs, nc, ds int

	dsudoku = s.full(n)

	bs = int(math.Sqrt(float64(n))) // block size
	nc = n * n // number cells
	ds = 4 * nc // dlx data size

	root = dlx.GetRoot()

	// add constaints
	for ri, r := range *dsudoku {
		for  ci, v := range r {
			if v >= 1 && int(v) <= n {
				d = make(dlx.RowData,ds)

				// one digit by cell
				d[0 * nc + ri * n + ci] = 1
				// one digit by column
				d[1 * nc + ri * n + int(v) - 1] = 1
				// one digit by row
				d[2 * nc + ci * n + int(v) - 1] = 1
				// one digit by box cell
				d[3 * nc + ((ri / bs) * bs + ci / bs) * n + int(v) - 1] = 1

				root.AddRow(d)
			} else if v == 0 {
				for val:= 1; val <= n; val++ {
					d = make(dlx.RowData,ds)

					// one digit by cell
					d[0 * nc + ri * n + ci] = 1
					// one digit by column
					d[1 * nc + ri * n + val - 1] = 1
					// one digit by row
					d[2 * nc + ci * n + val - 1] = 1
					// one digit by box cell
					d[3 * nc + ((ri / bs) * bs + ci / bs) * n + val - 1] = 1

					root.AddRow(d)
				}
			}
		}
	}

	// add names
	name = make(dlx.RowName,ds)
	for p:= 0; p < ds; p++ {
		if p < nc {
			// cell position
			name[p] = "(" + strconv.Itoa(p / n)  + "," + strconv.Itoa(p % n) + ")"
		} else if  p < 2 * nc {
			// cell value
			name[p] = strconv.Itoa(p % n + 1)
		} else {
			name[p] = ""
		}
	}

	root.SetColumnName(name)

	root.Solve(0,dlx.Minimum,maxsolutions,ps)
}
