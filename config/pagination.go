package config

import "goblog/pkg/config"

func init() {
	config.Add("pagination", config.MapStr{
		"perpage":   10,
		"url_query": "p",
	})
}
