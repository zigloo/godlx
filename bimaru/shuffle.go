package bimaru

import(
	"fmt"
	"math/rand"
	"github.com/zigloo/godlx/dlx"
)

type Shuffle struct {
	rows map[uint64](*dlx.RowData)
}

func GetShuffle() *Shuffle {
	var s *Shuffle

	s = new(Shuffle)
	s.rows = make(map[uint64](*dlx.RowData))

	return s
}

func (s *Shuffle) AddRow(d *dlx.RowData) {
		s.rows[uint64(len(s.rows))] = d
}

func (s *Shuffle) Permut(n uint64) {
	var r *rand.Rand
	var from, to uint64
	var size uint64
	var temp *dlx.RowData

	r = rand.New(rand.NewSource(25))

	if n > 0 {
		size = uint64(len(s.rows))
		if size > 2 {
			for i:=0; uint64(i) < n; i++ {
				for {
					from = uint64(r.Int63n(int64(size)))
					to = uint64(r.Int63n(int64(size)))
					fmt.Println(i+1,":",from,"->",to)
					if from != to {
						fmt.Println(s.rows[from],s.rows[to])
						temp = s.rows[from]
						s.rows[from] = s.rows[to]
						s.rows[to] = temp
						fmt.Println(s.rows[from],s.rows[to])
						break
					}
				}
			}
		}
	}

}

func (s *Shuffle) AddToRoot(r *dlx.Column) {
	for i:=0; i < len(s.rows); i++ {
		r.AddRow(*s.rows[uint64(i)])
	}
}

func (s *Shuffle) Print() {
	for i:=0; i < len(s.rows); i++ {
		fmt.Println(i, s.rows[uint64(i)])
	}
}
