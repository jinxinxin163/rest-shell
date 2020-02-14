/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/

package out

type OutputsResult struct {
	Status  int               `json:"status" description:"response status"`
	Error   string            `json:"error,omitempty" description:"debug information"`
	Outputs string 			  `json:"content,omitempty" description:"outputs string"`
}
