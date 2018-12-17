package captcha

import (
    "strconv"
)

func Captcha(pattern int, leftOper int, oper int, rightOper int) string {
	if oper == 2{
		return "1 - one"
	}
	if pattern == 2 {
		return ConvertNumberToStringCharacter(rightOper) + " + " + strconv.Itoa(leftOper)
	}
	return strconv.Itoa(leftOper) +" + "+ConvertNumberToStringCharacter(rightOper)
}

func ConvertNumberToStringCharacter(number int) string {
	if number == 0 {
		return "zero"
	}
	if number == 1 {
		return "one"
	}
	if number == 2 {
		return "two"
	}
	if number == 3 {
		return "three"
	}
	if number == 4 {
		return "four"
	}
	if number == 5 {
		return "five"
	}
	if number == 6 {
		return "six"
	}
	if number == 7 {
		return "seven"
	}
	if number == 8 {
		return "eight"
	}
	if number == 9 {
		return "nine"
	}

	return ""
}