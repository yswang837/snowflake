package snowflake

import (
	"fmt"
	"testing"
)

func TestGenID(t *testing.T) {
	startTime := "2022-07-16"
	machineId := "1"
	if err := Init(startTime, machineId); err != nil {
		return
	}
	fmt.Println(GenID("AA"))
}
