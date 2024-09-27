package math

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
