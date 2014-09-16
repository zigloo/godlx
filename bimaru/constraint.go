package bimaru

import (
	"fmt"
)

type constraint struct {
	d [][][]byte
}

func getConstraint(max byte) *constraint {
	var c *constraint

	c = new(constraint)
	// max is the number of bits
	c.d = make([][][]byte,max + 1)

	for n:= range c.d {
		c.d[n] = make([][]byte,1 << uint(n))
		// iterate on values of n bits
		for d:= range c.d[n] {
			c.d[n][d] = make([]byte,n + 1)
			for p:= 0; p < n; p++ {
				bit := 1 << uint(p)
				if d & bit != 0 {
					// index 0 contains number ones
					c.d[n][d][0]++
					// bits representation
					// higher bit on lower index
					c.d[n][d][n - p] = 1
				}
			}
		}
	}

	return c 
}

func (c *constraint) print() {

	fmt.Println("Constraint")
	fmt.Println("==========")

	fmt.Println("Number_bits value : [number_ones] bit_representation (higher_bit_first)")
	fmt.Println("-----------------------------------------------------------------------")

	for ni,n:= range c.d {
		for vi,v:= range n {
			fmt.Print(ni," ",vi," : ")
			for di,d:= range v { 
				if di == 0 {
					fmt.Print("[",d,"]")
				} else if d == 1 {
					fmt.Print("1")
				} else {
					fmt.Print("0")
				}
				fmt.Print(" ")
			}
			fmt.Println()
		}
	}
}
