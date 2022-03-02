// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package reward

import (
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/utils/units"
)

<<<<<<< HEAD:vms/platformvm/reward_test.go
func isPowerTwo(num int) bool {
	return (num != 0) && ((num & (num - 1)) == 0)
}

func TestRewardLongerDurationBonus(t *testing.T) {
=======
const (
	defaultMinStakingDuration = 24 * time.Hour
	defaultMaxStakingDuration = 365 * 24 * time.Hour

	defaultMinValidatorStake = 5 * units.MilliAvax
)

var defaultConfig = Config{
	MaxConsumptionRate: .12 * PercentDenominator,
	MinConsumptionRate: .10 * PercentDenominator,
	MintingPeriod:      365 * 24 * time.Hour,
	SupplyCap:          720 * units.MegaAvax,
}

func TestLongerDurationBonus(t *testing.T) {
	c := NewCalculator(defaultConfig)
>>>>>>> master:vms/platformvm/reward/calculator_test.go
	shortDuration := 24 * time.Hour
	totalDuration := 365 * 24 * time.Hour
	shortBalance := units.MegaAvax
	rewardPool := 500 * units.MegaAvax
	var TotalStake uint64 = 0
	for i := 0; i < int(totalDuration/shortDuration); i++ {
<<<<<<< HEAD:vms/platformvm/reward_test.go
		r := rewardEZC(
			uint(shortDuration.Seconds()),
			fromEZC(float64(shortBalance)),
			fromEZC(float64(TotalStake)),
			fromEZC(float64(rewardPool)),
		)
=======
		r := c.Calculate(shortDuration, shortBalance, 359*units.MegaAvax+shortBalance)
>>>>>>> master:vms/platformvm/reward/calculator_test.go
		shortBalance += r
		TotalStake += shortBalance
		rewardPool -= r
	}
<<<<<<< HEAD:vms/platformvm/reward_test.go
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
=======
	r := c.Calculate(totalDuration%shortDuration, shortBalance, 359*units.MegaAvax+shortBalance)
	shortBalance += r

	longBalance := units.KiloAvax
	longBalance += c.Calculate(totalDuration, longBalance, 359*units.MegaAvax+longBalance)

>>>>>>> master:vms/platformvm/reward/calculator_test.go
	if shortBalance >= longBalance {
		t.Fatalf("should promote stakers to stake longer")
	}
}

<<<<<<< HEAD:vms/platformvm/reward_test.go
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
=======
func TestRewards(t *testing.T) {
	c := NewCalculator(defaultConfig)
	tests := []struct {
		duration       time.Duration
		stakeAmount    uint64
		existingAmount uint64
		expectedReward uint64
	}{
		// Max duration:
		{ // (720M - 360M) * (1M / 360M) * 12%
			duration:       defaultMaxStakingDuration,
			stakeAmount:    units.MegaAvax,
			existingAmount: 360 * units.MegaAvax,
			expectedReward: 120 * units.KiloAvax,
		},
		{ // (720M - 400M) * (1M / 400M) * 12%
			duration:       defaultMaxStakingDuration,
			stakeAmount:    units.MegaAvax,
			existingAmount: 400 * units.MegaAvax,
			expectedReward: 96 * units.KiloAvax,
		},
		{ // (720M - 400M) * (2M / 400M) * 12%
			duration:       defaultMaxStakingDuration,
			stakeAmount:    2 * units.MegaAvax,
			existingAmount: 400 * units.MegaAvax,
			expectedReward: 192 * units.KiloAvax,
		},
		{ // (720M - 720M) * (1M / 720M) * 12%
			duration:       defaultMaxStakingDuration,
			stakeAmount:    units.MegaAvax,
			existingAmount: defaultConfig.SupplyCap,
			expectedReward: 0,
		},
		// Min duration:
		// (720M - 360M) * (1M / 360M) * (10% + 2% * MinimumStakingDuration / MaximumStakingDuration) * MinimumStakingDuration / MaximumStakingDuration
		{
			duration:       defaultMinStakingDuration,
			stakeAmount:    units.MegaAvax,
			existingAmount: 360 * units.MegaAvax,
			expectedReward: 274122724713,
		},
		// (720M - 360M) * (.005 / 360M) * (10% + 2% * MinimumStakingDuration / MaximumStakingDuration) * MinimumStakingDuration / MaximumStakingDuration
		{
			duration:       defaultMinStakingDuration,
			stakeAmount:    defaultMinValidatorStake,
			existingAmount: 360 * units.MegaAvax,
			expectedReward: 1370,
		},
		// (720M - 400M) * (1M / 400M) * (10% + 2% * MinimumStakingDuration / MaximumStakingDuration) * MinimumStakingDuration / MaximumStakingDuration
		{
			duration:       defaultMinStakingDuration,
			stakeAmount:    units.MegaAvax,
			existingAmount: 400 * units.MegaAvax,
			expectedReward: 219298179771,
		},
		// (720M - 400M) * (2M / 400M) * (10% + 2% * MinimumStakingDuration / MaximumStakingDuration) * MinimumStakingDuration / MaximumStakingDuration
		{
			duration:       defaultMinStakingDuration,
			stakeAmount:    2 * units.MegaAvax,
			existingAmount: 400 * units.MegaAvax,
			expectedReward: 438596359542,
		},
		// (720M - 720M) * (1M / 720M) * (10% + 2% * MinimumStakingDuration / MaximumStakingDuration) * MinimumStakingDuration / MaximumStakingDuration
		{
			duration:       defaultMinStakingDuration,
			stakeAmount:    units.MegaAvax,
			existingAmount: defaultConfig.SupplyCap,
			expectedReward: 0,
		},
>>>>>>> master:vms/platformvm/reward/calculator_test.go
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
<<<<<<< HEAD:vms/platformvm/reward_test.go
		stakeAmounts[i] = reward
		rewardPool -= reward
		TotalStake += Stake
	}

	for i := range stakeAmounts {
		if isPowerTwo(i) {
			percent := float64(stakeAmounts[i-1]) / float64(Stake) * 100
			if percentReward[power] != percent {
				t.Fatalf("reward caculate is wrong")
=======
		t.Run(name, func(t *testing.T) {
			r := c.Calculate(
				test.duration,
				test.stakeAmount,
				test.existingAmount,
			)
			if r != test.expectedReward {
				t.Fatalf("expected %d; got %d", test.expectedReward, r)
>>>>>>> master:vms/platformvm/reward/calculator_test.go
			}
			power += 1
		}
	}
}
