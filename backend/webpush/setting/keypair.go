package setting

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type keypair struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

func GetKeypair() (keypair, error) {
	f, err := os.Open("keypair.json")
	if err != nil {
		return keypair{}, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return keypair{}, err
	}

	var keys keypair
	err = json.Unmarshal(b, &keys)

	return keys, err
}
