/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/

package signals

import (
	"os"
)

var shutdownSignals = []os.Signal{os.Interrupt}
