package viperx

import (
	"fmt"
	kg "github.com/wordpress-plus/kit-common/kg"
	"testing"
)

func TestInitViper(_ *testing.T) {
	InitViper("config.yaml")
	fmt.Printf("kg.V: %v", kg.V)
	fmt.Printf("kg.C: %v", kg.C)
}
