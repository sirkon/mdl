package mad

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/sanity-io/litter"
	"github.com/sirkon/mad/testdata"
	"github.com/stretchr/testify/require"
)

type resp map[int]Comment

func (r *resp) Decode(dest interface{}, header String, d *Decoder, ctx Context) (Sufficient, error) {
	dd, ok := dest.(*resp)
	if !ok {
		return nil, fmt.Errorf("dest must be %T, got %T", r, dest)
	}
	var tmp resp
	if dd == nil || *dd == nil {
		tmp = resp(map[int]Comment{})
	} else {
		tmp = *dd
	}
	chunk := strings.Split(header.Value, "=")
	statLit := chunk[1]
	status64, err := strconv.ParseInt(statLit, 10, 64)
	if err != nil {
		return nil, err
	}
	status := int(status64)
	var cmt Comment
	if err := d.Decode(&cmt, ctx); err != nil {
		return nil, err
	}
	tmp[status] = cmt
	return &tmp, nil
}

func (r *resp) Required() bool {
	return true
}

func TestStatuses(t *testing.T) {
	data, err := testdata.Asset("statuses.md")
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDecoder(bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	var dest struct {
		S *resp `mad:"status=\d+"`
	}
	ctx := NewContext()
	if err := d.Decode(&dest, ctx); err != nil {
		t.Fatal(err)
	}
	require.Equal(t,
		resp{
			200: Comment("к успеху пришёл\n"),
			404: Comment("нихуя не нашёл\n"),
			500: Comment("ебать пиздец\n"),
		},
		*dest.S,
	)
}

func TestRealStructure(t *testing.T) {
	data, err := testdata.Asset("metric.md")
	if err != nil {
		t.Fatal(err)
	}

	var dest struct {
		Type    string `mad:"type"`
		Queries []Code `mad:"queries,syntax=sql"`
	}

	if err := Unmarshal(data, &dest, NewContext()); err != nil {
		t.Fatal(err)
	}

	require.Equal(t, "count(country, version)", dest.Type)
	require.Len(t, dest.Queries, 4)
}

// curjob test
type response map[string]Code

// StatusExtractor for quick and dirty extraction, autogenerated with https://github.com/sirkon/ldetool --go-string
// enabled
type StatusExtractor struct {
	rest    string
	Content string
}

// Extract ...
func (p *StatusExtractor) Extract(line string) (bool, error) {
	p.rest = line
	var pos int

	// Looking for '(' and then pass it
	pos = strings.IndexByte(p.rest, '(')
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find \033[1m%c\033[0m in `\033[1m%s\033[0m`", '(', string(p.rest))
	}

	// Take until ')' as Content(string)
	pos = strings.IndexByte(p.rest, ')')
	if pos >= 0 {
		p.Content = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field Content", ')', string(p.rest))
	}

	return true, nil
}

func (r *response) Decode(dest interface{}, header String, d *Decoder, ctx Context) (Sufficient, error) {
	dd, ok := dest.(*response)
	if !ok {
		return nil, fmt.Errorf("dest must be %T, got %T", r, dest)
	}
	var tmp response
	if dd == nil || *dd == nil {
		tmp = response{}
	} else {
		tmp = *dd
	}

	// extract statuses (dirty)
	e := &StatusExtractor{}
	_, err := e.Extract(header.Value)
	if err != nil {
		return nil, err
	}
	statuses := strings.Split(e.Content, ",")
	for i, status := range statuses {
		status = strings.TrimSpace(status)
		statuses[i] = status
		if _, ok := tmp[status]; ok {
			return nil, locerrf(header, "duplicate status %s", status)
		}
	}

	var c Code
	if err := d.Decode(&c, ctx); err != nil {
		return nil, err
	}
	for _, status := range statuses {
		tmp[status] = c
	}
	return &tmp, nil
}

func (r *response) Required() bool {
	return true
}

func TestCurJobLike(t *testing.T) {
	data, err := testdata.Asset("curjoblike.md")
	if err != nil {
		t.Fatal(err)
	}

	var dest map[string]struct {
		Requests  []Code    `mad:"requests,syntax=json"`
		Responses *response `mad:"\s*response\s*\(.*?\)\s*"`
	}

	if err := Unmarshal(data, &dest, NewContext()); err != nil {
		t.Fatal(err)
	}
	require.Len(t, dest, 1)
	require.Len(t, dest["Service.Method"].Requests, 2)
	require.Len(t, *(*map[string]Code)(dest["Service.Method"].Responses), 3)
	requiredKeys := []string{"OK", "ERROR", "NOT_AVAILABLE"}
	for _, requiredKey := range requiredKeys {
		_, ok := (*(*map[string]Code)(dest["Service.Method"].Responses))[requiredKey]
		if !ok {
			t.Errorf(
				"%s is not in %s",
				requiredKey,
				litter.Sdump(*(dest["Service.Method"].Responses)),
			)
		}
	}
}

// testing README.md sample
type SonInfo struct {
	Info Comment
	Code *Code
}

func (s *SonInfo) Decode(d *Decoder, ctx Context) error {
	newctx := ctx.New()
	newctx.Set("syntax", "sql") // only sql syntax is allowed

	// consume Info (comment) first
	if err := d.Decode(&s.Info, ctx); err != nil {
		return err
	}

	// consume Code if is there (pointer of pointer of something means it is not mandatory)
	s.Code = &Code{}
	_ = d.Decode(&s.Code, ctx)
	return nil
}

type Sons map[string]SonInfo

func (s Sons) Decode(dest interface{}, header String, d *Decoder, ctx Context) (Sufficient, error) {
	// check what the dest is and initialize it if needed
	v, ok := dest.(Sons)
	if !ok {
		return nil, fmt.Errorf("dest must be %T, got %T", s, dest)
	}
	if v == nil {
		v = Sons{}
	}

	// check if exactly this son has not been filled before
	if _, ok := v[header.Value]; ok {
		return nil, LocatedError{
			Lin: header.Lin,
			Col: header.Col,
			Err: fmt.Errorf("duplicate sone ID `%s`", header.Value),
		}
	}

	// decode son information
	sonInfo := SonInfo{}
	if err := d.Decode(&sonInfo, ctx.New()); err != nil {
		return nil, err
	}
	v[header.Value] = sonInfo

	// return created object
	return v, nil
}

func (s Sons) Required() bool {
	return false
}

type ExtendedFamily struct {
	ParentName string `mad:"parentName"`
	ParentAge  uint   `mad:"parentAge"`
	Parent     struct {
		Sons     Sons `mad:"son\s\d+"`
		Daughter struct {
			Grandson Comment `mad:"grandson"`
		} `mad:"daughter"`
	} `mad:"parent"`
}

func loadExtendedFamily(input []byte) (ef ExtendedFamily, err error) {
	err = Unmarshal(input, &ef, NewContext())
	return
}

func TestHard(t *testing.T) {
	// Extended family starts with the man named Peter, 57 y.o.
	data, err := testdata.Asset("allfeatures.md")
	if err != nil {
		t.Fatal(err)
	}

	ef, err := loadExtendedFamily(data)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, "Peter", ef.ParentName)
	require.Equal(t, uint(57), ef.ParentAge)
	require.Equal(t, "just born and they haven't gave him a name yet\n", string(ef.Parent.Daughter.Grandson))
	require.Len(t, ef.Parent.Sons, 2)
	require.Equal(t, "made a bit of SQL today\n", string(ef.Parent.Sons["son 1"].Info))
	require.Equal(t, "sql", string(ef.Parent.Sons["son 1"].Code.Syntax))
	require.Equal(t, "SELECT * FROM table\n", string(ef.Parent.Sons["son 1"].Code.Code))
	require.Equal(t, "is a musician\n", string(ef.Parent.Sons["son 2"].Info))
}

func TestEasy(t *testing.T) {
	data, err := testdata.Asset("easy_example.md")
	if err != nil {
		t.Fatal(err)
	}
	var dest struct {
		Name   string `mad:"name"`
		Age    int    `mad:"age"`
		Field1 struct {
			Code   Code    `mad:"code,syntax=python"`
			Field2 Comment `mad:"field2"`
		} `mad:"field1"`
	}
	if err := Unmarshal(data, &dest, NewContext()); err != nil {
		t.Fatal(err)
	}

	require.Equal(t, "Name", dest.Name)
	require.Equal(t, 22, dest.Age)
	require.Equal(t, "random content\n", string(dest.Field1.Field2))
	require.Equal(t, "python", dest.Field1.Code.Syntax)
	require.Equal(t, `print "hello world!"`+"\n", dest.Field1.Code.Code)
}