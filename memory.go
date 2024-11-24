package polkavm

type MemorySection struct {
}

type Memory struct {
	sections []*MemorySection
}

func (m *Memory) Write(addr uint32, val []byte) error {
	return nil
}

func (m *Memory) Read(addr uint32) ([]byte, error) {
	return nil, nil
}
