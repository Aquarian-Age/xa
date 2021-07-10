package pub

import (
	"fmt"
	"github.com/Aquarian-Age/xa/pkg/gz"
	"testing"
)

func TestZhis(t *testing.T) {
	arr := gz.GetJzArr()
	for i := 0; i < len(arr); i++ {
		gzs := arr[i]
		z := GetZhiS(gzs)
		fmt.Printf("%s-->%s ", arr[i], z)
	}

}
