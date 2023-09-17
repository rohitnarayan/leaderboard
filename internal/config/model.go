package config

import "time"

type Config struct {
	Server *ServerConfig
}

type ServerConfig struct {
	Name         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}
