package main

import (
	"context"
	"fmt"

	"github.com/alonegrowing/insight/pkg/treasure/log"

	"github.com/alonegrowing/insight/pkg/basic/dao/impl"
)

func main() {
	tmp := impl.DefaultPoemDao.GetPoemById(context.TODO(), 1)
	fmt.Println(tmp)

	data := impl.DefaultPoemDao.GetPoemList(context.TODO())
	for _, item := range data {
		fmt.Println(item.Author)
	}

	log.Errorf("alone:%v", data)
}
