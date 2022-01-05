// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/units"
)

var (
        fujiGenesisConfigJSON = `{
                "networkID": 5,
                "allocations": [
                        {
                                "ethAddr": "0x1a0e1c472f1693bf218e6561f83f5985f4edd7e7",
                                "avaxAddr": "X-fuji196zap8u8uw3hp7amhnrdr5a3ccwhwce476ygzv",
                                "initialAmount": 0,
                                "unlockSchedule": [
                                        {
                                                "amount": 40000000000000000
                                        }
                                ]
                        },
                        {
                                "ethAddr": "0xa72a690fe37ddaf9a10c5bd83db772a25b35e7b6",
                                "avaxAddr": "X-fuji1vag7ck7uagf3h7y784sa8ujxu5z8ct5hr9d4ud",
                                "initialAmount": 32000000000000000,
                                "unlockSchedule": []
                        },
                        {
                                "ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
                                "avaxAddr": "X-fuji1wrv92qg5x3dsqrtukdc8qxnpqust3qdakxgm4s",
                                "initialAmount": 32000000000000000,
                                "unlockSchedule": []
                        },
                        {
                                "ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
                                "avaxAddr": "X-fuji1qrmj7u9pquyy3mahzxeq0nnlnj2aceedjfqqrq",
                                "initialAmount": 32000000000000000,
                                "unlockSchedule": []
                        },
                        {
                                "ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
                                "avaxAddr": "X-fuji1cap3ru2ghc3jtdnuyey738ru8u5ekdadcvrtyk",
                                "initialAmount": 32000000000000000,
                                "unlockSchedule": []
                        },
                        {
                                "ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
                                "avaxAddr": "X-fuji18g2m7483k6swe46cpfmq96t09sp63pgv7judr4",
                                "initialAmount": 32000000000000000,
                                "unlockSchedule": []
                        },
                        {
                                "ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
                                "avaxAddr": "X-fuji1zwe0kxhg73x3ehgtkkz24k9czlfgztc45hgrg3",
                                "initialAmount": 32000000000000000,
                                "unlockSchedule": []
                        },
                        {
                                "ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
                                "avaxAddr": "X-fuji1fqcs4m9p8gdp7gckk30n8u68d55jk0hdumx30f",
                                "initialAmount": 32000000000000000,
                                "unlockSchedule": []
                        },
                        {
                                "ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
                                "avaxAddr": "X-fuji18lany6fjlzxc7vuqfd9x4k9wqp0yhk074p283d",
                                "initialAmount": 32000000000000000,
                                "unlockSchedule": []
                        },
                        {
                                "ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
                                "avaxAddr": "X-fuji18lany6fjlzxc7vuqfd9x4k9wqp0yhk074p283d",
                                "initialAmount": 32000000000000000,
                                "unlockSchedule": []
                        },
                        {
                                "ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
                                "avaxAddr": "X-fuji10d2fqjfl3ghl73z2ez65ufanxwwhccxugq8z2t",
                                "initialAmount": 32000000000000000,
                                "unlockSchedule": []
                        }
                ],
                "startTime": 1641371839,
                "initialStakeDuration": 31536000,
                "initialStakeDurationOffset": 5400,
                "initialStakedFunds": [
                        "X-fuji196zap8u8uw3hp7amhnrdr5a3ccwhwce476ygzv"
                ],
                "initialStakers": [
                        {
                                "nodeID": "NodeID-JuZmGx2NhZsjSsrmCsrSbCUezYELj34uC",
                                "rewardAddress": "X-fuji196zap8u8uw3hp7amhnrdr5a3ccwhwce476ygzv",
                                "delegationFee": 1000000
                        },
                        {
                                "nodeID": "NodeID-5UMaQhxhrzet968foe95VLp8KKWwT3tLk",
                                "rewardAddress": "X-fuji196zap8u8uw3hp7amhnrdr5a3ccwhwce476ygzv",
                                "delegationFee": 500000
                        },
                        {
                                "nodeID": "NodeID-LkiLFbmocMydnZ45jJgfoeLVTMN4nn83h",
                                "rewardAddress": "X-fuji196zap8u8uw3hp7amhnrdr5a3ccwhwce476ygzv",
                                "delegationFee": 250000
                        },
                        {
                                "nodeID": "NodeID-FRouddSdqsz9SqFddmcpX3kMcTUuczYWW",
                                "rewardAddress": "X-fuji196zap8u8uw3hp7amhnrdr5a3ccwhwce476ygzv",
                                "delegationFee": 125000
                        },
                        {
                                "nodeID": "NodeID-LmktUgqWHu3cDdLAAeCaLGabgbpoCwBz5",
                                "rewardAddress": "X-fuji196zap8u8uw3hp7amhnrdr5a3ccwhwce476ygzv",
                                "delegationFee": 62500
                        },
                        {
                                "nodeID": "NodeID-Bigfsv6ru1ZhEu6tBTofPizNYKfzhjraq",
                                "rewardAddress": "X-fuji196zap8u8uw3hp7amhnrdr5a3ccwhwce476ygzv",
                                "delegationFee": 31250
                        },
                        {
                                "nodeID": "NodeID-9a3Vo6ufbhvMzfu41gyJ9DDGGf34BSm8n",
                                "rewardAddress": "X-fuji196zap8u8uw3hp7amhnrdr5a3ccwhwce476ygzv",
                                "delegationFee": 20000
                        }
                ],
		"cChainGenesis": "{\"config\":{\"chainId\":43113,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"a72a690fe37ddaf9a10c5bd83db772a25b35e7b6\":{\"balance\":\"0x295BE96E64066972000000\"},\"0100000000000000000000000000000000000000\":{\"code\":\"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\",\"balance\":\"0x295BE96E64066972000000\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",

                "message": "hi mom"
        }`
	// FujiParams are the params used for the fuji testnet
	FujiParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliAvax,
			CreateAssetTxFee:      10 * units.MilliAvax,
			CreateSubnetTxFee:     100 * units.MilliAvax,
			CreateBlockchainTxFee: 100 * units.MilliAvax,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement:  .8, // 80%
			MinValidatorStake:  1 * units.Avax,
			MaxValidatorStake:  3 * units.MegaAvax,
			MinDelegatorStake:  1 * units.Avax,
			MinDelegationFee:   20000, // 2%
			MinStakeDuration:   24 * time.Hour,
			MaxStakeDuration:   365 * 24 * time.Hour,
			StakeMintingPeriod: 365 * 24 * time.Hour,
		},
	}
)
