package nqueens

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/zigloo/godlx/dlx"
)

type Nqueens [][]byte

func CreateSolutionGrid(sol *dlx.Solution, k uint64) *Nqueens {
	var c, r *dlx.Column
	var i, row, col uint64
	var grid Nqueens
	var n string

	grid = make(Nqueens,k)
	for ri := range grid {
		grid[ri] = make([]byte,k)
	}

	for i = 0; i < k; i++ {
		r = sol.GetRow(i)
		c = r
		row = 0
		col = 0
		for {
			n = c.GetCenterName()
			if len(n) > 0 {
				if strings.HasPrefix(n,"R") {
					row,_ = strconv.ParseUint(strings.TrimPrefix(n,"R"),10,64)
				} else {
					col,_ = strconv.ParseUint(strings.TrimPrefix(n,"C"),10,64)
				}
			}
			c = c.GetNext()
			if c == r {
				break
			}
		}
		grid[row][col] = 1
	}

	return &grid
}

func PrintSolution() (dlx.GetNumberDLXSolutions, dlx.PrintDLXSolution) {
	var numberSolution uint64

	numberSolution = 0

	return func() uint64 {
			return numberSolution
		}, func (sol *dlx.Solution, k uint64) {
		var grid *Nqueens

		numberSolution++

		fmt.Print("--- #")
		fmt.Print(numberSolution)
		fmt.Println(" ---")

		grid = CreateSolutionGrid(sol,k)

		for _,rv := range *grid {
			for _,cv := range rv {
				if cv == 1 {
					fmt.Print("x")
				} else {
					fmt.Print("-")
				}
			}
			fmt.Println()
		}
	}
}

func GetPrint() dlx.PrintDLX {
	var ps dlx.PrintDLX

	ps.PrintHeader = dlx.GetPrint().PrintHeader
	ps.GetNumberSolutions, ps.PrintSolution = PrintSolution()
	ps.PrintFooter = dlx.GetPrint().PrintFooter

	return ps
}
