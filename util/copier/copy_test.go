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

func TestWithMapAny(t *testing.T) {
	from := struct {
		MapField    map[string]interface{} //  map[string]string is invalid
		StringField string
	}{
		MapField:    map[string]interface{}{"key": "value"},
		StringField: "{\n    \"key\": \"value\"\n}",
	}

	to := struct {
		MapField    string
		StringField map[string]interface{}
	}{}

	err := CopyWithTime(&to, &from)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %+v\n", to)
	}
}

func TestWithMapString(t *testing.T) {
	from := struct {
		MapField    map[string]string //  map[string]string is invalid
		StringField string
	}{
		MapField:    map[string]string{"key": "value"},
		StringField: "{\n    \"key\": \"value\"\n}",
	}

	to := struct {
		MapField    string
		StringField map[string]string
	}{}

	err := CopyWithTime(&to, &from)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %+v\n", to)
	}
}
