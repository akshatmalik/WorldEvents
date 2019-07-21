package main

import (
	"world_events/common"
)

func main() {

	var err error

	err = common.Init()
	if err != nil {
		common.Log.Infof("Init failed")
		return
	}

}
