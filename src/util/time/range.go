package time

import (
	"fmt"
	"strings"
	"time"
)

const rangeSeparator = "-"

var EmptyRange = Range{}

type Range struct {
	Start time.Time
	End   time.Time
}

func ParseRange(t string) (Range, error) {
	parts := strings.Split(t, rangeSeparator)
	if len(parts) != 2 {
		return EmptyRange, fmt.Errorf("time range should consist of 2 parts")
	}

	start, err := time.Parse(HourMinuteLayout, parts[0])
	if err != nil {
		return EmptyRange, err
	}
	end, err := time.Parse(HourMinuteLayout, parts[1])
	if err != nil {
		return EmptyRange, err
	}
	return Range{
		Start: start,
		End:   end,
	}, nil
}

func MustParseRange(t string) Range {
	out, _ := ParseRange(t)
	return out
}

func (t Range) String() string {
	if t == EmptyRange {
		return ""
	}
	return fmt.Sprintf(
		"%s-%s",
		t.Start.Format(HourMinuteLayout),
		t.End.Format(HourMinuteLayout),
	)
}

func (t Range) Less(other Range) bool {
	return t.Start.Before(other.Start) || (t.Start.Equal(other.Start) && t.End.Before(other.End))
}

func (t Range) Intersects(other Range) bool {
	return t.Start.Equal(other.End) || other.Start.Equal(t.End) ||
		t.Start.Before(other.End) && other.Start.Before(t.End)
}
