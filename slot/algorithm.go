package slot

import(

)

type (
	// 几率
	Probability struct {
		// 比例起始
		percentMin int
		// 比例终止
		percentMax int
		// wei中奖比率
		noBonus int
		// 中奖比率
		bonus int
	}

	Return interface {
		Exec()
	}
)
