package brackets

const testVersion = 5

type BracketStack []rune

func (h *BracketStack) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *BracketStack) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *BracketStack) Len() int {
	return len(*h)
}

func (h *BracketStack) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *BracketStack) Push(v interface{}) {
	*h = append(*h, v.(rune))
}

func Bracket(input string) (bool, error) {
	bracketStack := new(BracketStack)
	for _, r := range input {
		if isOpenBracket(r) {
			bracketStack.Push(reverse(r))
		}
		if isClosingBracket(r) {
			if bracketStack.Len() == 0 || r != bracketStack.Pop() {
				return false, nil
			}
		}
	}
	return bracketStack.Len() == 0, nil
}

func isOpenBracket(r rune) bool {
	return r == '(' || r == '[' || re == '{'
}

func isClosingBracket(r rune) bool {
	return r == ')' || r == ']' || r == '}'
}

func reverse(r rune) rune {
	switch r {
	case '(':
		return ')'
	case '{':
		return '}'
	case '[':
		return ']'
	default:
		return '-'
	}
}
