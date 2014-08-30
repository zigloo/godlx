package dlx

import (
	"fmt"
	"math"
)

// In DLX there are two types: column and data (column without name and size)
// But, as go doesn't understand hierarchy, only the column type is used
type Column struct {
	up, down, left, right, center *Column
	name string
	size uint64
}

// Constants to select the method to choose the next row
// First to choose the first row (chooseFirst function below)
// Minimum to choose the row with minimum data in columns (chooseMinimum function below)
const First bool = true
const Minimum bool = false

// Data row type of the problem to solve
type RowData []byte

// The names associated to the data row
// Those names are used to describe a solution
type RowName []string

// The root where data rows and names are attached
func GetRoot() *Column {
	var r *Column = new(Column)

	r.up = r
	r.down = r
	r.left = r
	r.right = r
	r.center = r
	r.name = ""
	r.size = 0

	return r
}

// Function tool to append a column to the right of a column 
func (root *Column) appendColumn() *Column {
	var c * Column = new(Column)

	c.left = root.left
	c.right = root
	c.right.left = c
	c.left.right = c
	c.up = c
	c.down = c
	c.name = ""
	c.size = 0

	return c
}

// Function tool to append a data on the bottom of a column
func (c *Column) appendData() *Column {
	var d *Column = new(Column)

	d.up = c.up
	d.down = c
	d.up.down = d
	d.down.up = d
	d.left = d
	d.right = d
	d.center = c
	c.size++

	return d
}

// Add a data row on the bottom of the columns
func (r *Column) AddRow(row RowData) {
	var c *Column
	var d, d_l *Column

	d_l = nil
	c = r.right

	for _,v := range row {
		if c == r {
			c = r.appendColumn()
		}
		if v == 1 {
			d = c.appendData()
			if d_l != nil {
				d.left = d_l
				d.right = d_l.right
				d.right.left = d
				d.left.right = d
			}
			d_l = d
		}
		c = c.right
	}
}

// DLX cover algorithm
func (col *Column) cover() {
	var r, c *Column

	col.right.left = col.left
	col.left.right = col.right

	r = col.down
	for {
		if r != col {
			c = r.right
			for {
				if c != r {
					c.down.up = c.up
					c.up.down = c.down
					c.center.size--
					c = c.right
				} else {
					break
				}
			}
			r = r.down
		} else {
			break
		}
	}
}

// DLX uncover algorithm
func (col *Column) uncover() {
	var r, c *Column

	r = col.up
	for {
		if r != col {
			c = r.left
			for {
				if c != r {
					c.center.size++
					c.down.up = c
					c.up.down = c
					c = c.left
				} else {
					break
				}
			}
			r = r.up
		} else {
			break
		}
	}
	col.right.left = col
	col.left.right = col
}

// Set the names of the columns
func (root *Column) SetColumnName(n RowName) {
	var c *Column

	c = root
	for _,p := range n {
		c = c.right
		if c == root {
			c = root.appendColumn()
		}
		c.name = p
	}
}

// chooseFirst search the first next row
func (r *Column) chooseFirst() *Column {
		return r.right
}

// chooseMinimum search a row with a minimum of data in columns 
func (r *Column) chooseMinimum() *Column {
	var p, c *Column
	var s uint64

	s = math.MaxUint64
	p = nil
	c = r.right
	for {
		if c != r {
			if c.size < s {
				s = c.size
				p = c
			}
			c = c.right
		} else {
			break
		}
	}

	return p
}

// Function to select the method to choose the next row
func (r *Column) choose(first bool) *Column {
	if first == First {
		return r.chooseFirst()
	} else {
		return r.chooseMinimum()
	}
}

// Core DLX algorithm
//   k is the actual search level
//   s contains the potential solution
//   maxlexel limits the search to the specified level
//   maxlevel is used when only the first maxlevel data of rows are mandatory and the others are used as constraints
//      0 for no level limit
//   first selects the method to choose the next row
//      true to choose the first row
//      false to choose a row with a minimum of data in columns
//   maxsolutions limits the number of solutions to search
//      0 to find all solutions
//   p a structure of functions to print the solutions (see print.go)

func (h *Column) search(k uint64, s *Solution, maxlevel uint64, first bool, maxsolutions uint64, p PrintDLX) {
	var r, c, col *Column

	if p.GetNumberSolutions() < maxsolutions || maxsolutions == 0 {
	if h.right == h || ( k == maxlevel && maxlevel > 0 ) {
		p.PrintSolution(s,k)
	} else {
		col = h.choose(first)
		col.cover()
		r = col.down
		for {
			if r != col {
				s.AddRow(k,r)
				c = r.right
				for {
					if c != r {
						c.center.cover()
						c = c.right
					} else {
						break
					}
				}
				h.search(k + 1, s, maxlevel, first, maxsolutions, p)
				r = s.GetRow(k)
				c = r.center
				c = r.left
				for {
					if c != r {
						c.center.uncover()
						c = c.left
					} else {
						break
					}
				}
				r = r.down
			} else {
				break
			}
		}
		col.uncover()
	}
	}
}

// A tool function to print the column's name and size
func (h *Column) PrintColumn() {
	var c *Column

	c = h.right
	for {
		if c != h {
			fmt.Print(c.name)
			fmt.Print(" ")
			c = c.right
		} else {
			break
		}
	}
	fmt.Println()

	c = h.right
	for {
		if c != h {
			fmt.Print(c.size)
			fmt.Print(" ")
			c = c.right
		} else {
			break
		}
	}
	fmt.Println()
}

// A tool function to print a data row
func (c *Column) PrintRowData(d RowData) {
	for _,p := range d {
		if p == 1 {
			fmt.Print("1")
		} else {
			fmt.Print("-")
		}
		fmt.Print(" ")
	}
	fmt.Println()
}

// A wrapper of the search algorithm
// It permits to print the solutions in the format specified by the PrintDLX structure
func (h *Column) Solve(maxlevel uint64, first bool, maxsolutions uint64, p PrintDLX) {
	var sol *Solution

	sol = GetSolution()

	p.PrintHeader()

	h.search(0,sol,maxlevel,first,maxsolutions,p)

	p.PrintFooter(maxsolutions,p.GetNumberSolutions())
}
