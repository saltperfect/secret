package secret

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/saltperfect/exercise/secret/encrypt"
)

func NewVault(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath: filepath,
		keyValues: make(map[string]string),
	}
}

type Vault struct {
	encodingKey string
	filepath string
	keyValues map[string]string
	mutex sync.Mutex
}

func (v *Vault) loadKeyValues() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()
	var sb strings.Builder
	_, err = io.Copy(&sb, f)
	if err != nil {
		return err
	}
	decryptedJSON, err := encrypt.Decrypt(v.encodingKey, sb.String())
	if err != nil {
		return err
	}
	r := strings.NewReader(decryptedJSON)
	decoder:= json.NewDecoder(r)

	err = decoder.Decode(&v.keyValues)
	if err != nil {
		return err
	}
	return nil
}

func (v *Vault) saveKeyValue() error {
	var sb strings.Builder
	encoder := json.NewEncoder(&sb)
	err := encoder.Encode(v.keyValues)
	if err != nil {
		return err
	}
	encryptedJSON, err := encrypt.Encrypt(v.encodingKey, sb.String())
	if err != nil {
		return err
	}

	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprint(f, encryptedJSON)
	if err != nil {
		return err
	}
	return nil
}
func (v *Vault) Get(key string) (string, error) {
	v.loadKeyValues()
	ret, ok := v.keyValues[key]
	if !ok {
		return "", fmt.Errorf("no key found")
	}
	return ret, nil
}

func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.loadKeyValues()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	err = v.saveKeyValue()
	if err != nil {
		return err
	}
	return nil
}