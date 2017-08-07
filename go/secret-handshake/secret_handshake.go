package secret

const testVersion = 2

func Handshake(uNum uint) []string {
	num := (int)(uNum) //too lazy to cast all the time
	code := []string{}
	dict := map[int]string{
		1: "wink",
		2: "double blink",
		4: "close your eyes",
		8: "jump",
	}
	reverse := false
	if num/16 > 0 {
		num -= 16
		reverse = true
	}
	for i := 8; i >= 1; i /= 2 {
		if num/i > 0 {
			num -= i
			code = append(code, dict[i])
		}
	}
	if !reverse {
		code = reverseString(code)
	}
	return code
}

func reverseString(in []string) []string {
	var temp = make([]string, len(in))
	copy(temp, in)
	for i := len(temp) - 1; i >= 0; i-- {
		in[len(temp)-1-i] = temp[i]
	}
	return in
}
