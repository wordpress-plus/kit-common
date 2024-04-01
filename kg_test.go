package kit

import (
	"fmt"
	"github.com/wordpress-plus/kit-common/gormx/source/gen/dal"
	g "github.com/wordpress-plus/kit-common/kg"
	"testing"
)

func TestInit(t *testing.T) {
	Init("config_test.yaml")
	fmt.Printf("zap: %v", g.L)

	dal.SetDefault(g.DB)
	find, err := dal.Q.ArchivedUpsTag.Find()
	if err != nil || find != nil {
		return
	}
}
