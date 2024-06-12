package peakecs

// Entity Manager

type EntityManager struct {
	Entities     []*Entity
	NextEntityID chan int
}

type Entity struct {
	ID         int
	Masked     uint64
	Components []Component
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		Entities:     make([]*Entity, 0),
		NextEntityID: make(chan int),
	}
}

func (em *EntityManager) CreateEntity() *Entity {
	entity := &Entity{ID: <-em.NextEntityID}
	em.Entities = append(em.Entities, entity)
	return entity
}

func (em *EntityManager) Run() {
	for i := 0; ; i++ {
		em.NextEntityID <- i
	}
}

func (em *EntityManager) RemoveEntity(entity *Entity) {
	for i, e := range em.Entities {
		if e == entity {
			em.Entities = append(em.Entities[:i], em.Entities[i+1:]...)
			break
		}
	}
}

func (e *Entity) Mask() uint64 {
	return e.Masked
}

func maskSlice(components []Component) uint64 {
	mask := uint64(0)
	for _, c := range components {
		mask = mask | c.Mask()
	}
	return mask
}

func (entity *Entity) AddComponents(components ...Component) {
	for _, c := range components {
		if entity.Masked&c.Mask() == c.Mask() {
			continue
		}
		entity.Components = append(entity.Components, c)
		entity.Masked = maskSlice(entity.Components)
	}

}

func (e *Entity) Get(mask uint64) Component {
	for _, c := range e.Components {
		if c.Mask() == mask {
			return c
		}
	}
	return nil
}

func (e *Entity) Remove(mask uint64) {
	modified := false
	for i, c := range e.Components {
		if c.Mask() == mask {
			copy(e.Components[i:], e.Components[i+1:])
			e.Components[len(e.Components)-1] = nil
			e.Components = e.Components[:len(e.Components)-1]
			modified = true
			break
		}
	}
	if modified {
		e.Masked = maskSlice(e.Components)
	}
}

type Filter struct {
	Include uint64
	Exclude uint64
}

func (em *EntityManager) FilterByMask(f *Filter) (entities []*Entity) {
	entities = make([]*Entity, len(em.Entities))
	index := 0
	for _, e := range em.Entities {
		observed := e.Mask()
		if f.Exclude > 0 && observed&f.Exclude == f.Exclude {
			continue
		}
		if observed&f.Include == f.Include {
			entities[index] = e
			index++
		}
	}
	return entities[:index]
}
