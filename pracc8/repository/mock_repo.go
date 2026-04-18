package repository

type MockRepo struct {
	User *User
	Err  error
}

func (m *MockRepo) GetUserByID(id int) (*User, error) {
	return m.User, m.Err
}

func (m *MockRepo) CreateUser(user *User) error {
	return m.Err
}

func (m *MockRepo) GetByEmail(email string) (*User, error) {
	return m.User, m.Err
}

func (m *MockRepo) UpdateUser(user *User) error {
	return m.Err
}

func (m *MockRepo) DeleteUser(id int) error {
	return m.Err
}
