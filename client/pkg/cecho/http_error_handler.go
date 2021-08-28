package cecho

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"google.golang.org/grpc/status"
	"net/http"
)

func HTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	// Issue #1426
	code := he.Code
	message := he.Message
	if m, ok := he.Message.(string); ok {
		if c.Echo().Debug {
			message = echo.Map{"message": m, "error": err.Error()}
		} else {
			message = echo.Map{"message": m}
		}
	} else if e, ok := he.Message.(error); ok {
		message = echo.Map{
			"message": e.Error(),
		}
	}

	isJSONError := false
	s, ok := status.FromError(err)
	if ok {
		result := gjson.Parse(s.Message())
		if result.IsObject() {
			code = int(result.Get("code").Int())
			errorsResult := result.Get("errors")
			if errorsResult.Exists() && errorsResult.IsObject() {
				message = errorsResult.Raw
			}
			messageResult := result.Get("message")
			if messageResult.Exists() {
				message, _ = sjson.Set(`{}`, "message", messageResult.Value())
			}
			isJSONError = true
		}
	}

	if ve, ok := err.(*validation.Errors); ok {
		code = http.StatusUnprocessableEntity
		message = ve
	}

	if ve, ok := err.(validation.Errors); ok {
		code = http.StatusUnprocessableEntity
		message = ve
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(he.Code)
		} else if isJSONError {
			c.Response().Header().Set("Content-Type", "application/json; charset=UTF-8")
			err = c.String(code, message.(string))
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
