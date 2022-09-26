package generateName

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Generate(name string, number string) (string, error) {

	if _, err := strconv.ParseFloat(name, 64); err == nil {
		return "", errors.New("look like a number")
	}

	if _, err := strconv.ParseFloat(number, 64); err != nil {
		return "", errors.New("look like a number")
	}

	var nameUPPER string = strings.ToUpper(name)
	new, _ := strconv.ParseFloat(number, 64)
	result := nameUPPER + "-" + fmt.Sprintf("%f", new)
	return result, nil

}
