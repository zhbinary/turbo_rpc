//Created by zhbinary on 2019-07-31.
//Email: zhbinary@gmail.com
package core

import "turbo_raft/types"

// Inidempotence
type Client struct {
	// Just one channel per client
	channel types.Channel
}

func (this *Client) SendOneWay(msg struct{}) (err error) {

}

func (this *Client) SendRPCSync(msg struct{}) (err error) {

}

func (this *Client) SendRPCAsync(msg struct{}, cb func()) (err error) {

}
