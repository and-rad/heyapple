package app

// User represents a human user of the system, identified
// by their e-mail address or id. The password is always
// stored as an encrypted hash.
type User struct {
	Email string
	Pass  string
	ID    int
}

// Authenticate is a query that authenticates a user by
// checking against e-mail and password. If successful,
// the user's id is stored in the command.
type Authenticate struct {
	Email string
	Pass  string
	ID    int
}

func (c *Authenticate) Fetch(db DB) error {
	if user, err := db.UserByName(c.Email); err != nil {
		return err
	} else if !NewCrypter().Match(user.Pass, c.Pass) {
		return ErrCredentials
	} else {
		c.ID = user.ID
	}

	return nil
}
