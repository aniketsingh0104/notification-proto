package testing

import (
	"context"
	pb "notification/proto"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	. "github.com/smartystreets/goconvey/convey"
)

func testSendEmailNotifications(t *testing.T, client pb.NotificationsSenderClient, clientEmail pb.NotificationsTypesClient) {
	ctx := context.Background()

	Convey("receives an event (Testing SendNotifications)", t, func() {

		Convey("invalid notification id should return error", func() {
			test := &empty.Empty{}
			serialized, err := ptypes.MarshalAny(test)
			if err != nil {
				t.Fatal("could not serialize empty")
			}
			newNotification := &pb.NotificationRequest{
				NotificationId:   "",
				NotificationData: serialized,
			}
			resp, _ := client.SendNotifications(ctx, newNotification)
			So(resp.ErrorStatus, ShouldEqual, "Invalid NotificationId")
		})

		Convey("event is valid, forwarding event (Testing SendNotificationEmails)", func() {

			newNotification := &pb.NotificationRequest{}

			Convey("receiving event with invalid type should return error", func() {
				requestData := &pb.ReqEventData{
					Recipients: []string{"yxhdg@gmail.com", "yxhdg@gmail.com"},
					HtmlBody:   "<h1></h1>",
					Subject:    "Hello",
				}
				responseData := &pb.ResEventData{}
				req, err := ptypes.MarshalAny(requestData)
				if err != nil {
					t.Fatal("Marshaling error requestData")
				}
				res, err := ptypes.MarshalAny(responseData)
				if err != nil {
					t.Fatal("Marshaling error responseData")
				}
				test := &pb.PushNotificationEventDataReq{
					EventType:    "xyz",
					RequestData:  req,
					ResponseData: res,
				}
				serialized, err := ptypes.MarshalAny(test)
				if err != nil {
					t.Fatal("Marshaling error test")
				}
				newNotification.NotificationId = RandomString(10)
				newNotification.NotificationData = serialized
				resp, _ := clientEmail.SendNotificationEmails(ctx, newNotification)
				So(resp.ErrorStatus, ShouldEqual, "Type mismatch")
			})

			Convey("receiving email event with no recipients email ids should return error", func() {
				requestData := &pb.ReqEventData{
					Recipients: []string{},
					HtmlBody:   "<h1></h1>",
					Subject:    "Hello",
				}
				responseData := &pb.ResEventData{}
				req, err := ptypes.MarshalAny(requestData)
				if err != nil {
					t.Fatal("Marshaling error requestData")
				}
				res, err := ptypes.MarshalAny(responseData)
				if err != nil {
					t.Fatal("Marshaling error responseData")
				}
				test := &pb.PushNotificationEventDataReq{
					EventType:    "Email",
					RequestData:  req,
					ResponseData: res,
				}
				serialized, err := ptypes.MarshalAny(test)
				if err != nil {
					t.Fatal("Marshaling error test")
				}
				newNotification.NotificationId = RandomString(10)
				newNotification.NotificationData = serialized
				resp, _ := clientEmail.SendNotificationEmails(ctx, newNotification)
				So(resp.ErrorStatus, ShouldEqual, "No valid Recipients")
			})

			Convey("getting invalid email template should return error", func() {
				test := &empty.Empty{}
				serialized, err := ptypes.MarshalAny(test)
				if err != nil {
					t.Fatal("could not serialize test")
				}
				newNotification.NotificationId = RandomString(10)
				newNotification.NotificationData = serialized
				_, err = clientEmail.SendNotificationEmails(ctx, newNotification)
				So(err, ShouldNotBeNil)
			})

			Convey("email template rendering error should return error", func() {
				test := &empty.Empty{}
				serialized, err := ptypes.MarshalAny(test)
				if err != nil {
					t.Fatal("could not serialize empty")
				}
				newNotification.NotificationId = RandomString(10)
				newNotification.NotificationData = serialized
				_, err = clientEmail.SendNotificationEmails(ctx, newNotification)
				So(err, ShouldNotBeNil)
			})
			Convey("getting valid email template and rendering successful, forwarding email", func() {

				Convey("Sending email with invalid structure should return error", func() {
					test := &empty.Empty{}
					serialized, err := ptypes.MarshalAny(test)
					if err != nil {
						t.Fatal("could not serialize empty")
					}
					newNotification.NotificationId = /*some random id*/ ""
					newNotification.NotificationData = serialized
					_, err = client.SendNotifications(ctx, newNotification)
					So(err, ShouldNotBeNil)
				})

				Convey("Sending email with no proper recipients", func() {

				})

				Convey("Send email with proper recipients email ids", func() {
					requestData := &pb.ReqEventData{
						Recipients: []string{"aniketsingh0104@gmail.com", "aniket_singh1998@yahoo.com"},
						HtmlBody:   "<h1></h1>",
						Subject:    "Hello",
					}
					responseData := &pb.ResEventData{}
					req, err := ptypes.MarshalAny(requestData)
					if err != nil {
						t.Fatal("Marshaling error requestData")
					}
					res, err := ptypes.MarshalAny(responseData)
					if err != nil {
						t.Fatal("Marshaling error responseData")
					}
					test := &pb.PushNotificationEventDataReq{
						EventType:    "Email",
						RequestData:  req,
						ResponseData: res,
					}
					serialized, err := ptypes.MarshalAny(test)
					if err != nil {
						t.Fatal("Marshaling error test")
					}
					newNotification.NotificationId = RandomString(10)
					newNotification.NotificationData = serialized
					resp, err := client.SendNotifications(ctx, newNotification)
					So(err, ShouldBeNil)
					So(resp.ErrorStatus, ShouldEqual, "OK")
					resp, err = clientEmail.SendNotificationEmails(ctx, newNotification)
					So(resp.ErrorStatus, ShouldEqual, "OK")
					So(err, ShouldBeNil)

				})

			})

		})

	})

}
