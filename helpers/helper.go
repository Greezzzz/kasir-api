package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateID() string {
	now := time.Now().UnixMilli()
	rand := rand.Intn(1000)
	return fmt.Sprintf("%d%03d", now, rand)
}
