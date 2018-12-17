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

func TestCaseOneOneOneThree(t *testing.T) {
	result := captcha.Captcha(1, 1, 1, 3)
	expect := "1 + three"
	if result != expect {
		t.Error("it should be " , expect)
	}
}

func TestCaseTwoOneOneOne(t *testing.T) {
	result := captcha.Captcha(2, 1, 1, 1)
	expect := "one + 1"
	if result != expect {
		t.Error("it should be " , expect)
	}
}

func TestCaseTwoOneOneTwo(t *testing.T) {
	result := captcha.Captcha(2, 1, 1, 2)
	expect := "two + 1"
	if result != expect {
		t.Error("it should be " , expect)
	}
}

func TestCaseOneTwoOneOne(t *testing.T) {
	result := captcha.Captcha(1, 2, 1, 1)
	expect := "2 + one"
	if result != expect {
		t.Error("it should be " , expect)
	}
}