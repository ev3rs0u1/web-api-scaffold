package requestor

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"web-api-scaffold/internal/pkg/errno"
	"web-api-scaffold/internal/pkg/responder"
)

var validate = validator.New()

type Binder interface {
	Input() interface{}
	Validate() (err error)
	Handler(ctx *fiber.Ctx) error
}

func Bind(b Binder) fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		switch v := bindError(ctx, b).(type) {
		case nil:
			return
		case *responder.Body:
			return ctx.JSON(v)
		case errno.Code:
			return ctx.JSON(responder.BuildBody(v))
		default:
			return ctx.JSON(responder.BuildBody(errno.CodeUnknownError).With(err))
		}
	}
}

func bindError(ctx *fiber.Ctx, b Binder) (err error) {
	if reflect.ValueOf(b.Input()).Kind() == reflect.Ptr {
		input := reflect.ValueOf(b.Input()).Elem()
		input.Set(reflect.Zero(input.Type()))
	}

	switch ctx.Method() {
	case fiber.MethodGet:
		if err = ctx.QueryParser(b.Input()); err != nil {
			return errno.ErrInvalidParams
		}
	case fiber.MethodPost:
		if err = ctx.BodyParser(b.Input()); err != nil {
			if err != fiber.ErrUnprocessableEntity {
				return errno.ErrInvalidParams
			}
		}
	}

	if err = validate.Struct(b.Input()); err != nil {
		return responder.BuildBody(errno.ErrInvalidParams).With(err)
	}

	if err = b.Validate(); err != nil {
		return err
	}

	return b.Handler(ctx)
}
