package key

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
)

// FromFile loads a private key from the provided path and parses it.
// The private key is returned if parsing succeeds.
func FromFile(path string) (*rsa.PrivateKey, error) {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %s", err.Error())
	}

	return Parse(key)
}
