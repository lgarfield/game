package slot

import (
	"math/rand"
)

type(
	Bonus struct {
		// 奖金全局可赢取筹码上限
		bonusUpper int

		// 奖金全局已赢取筹码
		bonusGot int

		// bonus win probability
		bonusProbability *BonusProbability

		// 奖金几率修正
		bonusPra []*Probability

		// Card's num probability
		cardProbability []*BonusCardProbability

		// chip Slice
		chipSlice []*ChipSlice
	}

	// base probability
	BonusProbability struct {
		bonus int
		noBonus int
	}

	// The probability of cards'num
	BonusCardProbability struct {
		num int
		probability int
	}

	// Bonus / per goal
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
		if b.bonusPra[i].percentMin <= bonusPercent && bonusPercent <= b.bonusPra[i].percentMax {
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
	sec, nowPro, cardsNum := rand.Intn(100), 0, 0
	for _, value := range b.cardProbability {
		nowPro += value.probability
		if sec < nowPro {
			cardsNum = value.num
			break
		}
	}

	// Create the bonus player can get
	bonusValue := combination(b.chipSlice, cardsNum)

	lastAmount := 0
	for _, val := range bonusValue {
		lastAmount += val.amount
	}

	if lastAmount > (b.bonusUpper - b.bonusGot) {
		// no bonud
		return
	}

	// 筹码设定
	return
}

func combination(cs []*ChipSlice, num int) (re []*ChipSlice) {
	if last := len(cs) - num; num >= last {
		for ; num > 0; num-- {
			i := len(cs)
			//r := rand.Intn(i)

			re = append(re, cs[i])
		}
	} else {
		for ; last > 0; last-- {
			i := len(cs)
			r := rand.Intn(i)

			cs = append(cs[:r], cs[r+1:]...)
		}

		re = cs
	}

	return
}
