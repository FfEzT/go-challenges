package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/go-playground/validator"
)

type Response struct {
  Status string `json:"status"`
  Error  string `json:"error,omitempty"`

  X       string `json:"X,omitempty"`
  Y       string `json:"Y,omitempty"`
  IsEqual string `json:"IsEqual,omitempty"`
}

type Request struct {
  X1 float64 `json:"X1" validate:"required"`
  X2 float64 `json:"X2" validate:"required"`
  X3 float64 `json:"X3" validate:"required"`
  Y1 float64 `json:"Y1" validate:"required"`
  Y2 float64 `json:"Y2" validate:"required"`
  Y3 float64 `json:"Y3" validate:"required"`
  E  uint    `json:"E"`
}

const (
  StatusOK    = "OK"
  StatusError = "ERROR"
)

func LimitHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(402)
  render.JSON(
    w,
    r,
    Error("Too many requests"),
  )
}

func Error(msg string) Response {
  return Response{
    Status: StatusError,
    Error:  msg,
  }
}

// TODO use Closures for logging
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
  var req Request
  err := render.DecodeJSON(r.Body, &req)

  // empty Body of req
  if errors.Is(err, io.EOF) {
    // TODO log.Error("request body is empty")
    render.JSON(w, r, Error("empty request"))
    return
  }
  // failed to decode request
  if err != nil {
    // TODO log.Error("failed to decode request", ...)
    render.JSON(w, r, Error("failed to decode request"))
    return
  }
  if err := validator.New().Struct(req); err != nil {
    // TODO validateErr := err.(validator.ValidationErrors)
    render.JSON(w, r, Error("invalid request"))
    return
  }

  // NOTE выставление кол-ва чисел после запятой
  formatString := "%." + strconv.FormatUint(uint64(req.E), 10) + "f"
  // NOTE Каюсь, я на всякий случай попробовал параллельно вычислять X и Y, но это сильно увеличивает время
  var (
    x string = fmt.Sprintf(formatString, calc(req.X1, req.X2, req.X3) )
    y string = fmt.Sprintf(formatString, calc(req.Y1, req.Y2, req.Y3) )
    e string
  )
  if x == y {
    e = "T"
  } else {
    e = "F"
  }

  render.JSON(
    w,
    r,
    Response{
      Status: StatusOK,
      X: x,
      Y: y,
      IsEqual: e,
    },
  )
}

func calc(a, b, c float64) float64 {
  return (a * c) / b
}