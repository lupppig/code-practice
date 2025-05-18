package main

import (
	"testing"
)

func TestMultiply(t *testing.T) {

	var additiveTest = []struct {
		a      string
		b      string
		expect string
	}{
		{
			a:      "1",
			b:      "2",
			expect: "2",
		},
		{
			a:      "123",
			b:      "456",
			expect: "56088",
		},
		{
			a:      "498828660196",
			b:      "840477629533",
			expect: "419254329864656431168468",
		},
	}

	for _, mul := range additiveTest {
		got := Multiply2(mul.a, mul.b)
		if mul.expect != got {
			t.Errorf("expected %s got %s in test multiply string", mul.expect, got)
		}

	}
}
