package server

import (
	api "core-auth/internal/app/apis"
)

// NewServer Return new server instance
func NewServer() {
	api.CreateAPIApp()
	// server := core.NewServer(agApp, &(miscApp{}))
	// _ = server.Init()
	// return server
}
