package main

import (
	"github.com/smcduck/xnet/xaddr"
	"github.com/perlin-network/noise/crypto/ed25519"
	"github.com/perlin-network/noise/network"
	"github.com/perlin-network/noise/network/discovery"
	"github.com/pkg/errors"
	"github.com/perlin-network/noise/crypto"
)

type (
	Bitq struct {
		net *network.Network
		keypair *crypto.KeyPair
	}
)

func NewBitq(listen string) (*Bitq, error) {
	ip, port, err := xaddr.ParseHostAddrOnline(listen)
	if err != nil {
		return nil, errors.Wrap(err, "ParseHostAddrOnline()")
	}

	keys := ed25519.RandomKeyPair()

	builder := network.NewBuilder()
	builder.SetKeys(keys)
	builder.SetAddress(network.FormatAddress("xkcp", ip.String(), uint16(port)))

	// Register peer discovery plugin.
	builder.AddPlugin(new(discovery.Plugin))

	// Add custom message queue plugin.
	builder.AddPlugin(new(ChatPlugin))

	net, err := builder.Build()
	if err != nil {
		return nil, errors.Wrap(err, "Build()")
	}

	return &Bitq{net:net, keypair:keys}, nil
}

func (b *Bitq) KeyPair() (public, private []byte) {
	return b.keypair.PublicKey, b.keypair.PrivateKey
}

func (b *Bitq) Bootstrap(peerAddr ...string) {
	go b.net.Listen()

	if len(peerAddr) > 0 {
		b.net.Bootstrap(peerAddr...)
	}
}

func (b *Bitq) Peers() ([]string, error) {
	return nil, nil
}

func (b *Bitq) CacheSize() (int, error) {
	return 0, nil
}

func (b *Bitq) Flush() error {
	return nil
}

func (b *Bitq) Close() error {
	return nil
}