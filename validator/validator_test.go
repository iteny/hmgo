package validator

import "testing"

func TestEmail(t *testing.T) {
	t.Parallel()
	var tests = []string{
		"",
		"@qq.com",
		"iteny@",
		"invalid.com",
		"iteny @qq.com",
		"iteny->span@qq.com",
		"iteny@qq.coffee..com",
		"iteny=span@qq.com",
		"foo@bar.com",
		"x@x.x",
		"foo@bar.com.au",
		"foo+bar@bar.com",
		"foo@bar.coffee",
		"foo@bar.bar.coffee",
		"foo@bar.中文网",
		"test|123@m端ller.com",
		"hans@m端ller.com",
		"hans.m端ller@test.com",
		"NathAn.daVIeS@DomaIn.cOM",
		"NATHAN.DAVIES@DOMAIN.CO.UK",
	}
	for _, test := range tests {
		result := Email(test)
		if result {
			t.Errorf("Execute Email(%c[1;0;34m%q%c[0m) => %c[1;0;32m%v%c[0m ", 0x1B, test, 0x1B, 0x1B, result, 0x1B)
		} else {
			t.Errorf("Execute Email(%c[1;0;34m%q%c[0m) => %c[1;0;31m%v%c[0m ", 0x1B, test, 0x1B, 0x1B, result, 0x1B)
		}
	}
}
func TestLength(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		param string
		min   int
		max   int
	}{
		{"asdfasdf", 1, 3},
		{"213", 1, 3},
		{"2134", 1, 3},
		{"", 1, 3},
		{"asdfasdfasdf111111111234", 1, 20},
		{"", 0, 3},
		{"1", 0, 3},
		{"阿斯蒂", 0, 3},
	}
	for _, test := range tests {
		result := ByteLength(test.param, test.min, test.max)
		if result {
			t.Errorf("Execute Email(%c[1;0;34m%q,min=%d,max=%d%c[0m) => %c[1;0;32m%v%c[0m ", 0x1B, test.param, test.min, test.max, 0x1B, 0x1B, result, 0x1B)
		} else {
			t.Errorf("Execute Email(%c[1;0;34m%q,min=%d,max=%d%c[0m) => %c[1;0;31m%v%c[0m ", 0x1B, test.param, test.min, test.max, 0x1B, 0x1B, result, 0x1B)
		}
	}
}
