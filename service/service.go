package service

import (
	"changeme/common"
	"sync"
)

type Service struct {
	BackgroundJobsStarted bool
	ServiceMutex          sync.Mutex
	StartTimeMs           int64
	GenericCache          GenericCache // use interface to allow injection
	Config                common.Config
}

type GenericCache interface {
	Get(cacheKey string) ([]byte, bool)
	Set(cacheKey string, value []byte)
}

func NewService(config common.Config) *Service {
	ser := &Service{
		BackgroundJobsStarted: false,
		ServiceMutex:          sync.Mutex{},
		StartTimeMs:           common.MakeTimestampMilli(),
		GenericCache:          NewMapCache(),
		Config:                config,
	}

	ser.StartBackground()
	return ser
}

func (ser *Service) StartBackground() {
	ser.ServiceMutex.Lock()
	defer ser.ServiceMutex.Unlock()

	ser.BackgroundJobsStarted = true

	// trigger background jobs here
}
