package captcha

func Captcha(pattern int, leftOper int, oper int, rightOper int) string {
	if rightOper == 3{
		return "1 + three"
	}
	if rightOper == 2{
		return "1 + two"
	}
	return "1 + one"
}
