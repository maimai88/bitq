package main

import (
	"time"
)

type Replyer struct {

}

type ReplyerCb func(service string, requestData []byte) (replyData []byte)

func (b *Bitq) Reply(service string, cb ReplyerCb) (*Replyer, error) {
	return nil, nil
}

// Sync.
func (b *Bitq) Request(service string, data []byte, timeout time.Duration) (replyData []byte, err error) {
	return nil, nil
}

func (r *Replyer) Close() error {
	return nil
}