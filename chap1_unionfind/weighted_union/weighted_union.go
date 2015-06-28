/**********************************
/ Sedgewick's algorithm edition 4
/ Chapter 1 Weighted Union
*********************************/
package weighted_union

type Sites struct {
	id     []int
	weight []int
	number int
}

func Init(n int) *Sites {
	sites := &Sites{make([]int, n), make([]int, n), n}
	for i := range sites.id {
		sites.id[i] = i
		sites.weight[i] = 1
	}
	return sites
}

func (s *Sites) Union(p, q int) {
	pId := s.Find(p)
	qId := s.Find(q)

	if pId == qId {
		return
	}

	if s.weight[pId] < s.weight[qId] {
		s.id[pId] = qId
        s.weight[qId] = s.weight[pId] + s.weight[qId]
	} else {
		s.id[qId] = pId
        s.weight[pId] = s.weight[pId] + s.weight[qId]
	}
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
