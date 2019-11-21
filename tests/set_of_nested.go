package tests

import (
	"io"

	. "github.com/smartystreets/goconvey/convey"
)

// SetOfNested это набор вложенных тестов
type SetOfNested struct {
	i     int
	tests []NestedTest
}

// NewSetOfNested возвращает набор вложенных тестов
func NewSetOfNested() *SetOfNested {
	return &SetOfNested{}
}

// Add добавляет тест в набор
func (t *SetOfNested) Add(test NestedTest) *SetOfNested {
	t.tests = append(t.tests, test)
	return t
}

// Next возвращает следующий тест
func (t *SetOfNested) Next() (NestedTest, error) {
	if t.i >= len(t.tests) {
		return nil, io.EOF
	}
	test := t.tests[t.i]
	t.i++
	return test, nil
}

// Run запускает набор вложенных тестов
func (t *SetOfNested) Run(c C) {
	test, err := t.Next()
	if err == io.EOF {
		return
	}
	c.So(test, ShouldNotBeNil)

	Convey(test.Name(), func() {
		test.Run(c)
		t.Run(c)
	})
}
