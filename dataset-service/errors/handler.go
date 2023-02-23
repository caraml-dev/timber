package errors

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/caraml-dev/timber/common/log"
)

// Handler is error handler that will mutate the http status code based on error's type
func Handler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
	log.Error(err)
	newError := &runtime.HTTPStatusError{
		HTTPStatus: http.StatusInternalServerError,
		Err:        err,
	}

	switch err.(type) {
	case *NotFoundError:
		newError = &runtime.HTTPStatusError{
			HTTPStatus: http.StatusNotFound,
			Err:        err,
		}
	case *ConflictError:
		newError = &runtime.HTTPStatusError{
			HTTPStatus: http.StatusConflict,
			Err:        err,
		}
	case *InvalidInputError:
		newError = &runtime.HTTPStatusError{
			HTTPStatus: http.StatusBadRequest,
			Err:        err,
		}
	default:
	}

	// using default handler to do the rest of heavy lifting of marshaling error and adding headers
	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, request, newError)
}
