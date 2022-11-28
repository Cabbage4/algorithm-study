package sm

type SM struct {
	first *node
	last  *node

	mp map[int]*node
}

type node struct {
	key  int
	val  int
	next *node
	pre  *node
}

func New() *SM {
	return &SM{
		mp: make(map[int]*node),
	}
}

func (s *SM) Size() int {
	return len(s.mp)
}

func (s *SM) Add(key, value int) {
	v, ok := s.mp[key]
	if !ok {
		nd := &node{key: key, val: value}
		s.mp[key] = nd

		if len(s.mp) == 1 {
			s.first = nd
			s.last = nd
			return
		}

		nd.next = s.first
		s.first.pre = nd
		s.first = nd
		return
	}
	s.changeOrder(v)
}

func (s *SM) Get(key int) int {
	v, ok := s.mp[key]
	if !ok {
		return -1
	}

	s.changeOrder(v)
	return v.val
}

func (s *SM) Range(f func(key, value int)) {
	tmp := s.first
	for tmp != nil {
		f(tmp.key, tmp.val)
		tmp = tmp.next
	}
}

func (s *SM) ReserveRange(f func(key, value int)) {
	tmp := s.last
	for tmp != nil {
		f(tmp.key, tmp.val)
		tmp = tmp.pre
	}
}

func (s *SM) changeOrder(v *node) {
	if s.first == v {
		return
	}

	if s.last == v {
		s.last = nil
		v.pre, v.next = nil, s.first
		s.first = v
		return
	}

	v.next.pre, v.pre.next = v.pre, v.next
	v.pre, v.next = nil, s.first
	s.first.pre = v
	s.first = v
}
