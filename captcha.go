package captcha

func Captcha(pattern int, leftOper int, oper int, rightOper int) string {
	return "1 + "+ConvertNumberToStringCharacter(rightOper)
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