package conversions

import (
	"time"
)

type SegmentDuration time.Duration

func (d *SegmentDuration) ParseDuration() float32 {
	// time.Duration
	return float32(*d) / 1000.0
}

func (l *SegmentDuration) String() string { return "" }
