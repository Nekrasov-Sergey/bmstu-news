package config

import (
	"context"
)

//ключ к значению нашего конфига внутри параметров которые лежат внутри контекста
var configContextKey = struct{}{}

// FromContext Возвращает конфиг из контекста
func FromContext(ctx context.Context) *Config {
	cfgRaw := ctx.Value(configContextKey)
	cfg, ok := cfgRaw.(*Config)
	if ok {
		return cfg
	}
	return nil
}

// WrapContext Обогащает контекст конфигом
func WrapContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, configContextKey, cfg)
}
