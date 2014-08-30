package dlx

import (
	"github.com/zigloo/godlx/dlx"
)

func Example_empty() {
	var empty *dlx.Column

	empty = dlx.GetRoot()
	empty.Solve(0,dlx.Minimum,0,dlx.GetPrint())
	// Output:
	// --- #1 ---
	// ### (1) ###
}

func Example_minimum() {
	var root *dlx.Column

	root = dlx.GetRoot()
	root.AddRow(dlx.RowData{0,0,1,0,1,1,0})
	root.AddRow(dlx.RowData{1,0,0,1,0,0,1})
	root.AddRow(dlx.RowData{0,1,1,0,0,1,0})
	root.AddRow(dlx.RowData{1,0,0,1,0,0,0})
	root.AddRow(dlx.RowData{0,1,0,0,0,0,1})
	root.AddRow(dlx.RowData{0,0,0,1,1,0,1})
	root.SetColumnName(dlx.RowName{"A","B","C","D","E","F","G"})
	root.Solve(0,dlx.Minimum,0,dlx.GetPrint())
	// Output:
	// --- #1 ---
	// A D
	// E F C
	// B G
	// ### (1) ###
}

func Example_first() {
	var root *dlx.Column

	root = dlx.GetRoot()
	root.AddRow(dlx.RowData{0,0,1,0,1,1,0})
	root.AddRow(dlx.RowData{1,0,0,1,0,0,1})
	root.AddRow(dlx.RowData{0,1,1,0,0,1,0})
	root.AddRow(dlx.RowData{1,0,0,1,0,0,0})
	root.AddRow(dlx.RowData{0,1,0,0,0,0,1})
	root.AddRow(dlx.RowData{0,0,0,1,1,0,1})
	root.SetColumnName(dlx.RowName{"A","B","C","D","E","F","G"})
	root.Solve(0,dlx.First,0,dlx.GetPrint())
	// Output:
	// --- #1 ---
	// A D
	// B G
	// C E F
	// ### (1) ###
}

func Example_nosolution() {
	var root *dlx.Column

	root = dlx.GetRoot()
	root.AddRow(dlx.RowData{0,0,1,0,1,1,0})
	root.AddRow(dlx.RowData{0,0,0,1,1,0,1})
	root.SetColumnName(dlx.RowName{"A","B","C","D","E","F","G"})
	root.Solve(0,dlx.First,0,dlx.GetPrint())
	// Output:
	// ### (0) ###
}
