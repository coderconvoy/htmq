package htmq

type TagFilter func(*Tag) bool

//GetFirst takes a filter and returns the first (Depth first) Element to fulfil it,
func (t *Tag) GetFirst(f TagFilter, maxD int) *Tag {
	if f(t) {
		return t
	}
	if maxD == 0 { //negs go full depth -- careful of loops
		return nil
	}
	for _, c := range t.Children {
		res := c.GetFirst(f, maxD-1)
		if res != nil {
			return res
		}
	}
	return nil
}

func (t *Tag) GetAll(f TagFilter, maxD int) []*Tag {
	res := []*Tag{}
	if f(t) {
		res = append(res, t)
	}
	if maxD == 0 {
		return res
	}
	for _, c := range t.Children {
		cres := c.GetAll(f, maxD-1)
		res = append(res, cres...)
	}
	return res
}

func ByAnd(tfs ...TagFilter) TagFilter {
	return func(t *Tag) bool {
		for _, f := range tfs {
			if !f(t) {
				return false
			}
		}
		return true
	}
}

func ByAttr(k, v string) TagFilter {
	return func(t *Tag) bool {
		for _, a := range t.Attrs {
			if a.Name == k && a.Val == v {
				return true
			}
		}
		return false
	}
}

func ByType(tp string) TagFilter {
	return func(t *Tag) bool {
		return t.TType == tp
	}
}
