package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/skeptycal/siunits"
	. "github.com/skeptycal/types"
)

const (
	defaultMaxPaddingLevel = 80
	defaultPaddingLevel    = 1
	defaultPaddingString   = " "
)

type (
	// Any = types.Any

	// AnyMap map[Any]Any

	dict struct {
		name      string
		protected bool
		m         map[Any]Any
	}

	Dict interface {
		Keys() []Any
		Values() []Any
		GetSetter
	}
)

func NewDict(name string, protected bool) *dict {

	return &dict{
		name:      name,
		protected: protected,
		m:         make(AnyMap),
	}
}

func (d dict) Keys() []Any {
	keys := make([]Any, len(d.m))
	for k := range d.m {
		keys = append(keys, k)
	}
	return keys
}

func (d dict) Values() []Any {
	values := make([]Any, len(d.m))
	for v := range d.m {
		values = append(values, v)
	}
	return values
}

func (d dict) Get(key Any) (Any, error) {
	if v, ok := d.m[key]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("key not found: %v", key)
}

func (d dict) Set(key, value Any) error {
	d.m[key] = value
	return nil
}

func (d dict) Lock() {
	d.protected = true
}

func (d dict) Unlock() {
	d.protected = false
}

type (
	Padder interface {
		String() string
		Max() int
		Add(value int) error
		Set(level, value int) error

		// Char() string
		// PadLevel(n int) string
		// SetPaddingString(s string)
	}

	padding struct {
		level   int
		char    string
		current string
		m       map[int]int
	}
)

// NewPadder creates an indentation padder with 'levels' number
// of indentation levels using padString as the repeating unit.
//
// If levels < 0 or greater than the max, it is set to 1.
// If padString is "", it is set to a single space (" ").
func NewPadder(levels int, padString string) Padder {

	if levels < 0 || levels > defaultMaxPaddingLevel {
		levels = 0
	}

	if padString == "" {
		padString = defaultPaddingString
	}

	return &padding{
		level: levels,
		char:  defaultPaddingString,
		m:     make(map[int]int, levels),
	}
}

func (p *padding) String() string {
	if p.current == "" {
		p.setCurrent()
	}
	return p.current
}

func (p *padding) Keys() []int {
	keys := make([]int, len(p.m))
	for k := range p.m {
		keys = append(keys, k)
	}
	return keys
}

func (p *padding) Values() []int {
	values := make([]int, len(p.m))
	for _, v := range p.m {
		values = append(values, v)
	}
	return values
}

func (p *padding) Max() int { return len(p.m) }

func (p *padding) Add(value int) error {
	level := p.Max() + 1

	if _, ok := p.m[level]; ok {
		return fmt.Errorf("error adding padding value: (%v= %v)", level, value)
	}
	p.level = level
	p.m[level] = value
	return nil
}

func (p *padding) Set(level, value int) error {
	if level < 0 {
		return fmt.Errorf("padding level cannot be less than 0: %v", level)
	}

	if level > p.Max() {
		return p.Add(value)
	}

	if _, ok := p.m[level]; !ok {
		return fmt.Errorf("cannot lookup padding level: %v", level)
	}

	p.m[level] = value
	return nil
}

func (p *padding) setCurrent() {
	count := 0

	for _, v := range p.m {
		count += v
	}

	p.current = strings.Repeat(p.char, count)
}

func (p *padding) PadLevel(n int) string {
	if n > 0 && n < p.Max()+1 {
		p.level = n
		p.setCurrent()
	}
	return p.String()
}

// func (p *padding) Char() string {
// 	if p.char == "" {
// 		p.char = defaultPaddingString
// 	}
// 	return p.char
// }

// func (p *padding) SetPaddingString(s string) {
// 	sb := strings.Builder{}
// 	defer sb.Reset()

// 	for _, r := range s {
// 		if unicode.IsGraphic(r) {
// 			sb.WriteRune(r)
// 		}
// 	}
// }

// func (p *padding) Up() {
// 	if p.level < p.Max() {
// 		p.level += 1
// 		p.setCurrent()
// 	}
// }
// func (p *padding) Down() {
// 	if p.level > 0 {
// 		p.level -= 1
// 		p.setCurrent()
// 	}
// }

func main() {
	n := 1
	for i := 1; i < siunits.MaxIntLen; i++ {
		fmt.Printf("%v: %v (len: %v)\n", i, n, siunits.IntLen(n))
		// a := siunits.IntLen(n)
		n *= 10
	}

	pad := &padding{}

	pad.Set(0, 15)
	pad.Set(1, 15)
	p1 := pad.PadLevel(0)
	p2 := pad.PadLevel(1)

	fmt.Println()
	for i, test := range siunits.IntLenTests() {
		v := reflect.ValueOf(test)
		fmt.Print(p1)
		var header = fmt.Sprintf("test number %3d:", i)
		// var padlevel = 0
		// var padding[padlevel] = 15
		for j := 0; j < v.NumField(); j++ {

			f := v.Field(j)
			fmt.Print(p2)
			fmt.Printf("%15s%20s ...field %2d: %20v - length: %v\n", "", header, j, f, len(f.String()))

		}

	}

}
