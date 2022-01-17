// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"math/big"
)

const secondsPerYear = 31556952

var (
	// from 2022 - 01 - 01 , next element for 1 year and end to 2031 - 01 - 01
	/*
		//timestampStakes = [16]uint64{
		//	1640995200, 1672531200,
		//	1704067200, 1735689600,
		//	1767225600, 1798761600,
		//	1830297600, 1861920000,
		//	1893456000, 1924992000,
		//	1956528000, 1988150400,
		//	2019686400, 2051222400,
		//	2082758400, 2114380800,
		//}

		//allocateStakeBasedTimestamps = [16]big.Int{
		//	*big.NewInt(75000000000),
		//	*big.NewInt(63750000000),
		//	*big.NewInt(54187500000),
		//	*big.NewInt(46059375000),
		//	*big.NewInt(39150468750),
		//	*big.NewInt(33277898440),
		//	*big.NewInt(28286213670),
		//	*big.NewInt(24043281620),
		//	*big.NewInt(20436789380),
		//	*big.NewInt(17371270970),
		//	*big.NewInt(14765580330),
		//	*big.NewInt(12550743280),
		//	*big.NewInt(10668131790),
		//	*big.NewInt(90679120170),
		//	*big.NewInt(77077252150),
		//	*big.NewInt(65515664330),
		//}

		//maxRangeStakes = [8]big.Int{
		//	*big.NewInt(8000000000),
		//	*big.NewInt(16000000000),
		//	*big.NewInt(32000000000),
		//	*big.NewInt(64000000000),
		//	*big.NewInt(128000000000),
		//	*big.NewInt(256000000000),
		//	*big.NewInt(376000000000),
		//	*big.NewInt(512000000000),
		//}
		//
		//percentRewardStakes = [9]int64{
		//	50, 45, 40, 35, 30, 25, 25, 18, 12,
		//}
	*/

	ratOne  = big.NewRat(1, 1)
	ratZero = new(big.Rat)
	bigTwo  = big.NewInt(2)

	baseTS, _ = new(big.Int).SetString("1000000000000000000000000", 10)

	pcBase = big.NewRat(50, 100)
	pcStep = big.NewRat(-4, 100)

	percentRewardPool = big.NewRat(15, 100)

	etherUnit, _   = new(big.Int).SetString("1000000000000000000", 10)
	etherUnitFloat = new(big.Float).SetInt(etherUnit)
)

// x = ts / baseTS
// b = floor(log2(x))
// a = 2^b
// r = b-1+x/a
// p = 65% - 5%*r
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
