package kit

import (
	"fmt"
	// "github.com/wordpress-plus/kit-common/gormx/source/g/dal"
	"github.com/wordpress-plus/kit-common/kg"
	"testing"
)

func TestInit(t *testing.T) {
	Init("config_test.yaml")
	fmt.Printf("zap: %v", kg.L)

	//dal.SetDefault(kg.DB)
	//find, err := dal.Q.ArchivedUpsTag.Find()
	//if err != nil || find != nil {
	//	return
	//}
}
