package auth

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"time"
)

type Authz struct {
	*casbin.SyncedEnforcer
}

// NewAuthz 创建一个使用 casbin 完成授权的授权器.
func NewAuthz(db *gorm.DB) (*Authz, error) {
	// Initialize a Gorm dbAdapter and use it in a Casbin enforcer
	dbAdapter, err := adapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	m, _ := model.NewModelFromFile("pkg/auth/model.conf")
	// Initialize the enforcer.
	enforcer, err := casbin.NewSyncedEnforcer(m, dbAdapter)
	if err != nil {
		return nil, err
	}

	// Load the policy from DB.
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	//将policy定时从数据库同步至缓存，以提高授权速度
	enforcer.StartAutoLoadPolicy(5 * time.Second)

	a := &Authz{enforcer}

	return a, nil
}

// Authorize 用来进行授权.
func (a *Authz) Authorize(sub, obj, act string) (bool, error) {
	return a.Enforce(sub, obj, act)
}
