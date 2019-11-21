package tests

import (
	. "github.com/smartystreets/goconvey/convey"
)

// NestedTest это интерфейс вложенного теста
type NestedTest interface {
	// Name возвращает название вложенного теста
	Name() string

	// Run запускает вложенный тест
	Run(c C)
}

// nestedTest это вложенный тест
type nestedTest struct {
	name string
	f    func(c C)
}

// NewNestedTest возвращает новый вложенный тест
func NewNestedTest(name string, f func(c C)) NestedTest {
	return &nestedTest{
		name: name,
		f:    f,
	}
}

// Name возвращает название вложенного теста
func (t *nestedTest) Name() string {
	return t.name
}

// Run запускает вложенный тест
func (t *nestedTest) Run(c C) {
	c.So(t.f, ShouldNotBeNil)
	t.f(c)
}
