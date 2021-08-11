package session

import "fmt"

// 中间件让用户去选择使用哪个版本

var (
	sessionMgr SessionMgr
)

func Init(provider, addr string, options ...string) (err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedissionMgr()
	default:
		fmt.Errorf("不支持")
		return
	}
	err = sessionMgr.Init(addr, options...)
	return
}
