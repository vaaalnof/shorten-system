package http

import (
	"shortener-service/internal/model"
	"shortener-service/internal/usecase/report"

	"github.com/gofiber/fiber/v2"
)

type ReportController struct {
	useCase *report.ReportUseCase
}

func NewReportController(
	useCase *report.ReportUseCase,
) *ReportController {

	return &ReportController{
		useCase: useCase,
	}
}

// =====================================================
// REPORT SUMMARY
// =====================================================

func (c *ReportController) Summary(
	ctx *fiber.Ctx,
) error {

	request := &model.GetReportSummaryRequest{
		ID: ctx.Params(
			"id",
		),
	}

	response, err := c.useCase.Summary(
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
// REPORT CHART
// =====================================================

func (c *ReportController) Chart(

	ctx *fiber.Ctx,

) error {

	request := &model.GetReportChartRequest{
		ID: ctx.Params(
			"id",
		),
	}
	response, err := c.useCase.Chart(
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
// REPORT REFERRERS
// =====================================================

func (c *ReportController) Referrers(
	ctx *fiber.Ctx,
) error {

	request := &model.GetReportReferrersRequest{
		ID: ctx.Params(
			"id",
		),
	}

	response, err := c.useCase.Referrers(
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
// REPORT COUNTRIES
// =====================================================

func (c *ReportController) Countries(
	ctx *fiber.Ctx,
) error {

	request := &model.GetReportCountriesRequest{
		ID: ctx.Params(
			"id",
		),
	}

	response, err := c.useCase.Countries(
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
// REPORT DEVICES
// =====================================================

func (c *ReportController) Devices(
	ctx *fiber.Ctx,
) error {

	request := &model.GetReportDevicesRequest{
		ID: ctx.Params(
			"id",
		),
	}

	response, err := c.useCase.Devices(
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
// REPORT BROWSERS
// =====================================================

func (c *ReportController) Browsers(
	ctx *fiber.Ctx,
) error {

	request := &model.GetReportBrowsersRequest{
		ID: ctx.Params(
			"id",
		),
	}

	response, err := c.useCase.Browsers(
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
// TOP LINKS
// =====================================================

func (c *ReportController) TopLinks(
	ctx *fiber.Ctx,
) error {

	request := &model.GetTopLinksRequest{}

	response, err := c.useCase.TopLinks(
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
