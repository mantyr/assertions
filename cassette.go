package assertions

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/smartystreets/assertions"
	yaml "gopkg.in/yaml.v3"
)

func ShouldEqualYamlCassette(
	actual interface{},
	expected ...interface{},
) string {
	if actual == nil {
		return "This assertion requires the actual value is not nil"
	}
	if fail := need(1, expected); fail != success {
		return fail
	}
	address, ok := expected[0].(string)
	if !ok {
		return "This assertion requires the expected value is string"
	}
	buf := bytes.NewBuffer([]byte{})
	enc := yaml.NewEncoder(buf)
	err := enc.Encode(actual)
	if err != nil {
		return err.Error()
	}
	data, err := os.ReadFile(address)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err := os.WriteFile(address, buf.Bytes(), 0666)
			if err != nil {
				return err.Error()
			}
			return fmt.Sprintf(
				`A cassette with the name "%s" was not found and was created with the current value.`,
				expected[0],
			)
		}
		return err.Error()
	}
	return assertions.ShouldEqual(buf.String(), string(data))
}

func ShouldNotEqualYamlCassette(
	actual interface{},
	expected ...interface{},
) string {
	if actual == nil {
		return "This assertion requires the actual value is not nil"
	}
	if fail := need(1, expected); fail != success {
		return fail
	}
	address, ok := expected[0].(string)
	if !ok {
		return "This assertion requires the expected value is string"
	}
	buf := bytes.NewBuffer([]byte{})
	enc := yaml.NewEncoder(buf)
	err := enc.Encode(actual)
	if err != nil {
		return err.Error()
	}
	data, err := os.ReadFile(address)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Sprintf(
				`A cassette with the name "%s" was not found.`,
				expected[0],
			)
		}
		return err.Error()
	}
	return assertions.ShouldNotEqual(buf.String(), string(data))
}
