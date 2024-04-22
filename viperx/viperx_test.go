package viperx

import (
	"fmt"
	"testing"

	kg "github.com/micro-services-roadmap/kit-common/kg"
)

func TestInitViper(_ *testing.T) {
	InitViper("config.yaml")
	fmt.Printf("kg.V: %v", kg.V)
	fmt.Printf("kg.C: %v", kg.C)
}
