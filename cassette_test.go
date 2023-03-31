package assertions

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type shouldEqualYamlCassette struct {
	Name   string
	Value  string
	Values []string
}

func TestShouldEqualYamlCasette(t *testing.T) {
	Convey("Checkgin ShouldEqualYamlCasette", t, func() {
		s := shouldEqualYamlCassette{
			Name:  "ShouldEqualYamlCassette",
			Value: "ShouldEqualYamlCassette",
			Values: []string{
				"ShouldEqualYamlCassette1",
				"ShouldEqualYamlCassette2",
			},
		}
		Convey("Equal true", func() {
			So(s, ShouldEqualYamlCassette, "./testdata/shouldEqualYamlCassette.1.cassette.yaml")
		})
		Convey("Equal false", func() {
			So(
				ShouldEqualYamlCassette(s, "./testdata/shouldEqualYamlCassette.2.cassette.yaml"),
				ShouldNotEqual,
				"",
			)
		})
	})
}

func TestShouldNotEqualYamlCassette(t *testing.T) {
	Convey("Checking ShouldNotEqualYamlCassette", t, func() {
		s := shouldEqualYamlCassette{
			Name:  "ShouldEqualYamlCassette",
			Value: "ShouldEqualYamlCassette",
			Values: []string{
				"ShouldEqualYamlCassette1",
				"ShouldEqualYamlCassette2",
			},
		}
		So(s, ShouldNotEqualYamlCassette, "./testdata/shouldEqualYamlCassette.2.cassette.yaml")
	})
}
