package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (part: %s) - index: expected (%d) <> received (%d)."

func TestIndex(t *testing.T) {
	t.Parallel() // execute the test in parallel
	testCases := []struct {
		text     string
		part     string
		expected int
	}{
		{
			text:     "The first text",
			part:     "The",
			expected: 0,
		},
		{
			text:     "",
			part:     "",
			expected: 0,
		},
		{
			text:     "Hi",
			part:     "hi",
			expected: -1,
		},
		{
			text:     "test",
			part:     "s",
			expected: 2,
		},
	}
	for _, tC := range testCases {
		t.Logf("Mass: %v", tC)
		actual := strings.Index(tC.text, tC.part)

		if actual != tC.expected {
			t.Errorf(msgIndex, tC.text, tC.part, tC.expected, actual)
		}
	}
}
