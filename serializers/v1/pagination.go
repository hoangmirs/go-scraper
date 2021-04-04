package v1serializers

import (
	"github.com/google/jsonapi"

	"github.com/beego/beego/v2/core/utils/pagination"
)

type Pagination struct {
	Paginator *pagination.Paginator
}

func (serializer *Pagination) Meta() (meta *jsonapi.Meta) {
	return &jsonapi.Meta{
		"page": serializer.Paginator.Page(),
		"pages":        serializer.Paginator.PageNums(),
		"records":      serializer.Paginator.Nums(),
	}
}

func (serializer *Pagination) Links() (links *jsonapi.Links) {
	currentPage := serializer.Paginator.Page()

	return &jsonapi.Links{
		"self":  serializer.Paginator.PageLink(currentPage),
		"first": serializer.Paginator.PageLinkFirst(),
		"prev":  serializer.Paginator.PageLinkPrev(),
		"next":  serializer.Paginator.PageLinkNext(),
		"last":  serializer.Paginator.PageLinkLast(),
	}
}
