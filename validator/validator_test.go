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
		result := Length(test.param, test.min, test.max)
		if result {
			t.Errorf("Execute Length(%c[1;0;34m%q,min=%d,max=%d%c[0m) => %c[1;0;32m%v%c[0m ", 0x1B, test.param, test.min, test.max, 0x1B, 0x1B, result, 0x1B)
		} else {
			t.Errorf("Execute Length(%c[1;0;34m%q,min=%d,max=%d%c[0m) => %c[1;0;31m%v%c[0m ", 0x1B, test.param, test.min, test.max, 0x1B, 0x1B, result, 0x1B)
		}
	}
}
func TestByteLength(t *testing.T) {
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
			t.Errorf("Execute ByteLength(%c[1;0;34m%q,min=%d,max=%d%c[0m) => %c[1;0;32m%v%c[0m ", 0x1B, test.param, test.min, test.max, 0x1B, 0x1B, result, 0x1B)
		} else {
			t.Errorf("Execute ByteLength(%c[1;0;34m%q,min=%d,max=%d%c[0m) => %c[1;0;31m%v%c[0m ", 0x1B, test.param, test.min, test.max, 0x1B, 0x1B, result, 0x1B)
		}
	}
}
func TestNull(t *testing.T) {
	t.Parallel()
	var tests = []string{
		"",
		"asdfasdf",
	}
	for _, test := range tests {
		result := Null(test)
		if result {
			t.Errorf("Execute Null(%c[1;0;34m%q%c[0m) => %c[1;0;32m%v%c[0m ", 0x1B, test, 0x1B, 0x1B, result, 0x1B)
		} else {
			t.Errorf("Execute Null(%c[1;0;34m%q%c[0m) => %c[1;0;31m%v%c[0m ", 0x1B, test, 0x1B, 0x1B, result, 0x1B)
		}
	}
}
func TestEnglish(t *testing.T) {
	t.Parallel()
	var tests = []string{
		"\n",
		"\r",
		"Ⅸ",
		"",
		"   fooo   ",
		"abc!!!",
		"abc1",
		"abc〩",
		"abc",
		"소주",
		"ABC",
		"FoObAr",
		"소aBC",
		"소",
		"달기&Co.",
		"〩Hours",
		"\ufff0",
		"\u0070", //UTF-8(ASCII): p
		"\u0026", //UTF-8(ASCII): &
		"\u0030", //UTF-8(ASCII): 0
		"123",
		"0123",
		"-00123",
		"0",
		"-0",
		"123.123",
		" ",
		".",
		"-1¾",
		"1¾",
		"〥〩",
		"모자",
		"ix",
		"۳۵۶۰",
		"1--",
		"1-1",
		"-",
		"--",
		"1++",
		"1+1",
		"+",
		"++",
		"+1",
	}
	for _, test := range tests {
		result := English(test)
		if result {
			t.Errorf("Execute English(%c[1;0;34m%#v%c[0m) => %c[1;0;32m%v%c[0m ", 0x1B, test, 0x1B, 0x1B, result, 0x1B)
		} else {
			t.Errorf("Execute English(%c[1;0;34m%#v%c[0m) => %c[1;0;31m%v%c[0m ", 0x1B, test, 0x1B, 0x1B, result, 0x1B)
		}
	}
}
func TestEnglishNumeric(t *testing.T) {
	t.Parallel()
	var tests = []string{
		"\n",
		"\r",
		"Ⅸ",
		"",
		"   fooo   ",
		"abc!!!",
		"abc1",
		"abc〩",
		"abc",
		"소주",
		"ABC",
		"FoObAr",
		"소aBC",
		"소",
		"달기&Co.",
		"〩Hours",
		"\ufff0",
		"\u0070", //UTF-8(ASCII): p
		"\u0026", //UTF-8(ASCII): &
		"\u0030", //UTF-8(ASCII): 0
		"123",
		"0123",
		"-00123",
		"0",
		"-0",
		"123.123",
		" ",
		".",
		"-1¾",
		"1¾",
		"〥〩",
		"모자",
		"ix",
		"۳۵۶۰",
		"1--",
		"1-1",
		"-",
		"--",
		"1++",
		"1+1",
		"+",
		"++",
		"+1",
	}
	for _, test := range tests {
		result := EnglishNumeric(test)
		if result {
			t.Errorf("Execute EnglishNumeric(%c[1;0;34m%#v%c[0m) => %c[1;0;32m%v%c[0m ", 0x1B, test, 0x1B, 0x1B, result, 0x1B)
		} else {
			t.Errorf("Execute EnglishNumeric(%c[1;0;34m%#v%c[0m) => %c[1;0;31m%v%c[0m ", 0x1B, test, 0x1B, 0x1B, result, 0x1B)
		}
	}
}
func TestNumeric(t *testing.T) {
	t.Parallel()
	var tests = []string{
		"\n",
		"\r",
		"Ⅸ",
		"",
		"   fooo   ",
		"abc!!!",
		"abc1",
		"abc〩",
		"abc",
		"소주",
		"ABC",
		"FoObAr",
		"소aBC",
		"소",
		"달기&Co.",
		"〩Hours",
		"\ufff0",
		"\u0070", //UTF-8(ASCII): p
		"\u0026", //UTF-8(ASCII): &
		"\u0030", //UTF-8(ASCII): 0
		"123",
		"0123",
		"-00123",
		"0",
		"-0",
		"123.123",
		" ",
		".",
		"-1¾",
		"1¾",
		"〥〩",
		"모자",
		"ix",
		"۳۵۶۰",
		"1--",
		"1-1",
		"-",
		"--",
		"1++",
		"1+1",
		"+",
		"++",
		"+1",
	}
	for _, test := range tests {
		result := Numeric(test)
		if result {
			t.Errorf("Execute Numeric(%c[1;0;34m%#v%c[0m) => %c[1;0;32m%v%c[0m ", 0x1B, test, 0x1B, 0x1B, result, 0x1B)
		} else {
			t.Errorf("Execute Numeric(%c[1;0;34m%#v%c[0m) => %c[1;0;31m%v%c[0m ", 0x1B, test, 0x1B, 0x1B, result, 0x1B)
		}
	}
}
func TestNumericNoHeadZero(t *testing.T) {
	t.Parallel()
	var tests = []string{
		"01",
		"1000",
		"0011",
		"1100000",
	}
	for _, test := range tests {
		result := NumericNoHeadZero(test)
		if result {
			t.Errorf("Execute NumericNoHeadZero(%c[1;0;34m%#v%c[0m) => %c[1;0;32m%v%c[0m ", 0x1B, test, 0x1B, 0x1B, result, 0x1B)
		} else {
			t.Errorf("Execute NumericNoHeadZero(%c[1;0;34m%#v%c[0m) => %c[1;0;31m%v%c[0m ", 0x1B, test, 0x1B, 0x1B, result, 0x1B)
		}
	}
}
func TestTime(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		param  string
		format string
	}{
		{"2016-12-31 11:00", RF3339Js},
		{"2016-12-31 11:00:00", RF3339Js},
		{"2016-12-31T11:00", RF3339Js},
		{"2016-12-31T11:00:00", RF3339Js},
		{"2016-12-31T11:00:00Z", RF3339Js},
		{"2016-12-31T11:00:00+01:00", RF3339Js},
		{"2016-12-31T11:00:00-01:00", RF3339Js},
		{"2016-12-31T11:00:00.05Z", RF3339Js},
		{"2016-12-31T11:00:00.05-01:00", RF3339Js},
		{"2016-12-31T11:00:00.05+01:00", RF3339Js},
		{"2016-12-31T11:00:00", RF3339T},
		{"2016-12-31T11:00:00Z", RF3339T},
		{"2016-12-31T11:00:00+01:00", RF3339T},
		{"2016-12-31T11:00:00-01:00", RF3339T},
		{"2016-12-31T11:00:00.05Z", RF3339T},
		{"2016-12-31T11:00:00.05-01:00", RF3339T},
		{"2016-12-31T11:00:00.05+01:00", RF3339T},
	}
	for _, test := range tests {
		result := Time(test.param, test.format)
		if result {
			t.Errorf("Execute Time(%c[1;0;34m%q,format=%#v%c[0m) => %c[1;0;32m%v%c[0m ", 0x1B, test.param, test.format, 0x1B, 0x1B, result, 0x1B)
		} else {
			t.Errorf("Execute Time(%c[1;0;34m%q,format=%#v%c[0m) => %c[1;0;31m%v%c[0m ", 0x1B, test.param, test.format, 0x1B, 0x1B, result, 0x1B)
		}
	}
}
