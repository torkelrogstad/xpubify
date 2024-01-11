package main

import "testing"

func TestToXpub(t *testing.T) {
	type test struct {
		input, output string
	}

	for _, testcase := range []test{
		{
			"xpub6CXR8Z7GJHdiZEmHaS3YFvbEubePmWmijribhcfa2sYE3N36VsWYFk3MUyxugSH769zgsKz58Dm4CT9t5yyvu16NotGKjHDhiKPTDUvEqBP",
			"xpub6CXR8Z7GJHdiZEmHaS3YFvbEubePmWmijribhcfa2sYE3N36VsWYFk3MUyxugSH769zgsKz58Dm4CT9t5yyvu16NotGKjHDhiKPTDUvEqBP",
		},
		{
			"zpub6rBwjtT6beigFq9XF9cng6nFFXwHekkia5m3GQTLntHz9ZfZ1BqfVsMdXPt5gFawuSEJNHBC3YU9y2P1XNoxVUTaYZfAu6rgFmWjzhjwfdf",
			"xpub6CXR8Z7GJHdiZEmHaS3YFvbEubePmWmijribhcfa2sYE3N36VsWYFk3MUyxugSH769zgsKz58Dm4CT9t5yyvu16NotGKjHDhiKPTDUvEqBP",
		},
	} {
		t.Run(testcase.input, func(t *testing.T) {

			converted, err := toXpub(testcase.input)
			if err != nil {
				t.Fatal(err)
			}

			if converted != testcase.output {
				t.Errorf("expected:\t%s", testcase.output)
				t.Errorf("got:\t\t%s", converted)
			}
		})
	}
}
