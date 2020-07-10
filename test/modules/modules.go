package modules

import (
	"fmt"
	"sync"
)

type ModuleInfo struct {
	Name    string
	New     func() (StringBank, error) // 返回模块 是否是单例由模块决定
	Release func() error
}

type Module interface {
	ModuleInfo() ModuleInfo
}

type StringBank interface {
	Size() int
	Get(index int) string
}

func RegisterModule(instance Module) {
	mod := instance.ModuleInfo()
	if mod.Name == "" {
		panic("module Name missing")
	}

	if mod.New == nil {
		panic("missing ModuleInfo.New")
	}

	modulesMu.Lock()
	defer modulesMu.Unlock()

	if _, ok := modules[mod.Name]; ok {
		panic(fmt.Sprintf("module already registered: %s", mod.Name))
	}
	modules[mod.Name] = mod
}

func GetModule(id string) (ModuleInfo, error) {
	modulesMu.RLock()
	defer modulesMu.RUnlock()

	m, ok := modules[id]
	if !ok {
		return ModuleInfo{}, fmt.Errorf("module not registered: %s", id)
	}
	return m, nil
}

var (
	modules   = make(map[string]ModuleInfo)
	modulesMu sync.RWMutex
)
