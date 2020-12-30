package shared

import (
	"context"

	"github.com/alonegrowing/insight/pkg/core/model"
)

type PoemShared interface {
	GetPoemByID(ctx context.Context, id int64) *model.Poem
	GetPoemList(ctx context.Context) []*model.Poem
}
