// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/sampler"
)

// getIPs returns the beacon IPs for each network
func getIPs(networkID uint32) []string {
	switch networkID {
	case constants.MainnetID:
		return []string{
			"134.209.106.66:9651",
			"134.209.106.63:9651",
			"134.209.106.229:9651",
			"134.209.106.165:9651",
			"134.209.106.94:9651",
			"134.209.106.159:9651",
			"134.209.106.231:9651",
			"134.209.106.183:9651",
			"134.209.106.134:9651",
			"134.209.106.203:9651",
			"139.59.12.197:9651",
			"188.166.174.5:9651",
			"139.59.213.217:9651",
			"161.35.99.231:9651",
		}
	case constants.FujiID:
		return []string{
			"159.65.140.194:9651",
			"159.65.134.118:9651",
			"167.99.79.166:9651",
			"206.189.81.122:9651",
			"139.59.251.132:9651",
			"167.99.66.205:9651",
		}
	default:
		return nil
	}
}

// getNodeIDs returns the beacon node IDs for each network
func getNodeIDs(networkID uint32) []string {
	switch networkID {
	case constants.MainnetID:
		return []string{
			"NodeID-4YrMwoyo2ypQNhiEYGRMQmUB2QYTFxuTL",
			"NodeID-JeAspnL1KmYfFoAH1UdnghqaKSfqdd9ur",
			"NodeID-4dWu6veRSyLhmGv7dvCSCz7Pttgxv2F9e",
			"NodeID-MqjT6x5UgRnvLwb7RLtTGXTvHsSFed4QC",
			"NodeID-CavKfAxmQ66eqTMTSGMLWPqDR5jN87inE",
			"NodeID-F4xhEWL33XcjsXxUX5xsLHPgDY7Ue7tuV",
			"NodeID-MK8ZzV5gtqfKJM3hyFh2RoTLjJ7XcusMC",
			"NodeID-8D9QFsX3dZytrxNUmHSdBapNH4Mu8KUwu",
			"NodeID-Cur1e5VMWdAMSGqm1NrtWCpS5XSGytgGN",
			"NodeID-4W6kBgRjR6n3zxHzbyZ4wfyF2Zvi8XbTk",
			"NodeID-8xmgMU8yuSD3dvGZufTgyRFRtGsERU688",
			"NodeID-JUTXcCna617ZQiDwCny58KU7Wd5d7Q8Ui",
			"NodeID-JZuuuxp9AmigvLiVhr6vfAddP7b6AwXsj",
			"NodeID-CEQabJHQMkEXqNGTeVipt3JkKjNFv6bUT",
		}
	case constants.FujiID:
		return []string{
			"NodeID-5UMaQhxhrzet968foe95VLp8KKWwT3tLk",
			"NodeID-LkiLFbmocMydnZ45jJgfoeLVTMN4nn83h",
			"NodeID-FRouddSdqsz9SqFddmcpX3kMcTUuczYWW",
			"NodeID-LmktUgqWHu3cDdLAAeCaLGabgbpoCwBz5",
			"NodeID-Bigfsv6ru1ZhEu6tBTofPizNYKfzhjraq",
			"NodeID-9a3Vo6ufbhvMzfu41gyJ9DDGGf34BSm8n",
		}
	default:
		return nil
	}
}

// SampleBeacons returns the some beacons this node should connect to
func SampleBeacons(networkID uint32, count int) ([]string, []string) {
	ips := getIPs(networkID)
	ids := getNodeIDs(networkID)

	if numIPs := len(ips); numIPs < count {
		count = numIPs
	}

	sampledIPs := make([]string, 0, count)
	sampledIDs := make([]string, 0, count)

	s := sampler.NewUniform()
	_ = s.Initialize(uint64(len(ips)))
	indices, _ := s.Sample(count)
	for _, index := range indices {
		sampledIPs = append(sampledIPs, ips[int(index)])
		sampledIDs = append(sampledIDs, ids[int(index)])
	}

	return sampledIPs, sampledIDs
}
