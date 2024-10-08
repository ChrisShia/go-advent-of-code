package maths

//For Reference:
//int64 max : 	9223372036854775807
//20! = 		2432902008176640000
//50Choose25 =  126410606437752
//60Choose30 =  118264581564861424
//60Choose30 =	7116179530141809008

func GCD(i, j int) int {
	//TODO: https://proofwiki.org/wiki/GCD_for_Negative_Integers
	//TODO: needs work for positive negative cases
	if i == 0 {
		return j
	}
	for j > 0 {
		if i > j {
			i = i - j
		} else {
			j = j - i
		}
	}
	return i
}

func HCF(i, j int) int {
	//TODO: Needs work for negative cases (overflows)
	if j == 0 {
		return i
	} else if i > j {
		return HCF(i-j, j)
	}
	return HCF(i, j-i)
}

func LCM(i, j int) int {
	//TODO: needs work for negative cases https://proofwiki.org/wiki/GCD_for_Negative_Integers
	if i == 0 || j == 0 {
		return 0
	}
	if i == j {
		return i
	}
	p := i * j
	return p / GCD(i, j)
}

func Lcm(i ...int) int {
	//TODO: needs work on maybe ordering of calculation based on size, not sure
	//TODO: some calculations are maybe implicitly repeated
	if len(i) == 1 {
		return i[0]
	}
	return LCM(i[0], Lcm(i[1:]...))
}

func Mod(i, j int) int {
	return i % j
}

func NChooseK(n, k int) int {
	if n < k {
		return 0
	}
	if k == 0 {
		return 1
	}
	if n < 21 {
		return nChooseK(n, k)
	}
	diff := n - k
	if diff > k {
		return divisionBySimplification(diff+1, n, k)
	}
	return divisionBySimplification(k+1, n, diff)
}

func nChooseK(n int, k int) int {
	if k == 0 {
		return 1
	}
	diff := n - k
	if diff > k {
		return sequentialMultiplicationFromTo(diff+1, n) / factorial(k)
	}
	return sequentialMultiplicationFromTo(k+1, n) / factorial(diff)
}

func divisionBySimplification(numeratorFrom, numeratorTo, denominator int) int {
	numerator := make([]int, 0)
	for i := numeratorFrom; i <= numeratorTo; i++ {
		numerator = append(numerator, i)
	}
	denom := make([]int, 0)
	for i := 1; i <= denominator; i++ {
		denom = append(denom, i)
	}
	simplifiedNumerator, _ := simplifyNumeratorDenominatorSequences(numerator, denom)
	return multiply(simplifiedNumerator...)
}

// TODO: overflows
func multiply(ints ...int) int {
	if len(ints) == 0 {
		return 0
	}
	if len(ints) == 1 {
		return ints[0]
	}
	return ints[0] * multiply(ints[1:]...)
}

func simplifyNumeratorDenominatorSequences(numerator, denominator []int) ([]int, []int) {
	numer := make([]int, len(numerator))
	copy(numer, numerator)
	denom := make([]int, len(denominator))
	copy(denom, denominator)
	for i_ := 0; i_ < len(denom); i_++ {
		if denom[i_] == 1 {
			continue
		}
		for j := 0; j < len(numer); j++ {
			if numer[j] == 1 {
				continue
			}
			hcf := HCF(denom[i_], numer[j])
			if hcf == 1 {
				continue
			}
			numer[j] = numer[j] / hcf
			if hcf == denom[i_] {
				denom[i_] = 1
			} else {
				denom[i_] = denom[i_] / hcf
			}
		}
	}
	return numer, denom
}

func sequentialMultiplicationFromTo(from, to int) int {
	if from == to {
		return to
	}
	if from == to-1 {
		return from * to
	}
	return sequentialMultiplicationFromTo(from, to-1) * to
}

func _factorial(n int) int {
	if n > 20 {
		return 0
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 6
	}
	coeff := make([]int, 2, n-1)
	coeff[0], coeff[1] = 1, 1
	evenBuf := make([]int, 0, n-1)
	oddBuf := make([]int, 0, n-1)
	for i, j := 4, 2; i <= n; i, j = i+2, j+1 {
		//EVEN
		evenBuf = evenBuf[:len(coeff)+1]
		copy(evenBuf[1:], coeff)
		for k := 0; k < len(evenBuf); k++ {
			evenBuf[k] *= j
		}
		//ODD
		if i+1 <= n {
			oddBuf = oddBuf[:len(evenBuf)+1]
			copy(oddBuf[1:], evenBuf)
			for k := 0; k < len(oddBuf); k++ {
				oddBuf[k] *= j
			}
			for k := 0; k < len(evenBuf); k++ {
				oddBuf[k] += evenBuf[k]
			}
			copy(coeff[:len(oddBuf)], oddBuf)
			coeff = coeff[:len(oddBuf)]
		} else {
			copy(coeff[:len(evenBuf)], evenBuf)
			coeff = coeff[:len(evenBuf)]
		}
		clear(evenBuf)
		clear(oddBuf)
	}
	return sumPowersOfTwo(coeff)
}

func sumPowersOfTwo(coefficients []int) int {
	sum := 0
	for i := 0; i < len(coefficients); i++ {
		sum += coefficients[i] * BitShiftPowOfTwo(i+1)
	}
	return sum
}

func BitShiftPowOfTwo(n int) int {
	if n == 0 {
		return 1
	}
	return 2 << (n - 1)
}

func PowOfTwo(n int) int {
	if n == 0 {
		return 1
	}
	res := 1
	for i := 0; i < n; i++ {
		res *= 2
	}
	return res
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return sequentialMultiplicationFromTo(1, n)
}
