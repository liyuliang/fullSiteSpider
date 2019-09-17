package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
)

type Dom struct {
	node *goquery.Document
}

func InitDom(html string) (*Dom, error) {

	p := new(Dom)

	htmlBuf := bytes.NewBufferString(html)
	dom, err := goquery.NewDocumentFromReader(htmlBuf)
	if err != nil {

		return p, err
	}

	p.node = dom
	return p, nil
}

func (d *Dom) Find(selector string) *goquery.Selection {
	return d.node.Find(selector)
}

func (d *Dom) FindAll(selector string) []*goquery.Selection {
	var doms []*goquery.Selection
	d.node.Find(selector).Each(func(i int, s *goquery.Selection) {
		doms = append(doms, s)
	})
	return doms
}
