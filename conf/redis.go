package conf

import (
	"Dorm/cache"
)

func redis() {
	cache.Init()
}
