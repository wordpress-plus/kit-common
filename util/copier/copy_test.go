package copier

import (
	"fmt"
	"github.com/micro-services-roadmap/kit-common/util"
	"testing"
	"time"
)

// TestB not work for field directly
func TestB(t *testing.T) {
	now := time.Now()
	var dst string

	// Perform the copy with custom logic
	if err := CopyWithTime(&dst, &now); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Destination: %+v\n", dst)
	}
}

type Source struct {
	Name          string
	CreatedAt     time.Time
	CreatedAtZero time.Time
	UpdatedAt     *time.Time
	UpdatedAtNil  *time.Time
}

type Destination struct {
	Name          string
	CreatedAt     string
	UpdatedAt     *string
	UpdatedAtNil  string
	CreatedAtZero *string
}

func Test2Str(t *testing.T) {

	now := time.Now()
	src := Source{
		Name:      "Example",
		CreatedAt: now,
		UpdatedAt: &now,
	}
	dst := Destination{}

	// Perform the copy with custom logic
	if err := CopyWithTime(&dst, &src); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Destination: %+v\n", dst)
	}
}

func Test2Time(t *testing.T) {

	now := util.FormatVal(time.Now())
	src := Destination{
		Name:      "Example",
		CreatedAt: now,
		UpdatedAt: &now,
	}
	dst := Source{}
	// Perform the copy with custom logic
	if err := CopyWithTime(&dst, &src); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Destination: %+v\n", dst)
	}
}
