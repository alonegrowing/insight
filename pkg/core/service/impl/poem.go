package impl

import (
	"context"
	"strings"

	"github.com/alonegrowing/insight/pkg/basic/dao"
	"github.com/alonegrowing/insight/pkg/basic/dao/impl"
	"github.com/alonegrowing/insight/pkg/core/model"
	"github.com/alonegrowing/insight/pkg/core/service"
)

type PoemServiceImpl struct {
	poemDao dao.PoemDao
}

var DefaultPoemService service.PoemService

func init() {
	DefaultPoemService = NewPoemServiceImpl()
}

func NewPoemServiceImpl() *PoemServiceImpl {
	return &PoemServiceImpl{
		poemDao: impl.DefaultPoemDao,
	}
}

func (r *PoemServiceImpl) GetPoemByID(ctx context.Context, id int64) *model.Poem {
	return r.formater(r.poemDao.GetPoemById(ctx, id))
}

func (r *PoemServiceImpl) GetPoemList(ctx context.Context) []*model.Poem {
	data := r.poemDao.GetPoemList(ctx)
	poems := make([]*model.Poem, 0)
	for _, item := range data {
		poems = append(poems, r.formater(item))
	}
	return poems
}

func (r *PoemServiceImpl) formater(poem *dao.Poem) *model.Poem {
	return &model.Poem{
		Id:        poem.Id,
		Token:     poem.Token,
		Author:    poem.Author,
		Title:     poem.Title,
		Intro:     poem.Intro,
		Imgs:      strings.Split(poem.Imgs, ","),
		Content:   poem.Content,
		Status:    poem.Status,
		UpdatedAt: poem.UpdatedAt,
		CreatedAt: poem.CreatedAt,
	}
}
