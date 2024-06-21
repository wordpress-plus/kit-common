package ipx

import (
	"fmt"
	"testing"
)

func TestParseAgent(t *testing.T) {

	agent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"
	fmt.Println(ParseAgent(agent))
}
