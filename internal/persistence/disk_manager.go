package persistence

type DiskManager struct {
	dataVolumePath string
}

func NewDiskManager(dataVolumePath string) *DiskManager {
	return &DiskManager{
		dataVolumePath: dataVolumePath,
	}
}

func (manager *DiskManager) Init() {
	
}
