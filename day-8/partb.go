package main

func isStartingNode(s string) bool {
	bytes := []byte(s)
	if len(bytes) == 0 {
		return false
	}
	return byteSliceEndsInA(bytes)
}

func byteSliceEndsInA(bs []byte) bool {
	return 'A' == bs[len(bs)-1]
}

func stringEndsInZ(s string) bool {
	return 'Z' == s[len(s)-1]
}
