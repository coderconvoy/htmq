package htmq

import (
	"fmt"
	"strings"
	"testing"
)

func test_new(t *testing.T) {
	a := NewTag("button", "id", "btn_hello")
	a.AddChildren(NewTag("img", "id", "img_4", "class", "maker"))
	a.AddChildren(NewTag("button", "POOOO"))
	fmt.Println(a)
}

func test_Page(t *testing.T) {
	fmt.Println("PAGE")
	p, b := NewPage("gofish", "s/poo/g.css,poopopp.css")
	fmt.Println(p)
	fmt.Println("BODY")
	fmt.Println(b)

}

func Test_GetBy(t *testing.T) {
	p, b := NewPage("gofish", "s/poo/g.css,popop.css,pooooo.css")
	b2 := p.GetFirst(ByType("body"), 10)
	if b2 != b {
		t.Log("b2, not found same body")
		t.Log(p)
		t.Fail()
	}

	csss := p.GetAll(ByType("link"), 10)
	if len(csss) != 3 {
		t.Log("Should have 3 links")
		t.Log(p)
		t.Fail()
	}

	red := p.GetFirst(ByAttr("href", "popop.css"), -1)
	if red == nil {
		t.Log("No red item found")
		t.Log(p)
		t.FailNow()
	}
	if r, _ := red.Attr("href"); r != "popop.css" {
		t.Log("href of red element should be popop.css")
		t.Log(p)
		t.Fail()
	}
}

func Test_Button(t *testing.T) {
	but := QBut("Hello", "testf()", "^Goodbye")
	s := but.String()
	if !strings.Contains(s, "Goodbye") {
		t.Log("Button no HelloGoodbye")
		t.Log(s)
		t.Fail()
	}

	but = QBut("", "forr()", "id", "forrr", "!/s/s.svg", "^SSFORR")
	fmt.Println(but)
}
