package utils

var digitToByteMapping = []byte("0123456789")

type CalibrationLine struct {
	value     []byte
	digitList *DigitList
}

func NewCalibrationLine(byteLine []byte) CalibrationLine {
	digitList := newDigitList()
	return CalibrationLine{byteLine, digitList}
}

func (cl *CalibrationLine) Number() float64 {
	cl.identifyDigitsInCalLine()
	return cl.digitList.number()
}

func (cl *CalibrationLine) identifyDigitsInCalLine() {
	var digitsFound []int
	for _, char := range cl.value {
		for key, b := range digitToByteMapping {
			if b == char {
				digitsFound = append(digitsFound, key)
			}
		}
	}
	numberOfDigitsFound := len(digitsFound)
	if numberOfDigitsFound == 0 {
		return
	}
	if numberOfDigitsFound == 1 {
		cl.digitList.appendDigitN(digitsFound[0], 2)
		return
	}
	if numberOfDigitsFound > 2 {
		digitsToAppend := []int{digitsFound[0], digitsFound[numberOfDigitsFound-1]}
		cl.digitList.appendDigits(digitsToAppend)
		return
	}
	cl.digitList.appendDigits(digitsFound)
}
