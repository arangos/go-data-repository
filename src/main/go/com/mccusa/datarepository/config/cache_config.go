package config

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// Cache is a global in-memory cache with default expiration and cleanup interval.
var Cache = cache.New(5*time.Minute, 15*time.Minute)
