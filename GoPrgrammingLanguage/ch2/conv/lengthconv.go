package conv

import "fmt"

type Foot float64
type Meter float64

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%g ft", f)
}
