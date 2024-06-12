package peakecs

type Component interface {
	Mask() uint64
	Name() string
}
