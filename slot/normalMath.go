package slot

import (
	"math/rand"
)

type(
	Normal struct {
		// 单人可赢取筹码上限
		singleUpper uint16

		// 单人已赢取筹码
		singleGot uint16

		// 单人几率修正
		singlePra []*Probability

		// 全局可赢取筹码上限
		totalUpper uint16

		// 全局已赢取筹码
		totalGot uint16

		// 全局几率修正
		totalPra []*Probability

		// 普通模式几率
		normalPra []*NormalProbability

		// 当前第几次
		count int32

		// single bet in this game
		bet int

		// num of line in this game
		line int
	}

	NormalReturn struct {

	}

	NormalProbability struct {
		// 次数
		count uint8
		// 未中奖
		noBonus uint8
		// 小奖
		smallBonus uint8
		// 大奖
		bigBonus uint8
	}

	SmallInterface struct {
		row []*ProInterface{}
	}

	BigBonusInterface struct {
		row []*ProInterface{}
	}

	ProInterface struct {
		least float32
		most float32
		Pra int
	}
)

func (n *Normal)Exec() (ret *NormalReturn, err error) {
	// 单人修正
	var singlePercent uint8
	singlePercent = n.singleGot / n.singleUpper

	for i := 0; i < len(singlePra); i++ {
		if  n.singlePra[i].percentMin <= singlePercent <= n.singlePra[i].percentMax {
			singleNoBonus := n.singlePra[i].noBonus
			singleBonus := n.singlePra[i].bonus / 2
			break
		}
	}

	var firstNormalPra []*NormalProbability
	for _, value := range n.normalPra {
		value.noBonus += singleNoBonus
		value.smallBonus += singleBonus
		value.bigBonus += singleBonus

		firstNormalPra = append(firstNormalPra, value)
	}

	// 全局修正
	var totalPercent uint8
	totalPercent = n.totalGot / n.totalUpper

	for i := 0; i < len(totalPra); i ++ {
		if n.totalPra[i].percentMin <= totalPercent <= n.totalPra[i].percentMax {
			totalNoBonus := n.totalPra[i].noBonus
			totalBonus := n.totalPra[i].bonus / 2
			break
		}
	}

	var secondNormalPra []*NormalProbability
	for _, value := range firstNormalPra {
		value.noBonus += totalNoBonus
		value.smallBonus += totalBonus
		value.bigBonus += totalBonus

		secondNormalPra = append(secondNormalPra, value)
	}

	// 筹码设定
	// Get current winning rate
	if len := len(secondNormalPra); count >= len {
		currentProbability := secondNormalPra[len - 1]
	} else {
		currentProbability := secondNormalPra[count - 1]
	}

	// Get current bonus multiple
	currentLastCorns := singleUpper - singleGot
	maxMultiple := currentLastCorns / (bet * line)

	if maxMultiple < BigBonusInterface[0].least {
		currentProbability.smallBonus += currentProbability.bigBonus
	}

	for _, value := range BigBonusInterface {
		if value.least <= maxMultiple <= value.most {

		}
	}

	for _, value := range SmallBonusInterface {

	}

	r := rand.New(rand.NewSource(99))
	if fir := r.Intn(100); fir >= currentProbability.noBonus + currentProbability.smallBonus {
		// Get the big bonus, create reward now
		// first get the bonus interval
		if sec := r.Intn(100);

	} else if fir >= currentProbability.noBonus {
		// Get the small bonus, create reward now
		// first get the bonus interval
	} else {
		reward := 0
	}

	// 免费、奖金、技能、连赢等模式是否触发

}
