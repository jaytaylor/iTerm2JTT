package iterm2

import (
	"testing"
)

func TestTipEmpty(t *testing.T) {
	testCases := []struct {
		tip   *Tip
		empty bool
	}{
		{
			tip:   nil,
			empty: true,
		},
		{
			tip:   &Tip{},
			empty: true,
		},
		{
			tip: &Tip{
				ID: 1,
			},
			empty: false,
		},
		{
			tip: &Tip{
				Title: "myTitle",
			},
			empty: false,
		},
	}
	for i, testCase := range testCases {
		if actual := testCase.tip.Empty(); actual != testCase.empty {
			t.Errorf("[test-case #%v] Expected tip.Empty()=%v but actual=%v", i, testCase.empty, actual)
		}
	}
}
