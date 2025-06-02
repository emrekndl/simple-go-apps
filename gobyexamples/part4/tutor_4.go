package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/rand/v2"
	"os"
	"regexp"
	s "strings"
	"text/template"
	"time"
)

var p = fmt.Println

func _stringFunctions() {
	p("Contains:", s.Contains("test", "es"))
	p("Count:", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))
	p("Join:", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:", s.Repeat("a", 5))
	p("Replace:", s.Replace("foo", "o", "0", -1))
	p("Replace:", s.Replace("foo", "o", "0", 1))
	p("Split:", s.Split("a-b-c-d-e", "-"))
	p("ToLower:", s.ToLower("TEST"))
	p("ToUpper:", s.ToUpper("test"))
	p("Trim:", s.Trim("!hello!", "!"))
	p("TrimLeft:", s.TrimLeft("!hello!", "!"))
	p("TrimRight:", s.TrimRight("!hello!", "!"))
	p("TrimSpace:", s.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
	p("Fields:", s.Fields("  foo bar  baz   "))
	p("ContainsAny:", s.ContainsAny("team", "i"))
	p("ContainsRune:", s.ContainsRune("team", 't'))
	p("IndexAny:", s.IndexAny("team", "i"))
	p("IndexRune:", s.IndexRune("team", 'i'))
}

type point struct {
	x, y int
}

func _stringFormatting() {
	p := point{1, 2}

	fmt.Printf("Struct1(v): %v\n", p)
	fmt.Printf("Struct2(+v): %+v\n", p)
	fmt.Printf("Struct3(#v): %#v\n", p)

	fmt.Printf("Type(T): %T\n", p)
	fmt.Printf("Bool(t): %t\n", true)
	fmt.Printf("Int(d): %d\n", 123)
	fmt.Printf("Binary(b): %b\n", 14)
	fmt.Printf("Char(c): %c\n", 33)
	fmt.Printf("Hex(x): %x\n", 456)

	fmt.Printf("Float1(f): %f\n", 78.9)
	fmt.Printf("Float2(e): %e\n", 123400000.0)
	fmt.Printf("Float3(E): %E\n", 123400000.0)

	fmt.Printf("Str1(s): %s\n", "\"string\"")
	fmt.Printf("Str2(q): %q\n", "\"string\"")
	fmt.Printf("Str3(x): %x\n", "hex this")
	fmt.Printf("Pointer(p): %p\n", &p)

	fmt.Printf("Width1(6d): |%6d|%6d|\n", 12, 345)
	fmt.Printf("Width2(6.2f): |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("Width3(-6.2f): |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("Width4(6s): |%6s|%6s|\n", "foo", "b")
	fmt.Printf("Width5(-6s): |%-6s|%-6s|\n", "foo", "b")

	ss := fmt.Sprintf("Sprintf a %s\n", "string")
	fmt.Println(ss)
	fmt.Fprintf(os.Stderr, "Fprintf io: an %s\n", "error")
}

func _textTemplates() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value is {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 123)
	t1.Execute(os.Stdout, []string{
		"go",
		"rust",
		"c++",
		"c",
	})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Value is {{.}}\n")
	t2.Execute(os.Stdout, struct{ name string }{"furreal"})
	t2.Execute(os.Stdout, map[string]string{
		"name": "Umay",
	})
	// - means trim whitespaces
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "Range: {{range .}} {{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{
		"go",
		"rust",
		"c++",
		"c",
	})
}

func _regularExpressions() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println("Matchstring: ", r.MatchString("peach"))
	fmt.Println("FindString: ", r.FindString("peach punch"))
	fmt.Println("FindStringIndex: ", r.FindStringIndex("peach punch"))
	fmt.Println("FindStringSubmatch: ", r.FindStringSubmatch("peach punch"))
	fmt.Println("FindStringSubmatchIndex: ", r.FindStringSubmatchIndex("peach punch"))
	fmt.Println("FindAllString: ", r.FindAllString("peach punch pinch", -1))
	fmt.Println("FindAllStringSubmatchIndex: ", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println("FindAllString 2: ", r.FindAllString("peach punch pinch", 2))

	fmt.Println("Byte: ", r.Match([]byte("peach")))

	// MustCompile panics instead of returning an error
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("MustCompile: ", r)

	fmt.Println("ReplaceAllString: ", r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("ReplaceAllFunc: ", string(out))

}

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func _json() {
	boolJ, _ := json.Marshal(true)
	fmt.Println("toJson(bool): ", string(boolJ))

	intJ, _ := json.Marshal(1)
	fmt.Println("toJson(int): ", string(intJ))

	floatJ, _ := json.Marshal(2.34)
	fmt.Println("toJson(float): ", string(floatJ))

	strJ, _ := json.Marshal("gopher")
	fmt.Println("toJson(string): ", string(strJ))

	slcD := []string{"peach", "apple", "pear"}
	slcJ, _ := json.Marshal(slcD)
	fmt.Println("toJson(slice): ", string(slcJ))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapJ, _ := json.Marshal(mapD)
	fmt.Println("toJson(map): ", string(mapJ))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"peach", "apple", "pear"},
	}
	res1J, _ := json.Marshal(res1D)
	fmt.Println("toJson(struct1): ", string(res1J))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"peach", "apple", "pear"},
	}
	res2J, _ := json.Marshal(res2D)
	fmt.Println("toJson(struct2): ", string(res2J))

	// marshall indent
	res2Jindent, _ := json.MarshalIndent(res2D, "", "  ")
	fmt.Println("toJsonIndent(struct2-indent): ", string(res2Jindent))

	bytJ := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// json unknown values type for interface{}
	var dat map[string]interface{}

	if err := json.Unmarshal(bytJ, &dat); err != nil {
		panic(err)
	}
	fmt.Println("jsonTo(map): ", dat)

	num := dat["num"].(float64)
	fmt.Println("jsonTo(num-float64-type-casting): ", num)

	// interface{} to []interface{} casting
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println("jsonTo(str-string-type-casting): ", str1)

	str := `{"page": 1, "fruits": ["peach", "apple", "pear"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println("jsonTo(struct2): ", res)
	fmt.Println("jsonTo(struct2).fruits.0: ", res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant-> id:%v, name:%v, origin:%v", p.Id, p.Name, p.Origin)
}

func _xml() {
	coffee := &Plant{
		Id:   27,
		Name: "Coffee",
	}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.Marshal(coffee)
	fmt.Println("toXml(struct): ", string(out))
	out, _ = xml.MarshalIndent(coffee, " ", "    ")
	fmt.Println("toXmlIndent(struct-indent): ", xml.Header+string(out))

	var plant Plant
	if err := xml.Unmarshal(out, &plant); err != nil {
		panic(err)
	}
	fmt.Println("xmlTo(struct): ", plant)

	tomato := &Plant{
		Id:   28,
		Name: "Tomato",
	}
	tomato.Origin = []string{"Italy", "France"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "    ")
	fmt.Println("toXmlIndent(struct-nesting-indent): ", xml.Header+string(out))
}

func _time() {
	p := fmt.Println

	now := time.Now()
	p("Now: ", now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p("Then: ", then)

	p("Then.Year(): ", then.Year())
	p("Then.Month(): ", then.Month())
	p("Then.Day(): ", then.Day())
	p("Then.Hour(): ", then.Hour())
	p("Then.Minute(): ", then.Minute())
	p("Then.Second(): ", then.Second())
	p("Then.Nanosecond(): ", then.Nanosecond())
	p("Then.Location(): ", then.Location())

	p("Then.Weekday(): ", then.Weekday())
	p("Then.YearDay(): ", then.YearDay())
	// These methods compare two times
	p("Then.Before(now): ", then.Before(now))
	p("Then.After(now): ", then.After(now))
	p("Then.Equal(now): ", then.Equal(now))

	diff := now.Sub(then)
	p("Now.Sub(then)-diff: ", diff)

	p("diff.Hours(): ", diff.Hours())
	p("diff.Minutes(): ", diff.Minutes())
	p("diff.Seconds(): ", diff.Seconds())
	p("diff.Nanoseconds(): ", diff.Nanoseconds())

	p("Now.Add(diff): ", now.Add(diff))
	p("Now.Add(-diff): ", now.Add(-diff))
}

func _epoch() {
	now := time.Now()
	fmt.Println("Now: ", now)

	fmt.Println("Unix Time: ", now.Unix())
	fmt.Println("UnixMilli Time: ", now.UnixMilli())
	fmt.Println("UnixNano Time: ", now.UnixNano())
	fmt.Println("Format Unix Time: ", time.Unix(now.Unix(), 0))
	fmt.Println("Format UnixNano Time: ", time.Unix(now.UnixNano(), 0))
}

func _timeFormatingParsing() {
	p := fmt.Println

	t := time.Now()
	p("Format(RFC3339): ", t.Format(time.RFC3339))

	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	if e != nil {
		panic(e)
	}
	p("Parse(RFC3339): ", t1)

	p("Format(15:04:05): ", t.Format("15:04:05"))
	p("Format(3:04PM)", t.Format("3:04PM"))
	p("Format(Mon Jan _2 15:04:05 2006): ", t.Format("Mon Jan _2 15:04:05 2006"))
	p("Format(2006-01-02T15:04:05Z07:00): ", t.Format("2006-01-02T15:04:05Z07:00"))

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	if e != nil {
		panic(e)
	}
	p("Parse(3 04 PM): ", t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", t2.Year(), t2.Month(), t2.Day(), t2.Hour, t2.Minute, t2.Second)

	// ansic := "Mon Jan _2 15:04:05 2006"
	// _, e = time.Parse(ansic, "8:41PM")
	// if e != nil {
	// panic(e)
	// }
}

func _randomNumbers() {
	fmt.Print("Random Int(0<=x<100): ", rand.IntN(100), ", ")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	fmt.Println("Random Float64(0.0<=x<1.0): ", rand.Float64())

	fmt.Print("Random Float64(5.0<=x<10.0): ", (rand.Float64()*5)+5, ", ")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	seed1 := rand.NewPCG(42, 1024)
	rand1 := rand.New(seed1)
	fmt.Print("Random new seed2 PCG(42, 1024): ", rand1.IntN(100), ", ")
	fmt.Print(rand1.IntN(100))
	fmt.Println()

	seed2 := rand.NewPCG(42, 1024)
	rand2 := rand.New(seed2)
	fmt.Print("Random new seed2 PCG(42, 1024): ", rand2.IntN(100), ", ")
	fmt.Print(rand2.IntN(100))
	fmt.Println()
}

func main() {
	// String Functions
	_stringFunctions()
	fmt.Println("-------------")
	// String Formatting
	_stringFormatting()
	fmt.Println("-------------")
	// Text Templates
	_textTemplates()
	fmt.Println("-------------")
	// Regular Expressions
	_regularExpressions()
	fmt.Println("-------------")
	// JSON
	_json()
	fmt.Println("-------------")
	// XML
	_xml()
	fmt.Println("-------------")
	// Time
	_time()
	fmt.Println("-------------")
	// Epoch
	_epoch()
	fmt.Println("-------------")
	// Time Formating and Parsing
	_timeFormatingParsing()
	fmt.Println("-------------")
	// Random Numbers
	_randomNumbers()
	fmt.Println("-------------")
}
