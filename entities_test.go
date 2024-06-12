package peakecs

import "testing"

func TestEntityAdd(t *testing.T) {
	em := NewEntityManager()
	go em.Run()
	em.CreateEntity()
	expected := 1
	output := len(em.Entities)
	if expected != output {
		t.Fatalf("Expected number of entities: %d, got %d", expected, output)
	}
}

func TestEntityRemove(t *testing.T) {
	em := NewEntityManager()
	go em.Run()
	e := em.CreateEntity()
	em.RemoveEntity(e)
	expected := 0
	output := len(em.Entities)
	if expected != output {
		t.Fatalf("Expected number of entities: %d, got %d", expected, output)
	}
}
