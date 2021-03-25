package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

func httpGet(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
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

	return doc, err
}

func unicodeTextToString(uni string) (string, error) {
	rawBs := make([]byte, 0)
	for _, part := range strings.Split(uni, " ") {
		partNum, err := strconv.ParseInt(strings.TrimLeft(part, "U+"), 16, 32)
		if err != nil {
			return "", err
		}
		partStr := string(rune(partNum))
		rawBs = append(rawBs, []byte(partStr)...)
	}

	sb := strings.Builder{}
	for _, num := range rawBs {
		sb.WriteString(fmt.Sprintf("\\x%02x", num))
	}
	return sb.String(), nil
}

type emojiData struct {
	Group    string
	Subgroup string
	Name     string
	Var      string
	Keyword  string

	Unicode string
	UTF8    string
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
