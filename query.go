package nounproject

import (
	"bytes"
	"fmt"
	"net/url"
)

// Query represents a query object
type Query map[string]string

// Pagination represents a pagination option object
type Pagination struct {
	Limit  int
	Offset int
	Page   int
}

// Merge function that merge another key-value map
func (q Query) Merge(m map[string]string) {
	for k, v := range m {
		q[k] = v
	}
}

// Add function that add a single key-value set
func (q Query) Add(k string, v string) {
	q[k] = v
}

// String function that make a query string
func (q Query) String() string {
	if len(q) == 0 {
		return ""
	}

	var buf bytes.Buffer
	for k, v := range q {
		buf.WriteString(url.QueryEscape(k))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(v))
		buf.WriteByte('&')
	}
	s := buf.String()
	return s[0 : len(s)-1]
}

// Query function that make a pagination query map
func (p *Pagination) Query() map[string]string {
	q := map[string]string{}
	if p == nil {
		return q
	}

	q["limit"] = fmt.Sprint(p.Limit)
	q["offset"] = fmt.Sprint(p.Offset)
	q["page"] = fmt.Sprint(p.Page)
	return q
}
