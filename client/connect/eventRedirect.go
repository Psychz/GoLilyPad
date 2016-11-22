package connect

import (
	"github.com/Psychz/GoLilyPad/packet/connect"
)

type EventRedirect struct {
	Server string
	Player string
}

func WrapEventRedirect(packet *connect.PacketRedirectEvent) (this *EventRedirect) {
	this = new(EventRedirect)
	this.Server = packet.Server
	this.Player = packet.Player
	return
}
