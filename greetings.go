package greetings

/* to collect related functions. */

import (
	"errors"
	"fmt"
)
/* In Go, a function whose name starts with a capital letter can
be called by a function NOT in the same package.
This is known in Go as an exported name.
*/

func Hello(name string) (string, error) {

	/* In Go, the := operator is a shortcut for declaring and initializing
	 a variable in one line

	 	var message string
		message = fmt.Sprintf("Hi, %v. Welcome!", name)

	 Use the fmt package's Sprintf function to create a greeting message.
	 The first argument is a format string, and Sprintf substitutes
	 the name parameter's value for the %v format verb.
	 Inserting the value of the name parameter completes the greeting text.
	*/
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message, nil
}
