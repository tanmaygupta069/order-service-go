package order

import (
	"fmt"

	"github.com/google/uuid"
)

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	if err!=nil{
		fmt.Print(err.Error())
	}
	return err == nil
}