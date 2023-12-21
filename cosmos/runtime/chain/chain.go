// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package chain

import (
	"context"

	"github.com/berachain/polaris/eth/core"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// WrappedBlockchain is a struct that wraps the core blockchain with additional
// application context.
type WrappedBlockchain struct {
	core.Blockchain     // chain is the core blockchain.
	app             App // App is the application context.

}

// New creates a new instance of WrappedBlockchain with the provided core blockchain
// and application context.
func New(chain core.Blockchain, app App) *WrappedBlockchain {
	return &WrappedBlockchain{Blockchain: chain, app: app}
}

// WriteGenesisState writes the genesis state to the blockchain.
// It uses the provided context as the application context.
func (wbc *WrappedBlockchain) WriteGenesisState(
	ctx context.Context, genState *core.Genesis,
) error {
	wbc.PreparePlugins(ctx)
	return wbc.WriteGenesisBlock(genState.ToBlock())
}

// InsertBlockAndSetHead inserts a block into the blockchain and sets
// it as the head. It uses the provided context as the application context.
func (wbc *WrappedBlockchain) InsertBlockAndSetHead(
	ctx context.Context, block *ethtypes.Block,
) error {
	wbc.PreparePlugins(ctx)
	return wbc.Blockchain.InsertBlockAndSetHead(block)
}

// InsertBlockWithoutSetHead inserts a block into the blockchain without setting it
// as the head. It uses the provided context as the application context.
func (wbc *WrappedBlockchain) InsertBlockWithoutSetHead(
	ctx context.Context, block *ethtypes.Block,
) error {
	wbc.PreparePlugins(ctx)
	return wbc.Blockchain.InsertBlockWithoutSetHead(block)
}
