package htmq

import "strings"

func QSelect(name string, options ...string) *Tag {
	ops := []*Tag{}
	id := ""
	for _, v := range options {
		if strings.HasPrefix(v, "#") {
			id = strings.TrimPrefix(v, "#")
			continue
		}
		ops = append(ops, NewTextTag("option", v, "value", v))
	}
	res := NewParent("select", ops, "id", id, "name", name)
	if id != "" {
		res.SetAttr("id", id)
	}

	return res
}

func QInput(ttype, name string, options ...string) *Tag {
	return NewTag("input", append(options, "type", ttype, "name", name)...)
}

func QSubmit(text string) *Tag {
	return NewTag("input", "type", "submit", "text", text, "value", text)
}

func QLink(href, inner string) *Tag {
	return NewTextTag("a", inner, "href", href)
}

func QImg(src string, options ...string) *Tag {
	return NewTag("img", append(options, "src", src)...)
}

func QForm(action string, chids []*Tag, options ...string) *Tag {
	return NewParent("form", chids, append(options, "method", "post", "action", action)...)
}

func QMulti(ptype, ttype string, chids ...string) *Tag {
	ops := []*Tag{}
	for _, v := range chids {
		ops = append(ops, NewTextTag(ttype, v))
	}
	return NewParent(ptype, ops)
}

func QScript(ss ...string) *Tag {
	inner := "\n"
	for _, v := range ss {
		inner += v + "\n// --- --- --- --- ---\n"
	}
	return NewTextTag("script", inner)
}

func QBut(inner string, onclick string, ss ...string) *Tag {
	//TODO,Look for images in ss
	return NewTextTag("button", inner, "onclick", onclick, ss...)
}

func (t *Tag) Wrap(ttype string, ss ...string) *Tag {
	return NewParent(ttype, []*Tag{t}, ss...)
}
