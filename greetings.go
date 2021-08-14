package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Hello(names []string) ( map[string]string, error) {
    // If no name was given, return an error with a message.
    messages := make(map[string]string)

    for _, name := range names {
   	if name == "" {
		return nil, errors.New("empty name")
    	}
    	// Create a message using a random format.
   	message := fmt.Sprintf(randomFormat(), name)
    	messages[name] = message
    }

    return messages, nil

}
	// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
