package utils

import (
	"fmt"
	"regexp"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

func validateName(name string) bool {
	// use the regexp to check the length between 2-32 and all letters only

	r := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{2,32}$`)

	if r.MatchString(name) {
		return true
	}
	return false
}

// validate the email
func validateEmail(email string) bool {
	// check for the email
	r := regexp.MustCompile(`^[a-zA-Z0-9!#$%^&*+.]+@[a-z]+\.(com.np|com|org)$`)

	if r.MatchString(email) {
		return true
	}
	return false
}

// validate for the password
func validatePassword(password string) bool {
	// check for the email
	format := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*+.]{8,32}$`)

	var capitalLetters = regexp.MustCompile(`[A-Z]`)
	var smallLetters = regexp.MustCompile(`[a-z]`)
	var numbers = regexp.MustCompile(`[0-9]`)
	var specialCharacters = regexp.MustCompile(`[!@#$%^&*+]`)

	fmt.Println(capitalLetters.MatchString(password))
	fmt.Println(smallLetters.MatchString(password))
	fmt.Println(numbers.MatchString(password))
	fmt.Println(format.MatchString(password))
	fmt.Println(specialCharacters.MatchString(password))

	// all data must be satisfied
	if format.MatchString(password) && capitalLetters.MatchString(password) && smallLetters.MatchString(password) && numbers.MatchString(password) && specialCharacters.MatchString(password) {
		return true
	}
	return false
}

// validate the address
func validateAddress(address string) bool {
	var validAddress = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9,._-]+( [a-zA-Z0-9,.-_]+)*$`)
	validated := false
	if validAddress.MatchString(address) && (len(address) < 100 && len(address) > 10) {
		validated = true
	}
	return validated
}

func ValidateUserRegister(credentials *models.Register) (bool, string){
	// all of the above data must be checked properly

	checkFirstName := validateName(credentials.FirstName)
	checkLastName := validateName(credentials.LastName)
	checkAddress := validateAddress(credentials.Address)
	checkPassword := validatePassword(credentials.Password)
	checkEmail := validateEmail(credentials.Email)


	// if there is error in name
	if !checkFirstName  &&  !checkLastName{
		// error in the name
		return false, "First and last name must be between 2-32"
	}

	// if there is error in the address
	if !checkAddress{
		// 
		return false, "Address must be in proper format."
	}

	// if there is error in the email
	if !checkEmail{
		//
		return false, "Error in the mail."
	}

	// if there is error in the password
	if !checkPassword{
		return false, "Password must be in proper format."
	}



	// if its correct at the last then
	return true, ""


}

func ValidateUserLogin(credentials *models.Login)  (bool, string){

	checkPassword := validatePassword(credentials.Password)
	checkEmail := validateEmail(credentials.Email)

	// if error in the mail
	if !checkEmail{
		return false, "Error in the mail."
	}

	// if there is error in the password
	if !checkPassword{
		return false, "Password must be in proper format."
	}

	// if all is correct
	return true, ""
}