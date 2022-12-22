package stribog

// X-преобразование
func AddXor512(seq1, seq2, res *[]byte) {

	for i := 0; i < 64; i++ {
		(*res)[i] = (*seq1)[i] ^ (*seq2)[i]
	}

}

func AddModulo512(seq1, seq2, res *[]byte) {

	var t uint16 = 0

	for i := 63; i >= 0; i-- {
		t = uint16((*seq1)[i]+(*seq2)[i]) + (t >> 8)
		(*res)[i] = (byte)(t & 0xFF)
	}
}

// P-преобразование
func P_transformation(seq *[]byte) {

	var temp []byte
	temp = make([]byte, 64)

	for i := 0; i < Length; i++ {
		temp[i] = (*seq)[Tau[i]]
	}
	for i := 0; i < Length; i++ {
		(*seq)[i] = temp[i]
	}
}

// S-преобразование
func S_transformation(seq *[]byte) {

	for i := 0; i < Length; i++ {
		(*seq)[i] = Pi[(*seq)[i]]
	}

}

// L-преобразование
func L_transformation(seq *[]byte) {

	var v uint64

	temp := make([]byte, 64)
	for i := 0; i < Length; i++ {
		temp[i] = (*seq)[i]
	}

	for i := 0; i < 8; i++ {

		v = 0
		for j := 0; j < 8; j++ {

			for k := 0; k < 8; k++ {

				if (temp[j*8+i] & 0x80 >> k) != 0 {
					v ^= A[j*8+k]
				}
			}
		}

		for j := 0; j < 8; j++ {
			// Возможно не надо приводить к byte
			(*seq)[i*8+j] = byte(v >> (7 - j) * 8)
		}
	}
}
