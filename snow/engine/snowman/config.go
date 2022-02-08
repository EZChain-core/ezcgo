// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/EZChain-core/ezcgo/snow/consensus/snowball"
	"github.com/EZChain-core/ezcgo/snow/consensus/snowman"
	"github.com/EZChain-core/ezcgo/snow/engine/snowman/bootstrap"
)

// Config wraps all the parameters needed for a snowman engine
type Config struct {
	bootstrap.Config

	Params    snowball.Parameters
	Consensus snowman.Consensus
}
