package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/naoya7076/golang-api/api/middlewares"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := middlewares.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParams:
		statusCode = http.StatusBadRequest
	case RequiredAuthorizationHeader, Unauthorized:
		statusCode = http.StatusUnauthorized
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
