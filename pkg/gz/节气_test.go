package gz

import (
	"fmt"
	"github.com/starainrt/astro/calendar"
	"testing"
)

func TestGanZhi_JieQi(t *testing.T) {
	year := 2022
	lichu, jieqi, zhongqi := getJie12T(year)
	fmt.Println(lichu)
	fmt.Println(jieqi)
	fmt.Println(zhongqi)

	lichut := calendar.JieQi(year, calendar.JQ_立春)
	fmt.Println(lichut)
}
