package utils

import (
	"fmt"
	"os"
)

func ReadFile(path string) string {
	data, err := os.ReadFile(fmt.Sprintf("../%s", path))

	if err != nil {
		panic(err)
	}

	return string(data)
}
