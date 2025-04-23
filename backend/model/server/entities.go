package server

import (
	"backend/database/gen"

	"github.com/oklog/ulid/v2"
)

type SeverType string

var ServerTypeEnum = struct {
	VANILLA    SeverType
	CUSTOM     SeverType
	CURSEFORGE SeverType
	FTBA       SeverType
	MODRINTH   SeverType
}{
	VANILLA:    "VANILLA",
	CUSTOM:     "CUSTOM",
	CURSEFORGE: "CURSEFORGE",
	FTBA:       "FTBA",
	MODRINTH:   "MODRINTH",
}


type Server struct {
	gen.Server
	ContainerId string `json:"containerId"`
}

type ServerConfigDTO struct {
	Type        SeverType `json:"type"`
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	MaxPlayers  int32     `json:"maxPlayers"`
	WhiteList   []string  `json:"whitelist"`
	Memory      int32     `json:"memory"`
	UseAikaFlag bool      `json:"useAikarFlag"`
}

type CreationServer struct {
	ID ulid.ULID `json:"id"`
}
