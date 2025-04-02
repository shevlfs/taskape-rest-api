package handlers

import (
	"context"
	"strconv"
	"taskape-rest-api/internal/dto"
	proto "taskape-rest-api/proto"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/metadata"
)

const timestampFormat = "2006-01-02T15:04:05Z07:00"

type EventHandler struct {
	BackendClient proto.BackendRequestsClient
}

func (h *EventHandler) GetUserEvents(c *fiber.Ctx) error {
	userID := c.Params("userID")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User ID is required",
			"events":  []interface{}{},
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authorization token is required",
			"events":  []interface{}{},
		})
	}

	limit, _ := strconv.Atoi(c.Query("limit", "50"))
	includeExpired := c.Query("include_expired", "false") == "true"

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.GetUserEvents(ctx, &proto.GetUserEventsRequest{
		UserId:         userID,
		Limit:          int32(limit),
		IncludeExpired: includeExpired,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch events: " + err.Error(),
			"events":  []interface{}{},
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": resp.Error,
			"events":  []interface{}{},
		})
	}

	events := make([]dto.EventResponse, len(resp.Events))
	for i, event := range resp.Events {
		events[i] = convertToEventResponse(event)
	}
	println("giving out events, count:", len(events))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"events":  events,
		"message": nil,
	})
}

func (h *EventHandler) LikeEvent(c *fiber.Ctx) error {
	eventID := c.Params("eventID")
	if eventID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Event ID is required",
		})
	}

	var request dto.LikeEventRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request format: " + err.Error(),
		})
	}

	if request.UserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User ID is required",
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.LikeEvent(ctx, &proto.LikeEventRequest{
		EventId: eventID,
		UserId:  request.UserID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to like event: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":     false,
			"message":     resp.Error,
			"likes_count": 0,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":     true,
		"likes_count": resp.LikesCount,
	})
}

func (h *EventHandler) UnlikeEvent(c *fiber.Ctx) error {
	eventID := c.Params("eventID")
	if eventID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Event ID is required",
		})
	}

	userID := c.Query("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User ID query parameter is required",
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.UnlikeEvent(ctx, &proto.UnlikeEventRequest{
		EventId: eventID,
		UserId:  userID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to unlike event: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":     false,
			"message":     resp.Error,
			"likes_count": 0,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":     true,
		"likes_count": resp.LikesCount,
	})
}

func (h *EventHandler) GetEventComments(c *fiber.Ctx) error {
	eventID := c.Params("eventID")
	if eventID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":  false,
			"message":  "Event ID is required",
			"comments": []interface{}{},
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success":  false,
			"message":  "Authorization token is required",
			"comments": []interface{}{},
		})
	}

	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.GetEventComments(ctx, &proto.GetEventCommentsRequest{
		EventId: eventID,
		Limit:   int32(limit),
		Offset:  int32(offset),
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success":  false,
			"message":  "Failed to fetch comments: " + err.Error(),
			"comments": []interface{}{},
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":  false,
			"message":  resp.Error,
			"comments": []interface{}{},
		})
	}

	comments := make([]dto.EventCommentResponse, len(resp.Comments))
	for i, comment := range resp.Comments {
		comments[i] = convertToEventCommentResponse(comment)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":     true,
		"comments":    comments,
		"total_count": resp.TotalCount,
	})
}

func (h *EventHandler) AddEventComment(c *fiber.Ctx) error {
	eventID := c.Params("eventID")
	if eventID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Event ID is required",
		})
	}

	var request dto.AddEventCommentRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request format: " + err.Error(),
		})
	}

	if request.UserID == "" || request.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User ID and content are required",
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.AddEventComment(ctx, &proto.AddEventCommentRequest{
		EventId: eventID,
		UserId:  request.UserID,
		Content: request.Content,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to add comment: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": resp.Error,
		})
	}

	comment := convertToEventCommentResponse(resp.Comment)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"comment": comment,
	})
}

func (h *EventHandler) DeleteEventComment(c *fiber.Ctx) error {
	eventID := c.Params("eventID")
	commentID := c.Params("commentID")
	if eventID == "" || commentID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Event ID and Comment ID are required",
		})
	}

	userID := c.Query("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User ID query parameter is required",
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.DeleteEventComment(ctx, &proto.DeleteEventCommentRequest{
		CommentId: commentID,
		UserId:    userID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete comment: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": resp.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func convertToEventResponse(event *proto.Event) dto.EventResponse {
	response := dto.EventResponse{
		ID:             event.Id,
		UserID:         event.UserId,
		TargetUserID:   event.TargetUserId,
		Type:           convertEventTypeToString(event.Type),
		Size:           convertEventSizeToString(event.Size),
		CreatedAt:      event.CreatedAt.AsTime().Format(timestampFormat),
		TaskIDs:        event.TaskIds,
		StreakDays:     int(event.StreakDays),
		LikesCount:     int(event.LikesCount),
		CommentsCount:  int(event.CommentsCount),
		LikedByUserIDs: event.LikedByUserIds,
	}

	if event.ExpiresAt != nil {
		expiresAt := event.ExpiresAt.AsTime().Format(timestampFormat)
		response.ExpiresAt = &expiresAt
	}

	return response
}

func convertToEventCommentResponse(comment *proto.EventComment) dto.EventCommentResponse {
	response := dto.EventCommentResponse{
		ID:        comment.Id,
		EventID:   comment.EventId,
		UserID:    comment.UserId,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.AsTime().Format(timestampFormat),
		IsEdited:  comment.IsEdited,
	}

	if comment.EditedAt != nil {
		editedAt := comment.EditedAt.AsTime().Format(timestampFormat)
		response.EditedAt = &editedAt
	}

	return response
}

func convertEventTypeToString(eventType proto.EventType) string {
	switch eventType {
	case proto.EventType_NEW_TASKS_ADDED:
		return "new_tasks_added"
	case proto.EventType_NEWLY_RECEIVED:
		return "newly_received"
	case proto.EventType_NEWLY_COMPLETED:
		return "newly_completed"
	case proto.EventType_REQUIRES_CONFIRMATION:
		return "requires_confirmation"
	case proto.EventType_N_DAY_STREAK:
		return "n_day_streak"
	case proto.EventType_DEADLINE_COMING_UP:
		return "deadline_coming_up"
	default:
		return "unknown"
	}
}

func convertEventSizeToString(eventSize proto.EventSize) string {
	switch eventSize {
	case proto.EventSize_SMALL:
		return "small"
	case proto.EventSize_MEDIUM:
		return "medium"
	case proto.EventSize_LARGE:
		return "large"
	default:
		return "medium"
	}
}

func (h *EventHandler) GetUserRelatedEvents(c *fiber.Ctx) error {
	targetUserID := c.Params("userID")
	if targetUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User ID is required",
			"events":  []interface{}{},
		})
	}

	requesterID := c.Query("requester_id", "")
	if requesterID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "requester_id query parameter is required",
			"events":  []interface{}{},
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authorization token is required",
			"events":  []interface{}{},
		})
	}

	limit, _ := strconv.Atoi(c.Query("limit", "50"))
	includeExpired := c.Query("include_expired", "false") == "true"

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.GetUserRelatedEvents(ctx, &proto.GetUserRelatedEventsRequest{
		TargetUserId:   targetUserID,
		RequesterId:    requesterID,
		Limit:          int32(limit),
		IncludeExpired: includeExpired,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch events: " + err.Error(),
			"events":  []interface{}{},
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": resp.Error,
			"events":  []interface{}{},
		})
	}

	events := make([]dto.EventResponse, len(resp.Events))
	for i, event := range resp.Events {
		events[i] = convertToEventResponse(event)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"events":  events,
		"message": nil,
	})
}
