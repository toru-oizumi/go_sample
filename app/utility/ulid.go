package utility

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GetUlid() string {
	t := time.Now()

	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)

	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

//Validation ulid string check size
func ValidateUlid(id string) error {
	if len(id) != ulid.EncodedSize {
		return fmt.Errorf("unique ID generator validation error: length is not match")
	}
	return nil
}
