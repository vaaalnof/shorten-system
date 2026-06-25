package http

import (
	"shortener-service/internal/model"
	"shortener-service/internal/usecase/shorturl"

	"github.com/gofiber/fiber/v2"
)

type URLController struct {
	useCase *shorturl.URLUseCase
}

func NewURLController(
	useCase *shorturl.URLUseCase,
) *URLController {

	return &URLController{
		useCase: useCase,
	}
}

// =====================================================
// CREATE URL
// =====================================================

func (c *URLController) Create(
	ctx *fiber.Ctx,
) error {

	request := new(
		model.AddURLRequest,
	)

	if err := ctx.BodyParser(
		request,
	); err != nil {

		return err
	}

	response, err := c.useCase.Add(
		ctx.UserContext(),
		request,
	)

	if err != nil {
		return err
	}

	return ctx.Status(
		fiber.StatusCreated,
	).JSON(
		response,
	)
}

// =====================================================
// GET URL
// =====================================================

func (c *URLController) Get(
	ctx *fiber.Ctx,
) error {

	request := &model.GetURLRequest{
		ID: ctx.Params(
			"id",
		),
	}

	response, err := c.useCase.Get(
		ctx.UserContext(),
		request,
	)

	if err != nil {
		return err
	}

	return ctx.JSON(
		response,
	)
}

// =====================================================
// REDIRECT URL
// =====================================================

func (c *URLController) Redirect(
	ctx *fiber.Ctx,
) error {

	request := &model.RedirectURLRequest{
		ShortCode: ctx.Params(
			"short_code",
		),
	}

	originalURL, err := c.useCase.Redirect(
		ctx.UserContext(),
		request,
	)

	if err != nil {

		ctx.App().Config().ErrorHandler(
			ctx,
			err,
		)
		return err

	}

	return ctx.Redirect(
		originalURL,
		fiber.StatusTemporaryRedirect,
	)
}

// =====================================================
// UPDATE PASSWORD
// =====================================================

func (c *URLController) UpdatePassword(
	ctx *fiber.Ctx,
) error {

	request := &model.UpdateURLPasswordRequest{
		ID: ctx.Params(
			"id",
		),
	}

	if err := ctx.BodyParser(
		request,
	); err != nil {

		return err
	}

	response, err := c.useCase.UpdatePassword(
		ctx.UserContext(),
		request,
	)

	if err != nil {

		return err
	}

	return ctx.JSON(
		response,
	)
}

// =====================================================
// REMOVE PASSWORD
// =====================================================

func (c *URLController) RemovePassword(
	ctx *fiber.Ctx,
) error {

	request := &model.RemoveURLPasswordRequest{
		ID: ctx.Params(
			"id",
		),
	}

	if err := ctx.BodyParser(
		request,
	); err != nil {

		return err
	}

	response, err := c.useCase.RemovePassword(
		ctx.UserContext(),
		request,
	)

	if err != nil {

		return err
	}

	return ctx.JSON(
		response,
	)
}

// =====================================================
// VERIFY PASSWORD
// =====================================================

func (c *URLController) VerifyPassword(
	ctx *fiber.Ctx,
) error {

	request := &model.VerifyURLPasswordRequest{
		ShortCode: ctx.Params(
			"short_code",
		),
	}

	if err := ctx.BodyParser(
		request,
	); err != nil {

		return err
	}

	originalURL, err := c.useCase.VerifyPassword(
		ctx.UserContext(),
		request,
	)

	if err != nil {

		return err
	}

	return ctx.Redirect(
		originalURL,
		fiber.StatusTemporaryRedirect,
	)
}

// =====================================================
// LIST URLS
// =====================================================

func (c *URLController) List(
	ctx *fiber.Ctx,
) error {

	request := &model.ListURLsRequest{}

	if err := ctx.QueryParser(
		request,
	); err != nil {

		return err
	}

	response, err := c.useCase.List(
		ctx.UserContext(),
		request,
	)

	if err != nil {

		return err
	}

	return ctx.JSON(
		response,
	)
}

// =====================================================
// DELETE URL
// =====================================================

func (c *URLController) Delete(
	ctx *fiber.Ctx,
) error {

	request := &model.DeleteURLRequest{
		ID: ctx.Params(
			"id",
		),
	}

	response, err := c.useCase.Delete(
		ctx.UserContext(),
		request,
	)

	if err != nil {

		return err
	}

	return ctx.JSON(
		response,
	)
}

// =====================================================
// UPDATE URL
// =====================================================

func (c *URLController) Update(
	ctx *fiber.Ctx,
) error {

	request := &model.UpdateURLRequest{
		ID: ctx.Params(
			"id",
		),
	}

	if err := ctx.BodyParser(
		request,
	); err != nil {

		return err
	}

	response, err := c.useCase.Update(
		ctx.UserContext(),
		request,
	)

	if err != nil {

		return err
	}

	return ctx.JSON(
		response,
	)
}
