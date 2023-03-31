package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testSetOfNested struct {
	state []string
}

func (t *testSetOfNested) First(c C) {
	t.state = append(t.state, "first")
}

func (t *testSetOfNested) Second(c C) {
	t.state = append(t.state, "second")
}

func (t *testSetOfNested) Third(c C) {
	t.state = append(t.state, "third")
}

func TestSetOfNested(t *testing.T) {
	Convey("Проверяем SetOfNested - набор вложенных тестов", t, func() {
		tests := NewSetOfNested()
		So(tests, ShouldNotBeNil)

		Convey("Проверяем что тесты выполняются последовательно", func(c C) {
			test := &testSetOfNested{}
			tests.Add(NewNestedTest("first", test.First))
			tests.Add(NewNestedTest("second", test.Second))
			tests.Add(NewNestedTest("third", test.Third))
			tests.Run(c)
			So(
				test.state,
				ShouldResemble,
				[]string{
					"first",
					"second",
					"third",
				},
			)
		})
	})
}
