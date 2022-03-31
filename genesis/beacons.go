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
			"178.128.61.28:9651",
			"178.128.219.27:9651",
			"178.128.221.4:9651",
			"178.128.222.159:9651",
			"178.128.219.5:9651",
			"178.128.221.157:9651",
			"178.128.213.165:9651",
			"157.230.36.32:9651",
			"178.128.209.191:9651",
			"157.230.44.79:9651",
			"137.184.70.139:9651",
			"138.68.168.230:9651",
			"104.248.134.92:9651",
			"139.59.31.231:9651",
		}
	case constants.FujiID:
		return []string{
			//"159.65.140.194:9651",
			//"159.65.134.118:9651",
			//"167.99.79.166:9651",
			//"206.189.81.122:9651",
			//"139.59.251.132:9651",
			//"167.99.66.205:9651",
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
			"NodeID-9tw3VEy9ec4hALt5twYQfrTnwUxG2Y427",
			"NodeID-H7hK3RQk8TwtNTgmsTASFz2s6Zyte15F3",
			"NodeID-CrZXSE5RS8iL8VcMebdTkhH1cx5QGU5Bt",
			"NodeID-FPPBLNq93pSG2SSosuumbp2hAY7jtXuES",
			"NodeID-Ensvn756qCTFJkQLBm1uwYi532u7xCg4f",
			"NodeID-2AeeQmA6sxAV7TN3vdNGsZ7c9L16fCr1z",
			"NodeID-6bi69cif7p3ce87MGLsfmEwuUCF6CVAP5",
			"NodeID-7hRtzmasYGCkevhvdckaRBzeefWkS88tR",
			"NodeID-8QFUSzmv3yB3SF8U3mTCrSemfEwJrwjCn",
			"NodeID-J1Zfm4ErKDg4qGmqeAU1nQBHQexTYYai9",
			"NodeID-AQ2NrUd6K7S7tWEfvdz7gPwFi6PRsrT4S",
			"NodeID-PEXNS9A69Hqg9biqJRfbRp8v4gT2uzVKe",
			"NodeID-DmCPDFFyMhdf98SwXDn7NehyZBU4P2joo",
			"NodeID-5xykNSMN2WcoYQpuqcwdNw6hA5DuBb8VP",
		}
	case constants.FujiID:
		return []string{
			//"NodeID-5UMaQhxhrzet968foe95VLp8KKWwT3tLk",
			//"NodeID-LkiLFbmocMydnZ45jJgfoeLVTMN4nn83h",
			//"NodeID-FRouddSdqsz9SqFddmcpX3kMcTUuczYWW",
			//"NodeID-LmktUgqWHu3cDdLAAeCaLGabgbpoCwBz5",
			//"NodeID-Bigfsv6ru1ZhEu6tBTofPizNYKfzhjraq",
			//"NodeID-9a3Vo6ufbhvMzfu41gyJ9DDGGf34BSm8n",
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
