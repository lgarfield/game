package slot

import (
	"math/rand"
)

type(

	Skill struct {
		// 技能全局可赢取筹码上限
		skillUpper int

		// 技能全局已赢取筹码
		skillGot int

		// Probability correct
		skillPra []*Probability

		// 技能几率
		skillProbability *SkillProbability

		// bet amount
		bet int

		// skill loop
		skillLoop []*SkillLoop
	}

	SkillProbability struct {
		bonus int
		noBonus int
	}

	// Skill Loop
	SkillLoop struct {
		min int
		max int
		bonus int
	}

	SkillReturn struct {

	}
)

func (s *Skill)Exec() (ret *SkillReturn, err error) {
	// 中奖几率修正
	skillPercent := s.skillGot / s.skillUpper

	for i := 0; i < len(s.skillPra); i++ {
		if s.skillPra[i].percentMin <= skillPercent <= s.skillPra[i].percentMax {
			s.skillProbability.bonus += s.skillPra[i].bonus
			s.skillProbability.noBonus += s.skillPra[i].noBonus
			break
		}
	}

	// Determin whether winning
	if fir := rand.Intn(100); fir < s.skillProbability.noBonus {
		// no bonus
		return
	}

	// Determine the bonus multiple
	sec, nowPro := rand.Intn(100), 0
	for _, value := range.s.skillLoop {
		nowPro += value.bonus
		if sec < nowPro {
			currentLoop := value
			break
		}
	}

	third := rand.Intn(101)
	diff := currentLoop.max - currentLoop.min
	multiple := third * diff / 100 + currentLoop.min

	// 筹码设定
	skillBonus := s.bet * multiple

	if skillBonus > (s.skillUpper - s.skillGot) {
		// no bonus
		return
	}

	// return the bonus amount
	return
}
