package connect

import (
	"github.com/Psychz/GoLilyPad/packet"
	"github.com/Psychz/GoLilyPad/packet/connect"
)

func NewCodecRegistry(connectClient Connect) (codec *packet.PacketCodecRegistry) {
	registryCodec := connect.PacketCodec.Copy()
	registryCodec.DecodeCodecs[connect.PACKET_RESULT] = connect.NewPacketResultCodec(connectClient)
	codec = registryCodec
	return
}
