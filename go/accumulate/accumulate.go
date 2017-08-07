package accumulate

const testVersion = 1

func Accumulate(in []string, f func(string) string) []string {
	out := make([]string, len(in))
	for i, item := range in {
		out[i] = f(item)
	}
	return out
}
