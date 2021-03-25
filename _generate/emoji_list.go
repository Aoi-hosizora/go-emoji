package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

const (
	unicodeEmojiListUrl = "https://unicode.org/emoji/charts/emoji-list.html"
)

func getEmojiList() ([]*emojiData, error) {
	doc, err := httpGet(unicodeEmojiListUrl)
	if err != nil {
		return nil, err
	}

	currGroup := ""
	currSubgroup := ""
	trs := doc.Find("tr")
	out := make([]*emojiData, 0, trs.Length())
	trs.Each(func(i int, sel *goquery.Selection) {
		if head := sel.Find("th.bighead"); head.Length() != 0 {
			currGroup = head.Text()
			currSubgroup = ""
			return
		}
		if head := sel.Find("th.mediumhead"); head.Length() != 0 {
			currSubgroup = head.Text()
			return
		}
		if sel.Find("td.code").Length() == 0 {
			return
		}

		uni := sel.Find("td.code").Text()
		name := sel.Find("td.name").First().Text()
		keyword := sel.Find("td.name").Last().Text()
		keyword = strings.ReplaceAll(keyword, " | ", ", ")
		raw, err := unicodeTextToString(uni)
		if err != nil {
			return
		}

		out = append(out, &emojiData{
			Group:    currGroup,
			Subgroup: currSubgroup,
			Name:     name,
			Var:      variable(name),
			Keyword:  keyword,

			Unicode: uni,
			UTF8:    raw,
		})
	})

	return out, nil
}
