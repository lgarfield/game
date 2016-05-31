package slot

import (
	"math/rand"
)

type(
	Free struct {
		// 免费全局可赢取筹码上限
		freeUpper int

		// 免费全局已赢取筹码
		freeGot int

		// probability correct
		freePra []*Probability

		// 免费中奖几率
		freeProbability *FreeProability

		// cards bonus probability
		cardProbability []*CardProbability

		// free loop
		freeLoop []*FreeLoop

		// bet amount
		bet

		// free multiple
		freeMultiple
	}

	// The probability of the num of cards
	CardProbability struct {
		num int
		probability int
		multiple int
	}

	FreeProbability struct {
		bonus int
		noBonus int
	}

	FreeLoop struct {
		min int
		max int
		bonus int
		count []int
	}

	FreeReturn struct {

	}
)

func (f *Free)exec() (ret *FreeReturn, err error) {
	// 几率修正
	freePercent := f.freeGot / f.freeUpper

	for i := 0; i < len(f.freePra); i++ {
		if f.freePra[i].percentMin <= freePercent <= f.freePra[i].percentMax {
			f.FreeProbability.bonus += f.freePra[i].bonus
			f.FreeProbability.noBonus += f.freePra[i].noBonus
			break
		}
	}

	// Determine whether winning
	if fir := rand.Intn(100); fir < f.FreeProbability.noBonus {
		// no bonus
		return
	}

	// Determine the bonus multiple
	sec := rand.Intn(100)
	nowPro := 0
	for _, value := range f.freeLoop {
		nowPro += value.bonus
		if sec < nowPro {
			currentFreeLoop := value
			break
		}
	}

	// Get the multiple and loop num.
	third := rand.Intn(101);
	diff := currentFreeLoop.max - currentFreeLoop.min
	multiple := third * diff / 100 + currentFreeLoop.min

	if num := len(currentFreeLoop.loop); num == 2 {
		forth := rand.Intn(2)
		loop := currentFreeLoop.loop[forth]
	} else {
		loop := currentFreeLoop.loop[0]
	}

	// Get the amount of bonus
	// first is chip bonus
	chipBonus := f.bet * multiple * f.freeMultiple

	var cardBonus int
	// second is card bonus
	for i := 0; i <= len(loop); i ++ {
		fifth := rand.Intn(100)

		num := 0
		for _, value = range f.cardProbability {
			num += value.probability
			if fifth < num {
				multiple := value.multiple
				break
			}
		}

		cardBonus += f.bet * multiple
	}

	// See all the bonus is exceed the amount can get
	if (chipBonus + cardBonus) > (f.freeUpper - f.freeGot) {
		// exceed, so no bonus
	}

	// return the correct bonus
}
