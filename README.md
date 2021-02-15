# validator
Struct values validator for Golang

### Install
```sh
$ go get -u github.com/alseiitov/validator
```

### Usage
| Flag        | Description                                                                                                                  |
| :---------- | :--------------------------------------------------------------------------------------------------------------------------- |
| `required`  | `string` length must be more than zero and it must contain not only spaces, tabs or newlines<br>`int` value must not be zero |
| `password`  | `string` must contain at least one lowercase, one uppercase, one number and one symbol                                       |
| `email`     | `string` must be a valid email address                                                                                       |
| `username`  | `string` must contain only letters, digits, dots, underscores and dashes                                                     |
| `min=value` | `string` length must be greater than or equal to given value<br>`int` value must be greater than or equal to given value     |
| `max=value` | `string` length must be less than or equal to given value<br>`int` value must be less than or equal to given value           |

If array have `required` flag it must contain at least one element. Other flags works for elements in array, but not for array itself.


### Example
```go
package main

import (
	"fmt"

	"./repos/validator"
)

type user struct {
	Username  string   `validator:"required,username,min=3,max=32"`
	Password  string   `validator:"required,password,min=7,max=32"`
	Email     string   `validator:"required,email"`
	Age       int      `validator:"required,min=18"`
	Interests []string `validator:"required,min=2, max=32"`
}

func main() {
	users := []user{
		user{
			Username:  "johndoe",
			Password:  "Password123@",
			Email:     "john_doe@protonmail.com",
			Age:       18,
			Interests: []string{"IT", "Music"},
		},
		user{
			Username:  "",
			Password:  "Password123@",
			Email:     "john_doe@protonmail.com",
			Age:       18,
			Interests: []string{"IT", "Music"},
		},
		user{
			Username:  "johndoe",
			Password:  "Password@",
			Email:     "john_doe@protonmail.com",
			Age:       18,
			Interests: []string{"IT", "Music"},
		},
		user{
			Username:  "johndoe",
			Password:  "Password123@",
			Email:     "john_doeprotonmail.com",
			Age:       18,
			Interests: []string{"IT", "Music"},
		},
		user{
			Username:  "johndoe",
			Password:  "Password123@",
			Email:     "john_doe@protonmail.com",
			Age:       17,
			Interests: []string{"IT", "Music"},
		},
		user{
			Username:  "johndoe",
			Password:  "Password123@",
			Email:     "john_doe@protonmail.com",
			Age:       18,
			Interests: []string{},
		},
		user{
			Username:  "johndoe",
			Password:  "Password123@",
			Email:     "john_doe@protonmail.com",
			Age:       18,
			Interests: []string{"A"},
		},
	}

	for _, u := range users {
		if err := validator.Validate(u); err != nil {
			fmt.Printf("%s\n\n", err.Error())
		} else {
			fmt.Printf("OK!\n\n")
		}
	}
}

```

```sh
$ go run main.go
OK!

Username is required

Password must contain at least one number

e-mail is invalid

Age value (17) is lower than minimum value (18)

Interests is required

Interests length (1) is lower than minimum length (2)

```
