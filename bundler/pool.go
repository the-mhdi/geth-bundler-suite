package bundler

import (
	"sync"
)

type poolConf struct {
}

type Pool struct {
	config *poolConf
	queue  *sync.Pool
}
