// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/utils/units"
)

func isPowerTwo(num int) bool {
	return (num != 0) && ((num & (num - 1)) == 0)
}

func TestRewardLongerDurationBonus(t *testing.T) {
	shortDuration := 24 * time.Hour
	totalDuration := 365 * 24 * time.Hour
	shortBalance := units.MegaAvax
	rewardPool := 500 * units.MegaAvax
	var TotalStake uint64 = 0
	for i := 0; i < int(totalDuration/shortDuration); i++ {
		r := rewardEZC(
			uint(shortDuration.Seconds()),
			fromEZC(float64(shortBalance)),
			fromEZC(float64(TotalStake)),
			fromEZC(float64(rewardPool)),
		)
		shortBalance += r
		TotalStake += shortBalance
		rewardPool -= r
	}
	r := rewardEZC(
		uint((totalDuration % shortDuration).Seconds()),
		fromEZC(float64(shortBalance)),
		fromEZC(float64(TotalStake)), fromEZC(float64(rewardPool)),
	)
	shortBalance += r

	// reset reward pool for compare
	rewardPool = 500 * units.MegaAvax
	longBalance := units.MegaAvax
	longBalance += rewardEZC(
		uint(totalDuration.Seconds()),
		fromEZC(float64(longBalance)),
		fromEZC(float64(0)),
		fromEZC(float64(rewardPool)),
	)
	if shortBalance >= longBalance {
		t.Fatalf("should promote stakers to stake longer")
	}
}

func TestRewardYearTokenomic(t *testing.T) {
	// 365 * 24 * time.Hour
	Duration := 31556952
	Stake := units.MegaAvax
	rewardPool := 500 * units.MegaAvax
	var TotalStake uint64 = 0
	reward := rewardEZC(
		uint(Duration),
		fromEZC(float64(Stake)),
		fromEZC(float64(TotalStake)),
		fromEZC(float64(rewardPool)),
	)

	if reward > units.MegaAvax/2 {
		t.Fatalf("it seem first reward for first stake wrong formulation")
	}

	// Total Stake	1	2	4	8	16	20	64	128	256	450	512	1024 ( mega avax )
	// Percent Gain	50% 46% 42%	38%	34%	30%	26%	22%	18%	14%	10%	6%
	// reset total and reward pool
	// how total max supply is 1000, mean 500 for genesis in wallet
	var stakeAmounts [256]uint64
	percentReward := [12]float64{50, 46, 42, 38, 34, 30, 26, 22, 18, 14, 10, 6}
	TotalStake = 0
	rewardPool = 500 * units.MegaAvax
	power := 0
	for i := 0; i < 256; i++ {
		reward := rewardEZC(
			uint(Duration),
			fromEZC(float64(Stake)),
			fromEZC(float64(TotalStake)),
			fromEZC(float64(rewardPool)),
		)
		stakeAmounts[i] = reward
		rewardPool -= reward
		TotalStake += Stake
	}

	for i := range stakeAmounts {
		if isPowerTwo(i) {
			percent := float64(stakeAmounts[i-1]) / float64(Stake) * 100
			if percentReward[power] != percent {
				t.Fatalf("reward caculate is wrong")
			}
			power += 1
		}
	}
}
