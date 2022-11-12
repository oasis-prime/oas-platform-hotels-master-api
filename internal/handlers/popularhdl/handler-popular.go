package popularhdl

import (
	"context"

	"github.com/oasis-prime/oas-platform-core/domain/masterdm"
	"github.com/oasis-prime/oas-platform-core/repositories/enums/langenums"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Handler struct {
	servPopular ports.PopularService
}

func NewHandler(
	servPopular ports.PopularService,
) *Handler {
	return &Handler{
		servPopular: servPopular,
	}
}

func (h *Handler) GetPopular(ctx context.Context, input model.GetPopularInput) (*model.PopularData, error) {
	condition := masterdm.GetCustomerPopularRequest{
		GetAllRequestBasic: masterdm.GetAllRequestBasic{
			Page:      input.Pagination.Page,
			PageSize:  input.Pagination.PageSize,
			Language:  (*langenums.Language)(&input.Language),
			IsContent: true,
		},
	}

	result, total, err := h.servPopular.GetAll(condition)
	if err != nil {
		return nil, err
	}

	data := []*model.Popular{}

	for _, v := range result {
		count := int(v.Count)
		createdAt := v.CreatedAt.Format("2006-01-02")
		updatedAt := v.UpdatedAt.Format("2006-01-02")
		data = append(data, &model.Popular{
			Name:      &v.PopularName,
			Image:     &v.Image,
			Link:      &v.Link,
			Count:     &count,
			Status:    (*string)(&v.Status),
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		})
	}

	display := model.PopularData{
		Data: data,
		Pagination: &model.PaginationType{
			Page:     input.Pagination.Page,
			PageSize: input.Pagination.PageSize,
			Total:    int(total),
		},
	}

	return &display, nil
}
