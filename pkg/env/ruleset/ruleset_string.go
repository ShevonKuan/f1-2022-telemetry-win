// Code generated by "stringer -type=Ruleset -linecomment"; DO NOT EDIT.

package ruleset

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PracticeAndQualifying-0]
	_ = x[Race-1]
	_ = x[TimeTrial-2]
	_ = x[TimeAttack-4]
	_ = x[CheckpointChallenge-6]
	_ = x[Autocross-8]
	_ = x[Drift-9]
	_ = x[AverageSpeedZone-10]
	_ = x[RivalDuel-11]
}

const (
	_Ruleset_name_0 = "Practice & QualifyingRaceTime Trial"
	_Ruleset_name_1 = "Time Attack"
	_Ruleset_name_2 = "Checkpoint Challenge"
	_Ruleset_name_3 = "AutocrossDriftAverage Speed ZoneRival Duel"
)

var (
	_Ruleset_index_0 = [...]uint8{0, 21, 25, 35}
	_Ruleset_index_3 = [...]uint8{0, 9, 14, 32, 42}
)

func (i Ruleset) String() string {
	switch {
	case i <= 2:
		return _Ruleset_name_0[_Ruleset_index_0[i]:_Ruleset_index_0[i+1]]
	case i == 4:
		return _Ruleset_name_1
	case i == 6:
		return _Ruleset_name_2
	case 8 <= i && i <= 11:
		i -= 8
		return _Ruleset_name_3[_Ruleset_index_3[i]:_Ruleset_index_3[i+1]]
	default:
		return "Ruleset(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
