package impl

import (
	"context"

	"github.com/alonegrowing/insight/pkg/business/shared"
	"github.com/alonegrowing/insight/pkg/business/shared/impl"
	core "github.com/alonegrowing/insight/pkg/core/model"
	"github.com/alonegrowing/insight/pkg/web/controller"
	"github.com/alonegrowing/insight/pkg/web/model"
)

type HomeControllerImpl struct {
	poemShared shared.PoemShared
}

var DefaultHomeController controller.HomeController

func init() {
	DefaultHomeController = NewHomeControllerImpl()
}

func NewHomeControllerImpl() *HomeControllerImpl {
	return &HomeControllerImpl{
		poemShared: impl.DefaultPoemShared,
	}
}

func (r *HomeControllerImpl) GetPoemById(ctx context.Context, id int64) *model.Poem {
	return r.formater(r.poemShared.GetPoemByID(ctx, id))
}

func (r *HomeControllerImpl) GetPoemList(ctx context.Context) []*model.Poem {
	data := r.poemShared.GetPoemList(ctx)
	poems := make([]*model.Poem, 0)
	for _, poem := range data {
		poems = append(poems, r.formater(poem))
	}
	return poems
}

func (r *HomeControllerImpl) formater(poem *core.Poem) *model.Poem {
	return &model.Poem{
		Id:        poem.Id,
		Token:     poem.Token,
		Title:     poem.Title,
		Author:    poem.Author,
		Intro:     poem.Intro,
		Imgs:      poem.Imgs,
		Content:   poem.Content,
		UpdatedAt: poem.UpdatedAt,
		CreatedAt: poem.CreatedAt,
	}
}
