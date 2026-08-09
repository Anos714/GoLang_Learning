//

package greet

import (
	"fmt"
	"strings"
)

// exported functions starts with capital letter
// other pkg can call this function
func Greet(name string) string {
	greet := normalizeName(name)
	return greet
}

// this function cannot be called from outside this package or cannot exported
// other pkg cannot call this function cause it starts with lowercase letter
func normalizeName(name string)string{
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Sprintf("Hello, Guest!")
	}
	return fmt.Sprintf("Hello, %s!", strings.ToTitle(name))
}
