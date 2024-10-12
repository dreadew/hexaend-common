package utils

import (
	"fmt"
	"math/rand/v2"
	"net/mail"
	"strings"

	"github.com/google/uuid"
)

func CheckEmail(str string) bool {
	_, err := mail.ParseAddress(str)

	return err == nil
}

func GenerateUuidOneSeq() string {
	return strings.Split(uuid.NewString(), "-")[0]
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func GenerateSimilarUsernames(baseUsername string, count, minRange, maxRange int) []string {
	suggestions := make([]string, count)

	for i := 0; i < count; i++ {
		generatedInt := randRange(minRange, maxRange)
		suggestions[i] = fmt.Sprintf("%s%d", baseUsername, generatedInt)
	}
	return suggestions
}

func Difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
