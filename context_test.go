package assertions

import (
	"context"
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShouldClosedControl(t *testing.T) {
	Convey("Контрольная проверка ShouldNotBeClosedBefore и ShouldBeClosedBefore через convey.C.So", t, func(c C) {
		ctx, cancel := context.WithCancel(context.Background())
		ch1, ch2, ch3 := make(chan struct{}), make(chan struct{}), make(chan struct{})
		go func() {
			time.Sleep(2 * time.Second)
			cancel()
			ch1 <- struct{}{}
		}()
		go func() {
			c.So(ctx, ShouldNotBeClosedBefore, time.Second)
			ch2 <- struct{}{}
		}()
		go func() {
			c.So(ctx, ShouldBeClosedBefore, 3*time.Second)
			ch3 <- struct{}{}
		}()
		<-ch1
		<-ch2
		<-ch3
	})
}

func TestShouldBeClosedBefore(t *testing.T) {
	Convey("Проверяем ShouldBeClosedBefore", t, func() {
		ctx, cancel := context.WithCancel(context.Background())
		Convey("Контекст закрыт", func() {
			cancel()
			So(
				ShouldBeClosedBefore(ctx, time.Second),
				ShouldEqual,
				"",
			)
		})
		Convey("Контекст открыт", func() {
			So(
				ShouldBeClosedBefore(ctx, time.Second),
				ShouldEqual,
				fmt.Sprintf(shouldClosedBefore, "context.Background.WithCancel", time.Second),
			)
		})
		Convey("Контекст не заполнен", func() {
			var context context.Context
			So(
				ShouldBeClosedBefore(context, time.Second),
				ShouldEqual,
				shouldUseContext,
			)
			So(
				ShouldBeClosedBefore(nil, time.Second),
				ShouldEqual,
				shouldUseContext,
			)
		})
		Convey("Не корректно заполненный аргумент", func() {
			So(
				ShouldBeClosedBefore(ctx, 123),
				ShouldEqual,
				shouldUseDuration,
			)
		})
		Convey("Контекст закрыт в указанный период", func(c C) {
			ch1, ch2, ch3 := make(chan struct{}), make(chan struct{}), make(chan struct{})
			go func() {
				time.Sleep(2 * time.Second)
				cancel()
				ch1 <- struct{}{}
			}()
			go func() {
				c.So(
					ShouldBeClosedBefore(ctx, time.Second),
					ShouldEqual,
					fmt.Sprintf(shouldClosedBefore, "context.Background.WithCancel", time.Second),
				)
				ch2 <- struct{}{}
			}()
			go func() {
				c.So(
					ShouldBeClosedBefore(ctx, 3*time.Second),
					ShouldEqual,
					"",
				)
				ch3 <- struct{}{}
			}()
			<-ch1
			<-ch2
			<-ch3
		})
	})
}

func TestShouldNotBeClosedBefore(t *testing.T) {
	Convey("Проверяем ShouldNotBeClosedBefore", t, func() {
		ctx, cancel := context.WithCancel(context.Background())
		Convey("Контекст закрыт", func() {
			cancel()
			So(
				ShouldNotBeClosedBefore(ctx, time.Second),
				ShouldEqual,
				fmt.Sprintf(shouldNotClosedBefore, "context.Background.WithCancel", time.Second),
			)
		})
		Convey("Контекст открыт", func() {
			So(
				ShouldNotBeClosedBefore(ctx, time.Second),
				ShouldEqual,
				"",
			)
		})
		Convey("Контекст не заполнен", func() {
			var context context.Context
			So(
				ShouldNotBeClosedBefore(context, time.Second),
				ShouldEqual,
				shouldUseContext,
			)
			So(
				ShouldNotBeClosedBefore(nil, time.Second),
				ShouldEqual,
				shouldUseContext,
			)
		})
		Convey("Не корректно заполненный аргумент", func() {
			So(
				ShouldNotBeClosedBefore(ctx, 123),
				ShouldEqual,
				shouldUseDuration,
			)
		})
		Convey("Контекст закрыт в указанный период", func(c C) {
			ch1, ch2, ch3 := make(chan struct{}), make(chan struct{}), make(chan struct{})
			go func() {
				time.Sleep(2 * time.Second)
				cancel()
				ch1 <- struct{}{}
			}()
			go func() {
				c.So(
					ShouldNotBeClosedBefore(ctx, time.Second),
					ShouldEqual,
					"",
				)
				ch2 <- struct{}{}
			}()
			go func() {
				c.So(
					ShouldNotBeClosedBefore(ctx, 3*time.Second),
					ShouldEqual,
					fmt.Sprintf(shouldNotClosedBefore, "context.Background.WithCancel", 3*time.Second),
				)
				ch3 <- struct{}{}
			}()
			<-ch1
			<-ch2
			<-ch3
		})
	})
}
