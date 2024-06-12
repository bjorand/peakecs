package peakecs

import "testing"

type SystemTest struct {
}

func NewSystemTest() System {
	return &SystemTest{}
}

func (s *SystemTest) Update(em *EntityManager) {
	for _, e := range em.FilterByMask(
		&Filter{
			Include: TestComponentMask},
	) {
		c := e.Get(TestComponentMask).(*TestComponent)
		c.A = "foobar2"

	}
}

func TestSystem(t *testing.T) {
	em := NewEntityManager()
	go em.Run()
	e1 := em.CreateEntity()
	e1.AddComponents(
		NewTestComponent("foobar"),
	)
	sm := NewSystemManager(em)
	sm.Add(
		NewSystemTest(),
	)
	sm.Update()
	expected := "foobar2"
	output := e1.Components[0].(*TestComponent).A
	if expected != output {
		t.Fatalf("Expected number of entities: %s, got %s", expected, output)
	}
}
