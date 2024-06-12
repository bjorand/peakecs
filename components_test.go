package peakecs

import "testing"

const (
	TestComponentMask uint64 = 1 << iota
	Test2ComponentMask
)

type TestComponent struct {
	A string
}

type Test2Component struct {
}

func (c *TestComponent) Mask() uint64 {
	return TestComponentMask
}

func (c *TestComponent) Name() string {
	return "test"
}

func (c *Test2Component) Mask() uint64 {
	return Test2ComponentMask
}

func (c *Test2Component) Name() string {
	return "test2"
}

func NewTestComponent(a string) Component {
	return &TestComponent{A: a}
}

func NewTest2Component() Component {
	return &Test2Component{}
}

func TestComponents(t *testing.T) {
	em := NewEntityManager()
	go em.Run()
	e1 := em.CreateEntity()
	e1.AddComponents(
		NewTestComponent("foobar"),
	)
	expected := 1
	output := len(e1.Components)
	if expected != output {
		t.Fatalf("Expected number of components: %d, got %d", expected, output)
	}
	e1.AddComponents(
		NewTest2Component(),
	)
	expected = 2
	output = len(e1.Components)
	if expected != output {
		t.Fatalf("Expected number of components: %d, got %d", expected, output)
	}
	e2 := em.CreateEntity()
	e2.AddComponents(
		NewTestComponent("foobar"),
	)
	e3 := em.CreateEntity()
	e3.AddComponents(
		NewTest2Component(),
	)
	entities := em.FilterByMask(
		&Filter{
			Include: TestComponentMask,
		},
	)
	expected = 2
	output = len(entities)
	if expected != output {
		t.Fatalf("Expected number of entities: %d, got %d", expected, output)
	}
	entities = em.FilterByMask(
		&Filter{
			Include: TestComponentMask,
			Exclude: Test2ComponentMask,
		},
	)
	expected = 1
	output = len(entities)
	if expected != output {
		t.Fatalf("Expected number of entities: %d, got %d", expected, output)
	}
	entities = em.FilterByMask(
		&Filter{
			Include: TestComponentMask | Test2ComponentMask,
		},
	)
	expected = 1
	output = len(entities)
	if expected != output {
		t.Fatalf("Expected number of entities: %d, got %d", expected, output)
	}
	entities = em.FilterByMask(
		&Filter{
			Include: TestComponentMask & Test2ComponentMask,
		},
	)
	expected = 3
	output = len(entities)
	if expected != output {
		t.Fatalf("Expected number of entities: %d, got %d", expected, output)
	}

}
