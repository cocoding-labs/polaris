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

package state

import (
	"github.com/berachain/stargazer/eth/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StorageKeyFor", func() {
	It("returns a prefix to iterate over a given account storage", func() {
		address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
		prefix := StorageKeyFor(address)
		Expect(prefix).To(HaveLen(1 + common.AddressLength))
		Expect(prefix[0]).To(Equal(keyPrefixStorage))
		Expect(prefix[1:]).To(Equal(address.Bytes()))
	})
})

var _ = Describe("SlotKeyFor", func() {
	It("returns a storage key for a given account and storage slot", func() {
		address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
		slot := common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
		key := SlotKeyFor(address, slot)
		Expect(key).To(HaveLen(1 + common.AddressLength + common.HashLength))
		Expect(key[0]).To(Equal(keyPrefixStorage))
		Expect(key[1 : 1+common.AddressLength]).To(Equal(address.Bytes()))
		Expect(key[1+common.AddressLength:]).To(Equal(slot.Bytes()))
	})
})

var _ = Describe("CodeHashKeyFo or a given account", func() {
	address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	key := CodeHashKeyFor(address)
	Expect(key).To(HaveLen(1 + common.AddressLength))
	Expect(key[0]).To(Equal(keyPrefixCode))
	Expect(key[1:]).To(Equal(address.Bytes()))
})

var _ = Describe("CodeKeyFor", func() {
	It("returns a code key for a given account", func() {
		address := common.HexToHash("0x1234567890abcdef1234567890abcdef12345678")
		key := CodeKeyFor(address)
		Expect(key).To(HaveLen(1 + common.HashLength))
		Expect(key[0]).To(Equal(keyPrefixCode))
		Expect(key[1:]).To(Equal(address.Bytes()))
	})
})
