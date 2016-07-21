package slot

import (
	"fmt"
	"math/rand"
)

type(
	Normal struct {
		// 单人可赢取筹码上限
		singleUpper int

		// 单人已赢取筹码
		singleGot int

		// 单人几率修正
		singlePra []*Probability

		// 全局可赢取筹码上限
		totalUpper int

		// 全局已赢取筹码
		totalGot int

		// 全局几率修正
		totalPra []*Probability

		// 普通模式几率
		normalPra []*NormalProbability

		// 当前第几次
		count int

		// single bet in this game
		bet int

		// num of line in this game
		line int

		// big Bonus
		bigBonus []*ProInterface

		// small bonus
		smallBonus []*ProInterface
	}

	NormalReturn struct {

	}

	NormalProbability struct {
		// 次数
		count int
		// 未中奖
		noBonus int
		// 小奖
		smallBonus int
		// 大奖
		bigBonus int
	}

	ProInterface struct {
		min int
		max int
		pra int
	}
)

func (n *Normal)Exec() (ret *NormalReturn, err error) {
	// Get the max number of bonus probability number
	lenth := len(n.normalPra)
	currentCountPra := n.count / n.normalPra[lenth-1].count

	// 单人修正
	var singlePercent int
	singlePercent = n.singleGot / n.singleUpper

	singleBonus := 0
	for i := 0; i < len(n.singlePra); i++ {
		if  n.singlePra[i].percentMin <= singlePercent && singlePercent <= n.singlePra[i].percentMax {
			//singleNoBonus := n.singlePra[i].noBonus
			singleBonus = n.singlePra[i].bonus / 2
			break
		}
	}

	var firstNormalPra []*NormalProbability
	for _, value := range n.normalPra {
		value.smallBonus = (value.smallBonus + singleBonus) * currentCountPra
		value.bigBonus = (value.bigBonus + singleBonus) * currentCountPra
		value.noBonus = 1 - value.smallBonus - value.bigBonus

		firstNormalPra = append(firstNormalPra, value)
	}

	// 全局修正
	var totalPercent int
	totalPercent = n.totalGot / n.totalUpper

	totalBonus := 0
	for i := 0; i < len(n.totalPra); i ++ {
		if n.totalPra[i].percentMin <= totalPercent && totalPercent <= n.totalPra[i].percentMax {
			//totalNoBonus := n.totalPra[i].noBonus
			totalBonus = n.totalPra[i].bonus / 2
			break
		}
	}

	var secondNormalPra []*NormalProbability
	for _, value := range firstNormalPra {
		value.smallBonus = (value.smallBonus + totalBonus) * currentCountPra
		value.bigBonus = (value.bigBonus + totalBonus) * currentCountPra
		value.noBonus = 1 - value.bigBonus - value.smallBonus

		secondNormalPra = append(secondNormalPra, value)
	}

	// 筹码设定
	// Get the max bonus rate the player can get
	var currentProbability *NormalProbability
	if len := len(secondNormalPra); n.count >= len {
		currentProbability = secondNormalPra[len - 1]
	} else {
		currentProbability = secondNormalPra[n.count - 1]
	}

	fir := rand.Intn(100)
	if fir < currentProbability.noBonus {
		// no bonus
		return
	}

	// bonus, choose small bonus or big bonus - one -- by multiple
	newMultiple := 0
	if multiple := n.singleUpper / (n.line * n.bet); multiple < n.bigBonus[0].min {
		// In the small bonus interval
		currentBonusLoop := getBonusLoop(n.smallBonus, multiple)

		newMultiple = getBonusMultiple(currentBonusLoop, multiple)
	} else {
		// choose small bonus or big bonus - two -- by fir
		if fir < (currentProbability.noBonus + currentProbability.smallBonus) {
			// in the small bonus
			newMultiple = getBonusMultiple(n.smallBonus, 0)
		} else {
			// In the big bonus
			currentBonusLoop := getBonusLoop(n.bigBonus, multiple)

			newMultiple = getBonusMultiple(currentBonusLoop, multiple)
		}
	}

	// return this game's bonus
	lastBonus := newMultiple * n.line * n.bet

	fmt.Print(lastBonus)

	return
}

func getBonusLoop(currentBonusLoop []*ProInterface, multiple int) []*ProInterface {
	biggerThanMu := 0

	for n := len(currentBonusLoop) - 1; n >= 0; n-- {
		value := currentBonusLoop[n]
		if value.min > multiple {
			biggerThanMu += value.pra
		} else {
			currentBonusLoop[n].pra += biggerThanMu
			break
		}
	}

	return currentBonusLoop
}

func getBonusMultiple(currentBonusLoop []*ProInterface, multiple int) int {
	first, nowPro := rand.Intn(100), 0

	var currentInterval *ProInterface
	for _, value := range currentBonusLoop {
		nowPro += value.pra
		if first < nowPro {
			currentInterval = value
			break
		}
	}

	second, diff := 0, 0
	if multiple != 0 {
		second, diff = rand.Intn(101), multiple - currentInterval.min
	} else {
		second, diff = rand.Intn(101), currentInterval.max - currentInterval.min
	}

	newMultiple := second * diff / 100 + currentInterval.min

	return newMultiple
}
