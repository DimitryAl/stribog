package hash

// X-преобразование
func AddXor512(seq1, seq2, res *[]byte) {

	for i := 0; i < Length; i++ {
		(*res)[i] = (*seq1)[i] ^ (*seq2)[i]
	}
}

func AddModulo512(seq1, seq2, res *[]byte) {

	var t int = 0

	for i := 0; i < Length; i++ {
		t = int((*seq1)[i]) + int((*seq2)[i]) + (t >> 8)
		(*res)[i] = (byte)(t & 0xFF)
	}
}

// P-преобразование
func P_transformation(seq *[]byte) {

	var temp *[]byte

	for i := 0; i < Length; i++ {
		(*temp)[i] = (*seq)[tau[i]]
	}
	seq = temp

}

// S-преобразование
func S_transformation(seq *[]byte) {

	for i := 0; i < Length; i++ {
		(*seq)[i] = pi[(*seq)[i]]
	}
}

// L-преобразование
func L_transformation(seq *[]byte) {

	var v uint64
	/*
	 * subvectors of 512-bit vector (64*8 bits)
	 * an subvector is start at [j*8], its componenst placed
	 * with step of 8 bytes (due to this function is composition
	 * of P and L) and have length of 64 bits (8*8 bits)
	 */
	for i := 0; i < 8; i++ {
		v = 0
		/*
		* subvectors of 512-bit vector (64*8 bits)
		* an subvector is start at [j*8], its componenst placed
		* with step of 8 bytes (due to this function is composition
		* of P and L) and have length of 64 bits (8*8 bits)
		 */
		for k := 0; k < 8; k++ {
			/* bit index of current 8-bit component */
			for j := 0; j < 8; j++ {
				/* check if current bit is set */
				//
				if ((*seq)[i*8+k] & (1 << (7 - j))) != 0 {
					v ^= A[k*8+j]
				}
			}
		}
		for k := 0; k < 8; k++ {
			// Возможно не надо приводить к byte
			(*seq)[i*8+k] = byte((v & (uint64(0xFF) << (7 - k) * 8)) >> (7 - k) * 8)
		}
	}
}
