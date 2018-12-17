package captcha_test

import (
	"testing"
	captcha "github.com/suwitpongk/captcha"
)

func TestCaseOneOneOneOne(t *testing.T) {
	result := captcha.Captcha(1, 1, 1, 1)
	expect := "1 + one"
	if result != expect {
		t.Error("it should be " , expect)
	}
}

func TestCaseOneOneOneTwo(t *testing.T) {
	result := captcha.Captcha(1, 1, 1, 2)
	expect := "1 + two"
	if result != expect {
		t.Error("it should be " , expect)
	}
}