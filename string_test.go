package dry

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StringMap(t *testing.T) {
	result := StringMap(strings.TrimSpace, []string{"  a  ", " b ", "c", "  d", "e  "})
	correct := []string{"a", "b", "c", "d", "e"}
	if len(result) != len(correct) {
		t.Fail()
	}
	for i := range result {
		if result[i] != correct[i] {
			t.Fail()
		}
	}
}

func Test_StringFilter(t *testing.T) {
	hFunc := func(s string) bool {
		return strings.HasPrefix(s, "h")
	}
	result := StringFilter(hFunc, []string{"cheese", "mouse", "hi", "there", "horse"})
	correct := []string{"hi", "horse"}
	if len(result) != len(correct) {
		t.Fail()
	}
	for i := range result {
		if result[i] != correct[i] {
			t.Fail()
		}
	}
}

func Test_StringFindBetween(t *testing.T) {
	const s = "Hello <em>World</em>!"

	between, remainder, found := StringFindBetween(s, "<em>", "</em>")
	if between != "World" {
		t.Fail()
	}
	if remainder != "!" {
		t.Fail()
	}
	if !found {
		t.Fail()
	}

	between, remainder, found = StringFindBetween(s, "l", "l")
	if between != "" {
		t.Fail()
	}
	if remainder != "o <em>World</em>!" {
		t.Fail()
	}
	if !found {
		t.Fail()
	}

	between, remainder, found = StringFindBetween(s, "<i>", "</i>")
	if between != "" {
		t.Fail()
	}
	if remainder != s {
		t.Fail()
	}
	if found {
		t.Fail()
	}
}

func Test_StringStripHTMLTags(t *testing.T) {
	const withHTML = "<div>Hello > World <br/> <im src='xxx'/>"
	skippedHTML := "Hello > World  "

	if StringStripHTMLTags(withHTML) != skippedHTML {
		t.Fail()
	}
}

func Test_StringReplaceHTMLTags(t *testing.T) {
	withHTML := "<div>Hello > World <br/> <im src='xxx'/>"
	replacedHTML := "xxHello > World xx xx"

	if StringReplaceHTMLTags(withHTML, "xx") != replacedHTML {
		t.Fail()
	}
}

func TestStringsCommonPrefix(t *testing.T) {
	tests := []struct {
		a, b, want string
	}{
		{
			"abc",
			"abcdef",
			"abc",
		},
		{
			"abc",
			"abc",
			"abc",
		},
		{
			"one",
			"two",
			"",
		},
		{
			"проверыа",
			"проверка",
			"провер",
		},
		{
			"п😹",
			"п😸",
			"п",
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run("StringsCommonPrefix_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, tt.want, StringsCommonPrefix(tt.a, tt.b))
		})
	}
}
