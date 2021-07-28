//Created by zhbinary on 2019-07-31.
//Email: zhbinary@gmail.com
package types

type Channel interface {
	Write([]byte)
	WriteAndFlush([]byte)
	Flush()
	Close()
	Read()
}
