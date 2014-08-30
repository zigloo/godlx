package nqueens

import (
	"strconv"
	"github.com/zigloo/godlx/dlx"
)

// Search the solutions of n Queens problem on a n x n 'chess' board using DLX algorithm
func SolveNQueens(n uint64, maxsolutions uint64, ps dlx.PrintDLX) {
	var root *dlx.Column
	var d dlx.RowData
	var name dlx.RowName
	var s,r,c,p uint64

	root = dlx.GetRoot()

	s = 2 * n + 2 * (2 * n - 1)
	d = make(dlx.RowData,s)
	name = make(dlx.RowName,s)

	// add constraints
	for r = 0; r < n; r++ {
		for c= 0; c < n; c++ {
			// row Column
			d[r] = 1
			// Column Column
			d[1 * n + c] = 1
			// upper diagonal Column
			d[2 * n + r + c] = 1
			// lower diagonal Column
			d[2 * n + (2 * n - 1) + (n - 1) + r - c] = 1

			//root.PrintRowData(d)
			root.AddRow(d)

			// reset RowData
			d = make(dlx.RowData,s)
		}
	}

	// add names
	for p = 0; p < s; p++ {
		if p < n {
			name[p] = "R" + strconv.FormatUint(p % n, 10)
		} else if p < 2 * n {
			name[p] = "C" + strconv.FormatUint(p % n, 10)
		} else {
			name[p] = ""
		}
	}

	root.SetColumnName(name)

	root.Solve(n,dlx.First,maxsolutions,ps)
}
