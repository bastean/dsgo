package user

type Verified struct {
	Value bool
}

func NewVerified(value bool) (*Verified, error) {
	valueObj := &Verified{
		Value: value,
	}

	return valueObj, nil
}
