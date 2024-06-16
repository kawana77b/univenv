package common

func Validate(v ...Validator) error {
	for _, val := range v {
		if err := val.Validate(); err != nil {
			return err
		}
	}
	return nil
}
