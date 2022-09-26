package generateName

import (
	"fmt"
	"strconv"
	"strings"
)

func Generate(name string, number string) string {

	if _, err := strconv.ParseFloat(name, 64); err == nil {
		fmt.Printf("%q looks like a number.\n", name)
	}

	if _, err := strconv.ParseFloat(number, 64); err != nil {
		fmt.Printf("%q not a number.\n", number)
	}

	var nameUPPER string = strings.ToUpper(name)
	new, _ := strconv.ParseFloat(number, 64)
	result := nameUPPER + "-" + fmt.Sprintf("%f", new)
	return result

}
