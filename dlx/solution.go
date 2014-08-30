package dlx

type Solution struct {
	rows map[uint64]*Column
}

func GetSolution() *Solution {
	var s *Solution

	s = new(Solution)
	s.rows = make(map[uint64]*Column)

	return s
}

func (s *Solution) AddRow(k uint64, r *Column) {
	s.rows[k] = r
}

func (s *Solution) DelRow(k uint64, r *Column) {
	delete(s.rows, k)
}

func (s *Solution) GetRow(k uint64) *Column {
	return s.rows[k]
}
