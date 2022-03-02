// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package reward

import (
	"math/big"
	"time"
)

const secondsPerYear = 31556952
var (

	ratOne  = big.NewRat(1, 1)
	ratZero = new(big.Rat)
	bigTwo  = big.NewInt(2)

	baseTS, _ = new(big.Int).SetString("1000000000000000000000000", 10)

	pcBase = big.NewRat(50, 100)
	pcStep = big.NewRat(-4, 100)

	percentRewardPool = big.NewRat(15, 100)

	etherUnit, _   = new(big.Int).SetString("1000000000000000000", 10)
	etherUnitFloat = new(big.Float).SetInt(etherUnit)
	_ Calculator = &calculator{}
)

type Calculator interface {
	Calculate(stakedDuration time.Duration, stakedAmount, currentSupply uint64) uint64
}

type calculator struct {
	maxSubMinConsumptionRate *big.Int
	minConsumptionRate       *big.Int
	mintingPeriod            *big.Int
	supplyCap                uint64
}

func NewCalculator(c Config) Calculator {
	return &calculator{
		maxSubMinConsumptionRate: new(big.Int).SetUint64(c.MaxConsumptionRate - c.MinConsumptionRate),
		minConsumptionRate:       new(big.Int).SetUint64(c.MinConsumptionRate),
		mintingPeriod:            new(big.Int).SetUint64(uint64(c.MintingPeriod)),
		supplyCap:                c.SupplyCap,
	}
}

// Reward returns the amount of tokens to reward the staker with.
//
// RemainingSupply = SupplyCap - ExistingSupply
// PortionOfExistingSupply = StakedAmount / ExistingSupply
// PortionOfStakingDuration = StakingDuration / MaximumStakingDuration
// MintingRate = MinMintingRate + MaxSubMinMintingRate * PortionOfStakingDuration
// Reward = RemainingSupply * PortionOfExistingSupply * MintingRate * PortionOfStakingDuration
func (c *calculator) Calculate(stakedDuration time.Duration, stakedAmount, currentSupply uint64) uint64 {
	bigStakedDuration := new(big.Int).SetUint64(uint64(stakedDuration))
	bigStakedAmount := new(big.Int).SetUint64(stakedAmount)
	bigCurrentSupply := new(big.Int).SetUint64(currentSupply)

	adjustedConsumptionRateNumerator := new(big.Int).Mul(c.maxSubMinConsumptionRate, bigStakedDuration)
	adjustedMinConsumptionRateNumerator := new(big.Int).Mul(c.minConsumptionRate, c.mintingPeriod)
	adjustedConsumptionRateNumerator.Add(adjustedConsumptionRateNumerator, adjustedMinConsumptionRateNumerator)
	adjustedConsumptionRateDenominator := new(big.Int).Mul(c.mintingPeriod, consumptionRateDenominator)

	reward := new(big.Int).SetUint64(c.supplyCap - currentSupply)
	reward.Mul(reward, adjustedConsumptionRateNumerator)
	reward.Mul(reward, bigStakedAmount)
	reward.Mul(reward, bigStakedDuration)
	reward.Div(reward, adjustedConsumptionRateDenominator)
	reward.Div(reward, bigCurrentSupply)
	reward.Div(reward, c.mintingPeriod)

	return reward.Uint64()
}

// x = ts / baseTS
// b = floor(log2(x))
// a = 2^b
// r = b-1+x/a
// p = 50% - 4%*r
func rewardPercent(ts *big.Int) *big.Rat {
	x := new(big.Rat).SetFrac(ts, baseTS)
	b := big.NewInt(int64(ts.BitLen() - baseTS.BitLen()))

	a := new(big.Int).Exp(bigTwo, b, nil)

	r := x.Quo(x, new(big.Rat).SetInt(a))
	r = r.Add(r, new(big.Rat).SetInt(b))
	r = r.Sub(r, ratOne)

	p := r.Mul(r, pcStep)
	p = p.Add(p, pcBase)

	if p.Sign() < 0 {
		p.Set(ratZero)
	}

	return p
}

func fromEther(v float64) (wei *big.Int) {
	f := big.NewFloat(v)
	f = f.Mul(f, etherUnitFloat)
	wei, _ = f.Int(nil)
	return wei
}

func toEther(wei *big.Int) (v float64) {
	v, _ = new(big.Rat).SetFrac(wei, etherUnit).Float64()
	return v
}

func fromEZC(v float64) (wei *big.Int) {
	return fromEther(v / 1000000000)
}

func toEZC(wei *big.Int) (v float64) {
	return toEther(wei) * 1000000000
}

/**
 * duration: number of seconds to stake
 * stake: the stake amount to be locked
 * totalStake: total locked stake before this stake is locked
 * rewardPool: the remain reward pool
 */
func reward(duration uint, stake *big.Int, totalStake *big.Int, rewardPool *big.Int) *big.Int {
	ts := new(big.Int).Add(totalStake, stake)
	if ts.Cmp(baseTS) < 0 {
		ts = ts.Set(baseTS)
	}

	p := rewardPercent(ts)

	// total rate (tr) = min(rp*20%, ts * p)
	tr := p.Mul(p, new(big.Rat).SetInt(ts))
	cap := new(big.Rat).SetInt(rewardPool)
	cap = cap.Mul(cap, percentRewardPool)
	if tr.Cmp(cap) > 0 {
		tr = cap
	}

	// reward = tr * s / ts
	reward := tr.Mul(tr, new(big.Rat).SetFrac(stake, ts))

	// reward *= duration / secondsPerYear
	reward = reward.Mul(reward, new(big.Rat).SetFrac64(int64(duration), secondsPerYear))

	return reward.Num().Div(reward.Num(), reward.Denom())
}

func rewardEZC(duration uint, stake *big.Int, totalStake *big.Int, rewardPool *big.Int) uint64 {
	return uint64(toEZC(reward(duration, stake, totalStake, rewardPool)))
}
