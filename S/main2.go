package S

type User struct {
	FirstName string
	LastName  string
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

type UserRepository struct {
	// Database connection or other storage-related fields
}

func (r *UserRepository) Save(u *User) error {
	// Save user to the database
	// ...
	return nil
}

//Now, the User struct is only responsible for managing user data, while the UserRepository handles database operations
