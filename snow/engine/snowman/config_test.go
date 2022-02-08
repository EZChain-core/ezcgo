// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/EZChain-core/ezcgo/database/memdb"
	"github.com/EZChain-core/ezcgo/snow/consensus/snowball"
	"github.com/EZChain-core/ezcgo/snow/consensus/snowman"
	"github.com/EZChain-core/ezcgo/snow/engine/common"
	"github.com/EZChain-core/ezcgo/snow/engine/common/queue"
	"github.com/EZChain-core/ezcgo/snow/engine/snowman/block"
	"github.com/EZChain-core/ezcgo/snow/engine/snowman/bootstrap"
)

func DefaultConfig() Config {
	blocked, _ := queue.NewWithMissing(memdb.New(), "", prometheus.NewRegistry())
	return Config{
		Config: bootstrap.Config{
			Config:  common.DefaultConfigTest(),
			Blocked: blocked,
			VM:      &block.TestVM{},
		},
		Params: snowball.Parameters{
			K:                     1,
			Alpha:                 1,
			BetaVirtuous:          1,
			BetaRogue:             2,
			ConcurrentRepolls:     1,
			OptimalProcessing:     100,
			MaxOutstandingItems:   1,
			MaxItemProcessingTime: 1,
		},
		Consensus: &snowman.Topological{},
	}
}
