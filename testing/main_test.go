package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSendEmailNotifications(t *testing.T) {

	Convey("receives an event (Testing SendNotifications)", t, func() {

		Convey("receiving event with invalid notification id", nil)

	})

	Convey("receives an event", t, func() {

		Convey("event is valid, forwarding event (Testing SendNotificationEmails)", func() {

			Convey("receiving event with invalid type", nil)

		})

	})

	Convey("receives an event", t, func() {

		Convey("event is valid, forwarding event (Testing SendNotificationEmails)", func() {

			Convey("receiving email event with no recipients email ids", nil)

		})

	})

	Convey("receives an event", t, func() {

		Convey("event is valid, forwarding event (Testing SendNotificationEmails)", func() {

			Convey("getting invalid email template", nil)

		})

	})

	Convey("receives an event", t, func() {

		Convey("event is valid, forwarding event (Testing SendNotificationEmails)", func() {

			Convey("email template rendering error", nil)

		})

	})

	Convey("receives an event", t, func() {

		Convey("event is valid, forwarding event", func() {

			Convey("getting valid email template, forwarding email", func() {

				Convey("email template rendering successful", func() {

					Convey("Send email with proper recipients email ids", nil)

				})

			})

		})

	})

}
