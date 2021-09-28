package main

import (
	"testing"
)

func TestShrinkString(t *testing.T) {
	type testVariant struct {
		Value  string
		Result string
	}

	testVariants := []*testVariant{
		{Value: "", Result: ""},
		{Value: "aaabbbcccccaaaaa", Result: "a8b3c5"},
		{Value: "zzzzcccUUUuu", Result: "U3c3u2z4"},
		{Value: "ЯЯЯБББддд", Result: "Б3Я3д3"},
		{Value: "ЯЯЯБББaaa", Result: "a3Б3Я3"},
		{Value: "ЯЯЯ", Result: "Я3"},
		{Value: "ZZZzz", Result: "Z3z2"},
	}
	for i, variant := range testVariants {
		result := ShrinkString(variant.Value)
		if result != variant.Result {
			t.Errorf("Output number %d: %q not equal to expected %q", i, result, variant.Result)
		}
	}
}
