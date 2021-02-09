package validator_test

import (
	"testing"

	"github.com/alseiitov/validator"
)

type user struct {
	Username string `validator:"required,username,min=3,max=32"`
	Password string `validator:"required,password,min=7,max=32"`
	Email    string `validator:"required,email"`
	Age      int    `validator:"required,min=18"`
}

type testCase struct {
	User     user
	Expected string
}

func TestValidate(t *testing.T) {
	testCases := []testCase{
		testCase{
			User: user{
				Username: "johndoe",
				Password: "Password123@",
				Email:    "john_doe@protonmail.com",
				Age:      18,
			},
			Expected: "",
		},
		testCase{
			User: user{
				Username: "j",
				Password: "Password123@",
				Email:    "john_doe@protonmail.com",
				Age:      18,
			},
			Expected: "Username length (1) is lower than minimum length (3)",
		},
		testCase{
			User: user{
				Username: "jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj",
				Password: "Password123@",
				Email:    "john_doe@protonmail.com",
				Age:      18,
			},
			Expected: "Username length (33) length is higher than maximim length (32)",
		},
		testCase{
			User: user{
				Username: "",
				Password: "Password123@",
				Email:    "john_doe@protonmail.com",
				Age:      18,
			},
			Expected: "Username is required",
		},
		testCase{
			User: user{
				Username: "johndoe",
				Password: "Password123",
				Email:    "john_doe@protonmail.com",
				Age:      18,
			},
			Expected: "Password must contain at least one symbol\n(!, @, #, ~, $, %, ^, &, *, (, ), +, |, _, )",
		},
		testCase{
			User: user{
				Username: "johndoe",
				Password: "Password@",
				Email:    "john_doe@protonmail.com",
				Age:      18,
			},
			Expected: "Password must contain at least one number",
		},
		testCase{
			User: user{
				Username: "johndoe",
				Password: "password123@",
				Email:    "john_doe@protonmail.com",
				Age:      18,
			},
			Expected: "Password must contain at least one uppercase letter",
		},
		testCase{
			User: user{
				Username: "johndoe",
				Password: "PASSWORD123@",
				Email:    "john_doe@protonmail.com",
				Age:      18,
			},
			Expected: "Password must contain at least one lowercase letter",
		},
		testCase{
			User: user{
				Username: "johndoe",
				Password: "Password123@",
				Email:    "john_doeprotonmail.com",
				Age:      18,
			},
			Expected: "e-mail is invalid",
		},
		testCase{
			User: user{
				Username: "johndoe",
				Password: "Password123@",
				Email:    "john_doe@protonmailcom",
				Age:      18,
			},
			Expected: "e-mail is invalid",
		},
		testCase{
			User: user{
				Username: "johndoe",
				Password: "Password123@",
				Email:    "john_doe@protonmail.c",
				Age:      18,
			},
			Expected: "e-mail is invalid",
		},
		testCase{
			User: user{
				Username: "johndoe",
				Password: "Password123@",
				Email:    "john_doe@protonmail.com",
				Age:      0,
			},
			Expected: "Age is required",
		},
		testCase{
			User: user{
				Username: "johndoe",
				Password: "Password123@",
				Email:    "john_doe@protonmail.com",
				Age:      17,
			},
			Expected: "Age value (17) is lower than minimum value (18)",
		},
	}

	for _, test := range testCases {
		err := validator.Validate(test.User)
		if err != nil {
			if test.Expected == "" {
				t.Log(test.User)
				t.Logf("\nDidn't expected error, but got:	%s", err.Error())
				t.Fail()
			}

			if err.Error() != test.Expected {
				t.Log(test.User)
				t.Logf("\nExpected:	%s\nbut got:	%s", test.Expected, err.Error())
				t.Fail()
			}
		} else {
			if test.Expected != "" {
				t.Log(test.User)
				t.Logf("\nTest passes, but expected error:	%s", test.Expected)
				t.Fail()
			}
		}
	}
}
