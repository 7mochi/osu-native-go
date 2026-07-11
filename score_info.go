package osunative

type ScoreInfo struct {
	MaxCombo           int
	Accuracy           float64
	CountMiss          int
	CountMeh           int
	CountOk            int
	CountGood          int
	CountGreat         int
	CountPerfect       int
	CountSmallTickMiss int
	CountSmallTickHit  int
	CountLargeTickMiss int
	CountLargeTickHit  int
	CountSliderTailHit int
	LegacyTotalScore   *int64
}
