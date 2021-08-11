package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

// 定义对象
type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

// 构造函数
func NewMemorySessionMgr() SessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}

func (s *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

// 创建session
func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	// 用uuid作为sessionId
	id := uuid.NewV4()
	// 转string
	sessionId := id.String()
	// 创建个session
	memorySession := NewMemorySession(sessionId)
	// 加入到大map
	s.sessionMap[sessionId] = session
	return memorySession, nil
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}

	return
}
