package dataservices

func (ms *Client) GetByID(i int) (*User, *Error) {
	for _, value := range Data {
		if value.Id == i {
			return &value, nil
		}
	}
	err := &Error{i, "User does not exist"}
	return nil, err
}
