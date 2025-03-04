package bundler

import (
	"sync"
)

type poolConf struct {
	poolType string //public - private

}

type Pool struct {
	config *poolConf
	queue  *sync.Pool
}
