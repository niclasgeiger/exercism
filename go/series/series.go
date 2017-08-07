package series

const testVersion = 2

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}
	out := make([]string, len(s)-n+1)
	for i := 0; i <= len(s)-n; i++ {
		out[i] = s[i : i+n]
	}
	return out
}

func UnsafeFirst(n int, s string) (out string) {
	defer func() {
		if r := recover(); r != nil {
			out = s + "0"
		}
	}()
	out = All(n, s)[0]
	return out
}

func First(n int, s string) (string, bool) {
	all := All(n, s)
	if all == nil {
		return s, false
	}
	return all[0], true

}
