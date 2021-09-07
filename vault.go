package secret

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/saltperfect/exercise/secret/cipher"
)

func NewVault(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

type Vault struct {
	encodingKey string
	filepath    string
	keyValues   map[string]string
	mutex       sync.Mutex
}

func (v *Vault) load() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()
	reader, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}
	err = v.readKeyValues(reader)
	return err
}

func (v *Vault) readKeyValues(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&v.keyValues)
	return err
}

func (v *Vault) save() error {

	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	writer, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	err = v.writeKeyValues(writer)
	return err
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(v.keyValues)
	return err
}

func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return "", err
	}
	ret, ok := v.keyValues[key]
	if !ok {
		return "", fmt.Errorf("no key found")
	}
	return ret, nil
}

func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	return v.save()
}
