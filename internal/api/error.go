package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const defaultStatusCode = http.StatusInternalServerError

type ProblemDetails struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

type (
	HttpErrMap      func(err error) (int, bool)
	HttpErrMappings []HttpErrMap
)

func (h HttpErrMappings) find(err error) (int, bool) {
	for _, v := range h {
		if status, ok := v(err); ok {
			return status, true
		}
	}
	return 0, false
}

type HttpErrHandler struct {
	httpErrMappings HttpErrMappings
	handle          func(err error, c echo.Context)
}

func NewHttpErrHandler(httpErrMappings HttpErrMappings) *HttpErrHandler {
	errHandler := &HttpErrHandler{
		httpErrMappings: httpErrMappings,
	}
	errHandler.setDefaultProblemDetailsHandle()
	return errHandler
}

func (h *HttpErrHandler) setDefaultProblemDetailsHandle() {
	problemDetailsHandle := func(err error, c echo.Context) {
		code := defaultStatusCode

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		mCode, ok := h.httpErrMappings.find(err)
		if !ok {
			handleErr(c, prepareProblemDetails(err, "internal-server-error", code, c))
			return
		}
		handleErr(c, prepareProblemDetails(err, "application-error", mCode, c))
	}
	h.handle = problemDetailsHandle
}

func (s *Server) useErrorHandler(httpErrHandler *HttpErrHandler) {
	s.echo.HTTPErrorHandler = httpErrHandler.handle
}

func handleErr(c echo.Context, pDetails ProblemDetails) {
	if c.Response().Committed {
		return
	}
	if c.Request().Method == http.MethodHead {
		if err := c.NoContent(pDetails.Status); err != nil {
			c.Logger().Error(err)
		}
		return
	}
	if err := c.JSON(pDetails.Status, pDetails); err != nil {
		c.Logger().Error(err)
	}
}

func prepareProblemDetails(err error,
	typ string,
	code int,
	c echo.Context) ProblemDetails {
	return ProblemDetails{
		Type:     typ,
		Title:    err.Error(),
		Status:   code,
		Detail:   err.Error(),
		Instance: c.Request().RequestURI,
	}
}
