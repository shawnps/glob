package match

import (
	"fmt"
	"github.com/gobwas/glob/runes"
	"unicode/utf8"
)

type List struct {
	List []rune
	Not  bool
}

func (self List) Match(s string) bool {
	r, w := utf8.DecodeRuneInString(s)
	if len(s) > w {
		return false
	}

	inList := runes.IndexRune(self.List, r) != -1
	return inList == !self.Not
}

func (self List) Len() int {
	return lenOne
}

func (self List) Index(s string, segments []int) (int, []int) {
	for i, r := range s {
		if self.Not == (runes.IndexRune(self.List, r) == -1) {
			return i, append(segments, utf8.RuneLen(r))
		}
	}

	return -1, nil
}

func (self List) String() string {
	var not string
	if self.Not {
		not = "!"
	}

	//	var list []string
	//	for _, r := range self.List {
	//		list = append(list, string(r))
	//	}

	return fmt.Sprintf("<list:%s[%s]>", not, string(self.List))
}
