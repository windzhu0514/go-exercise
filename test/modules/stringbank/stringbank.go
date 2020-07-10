package stringbank

import "go-exercise/test/modules"

func init() {
	modules.RegisterModule(module{})
}

type module struct{}

func (module) ModuleInfo() modules.ModuleInfo {
	return modules.ModuleInfo{
		Name: "string.bank",
		New: func() (modules.StringBank, error) {
			var sb Stringbank
			return &sb, nil
		},
		Release: func() error {
			return nil
		},
	}
}

type Stringbank struct {
	current     []byte
	allocations [][]byte
}

func (s *Stringbank) Size() int {
	return 100
}

// Get converts an index to the original string
func (s *Stringbank) Get(index int) string {
	return "sdfsdf"
}
