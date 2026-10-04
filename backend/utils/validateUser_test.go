package utils

import (
	"testing"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

func TestValidateUserRegister(t *testing.T){
	righttests := []models.Register{
		{
			FirstName: "Milan",
			LastName: "Bist",
			Address: "Mahakali -07 Dhap, Darchula",
			Email: "bistmilan46@gmail.com",
			Password: "Milbis123@#",
		},
		{
			FirstName: "John",
			LastName: "Doe",
			Address: "7th strret, LA",
			Email: "johndoe@gmail.com",
			Password: "JohnDoe!12",
		},
	}

	wrongTests := []struct {
    name string
    user models.Register
}{
    {
        name: "weak password",
        user: models.Register{
            FirstName: "Milan",
            LastName:  "Bist",
            Address:   "Mahakali -07 Dhap, Darchula",
            Email:     "bistmilan46@gmail.com",
            Password:  "milan123",
        },
    },
    {
        name: "empty first name",
        user: models.Register{
            FirstName: "",
            LastName:  "Doe",
            Address:   "7th street, LA",
            Email:     "johndoe@gmail.com",
            Password:  "JohnDoe!12",
        },
    },
    {
        name: "empty last name",
        user: models.Register{
            FirstName: "John",
            LastName:  "",
            Address:   "7th street, LA",
            Email:     "johndoe@gmail.com",
            Password:  "JohnDoe!12",
        },
    },
	{
		name: "not proper email",
        user: models.Register{
            FirstName: "John",
            LastName:  "Doe",
            Address:   "7th street, LA",
            Email:     "johndoe@gmail",
            Password:  "JohnDoe!12",
        },
	},
	{
		name: "not proper email",
        user: models.Register{
            FirstName: "John",
            LastName:  "Doe",
            Address:   "",
            Email:     "johndoe@gmail.com",
            Password:  "JohnDoe!12",
        },
	},
}

	// loop through the user data and test it seperately
	for _, value := range righttests{
		t.Run(value.FirstName,  func(t *testing.T) {
			valid, msg := ValidateUserRegister(value)
			if valid == false{
				t.Fatalf("Got an error as: %s", msg)
			}
		})
	}

	// loop through the wrong tests
	for _, value := range wrongTests{
		t.Run(value.name, func(t *testing.T) {
			valid, _ := ValidateUserRegister(value.user)
			if valid != false{
				t.Fatalf("%v failed.", value.name)
			}
		})
	}
}



func TestValidateUserLogin(t *testing.T){
	righttests := []models.Login{
		{
			Email: "bistmilan46@gmail.com",
			Password: "Milbis123@#",
		},
		{
			Email: "johndoe@gmail.com",
			Password: "JohnDoe!12",
		},
	}

	wrongTests := []struct {
    name string
    user models.Login
	}{
	{
		name: "not proper email",
        user: models.Login{
            Email:     "johndoe@gmail",
            Password:  "JohnDoe!12",
        },
	},
	{
		name: "not proper password",
        user: models.Login{
            Email:     "johndoe@gmail.com",
            Password:  "JohnDoe",
        },
	},
	}

	// loop through the user data and test it seperately
	for _, value := range righttests{
		t.Run(value.Email,  func(t *testing.T) {
			valid, msg := ValidateUserLogin(&value)
			if valid == false{
				t.Fatalf("Got an error as: %s", msg)
			}
		})
	}

	// loop through the wrong tests
	for _, value := range wrongTests{
		t.Run(value.name, func(t *testing.T) {
			valid, _ := ValidateUserLogin(&value.user)
			if valid != false{
				t.Fatalf("%v failed.", value.name)
			}
		})
	}
}