package minecraft

import (
	"github.com/Psychz/GoLilyPad/packet"
)

type Version struct {
	Name             string
	LoginClientCodec packet.PacketPipelineChild
	LoginServerCodec packet.PacketPipelineChild
	PlayClientCodec  packet.PacketPipelineChild
	PlayServerCodec  packet.PacketPipelineChild
	IdMap            *IdMap
}
