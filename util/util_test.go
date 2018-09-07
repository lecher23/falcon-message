package util

import (
	"testing"
	"fmt"
)

func TestHandleContent(t *testing.T) {
	testS := "[P0][OK][ONLINE][][网站响应过慢 all(#2) web_check tag=time_cost 0.104>=2][O1 2018-09-07 11:22:00]"
	ctn := HandleContent(testS)
	fmt.Println(ctn)
}
