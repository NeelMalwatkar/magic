package errors

type BaseError struct {
	Message string
}

func (e *BaseError) Error() string {
	return e.Message
}

type BadRequest struct{ BaseError }
type NotFound struct{ BaseError }
type ServiceUnavailable struct{ BaseError }
type Forbidden struct{ BaseError }
type Unauthorized struct{ BaseError }
type MethodNotAllowed struct{ BaseError }
type Conflict struct{ BaseError }
type Gone struct{ BaseError }
type UnsupportedMediaType struct{ BaseError }
type UnprocessableEntity struct{ BaseError }
type TooManyRequests struct{ BaseError }
type InternalServerError struct{ BaseError }
type BadGateway struct{ BaseError }
type GatewayTimeout struct{ BaseError }
type RequestTimeout struct{ BaseError }
type NotImplemented struct{ BaseError }
