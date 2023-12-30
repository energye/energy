//go:build darwin

package bar

import (
	"crypto/rand"
	"math/big"
)

var spinValues = []string{"🍒", "💎", "7️⃣", "🍊", "🔔", "⭐", "🍇", "🍀"}

func randomChoice(choices []string) (*string, error) {
	v, err := rand.Int(rand.Reader, big.NewInt(int64(len(choices))))
	if err != nil {
		return nil, err
	}
	return &choices[int(v.Uint64())], nil
}

func getRandomValue() (*string, error) {
	return randomChoice(spinValues)
}
