/**********************************
/ Sedgewick's algorithm edition 4
/ Chapter 1 Quick Union
*********************************/
package quick_union

type Sites struct {
	id     []int
	number int
}

func Init(n int) *Sites {
	sites := &Sites{make([]int, n), n}
	for i := range sites.id {
		sites.id[i] = i
	}
	return sites
}

func (s *Sites) Union(p, q int) {
	pId := s.Find(p)
	qId := s.Find(q)

	if pId == qId {
		return
	}

	s.id[pId] = qId
	s.number = s.number - 1
}

func (s Sites) Find(p int) int {
	for s.id[p] != p {
		p = s.id[p]
	}
	return p
}

func (s Sites) Connected(p, q int) bool {
	return s.Find(p) == s.Find(q)
}

func (s Sites) Count() int {
	return s.number
}
