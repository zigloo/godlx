package nqueens

import (
	"github.com/zigloo/godlx/nqueens"
)

func Example_onequeen() {
	nqueens.SolveNQueens(1,0,nqueens.GetPrint())
	// Output:
	// --- #1 ---
	// x
	// ### (1) ###
}

func Example_treequeens() {
	// No solution for 3 queens 
	nqueens.SolveNQueens(3,0,nqueens.GetPrint())
	// Output:
	// ### (0) ###
}

func Example_fivequeens() {
	// Ten solutions for 5 queens
	nqueens.SolveNQueens(5,0,nqueens.GetPrint())
	//  Output:
	// --- #1 ---
	// x----
	// --x--
	// ----x
	// -x---
	// ---x-
	// --- #2 ---
	// x----
	// ---x-
	// -x---
	// ----x
	// --x--
	// --- #3 ---
	// -x---
	// ---x-
	// x----
	// --x--
	// ----x
	// --- #4 ---
	// -x---
	// ----x
	// --x--
	// x----
	// ---x-
	// --- #5 ---
	// --x--
	// x----
	// ---x-
	// -x---
	// ----x
	// --- #6 ---
	// --x--
	// ----x
	// -x---
	// ---x-
	// x----
	// --- #7 ---
	// ---x-
	// x----
	// --x--
	// ----x
	// -x---
	// --- #8 ---
	// ---x-
	// -x---
	// ----x
	// --x--
	// x----
	// --- #9 ---
	// ----x
	// -x---
	// ---x-
	// x----
	// --x--
	// --- #10 ---
	// ----x
	// --x--
	// x----
	// ---x-
	// -x---
	// ### (10) ###
}

func Example_fivequeensfirstsolution() {
	// Limit 5 queens to the first solution
	nqueens.SolveNQueens(5,1,nqueens.GetPrint())
	//  Output:
	// --- #1 ---
	// x----
	// --x--
	// ----x
	// -x---
	// ---x-
	// ### (>=1) ###
}

func Example_fivequeenssevensolutions() {
	// Limit 5 queens to 7 solutions
	nqueens.SolveNQueens(5,7,nqueens.GetPrint())
	//  Output:
	// --- #1 ---
	// x----
	// --x--
	// ----x
	// -x---
	// ---x-
	// --- #2 ---
	// x----
	// ---x-
	// -x---
	// ----x
	// --x--
	// --- #3 ---
	// -x---
	// ---x-
	// x----
	// --x--
	// ----x
	// --- #4 ---
	// -x---
	// ----x
	// --x--
	// x----
	// ---x-
	// --- #5 ---
	// --x--
	// x----
	// ---x-
	// -x---
	// ----x
	// --- #6 ---
	// --x--
	// ----x
	// -x---
	// ---x-
	// x----
	// --- #7 ---
	// ---x-
	// x----
	// --x--
	// ----x
	// -x---
	// ### (>=7) ###
}

func Example_fivequeenslimitedtomorethansolutions() {
	// Ten solutions for 5 queens
	nqueens.SolveNQueens(5,11,nqueens.GetPrint())
	//  Output:
	// --- #1 ---
	// x----
	// --x--
	// ----x
	// -x---
	// ---x-
	// --- #2 ---
	// x----
	// ---x-
	// -x---
	// ----x
	// --x--
	// --- #3 ---
	// -x---
	// ---x-
	// x----
	// --x--
	// ----x
	// --- #4 ---
	// -x---
	// ----x
	// --x--
	// x----
	// ---x-
	// --- #5 ---
	// --x--
	// x----
	// ---x-
	// -x---
	// ----x
	// --- #6 ---
	// --x--
	// ----x
	// -x---
	// ---x-
	// x----
	// --- #7 ---
	// ---x-
	// x----
	// --x--
	// ----x
	// -x---
	// --- #8 ---
	// ---x-
	// -x---
	// ----x
	// --x--
	// x----
	// --- #9 ---
	// ----x
	// -x---
	// ---x-
	// x----
	// --x--
	// --- #10 ---
	// ----x
	// --x--
	// x----
	// ---x-
	// -x---
	// ### (10) ###
}
