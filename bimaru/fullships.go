package bimaru

import (
	"fmt"
)

type shippos struct {
	r,c,dr,dc,size int
}

type fullships struct {
	ships map[int]*shippos
}

func getShipPos(r,c,dr,dc,size int) *shippos {
	var s *shippos

	s = new(shippos)
	s.r = r
	s.c = c
	s.dr = dr
	s.dc = dc
	s.size = size

	return s
}

func getFullShips() *fullships {
	var fs *fullships

	fs = new(fullships)
	fs.ships = make(map[int]*shippos)

	return fs
}

func (fs *fullships) addShip(r,c,dr,dc,size int) {
	fs.ships[len(fs.ships)] = getShipPos(r,c,dr,dc,size)
}

func (fs *fullships) print() {
	var sp *shippos

	fmt.Println("Full ships")
	fmt.Println("==========")
	for p:= range fs.ships {
		sp = fs.ships[p]
		fmt.Println(sp.size,"[",sp.r,",",sp.c,"] => (",sp.dr,",",sp.dc,")")
	}
}

func (fs *fullships) hasShip(r,c int) bool {
	for p:= range fs.ships {
		if fs.ships[p].r == r && fs.ships[p].c == c {
				return true
		}
	}
	return false
}
