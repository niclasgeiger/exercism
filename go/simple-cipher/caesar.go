package cipher

type Caesar struct {
	Shift
}

func NewCaesar() *Caesar {
	return &Caesar{
		Shift{
			s: 3,
		},
	}
}

func (caesar Caesar) Encode(clear string) (out string) {
	return shiftHelp(clear, func(r rune) rune { return r + int32(caesar.s) })
}

func (caesar Caesar) Decode(encoded string) (out string) {
	return shiftHelp(encoded, func(r rune) rune { return r - int32(caesar.s) })
}
