package order

import (
	"fmt"

	"github.com/google/uuid"
)

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	fmt.Printf("error in validating uuid : %v",err)
	return err == nil
}