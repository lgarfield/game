package slot

import (
	"math/rand"
)

type(
	Forecast struct {
		// 连赢全局可赢取筹码上限
		forecastUpper int

		// 连赢已赢取筹码
		forecastGot int

		// 连赢几率
		forecastProbability *ForecastProbability

		// Probability correct
		forecastPra []*Probability

		// bet amount
		bet int

		// forecast loop
		forecastLoop []*ForecastLoop
	}

	ForecastProbability struct {
		bonus int
		noBonus int
	}

	ForecastLoop struct {
		min int
		max int
		bonus int
	}

	ForecastReturn struct {

	}
)

func (f *Forecast)Exec() (ret *ForecastReturn, err error){
	// 几率修正
	forecastPercent := f.forecastGot / f.forecastUpper

	for i := 0; i < len(f.forecastPra); i++ {
		if f.forecastPra[i].percentMin <= forecastPercent <= f.forecastPra[i].percentMax {
			f.forecastProbability.bonus += f.forecastPra[i].bonus
			f.forecastProbability.noBonus += f.forecastPra[i].noBonus
			break
		}
	}

	// Determine whether winning
	if fir := rand.Intn(100); fir < f.forecastProbability.noBonus {
		// no bonus
		return
	}

	// Determine the bonus multiple
	sec, nowPro := rand.Intn(100), 0
	for _, value := range f.forecastLoop {
		nowPro += value.bonus
		if sec < nowPro {
			currentLoop := value
			break
		}
	}

	third, diff := rand.Intn(101), currentLoop.max - currentLoop.min
	multiple := third * diff / 100 + currentLoop.min

	// 筹码设定
	forecastBonus := f.bet * multiple

	if forecastBonus > (f.forecastUpper - f.forecastGot) {
		// no bonus
		return
	}

	// reutrn the bonus amount
	return
}
