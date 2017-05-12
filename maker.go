package htmq

import "strings"

type Attr struct {
	Name string
	Val  string
}

//Use this to specify an empty boolean attribute
const EMPTY = "\\nil"

type Tag struct {
	TType    string
	Attrs    []Attr
	Children []*Tag
	Inner    string
}

func NewTag(kind string, s ...string) *Tag {
	res := &Tag{kind, []Attr{}, make([]*Tag, 0), ""}
	res.AddAttrs(s...)
	return res
}

func NewParent(kind string, children []*Tag, s ...string) *Tag {
	res := &Tag{kind, []Attr{}, children, ""}
	res.AddAttrs(s...)
	return res
}

func NewTextTag(kind string, inner string, s ...string) *Tag {
	res := &Tag{kind, []Attr{}, make([]*Tag, 0), inner}
	res.AddAttrs(s...)
	return res
}

//Deprecated
func NewText(s string) *Tag {
	return &Tag{
		TType: "text",
		Inner: s,
	}
}

// NewPage takes the standard requests, for a page, and returns a page object
//Optional string Params:
//title - The page title
//css - coma separated links to css,
//js - coma separated links to js,
//return the top page object, and the body
func NewPage(ss ...string) (*Tag, *Tag) {
	//fill params
	title, css, js := "", "", ""
	if len(ss) > 0 {
		title = ss[0]
	}
	if len(ss) > 1 {
		css = ss[1]
	}
	if len(ss) > 2 {
		js = ss[2]
	}

	//create bases
	dt := NewTag("!DOCTYPE", "--html")
	mh := NewTag("html")
	head := NewTag("head")
	body := NewTag("body")
	mh.AddChildren(head, body)
	head.AddChildren(
		NewTextTag("title", title),
		NewTag("meta", "charset", "utf-8"),
	)
	for _, s := range strings.Split(css, ",") {
		if s != "" {
			head.AddChildren(NewTag("link", "rel", "stylesheet", "type", "text/css", "href", s))
		}
	}

	for _, s := range strings.Split(js, ",") {
		if s != "" {
			head.AddChildren(NewTag("script", "src", s))
		}

	}

	return NewParent("page", []*Tag{dt, mh}), body

}

func (t *Tag) Attr(k string) (string, bool) {
	for _, el := range t.Attrs {
		if el.Name == k {
			return el.Val, true
		}
	}
	return EMPTY, false
}

func (t *Tag) SetAttr(k, v string) {
	for _, el := range t.Attrs {
		if el.Name == k {
			el.Val = v
			return
		}
	}
	t.Attrs = append(t.Attrs, Attr{k, v})
}

//AddAttrs is a function for the super lazy.
//Use "--" to indicate check true
//"!" to add an image lazily
//"^" to add text as a child
func (t *Tag) AddAttrs(s ...string) {

	mode := 0
	atname := ""
	for _, v := range s {
		if mode == 0 {
			//New Variable name
			if strings.HasPrefix(v, "--") {
				vv := strings.TrimPrefix(v, "--")
				t.SetAttr(vv, EMPTY)
				continue
			}
			if strings.HasPrefix(v, "!") {
				t.AddChildren(QImg(strings.TrimPrefix(v, "!")))
				continue
			}
			if strings.HasPrefix(v, "^") {
				t.AddChildren(NewText(strings.TrimPrefix(v, "^")))
				continue
			}
			atname = strings.TrimPrefix(v, "\\")
			mode = 1
			continue
		}
		//mode == 1 no add attr
		t.SetAttr(atname, v)
		mode = 0

	}

}

func (self *Tag) AddChildren(ts ...*Tag) {
	self.Children = append(self.Children, ts...)
}

func Childless(ttype string) bool {
	ttype = strings.ToLower(ttype)
	childless := []string{"input", "br", "img", "meta", "!doctype"}
	for _, s := range childless {
		if s == ttype {
			return true
		}
	}
	return false
}
func (self *Tag) String() string {
	return self.toString("")
}

func (self *Tag) toString(pre string) string {
	res := ""
	pre2 := pre
	if self.TType != "page" && self.TType != "text" {
		res = pre + "<" + self.TType
		for _, v := range self.Attrs {
			if v.Val == EMPTY {
				res += " " + v.Name
				continue
			}
			res += " " + v.Name + "=" + "\"" + v.Val + "\""
		}
		res += ">"
		pre2 = pre + " "
	}
	if Childless(self.TType) {
		return res + "\n"
	}

	res += self.Inner

	if len(self.Children) > 0 {
		res += "\n"

		for i := 0; i < len(self.Children); i++ {
			res += self.Children[i].toString(pre2)
		}
		res += pre
	}
	if self.TType != "page" && self.TType != "text" {
		res += "</" + self.TType + ">\n"
	}

	return res
}
