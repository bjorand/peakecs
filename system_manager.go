package peakecs

type SystemManager struct {
	Systems []System
	Em      *EntityManager
}

func NewSystemManager(em *EntityManager) *SystemManager {
	return &SystemManager{
		Systems: make([]System, 0),
		Em:      em,
	}
}

func (sm *SystemManager) Add(systems ...System) {
	sm.Systems = append(sm.Systems, systems...)
}

func (sm *SystemManager) Update() {
	for _, s := range sm.Systems {
		s.Update(sm.Em)
	}
}
