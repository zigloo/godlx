package sudoku

import (
	"fmt"
	"math"
	"strings"
	"strconv"
	"github.com/zigloo/godlx/dlx"
)

func CreateSolutionGrid(sol *dlx.Solution, k uint64) *Sudoku {
	var c, r *dlx.Column
	var i, row, col,val uint64
	var grid Sudoku
	var n string
	var size int

	size = int(math.Sqrt(float64(k)))

	grid = make(Sudoku,size)
	for ri := range grid {
		grid[ri] = make([]byte,size)
	}

	for i = 0; i < k; i++ {
		r = sol.GetRow(i)
		c = r
		row = 0
		col = 0
		for {
			n = c.GetCenterName()
			if len(n) > 0 {
				if strings.HasPrefix(n,"(") {
					coord := strings.Split(strings.TrimSuffix(strings.TrimPrefix(n,"("),")"),",")
					row,_ = strconv.ParseUint(coord[0],10,64)
					col,_ = strconv.ParseUint(coord[1],10,64)
				} else {
					val,_ = strconv.ParseUint(n,10,64)
				}
			}
			c = c.GetNext()
			if c == r {
				break
			}
		}
		grid[row][col] = byte(val)
	}

	return &grid
}

func PrintSolution() (dlx.GetNumberDLXSolutions, dlx.PrintDLXSolution) {
	var numberSolution uint64

	numberSolution = 0

	return func() uint64 {
			return numberSolution
		}, func(sol *dlx.Solution, k uint64) {
		var grid *Sudoku
		var size int

		numberSolution++

		fmt.Print("--- #")
		fmt.Print(numberSolution)
		fmt.Println(" ---")

		size = int(math.Sqrt(float64(k)))

		grid = CreateSolutionGrid(sol,k)

		grid.Print(size)
	}
}

func GetPrint() dlx.PrintDLX {
	var ps dlx.PrintDLX

	ps.PrintHeader = dlx.PrintHeaderDefault
	ps.GetNumberSolutions, ps.PrintSolution = PrintSolution()
	ps.PrintFooter = dlx.PrintFooterDefault

	return ps
}
