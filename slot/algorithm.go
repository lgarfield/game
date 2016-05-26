package slot

import(

)

type (
	// 几率
	Probability struct {
		// 比例起始
		percentMin uint8
		// 比例终止
		percentMax uint8
		// wei中奖比率
		noBonus int8
		// 中奖比率
		bonus int8
	}

	Return interface {
		Exec()
	}
)
