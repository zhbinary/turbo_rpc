//Created by zhbinary on 2019-07-31.
//Email: zhbinary@gmail.com
package core

import "turbo_raft/types"

// Idempotence
type XClient struct {
	// Many channels per client
	channels []types.Channel
}

func (this *XClient) SendOneWay(serviceName string, methodName string, msg struct{}) (err error) {

}

func (this *XClient) SendRPCSync(methodName string, msg struct{}) (err error) {

}

func (this *XClient) SendRPCAsync(methodName string, msg struct{}, cb func()) (err error) {

}
