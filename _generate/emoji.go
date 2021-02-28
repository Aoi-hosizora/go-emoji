package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
	"unicode"
)

type emojiData struct {
	Group    string
	Subgroup string
	Name     string
	Var      string
	Keyword  string

	Unicode string
	UTF8    []byte
}

const (
	unicodeEmojiListUrl = "https://unicode.org/emoji/charts/emoji-list.html"
)

func getEmojiList() ([]*emojiData, error) {
	req, err := http.NewRequest("GET", unicodeEmojiListUrl, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body := resp.Body
	defer body.Close()
	bs, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bs))
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

		out = append(out, &emojiData{
			Group:    currGroup,
			Subgroup: currSubgroup,
			Name:     name,
			Var:      variable(name),
			Keyword:  keyword,

			Unicode: uni,
			UTF8:    []byte{},
		})
	})
	return out, nil
}

func variable(s string) string {
	first := rune(0)
	for _, r := range s {
		first = r
		break
	}
	if unicode.IsDigit(first) {
		s = "Number" + s
	}

	rep := []string{
		"-", " ",
		":", " ",
		".", " ",
		"_", " ",
		",", " ",
		"!", " ",
		"#", " ",
		"*", "Asterisk",
		"&", " ",
		"“", " ",
		"”", " ",
		"(", " ",
		")", " ",

		"’", "",
		"⊛", "",
	}
	s = strings.NewReplacer(rep...).Replace(s)
	sp := strings.Split(s, " ")
	sb := strings.Builder{}
	for _, part := range sp {
		if part == "" {
			continue
		}
		sb.WriteString(capitalize(part))
	}
	return sb.String()
}

func capitalize(s string) string {
	for i, r := range s {
		f := string(unicode.ToUpper(r))
		return f + s[i+len(f):]
	}
	return ""
}
