package data_struct

// Ring 循环链表
type Ring struct {
	next, prev *Ring       //前驱和后驱节点
	Value      interface{} // 数据
}

// 初始化空循环链表
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// 创建指定大小为n的循环链表
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

func (r *Ring) Prev() *Ring {
	if r.prev == nil {
		return r.init()
	}
	return r.prev
}

// Get 获取第n个节点(n为负，向前遍历，否则向后遍历)
func (r *Ring) Get(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// Link 添加节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// Unlink 删除链表的后n个节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Get(n + 1))
}

// Len 获取链表长度
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}
