package manager

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	ENV_STORAGE_FILE = "STORAGE_FILE"
)

var singleton *PropertyManager = nil

type PropertyManager struct {
	filename string
}

func NewPropertyManager() *PropertyManager {
	return &PropertyManager{}
}

func GetPropertyManagerSingleton() (*PropertyManager, error) {
	if singleton == nil {
		singleton = NewPropertyManager()
		err := singleton.Init()
		if err != nil {
			return nil, err
		}
	}
	return singleton, nil
}

func (sm *PropertyManager) Init() error {
	sm.filename = os.Getenv(ENV_STORAGE_FILE)
	if sm.filename == "" {
		return errors.New("storage filename not found")
	}

	// check if exists
	info, err := os.Stat(sm.filename)
	if err != nil {
		return err
	}

	// check if writable
	mode := info.Mode()
	if mode&os.ModePerm == os.ModePerm {
		return errors.New("configured storage file is not writable")
	}

	return nil
}

func (sm *PropertyManager) GetProperty(key string) (string, error) {
	valueMap, err := getValueMap(sm.filename)
	if err != nil {
		return "", err
	}

	value := valueMap[key]
	return value, nil
}

// returns previous value & error
func (sm *PropertyManager) SetProperty(key string, value string) (string, error) {
	valueMap, err := getValueMap(sm.filename)
	if err != nil {
		return "", err
	}
	prev := valueMap[key]

	// write new value
	valueMap[key] = value

	// convert map to string: a=b\n
	b := new(bytes.Buffer)
	for key, value := range valueMap {
		fmt.Fprintf(b, "%s=%s\n", key, value)
	}

	err = os.WriteFile(sm.filename, b.Bytes(), 0666)
	if err != nil {
		return prev, err
	}

	return prev, nil
}

func getValueMap(filename string) (map[string]string, error) {
	valueMap := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return valueMap, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		txt := scanner.Text()
		before, after, found := strings.Cut(txt, "=")
		if found {
			valueMap[before] = after
		}
	}
	if err := scanner.Err(); err != nil {
		return valueMap, err
	}

	return valueMap, nil
}
