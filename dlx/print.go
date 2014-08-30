package dlx

import (
	"fmt"
)

// The solutions are printed three parts
//  - the header
//  - the solutions
//  - the footer
// GetNumberDLXSolutions is used to transfer the number of solutions found to the footer
type PrintDLXHeader func()
type PrintDLXSolution func(s *Solution, k uint64)
type PrintDLXFooter func(maxsolutions, n uint64)
type GetNumberDLXSolutions func() uint64

// A package type to simplify function's description
type PrintDLX struct {
	PrintHeader PrintDLXHeader
	PrintSolution PrintDLXSolution
	PrintFooter PrintDLXFooter
	GetNumberSolutions GetNumberDLXSolutions
}

// These two helpers functions are the only needed to print the DLX solutions
func (c *Column) GetNext() *Column {
	return c.right
}

func (c *Column) GetCenterName() string {
	return c.center.name
}

// A closure is a means to get the number of solutions from DLX search function
func PrintSolutionDefault() (GetNumberDLXSolutions, PrintDLXSolution) {
	var numberSolution uint64

	numberSolution = 0

	return func() uint64 {
		return numberSolution
	}, func(s *Solution, k uint64) {
		var c, r *Column
		var i uint64

		numberSolution++

		fmt.Print("--- #")
		fmt.Print(numberSolution)
		fmt.Println(" ---")
		for i = 0; i < k; i++ {
			r = s.GetRow(i)
			c = r
			for {
				fmt.Print(c.GetCenterName())
				c = c.GetNext()
				if c == r {
					break
				} else {
					// Sanitizing output because Examples testing don't strip trailing spaces
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}

// Some default functions to print the solutions found by DLX search
func PrintHeaderDefault() {}

func PrintFooterDefault(maxsolutions, n uint64) {
	fmt.Print("### (")
	// When we limit the number of solutions, we don't always know if we have found all solutions
	if maxsolutions > 0 && n == maxsolutions {
		fmt.Print(">=")
	}
	fmt.Print(n)
	fmt.Println(") ###")
}

// A function to simplify the use of DLX search or Solve function
func GetPrint() PrintDLX {
	var ps PrintDLX

	ps.PrintHeader = PrintHeaderDefault
	ps.GetNumberSolutions, ps.PrintSolution = PrintSolutionDefault()
	ps.PrintFooter= PrintFooterDefault

	return ps
}
