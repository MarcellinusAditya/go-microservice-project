package error

func ErrMapping(err error) bool {
	AllErrors := make([]error, 0)
	AllErrors = append(append(GeneralErrors[:], UserErrors[:]...))

	for _, item := range AllErrors {
		if err.Error() == item.Error() {
			return true
		}
	}
	return false
}