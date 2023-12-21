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

package txpool

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/berachain/polaris/cosmos/runtime/txpool/mocks"
	evmtypes "github.com/berachain/polaris/cosmos/x/evm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ethtypes "github.com/ethereum/go-ethereum/core/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("", func() {

	var (
		t       = GinkgoT()
		txPool  *mocks.GethTxPool
		sdkTx   *mocks.SdkTx
		mempool *Mempool
		ctx     = sdk.Context{}.WithIsCheckTx(true)
		wet     *evmtypes.WrappedEthereumTransaction
	)

	BeforeEach(func() {
		txPool = mocks.NewGethTxPool(t)
		sdkTx = mocks.NewSdkTx(t)
		mempool = &Mempool{txpool: txPool}
		wet, _ = evmtypes.WrapTx(ethtypes.NewTx(&ethtypes.LegacyTx{}))
	})

	When("we call insert", func() {
		When("the txpool does not error", func() {
			It("does not error", func() {
				sdkTx.On("GetMsgs").
					Return([]sdk.Msg{wet}).
					Once()
				txPool.On("Add", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Once()
				Expect(mempool.Insert(ctx, sdkTx)).ToNot(HaveOccurred())
			})
		})
		When("the txpool errors", func() {
			It("does error", func() {
				sdkTx.On("GetMsgs").Return(
					[]sdk.Msg{wet}).Once()
				txPool.On("Add",
					mock.Anything, mock.Anything, mock.Anything).
					Return([]error{errors.New("mock error")}).Once()
				Expect(mempool.Insert(ctx, sdkTx)).To(HaveOccurred())
			})
		})
		When("we use an sdkTx with no messages", func() {
			It("errors", func() {
				sdkTx.On("GetMsgs").Return([]sdk.Msg{}).Once()
				Expect(mempool.Insert(ctx, sdkTx)).To(HaveOccurred())
			})
		})
		When("we use an that is not an ethereum msg", func() {
			It("does not error", func() {
				sdkTx.On("GetMsgs").Return([]sdk.Msg{nil}).Once()
				Expect(mempool.Insert(ctx, sdkTx)).ToNot(HaveOccurred())
			})
		})
	})

	When("we call stats", func() {
		It("returns", func() {
			txPool.On("Stats").Return(16, 12).Once()
			Expect(mempool.CountTx()).To(Equal(28))
		})
	})
})
