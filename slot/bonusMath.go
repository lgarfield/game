package slot

import (

)

type Bonus struct {
	// 奖金全局可赢取筹码上限

	// 奖金全局已赢取筹码

	// 奖金几率修正
}

func (b *Bonus)Exec() (ret *BonusReturn, err error) {
	// 几率修正

	// 筹码设定
}
