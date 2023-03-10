// Code generated by "stringer -type=Track -linecomment"; DO NOT EDIT.

package track

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Melbourne-0]
	_ = x[PaulRicard-1]
	_ = x[Shanghai-2]
	_ = x[Sakhir-3]
	_ = x[Catalunya-4]
	_ = x[Monaco-5]
	_ = x[Montreal-6]
	_ = x[Silverstone-7]
	_ = x[Hockenheim-8]
	_ = x[Hungaroring-9]
	_ = x[Spa-10]
	_ = x[Monza-11]
	_ = x[Singapore-12]
	_ = x[Suzuka-13]
	_ = x[AbuDhabi-14]
	_ = x[Texas-15]
	_ = x[Brazil-16]
	_ = x[Austria-17]
	_ = x[Sochi-18]
	_ = x[Mexico-19]
	_ = x[Baku-20]
	_ = x[SakhirShort-21]
	_ = x[SilverstoneShort-22]
	_ = x[TexasShort-23]
	_ = x[SuzukaShort-24]
	_ = x[Hanoi-25]
	_ = x[Zandvoort-26]
	_ = x[Imola-27]
	_ = x[Partimao-28]
	_ = x[Jeddah-29]
	_ = x[Miami-30]
}

const _Track_name = "MelbournePaul RicardShanghaiSakhirCatalunyaMonacoMontrealSilverstoneHockenheimHungaroringSpaMonzaSingaporeSuzukaAbu DhabiTexasBrazilAustriaSochiMexicoBakuSakhir ShortSilverstone ShortTexas ShortSuzuka ShortHanoiZandvoortImolaPartimaoJeddahMiami"

var _Track_index = [...]uint8{0, 9, 20, 28, 34, 43, 49, 57, 68, 78, 89, 92, 97, 106, 112, 121, 126, 132, 139, 144, 150, 154, 166, 183, 194, 206, 211, 220, 225, 233, 239, 244}

func (i Track) String() string {
	if i >= Track(len(_Track_index)-1) {
		return "Track(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Track_name[_Track_index[i]:_Track_index[i+1]]
}
