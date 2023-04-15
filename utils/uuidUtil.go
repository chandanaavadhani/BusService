package utils

import (
	"strings"

	"github.com/pborman/uuid"
)

func UUID() string {
	uuidWithHyphen := uuid.NewRandom()
	// fmt.Println(uuidWithHyphen)
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	// fmt.Println("Your unique ID is : %s", uuid)
	return uuid
}
