package queuing

import "skynx.io/s-api-go/grpc/network/sxsp"

var RxControlQueue = make(chan *sxsp.Payload, 128)
var TxControlQueue = make(chan *sxsp.Payload, 128)
