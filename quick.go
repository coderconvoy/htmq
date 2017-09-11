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
func QText(s string) *Tag {
	return &Tag{
		TType: "text",
		Inner: s,
	}
}

func QInput(ttype, name string, options ...string) *Tag {
	return NewTag("input", append(options, "type", ttype, "name", name)...)
}

func QSubmit(text string, options ...string) *Tag {
	return NewTag("input", append(options, "type", "submit", "text", text, "value", text)...)
}

func QLink(href, inner string, options ...string) *Tag {
	return NewTextTag("a", inner, append(options, "href", href)...)
}

func QLinkRep(href, inner string, options ...string) *Tag {
	return NewTextTag("a", inner, append(options, "href", "javascript:void(0)", "onclick", "location.replace('"+href+"');return false;")...)
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

func QOption(realval, showval string) *Tag {
	return NewTextTag("option", showval, "value", realval)
}

func QScript(ss ...string) *Tag {
	inner := "\n"
	for _, v := range ss {
		inner += v + "\n// --- --- --- --- ---\n"
	}
	return NewTextTag("script", inner)
}

func QBut(inner string, onclick string, ss ...string) *Tag {

	res := NewTextTag("button", inner, append(ss, "onclick", onclick)...)
	return res

}

func (t *Tag) Wrap(ttype string, ss ...string) *Tag {
	return NewParent(ttype, []*Tag{t}, ss...)
}

func QUpload(action string, pre []*Tag, ss ...string) *Tag {
	a := append(pre,
		QInput("file", "uploadfile"),
		QInput("hidden", "token", "value", "{{.}}"),
		QSubmit("Upload File"),
	)
	return QForm(action, a, append(ss, "enctype", "multipart/form-data")...)
}
