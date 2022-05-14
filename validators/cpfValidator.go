package validators

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"gopkg.in/validator.v2"
)

func divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return quotient, remainder
}

func getDigitsAccumulator(stringDigits string) int {
	accumulator := 0
	multiplier := 2
	for i := 1; i <= len(stringDigits); i++ {
		digit, _ := strconv.Atoi(string(stringDigits[len(stringDigits)-i]))
		accumulator = (digit * multiplier) + accumulator
		multiplier += 1
	}
	return accumulator
}

func digitValidation(accumulator int, digit int) bool {
	_, remainder := divmod(int64(accumulator), 11)

	if remainder < 2 && digit != 0 {
		return false
	}
	if remainder >= 2 && int64(digit) != (11-remainder) {
		return false
	}
	return true
}

func isCPF(cpf string) bool {
	firstDigit, _ := strconv.Atoi(cpf[9:10])
	secondDigit, _ := strconv.Atoi(cpf[10:11])

	firstNineDigits := cpf[0:9]
	accumulatorFirstNineDigits := getDigitsAccumulator(firstNineDigits)

	isValidDigit := digitValidation(accumulatorFirstNineDigits, firstDigit)

	if !isValidDigit {
		return false
	}

	firstTenDigits := cpf[0:10]
	accumulatorFirstTenDigits := getDigitsAccumulator(firstTenDigits)

	return digitValidation(accumulatorFirstTenDigits, secondDigit)
}

func isIdentityNumberCustomValidator(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return errors.New("is not string")
	}
	if st.Len() != 11 {
		return errors.New("invalid length")
	}
	if !isCPF(fmt.Sprintf("%v", v)) {
		return errors.New("invalid")
	}
	return nil
}

func LoadValidators() {
	validator.SetValidationFunc("iscpf", isIdentityNumberCustomValidator)
}
