package bimaru

import (
	"github.com/zigloo/godlx/bimaru"
)

func Example_bimaru() {
	var b *bimaru.Bimaru

	b = bimaru.GetBimaru(10,10)

	b.Ship = []byte{4,3,3,2,2,2,1,1,1,1}
	b.Constraint[0] = []byte{1,2,3,4,2,3,1,2,0,2}
	b.Constraint[1] = []byte{1,0,2,2,2,0,2,3,5,3}

	//b.Grid[2][6] = b.One()
	//b.Grid[6][9] = b.One()
	//b.Grid[8][1] = b.Up()

	//b.Grid[6][1] = b.Up()
	//b.Grid[7][1] = b.Center()
	//b.Grid[8][1] = b.Center()
	//b.Grid[9][1] = b.Down()
	//b.Grid[1][1] = b.Left()
	//b.Grid[1][2] = b.Center()
	//b.Grid[1][3] = b.Right()
	//b.Grid[1][1] = b.Up()
	//b.Grid[2][1] = b.Center()
	//b.Grid[8][5] = b.Center()
	//b.Grid[8][6] = b.Center()
	//b.Grid[3][1] = b.Down()

	b.Print()

	b.SolveBimaru()
	// Output:
	//
}
