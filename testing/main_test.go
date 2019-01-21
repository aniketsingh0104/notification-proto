package main

import (
	"context"
	pb "notification/proto"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"
	. "github.com/smartystreets/goconvey/convey"
)

func testSendEmailNotifications(t *testing.T, client pb.NotificationsSenderClient) {
	ctx := context.Background()

	Convey("receives an event (Testing SendNotifications)", t, func() {

		Convey("invalid notification id should return error", func() {
			test := &empty.Empty{}
			serialized, err := proto.Marshal(test)
			if err != nil {
				t.Fatal("could not serialize empty")
			}
			newNotification := &pb.NotificationRequest{
				NotificationId: "",
				NotificationData: &any.Any{
					TypeUrl: proto.MessageName(test),
					Value:   serialized,
				},
			}
			resp, _ := client.SendNotifications(ctx, newNotification)
			So(resp.ErrorStatus, ShouldEqual, "Invalid NotificationId")
		})

		Convey("event is valid, forwarding event (Testing SendNotificationEmails)", func() {

			newNotification := &pb.NotificationRequest{}

			Convey("receiving event with invalid type should return error", func() {

			})

			Convey("receiving email event with no recipients email ids should return error", nil)

			Convey("getting invalid email template should return error", func() {
				test := &empty.Empty{}
				serialized, err := proto.Marshal(test)
				if err != nil {
					t.Fatal("could not serialize empty")
				}
				newNotification.NotificationId = /*some random id*/ ""
				newNotification.NotificationData = &any.Any{
					TypeUrl: proto.MessageName(test),
					Value:   serialized,
				}
				_, err = client.SendNotifications(ctx, newNotification)
				So(err, ShouldNotBeNil)
			})

			Convey("email template rendering error should return error", nil)
			Convey("getting valid email template and rendering successful, forwarding email", func() {

				Convey("Sending email with invalid structure should return error", func() {
					test := &empty.Empty{}
					serialized, err := proto.Marshal(test)
					if err != nil {
						t.Fatal("could not serialize empty")
					}
					newNotification.NotificationId = /*some random id*/ ""
					newNotification.NotificationData = &any.Any{
						TypeUrl: proto.MessageName(test),
						Value:   serialized,
					}
					_, err = client.SendNotifications(ctx, newNotification)
					So(err, ShouldNotBeNil)
				})

				Convey("Sending email with no proper recipients", nil)

				Convey("Send email with proper recipients email ids", nil)

			})

		})

	})

}
