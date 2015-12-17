package main

import "fmt"

type ErrInvalidSymbol struct {
	Symbol string
}

func (e ErrInvalidSymbol) Error() string {
	return fmt.Sprintf("rhyming: invalid symbol: (%v)", e.Symbol)
}

type Symbol uint8

const (
	AA Symbol = iota
	AA0
	AA1
	AA2
	AE
	AE0
	AE1
	AE2
	AH
	AH0
	AH1
	AH2
	AO
	AO0
	AO1
	AO2
	AW
	AW0
	AW1
	AW2
	AY
	AY0
	AY1
	AY2
	B
	CH
	D
	DH
	EH
	EH0
	EH1
	EH2
	ER
	ER0
	ER1
	ER2
	EY
	EY0
	EY1
	EY2
	F
	G
	HH
	IH
	IH0
	IH1
	IH2
	IY
	IY0
	IY1
	IY2
	JH
	K
	L
	M
	N
	NG
	OW
	OW0
	OW1
	OW2
	OY
	OY0
	OY1
	OY2
	P
	R
	S
	SH
	T
	TH
	UH
	UH0
	UH1
	UH2
	UW
	UW0
	UW1
	UW2
	V
	W
	Y
	Z
	ZH
	Invalid
)

func ParseSymbol(s string) (Symbol, error) {
	if len(s) == 0 {
		return Invalid, ErrInvalidSymbol{s}
	}
	if len(s) == 1 {
		if s[0] == 'B' {
			return B, nil
		} else if s[0] == 'D' {
			return D, nil
		} else if s[0] == 'F' {
			return F, nil
		} else if s[0] == 'G' {
			return G, nil
		} else if s[0] == 'K' {
			return K, nil
		} else if s[0] == 'L' {
			return L, nil
		} else if s[0] == 'M' {
			return M, nil
		} else if s[0] == 'N' {
			return N, nil
		} else if s[0] == 'P' {
			return P, nil
		} else if s[0] == 'R' {
			return R, nil
		} else if s[0] == 'S' {
			return S, nil
		} else if s[0] == 'T' {
			return T, nil
		} else if s[0] == 'V' {
			return V, nil
		} else if s[0] == 'W' {
			return W, nil
		} else if s[0] == 'Y' {
			return Y, nil
		} else if s[0] == 'Z' {
			return Z, nil
		}
	} else if len(s) == 2 {
		if s[0] == 'A' {
			if s[1] == 'A' {
				return AA, nil
			} else if s[1] == 'E' {
				return AE, nil
			} else if s[1] == 'H' {
				return AH, nil
			} else if s[1] == 'O' {
				return AO, nil
			} else if s[1] == 'W' {
				return AW, nil
			} else if s[1] == 'Y' {
				return AY, nil
			}
		} else if s[0] == 'C' {
			if s[1] == 'H' {
				return CH, nil
			}
		} else if s[0] == 'D' {
			if s[1] == 'H' {
				return DH, nil
			}
		} else if s[0] == 'E' {
			if s[1] == 'H' {
				return EH, nil
			} else if s[1] == 'R' {
				return ER, nil
			} else if s[1] == 'Y' {
				return EY, nil
			}
		} else if s[0] == 'H' {
			if s[1] == 'H' {
				return HH, nil
			}
		} else if s[0] == 'I' {
			if s[1] == 'H' {
				return IH, nil
			} else if s[1] == 'Y' {
				return IY, nil
			}
		} else if s[0] == 'J' {
			if s[1] == 'H' {
				return JH, nil
			}
		} else if s[0] == 'N' {
			if s[1] == 'G' {
				return NG, nil
			}
		} else if s[0] == 'O' {
			if s[1] == 'W' {
				return OW, nil
			} else if s[1] == 'Y' {
				return OY, nil
			}
		} else if s[0] == 'S' {
			if s[1] == 'H' {
				return SH, nil
			}
		} else if s[0] == 'T' {
			if s[1] == 'H' {
				return TH, nil
			}
		} else if s[0] == 'U' {
			if s[1] == 'H' {
				return UH, nil
			} else if s[1] == 'W' {
				return UW, nil
			}
		} else if s[0] == 'Z' {
			if s[1] == 'H' {
				return ZH, nil
			}
		}
	} else if len(s) == 3 {
		if s[2] == '0' {
			if s[0] == 'A' {
				if s[1] == 'A' {
					return AA0, nil
				} else if s[1] == 'E' {
					return AE0, nil
				} else if s[1] == 'H' {
					return AH0, nil
				} else if s[1] == 'O' {
					return AO0, nil
				} else if s[1] == 'W' {
					return AW0, nil
				} else if s[1] == 'Y' {
					return AY0, nil
				}
			} else if s[0] == 'E' {
				if s[1] == 'H' {
					return EH0, nil
				} else if s[1] == 'R' {
					return ER0, nil
				} else if s[1] == 'Y' {
					return EY0, nil
				}
			} else if s[0] == 'I' {
				if s[1] == 'H' {
					return IH0, nil
				} else if s[1] == 'Y' {
					return IY0, nil
				}
			} else if s[0] == 'O' {
				if s[1] == 'W' {
					return OW0, nil
				} else if s[1] == 'Y' {
					return OY0, nil
				}
			} else if s[0] == 'U' {
				if s[1] == 'H' {
					return UH0, nil
				} else if s[1] == 'W' {
					return UW0, nil
				}
			}
		} else if s[2] == '1' {
			if s[0] == 'A' {
				if s[1] == 'A' {
					return AA1, nil
				} else if s[1] == 'E' {
					return AE1, nil
				} else if s[1] == 'H' {
					return AH1, nil
				} else if s[1] == 'O' {
					return AO1, nil
				} else if s[1] == 'W' {
					return AW1, nil
				} else if s[1] == 'Y' {
					return AY1, nil
				}
			} else if s[0] == 'E' {
				if s[1] == 'H' {
					return EH1, nil
				} else if s[1] == 'R' {
					return ER1, nil
				} else if s[1] == 'Y' {
					return EY1, nil
				}
			} else if s[0] == 'I' {
				if s[1] == 'H' {
					return IH1, nil
				} else if s[1] == 'Y' {
					return IY1, nil
				}
			} else if s[0] == 'O' {
				if s[1] == 'W' {
					return OW1, nil
				} else if s[1] == 'Y' {
					return OY1, nil
				}
			} else if s[0] == 'U' {
				if s[1] == 'H' {
					return UH1, nil
				} else if s[1] == 'W' {
					return UW1, nil
				}
			}
		} else if s[2] == '2' {
			if s[0] == 'A' {
				if s[1] == 'A' {
					return AA2, nil
				} else if s[1] == 'E' {
					return AE2, nil
				} else if s[1] == 'H' {
					return AH2, nil
				} else if s[1] == 'O' {
					return AO2, nil
				} else if s[1] == 'W' {
					return AW2, nil
				} else if s[1] == 'Y' {
					return AY2, nil
				}
			} else if s[0] == 'E' {
				if s[1] == 'H' {
					return EH2, nil
				} else if s[1] == 'R' {
					return ER2, nil
				} else if s[1] == 'Y' {
					return EY2, nil
				}
			} else if s[0] == 'I' {
				if s[1] == 'H' {
					return IH2, nil
				} else if s[1] == 'Y' {
					return IY2, nil
				}
			} else if s[0] == 'O' {
				if s[1] == 'W' {
					return OW2, nil
				} else if s[1] == 'Y' {
					return OY2, nil
				}
			} else if s[0] == 'U' {
				if s[1] == 'H' {
					return UH2, nil
				} else if s[1] == 'W' {
					return UW2, nil
				}
			}
		}
	}
	return Invalid, ErrInvalidSymbol{s}
}

type PhoneType uint8

const (
	Vowel PhoneType = iota
	Stop
	Affricate
	Fricative
	Aspirate
	Liquid
	Nasal
	Semivowel
	InvalidPhoneType
)

func (s Symbol) Type() PhoneType {
	switch s {
	case AA, AA0, AA1, AA2, AE, AE0, AE1, AE2, AH, AH0, AH1, AH2,
		AO, AO0, AO1, AO2, AW, AW0, AW1, AW2, AY, AY0, AY1, AY2,
		EH, EH0, EH1, EH2, ER, ER0, ER1, ER2, EY, EY0, EY1, EY2,
		IH, IH0, IH1, IH2, IY, IY0, IY1, IY2, OW, OW0, OW1, OW2,
		OY, OY0, OY1, OY2, UH, UH0, UH1, UH2, UW, UW0, UW1, UW2:
		return Vowel
	case B, D, G, K, P, T:
		return Stop
	case CH, JH:
		return Affricate
	case DH, F, S, SH, TH, V, Z, ZH:
		return Fricative
	case HH:
		return Aspirate
	case L, R:
		return Liquid
	case M, N, NG:
		return Nasal
	case W, Y:
		return Semivowel
	}
	return InvalidPhoneType
}
