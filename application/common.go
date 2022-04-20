package application

import "github.com/rs/xid"

func generatorID() string {
	return xid.New().String()
}
