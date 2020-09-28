func isEqual(str1 string, str2 string) bool {

	var r uint8 = 0

	if len(str1) != len(str2) {
		return false
	}

	for i, _ := range str1 {
		r = r | (str1[i] ^ str2[i])
	}

	return r == 0

}
