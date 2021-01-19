package config

import "goblog/pkg/config"

func init() {
	config.Add("session", config.MapStr{
		"default":      config.Env("SESSION_DIRVER", "cookie"),
		"session_name": config.Env("SESSION_NAME", "session-id"),
	})
}
