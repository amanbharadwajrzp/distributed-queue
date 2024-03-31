package client

import "context"

type JobHandler interface {
	// Handle will be called to process each Job sent via Outboxer. JobPayload is the payload that was sent.
	Handle(ctx context.Context, payload IPayload) error

	// ZeroPayload should return the zero value of the payload for this handler. This is used to register the struct with encoder.
	// Handle payload will be of the type returned here.
	// Note: This should be returning the payload value, NOT the pointer to the value.
	// Note: With gob encoder, the payload has to be a non-scalar type. Otherwise you'll have to register all types separately.
	ZeroPayload() IPayload
}
