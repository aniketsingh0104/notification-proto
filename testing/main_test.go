package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNotifications(t *testing.T) {

	Convey("receives an event", t, func() {

		Convey("invalid notification id", func() {

		})

		Convey("forwarding notification as email notification", func() {

			Convey("invalid event type (type_url is invalid)", nil)

			Convey("invalid data", func() {

				Convey("recipients email ids are not present", nil)

			})

			Convey("get template", func() {

				Convey("invalid template", nil)

				Convey("render template", func() {

					Convey("rendering error", nil)

					Convey("send emails and check responses", func() {

						Convey("send email with proper recipients email addresses", nil)

						Convey("send invalid emails", nil)

					})

				})

			})

		})

	})

}
