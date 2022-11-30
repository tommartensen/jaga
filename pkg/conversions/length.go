package conversions

import "fmt"

type Length float32

func (l *Length) ToKilometers() float32 {
	return float32(*l) / 1000.0
}

func (l *Length) String() string {
	kilometers := l.ToKilometers()
	return fmt.Sprintf("%.2fkm", kilometers)
}
