package middlewares

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/naoya7076/golang-api/apperrors"
	"google.golang.org/api/idtoken"
)

var (
	googleClientID = os.Getenv("GOOGLE_CLIENT_ID")
)

type userNameKey struct{}

func GetUserName(ctx context.Context) string {
	id := ctx.Value(userNameKey{})
	if usernameStr, ok := id.(string); ok {
		return usernameStr
	}
	return ""
}

func SetUserName(req *http.Request, name string) *http.Request {
	ctx := req.Context()

	ctx = context.WithValue(ctx, userNameKey{}, name)
	req = req.WithContext(ctx)

	return req
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid req header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		bearer, idToken := authHeaders[0], authHeaders[1]
		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid req header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err = apperrors.CannotMakeValidator.Wrap(err, "internal auth error")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err = apperrors.Unauthorized.Wrap(err, "invalid auth token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		name, ok := payload.Claims["name"]
		if !ok {
			err = apperrors.Unauthorized.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		req = SetUserName(req, name.(string))

		next.ServeHTTP(w, req)
	})
}
