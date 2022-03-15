// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/platformvm/reward"
)

var (
	mainnetGenesisConfigJSON = `{
		"networkID": 1,
		"allocations": [
                        {
                                "ethAddr": "0x4CF9045563126432d58e8A89A932F378e6B41a08",
                                "avaxAddr": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "initialAmount": 417850000000000000,
                                "unlockSchedule": [
                                        {
                                                "amount": 150000000000000,
                                                "locktime": 1673840504
                                        }
                                ]
                        },
                        {
                                "ethAddr": "0x27A370Fc2AcD908ce7BD5E259638E90c6f6eb06F",
                                "avaxAddr": "X-ezc13hzpvfc957kus6ggfgfjg3hhzlt94te6wvndfr",
                                "initialAmount": 32000000000000000,
                                "unlockSchedule": []
                        }
                ],
		"startTime": 1642304504,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
                        "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju"
                ],
                "initialStakers": [
                        {
                                "nodeID": "NodeID-3YVrFDba5KwDYBCsMTWebZRFrN4FSWRvS",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-4YrMwoyo2ypQNhiEYGRMQmUB2QYTFxuTL",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-JeAspnL1KmYfFoAH1UdnghqaKSfqdd9ur",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-4dWu6veRSyLhmGv7dvCSCz7Pttgxv2F9e",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-MqjT6x5UgRnvLwb7RLtTGXTvHsSFed4QC",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-CavKfAxmQ66eqTMTSGMLWPqDR5jN87inE",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-F4xhEWL33XcjsXxUX5xsLHPgDY7Ue7tuV",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-MK8ZzV5gtqfKJM3hyFh2RoTLjJ7XcusMC",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-8D9QFsX3dZytrxNUmHSdBapNH4Mu8KUwu",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-Cur1e5VMWdAMSGqm1NrtWCpS5XSGytgGN",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-4W6kBgRjR6n3zxHzbyZ4wfyF2Zvi8XbTk",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-8xmgMU8yuSD3dvGZufTgyRFRtGsERU688",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-JUTXcCna617ZQiDwCny58KU7Wd5d7Q8Ui",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-JZuuuxp9AmigvLiVhr6vfAddP7b6AwXsj",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        },
                        {
                                "nodeID": "NodeID-CEQabJHQMkEXqNGTeVipt3JkKjNFv6bUT",
                                "rewardAddress": "X-ezc1x54tkpzezwgvd3m6mgsgfgyneh65jppdh9a9ju",
                                "delegationFee": 20000
                        }
                ],
                "cChainGenesis": "{\"config\":{\"chainId\":2612,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"27A370Fc2AcD908ce7BD5E259638E90c6f6eb06F\":{\"balance\":\"0x295BE96E64066972000000\"},\"0100000000000000000000000000000000000000\":{\"code\":\"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\",\"balance\":\"0x295BE96E64066972000000\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "From Snowflake to EZChain. Per consensum ad astra."
	}`
	// MainnetParams are the params used for mainnet
	MainnetParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliAvax,
			CreateAssetTxFee:      10 * units.MilliAvax,
			CreateSubnetTxFee:     1 * units.Avax,
			CreateBlockchainTxFee: 1 * units.Avax,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement: .8, // 80%
			MinValidatorStake: 2 * units.KiloAvax,
			MaxValidatorStake: 3 * units.MegaAvax,
			MinDelegatorStake: 25 * units.Avax,
			MinDelegationFee:  20000, // 2%
			MinStakeDuration:  2 * 7 * 24 * time.Hour,
			MaxStakeDuration:  365 * 24 * time.Hour,
			RewardConfig: reward.Config{
				MaxConsumptionRate: .12 * reward.PercentDenominator,
				MinConsumptionRate: .10 * reward.PercentDenominator,
				MintingPeriod:      365 * 24 * time.Hour,
				SupplyCap:          720 * units.MegaAvax,
			},
		},
	}
)
