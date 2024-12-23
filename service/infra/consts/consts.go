package consts

import "time"

var ALLOWSHAPES = []string{"quadrilateral", "circular", "triangle", "diamond", "pentagram", "petal", "love", "gear", "rabbit", "bear", "elephant", "cat", "deer", "cattle", "butterfly"}

const (
	SINGLE = "SINGLE"
)

// redis
const (
	REDIS_PRODUCT      = "product_"
	REDIS_CATEGOEY     = "category_"
	REDIS_DEFAULT_TIME = 300 * time.Second
)
