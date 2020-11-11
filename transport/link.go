package transport

import "v2ray.com/core/common/buf"

// Link is a utility for connecting between an inbound and an outbound proxy handler.
type Link struct {
	Reader buf.Reader
	Writer buf.Writer
}


/*
	dependency:
		common/buf
	defination:
		Link
			Reader
				type:
					buf.Reader
			Writer
				type:
					buf.Writer
*/