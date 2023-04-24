package notifications

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"proxy-handler/config"
	"proxy-handler/notifications/pb"
)

func RegisterRoutes(router fiber.Router) {
	conn, err := grpc.Dial(config.GetEnv("NOTIFICATION_SERVICE_HOST"), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewNotificationServiceClient(conn)

	router.Get(
		"/seen/:id<int>",
		func(ctx *fiber.Ctx) error {
			id, _ := ctx.ParamsInt("id")

			req := &pb.SeenNotificationRequest{
				Id: uint64(id),
			}

			res, err := client.SeenNotification(context.Background(), req)

			if err == nil {
				return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": res.Success,
				})
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "sa",
			})
		},
	)
	router.Get(
		"/send",
		func(ctx *fiber.Ctx) error {
			req := &pb.SendNotificationRequest{
				Title:     "title",
				Message:   "message",
				Recipient: 1,
				Sender:    1,
				Type:      "type",
				Avatar:    "avatar",
				Link:      "link",
				LinkText:  "linkText",
				Icon:      "icon",
			}

			res, err := client.SendNotification(context.Background(), req)

			if err == nil {
				return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": res.Success,
				})
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "sa",
			})
		},
	)
	router.Get(
		"/seen_all/:id<int>",
		func(ctx *fiber.Ctx) error {
			req := &pb.SeenAllNotificationsRequest{
				UserId: 1,
			}

			res, err := client.SeenAllNotifications(context.Background(), req)

			if err == nil {
				return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": res.Success,
				})
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "sa",
			})
		},
	)
	router.Get(
		"/get_all_notification/:user_id<int>/:limit<int>/:offset<int>",
		func(ctx *fiber.Ctx) error {
			userId, _ := ctx.ParamsInt("user_id")
			limit, _ := ctx.ParamsInt("limit")
			offset, _ := ctx.ParamsInt("offset")

			req := &pb.GetNotificationsRequest{
				UserId: uint64(userId),
				Limit:  uint32(limit),
				Offset: uint32(offset),
			}

			res, err := client.GetNotifications(context.Background(), req)

			if err == nil {
				return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
					"notifications": res.Notifications,
					"total":         res.Total,
					"message":       res.Success,
				})
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "sa",
			})
		},
	)
}
