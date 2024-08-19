package biz

import "myBlog/internal/myBlog/store"

type IBiz interface {
	Users() UserBiz
}

type biz struct {
	ds store.IStore
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

// Users 返回一个实现了 UserBiz 接口的实例.
func (b *biz) Users() UserBiz {
	return New(b.ds)
}
