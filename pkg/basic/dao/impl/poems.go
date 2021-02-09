package impl

import (
	"context"

	"github.com/alonegrowing/insight/pkg/basic/dao"
	"github.com/alonegrowing/insight/pkg/basic/macro"
	"github.com/alonegrowing/insight/pkg/basic/resource"
	"github.com/alonegrowing/insight/pkg/basic/util"
	"github.com/alonegrowing/insight/pkg/treasure/sql"
)

type PoemDaoImpl struct {
	tableName string
	pool      *sql.Group
}

var DefaultPoemDao dao.PoemDao

func init() {
	DefaultPoemDao = NewPoemDaoImpl()
}

func NewPoemDaoImpl() *PoemDaoImpl {
	return &PoemDaoImpl{
		tableName: "topic",
		pool:      resource.Get(macro.DBName),
	}
}

func (r *PoemDaoImpl) GetPoemById(ctx context.Context, id int64) (poem *dao.Poem) {
	poem = &dao.Poem{}
	if err := r.pool.Slave().Table(r.tableName).Where("id=?", id).Limit(1).Scan(&poem).Error; err != nil {
		util.PanicIfError(err)
	}
	return
}

func (r *PoemDaoImpl) GetPoemList(ctx context.Context) (poems []*dao.Poem) {
	poems = make([]*dao.Poem, 0)
	if err := r.pool.Slave().Table(r.tableName).Order("id desc").Find(&poems).Error; err != nil {
		util.PanicIfError(err)
	}
	return
}
