package slot

import (

)

type(
	Bonus struct {
		// 奖金全局可赢取筹码上限
		bonusUpper

		// 奖金全局已赢取筹码
		bonusGot

		// bonus win probability
		bonusProbability *BonusProbability

		// 奖金几率修正
		BonusPra []*Probability

		// chip Slice
		chipSlice []*ChipSlice
	}

	BonusProbability struct {
		bonus int
		noBonus int
	}

	ChipSlice struct {
		numbering int
		amount int
	}

	BonusReturn struct {

	}
)

func (b *Bonus)Exec() (ret *BonusReturn, err error) {
	// 几率修正
	bonusPercent := b.bonusGot / b.bonusUpper

	for i := 0; i < len(b.bonusPra); i++ {
		if b.bonusPra[i].percentMin <= bonusPercent <= b.bonusPra[i].percentMax {
			b.bonusProbability.bonus += b.bonusPra[i].bonus
			b.bonusProbability.noBonus += b.bonusPra[i].noBonus
			break
		}
	}

	// Determine whether winning
	if fir := rand.Intn(100); fir < b.bonusProbability.noBonus {
		// no bonus
		return
	}

	// Determine the bonus multiple
	sec := rand.Intn(100)


	// 筹码设定
}
