package pascal

const testVersion = 1

func Triangle(x int) [][]int {
	out := [][]int{{1}}
	for i := 1; i < x; i++ {
		pascal := []int{1}
		prev := out[i-1]
		for n := 0; n < len(prev)-1; n++ {
			pascal = append(pascal, prev[n]+prev[n+1])
		}
		pascal = append(pascal, 1)
		out = append(out, pascal)
	}
	return out
}
