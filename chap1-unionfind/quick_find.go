/**********************************
/ Sedgewick's algorithm edition 4
/ Chapter 1 Quick Find
*********************************/
package unionfind

type Sites struct {
	Id     []int
	Number int
}

func Init(n int) *Sites {
	sites := &Sites{make([]int, n), n}
	for i := range sites.Id {
		sites.Id[i] = i
	}
	return sites
}

func (s *Sites) Union(p, q int) {
	pId := s.Find(p)
	qId := s.Find(q)

	if pId == qId {
		return
	}

	for i := range s.Id {
		if s.Find(i) == pId {
			s.Id[i] = qId
		}
	}
	s.Number = s.Number - 1
}

func (s Sites) Find(p int) int {
	return s.Id[p]
}

func (s Sites) Connected(p, q int) bool {
	return s.Find(p) == s.Find(q)
}

func (s Sites) Count() int {
	return s.Number
}
