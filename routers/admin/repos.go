// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	"math"

	"github.com/gogits/gogs/models"
	"github.com/gogits/gogs/modules/base"
	"github.com/gogits/gogs/modules/middleware"
)

const (
	REPOS base.TplName = "admin/repo/list"
)

func pagination(ctx *middleware.Context, count int64, pageNum int) int {
	p := ctx.QueryInt("p")
	if p < 1 {
		p = 1
	}
	curCount := int64((p-1)*pageNum + pageNum)
	if curCount >= count {
		p = int(math.Ceil(float64(count) / float64(pageNum)))
	} else {
		ctx.Data["NextPageNum"] = p + 1
	}
	if p > 1 {
		ctx.Data["LastPageNum"] = p - 1
	}
	return p
}

func Repositories(ctx *middleware.Context) {
	ctx.Data["Title"] = ctx.Tr("admin.repositories")
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminRepositories"] = true

	pageNum := 50
	p := pagination(ctx, models.CountRepositories(), pageNum)

	var err error
	ctx.Data["Repos"], err = models.GetRepositoriesWithUsers(pageNum, (p-1)*pageNum)
	if err != nil {
		ctx.Handle(500, "GetRepositoriesWithUsers", err)
		return
	}
	ctx.HTML(200, REPOS)
}
