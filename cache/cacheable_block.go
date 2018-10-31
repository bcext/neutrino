package cache

import (
	"github.com/bcext/cashutil"
)

// CacheableBlock is a wrapper around the cashutil.Block type which provides a
// Size method used by the cache to target certain memory usage.
type CacheableBlock struct {
	*cashutil.Block
}

// Size returns size of this block in bytes.
func (c *CacheableBlock) Size() (uint64, error) {
	return uint64(c.Block.MsgBlock().SerializeSize()), nil
}
