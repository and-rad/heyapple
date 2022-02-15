package app_test

import (
	"heyapple/pkg/app"
	"os"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func Test_crypter_Encrypt(t *testing.T) {
	testdata := []struct {
		password string
	}{
		{"testpassword"},
		{"%/$§kjhJDO18§/&L+><"},
	}

	for idx, data := range testdata {
		c := app.NewCrypter()
		hash := c.Encrypt(data.password)
		if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data.password)); err != nil {
			t.Errorf("test case %d: hash mismatch", idx)
		}
	}
}

func Test_crypter_EncryptPanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("expected panic did not occur")
		}
	}()

	cost := os.Getenv("HEYAPPLE_ENCRYPT_COST")
	defer os.Setenv("HEYAPPLE_ENCRYPT_COST", cost)

	// cost is above bcrypt's max cost
	os.Setenv("HEYAPPLE_ENCRYPT_COST", "64")
	c := app.NewCrypter()
	c.Encrypt("Password")
}

func Test_crypter_Match(t *testing.T) {
	testdata := []struct {
		pass string
		hash string
		ok   bool
	}{
		{ //00//
			pass: "password123",
			hash: "$2a$10$ADm2JBRbt8UvB0uI7NNFBupOdTq7XKae6Dvc7NfVCnw89rPZr3.zK",
			ok:   true,
		},
		{ //00//
			pass: "Tr0ub4dor&3",
			hash: "$2a$10$CpVy94BcePvhBH3QS/mMnOtFVrfN0DvwdooEUc0T8tWdKNi3ayFXC",
			ok:   true,
		},
		{ //00//
			pass: "Tr0ub4dor&3",
			hash: "$2a$10$ADm2JBRbt8UvB0uI7NNFBupOdTq7XKae6Dvc7NfVCnw89rPZr3.zK",
			ok:   false,
		},
	}

	for idx, data := range testdata {
		c := app.NewCrypter()

		if ok := c.Match(data.hash, data.pass); ok != data.ok {
			t.Errorf("test case %d: unexpected match result \nhave: %v \nwant: %v", idx, ok, data.ok)
		}
	}
}
