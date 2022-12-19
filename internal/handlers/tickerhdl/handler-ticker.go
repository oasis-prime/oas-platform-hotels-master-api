package tickerhdl

import (
	"context"

	"github.com/oasis-prime/oas-platform-core/domain/masterdm"
	"github.com/oasis-prime/oas-platform-core/repositories/enums/langenums"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Handler struct {
	servTicker ports.TickerService
}

func NewHandler(
	servTicker ports.TickerService,
) *Handler {
	return &Handler{
		servTicker: servTicker,
	}
}

func (h *Handler) GetTickers(ctx context.Context, input model.GetTickersInput) (display *model.TickerData, err error) {
	order := ""
	if input.Pagination.OrderBy != nil {
		order = *input.Pagination.OrderBy
	}

	result, total, err := h.servTicker.TickerGetAll(masterdm.GetAllRequestBasic{
		Page:      input.Pagination.Page,
		PageSize:  input.Pagination.PageSize,
		Language:  (*langenums.Language)(&input.Language),
		Order:     order,
		IsContent: true,
	})

	if err != nil {
		return display, err
	}

	data := []*model.Ticker{}

	for _, v := range result {
		data = append(data, &model.Ticker{
			ID:          int(v.Code),
			Name:        &v.TickerName,
			Description: &v.TickerDescription,
			Image:       &v.Image,
			Link:        &v.Link,
			Count:       int(v.Count),
		})
	}

	display = &model.TickerData{
		Data: data,
		Pagination: &model.PaginationType{
			Page:     input.Pagination.Page,
			PageSize: input.Pagination.PageSize,
			Total:    int(total),
		},
	}

	return display, err
}

func (h *Handler) GetTicker(ctx context.Context, input model.GetTickerInput) (display *model.Ticker, err error) {
	language := (*langenums.Language)(&input.Language)

	record, err := h.servTicker.TickerGet(uint(input.ID), *language)
	if err != nil {
		return display, err
	}

	display = &model.Ticker{
		ID:               int(record.Code),
		Name:             &record.TickerName,
		Description:      &record.TickerDescription,
		Image:            &record.Image,
		Link:             &record.Link,
		Count:            int(record.Count),
		TotalSellingRate: float64(record.TotalSellingRate),
		TotalNet:         float64(record.TotalNet),
	}

	return display, err
}

func (h *Handler) Ticker(ctx context.Context, input model.TickerInput) (display *model.Ticker, err error) {

	return display, err
}
