package user_request

import "strconv"

func ValidateCpf(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}
	var checkFirstDigit int
	var checkSecondDigit int
	var nums1 int
	var nums2 int
	var diferent bool
	var prevDigit int
	nums1 = 10
	nums2 = 11
	for i, v := range cpf {
		num, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
		if prevDigit != num && i != 0 {
			diferent = true
		}
		prevDigit = num

		if i <= 8 {
			checkFirstDigit += (num * nums1)

			nums1 -= 1
		}
		if i <= 9 {
			checkSecondDigit += (num * (nums2))
			nums2 -= 1
		}
	}
	if !diferent {
		return false
	}
	res := (checkFirstDigit * 10) % 11
	firstDigit, err := strconv.Atoi(string(cpf[9]))
	if err != nil {
		return false
	}
	if res == 10 {
		res = 0
	}
	if res != firstDigit {
		return false
	}
	secondDigit, err := strconv.Atoi(string(cpf[10]))
	if err != nil {
		return false
	}
	res = (checkSecondDigit * 10) % 11
	if res == 10 {
		res = 0
	}
	if res != secondDigit {
		return false
	}
	return true
}
