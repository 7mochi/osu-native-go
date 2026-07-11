package osunative

type DifficultyAttributes struct {
	StarRating float64
	MaxCombo   int32
}

type OsuDifficultyAttributes struct {
	DifficultyAttributes
	AimDifficulty                float64
	AimDifficultSliderCount      float64
	SpeedDifficulty              float64
	SpeedNoteCount               float64
	FlashlightDifficulty         float64
	SliderFactor                 float64
	AimTopWeightedSliderFactor   float64
	SpeedTopWeightedSliderFactor float64
	AimDifficultStrainCount      float64
	SpeedDifficultStrainCount    float64
	NestedScorePerObject         float64
	LegacyScoreBaseMultiplier    float64
	MaximumLegacyComboScore      float64
	DrainRate                    float64
	HitCircleCount               int32
	SliderCount                  int32
	SpinnerCount                 int32
}

type TaikoDifficultyAttributes struct {
	DifficultyAttributes
	MechanicalDifficulty float64
	RhythmDifficulty     float64
	ReadingDifficulty    float64
	ColourDifficulty     float64
	StaminaDifficulty    float64
	MonoStaminaFactor    float64
	ConsistencyFactor    float64
	StaminaTopStrains    float64
}

type CatchDifficultyAttributes struct {
	DifficultyAttributes
}

type ManiaDifficultyAttributes struct {
	DifficultyAttributes
}

type PerformanceAttributes struct {
	Total float64
}

type OsuPerformanceAttributes struct {
	PerformanceAttributes
	Aim                          float64
	Speed                        float64
	Accuracy                     float64
	Flashlight                   float64
	EffectiveMissCount           float64
	SpeedDeviation               *float64
	ComboBasedEstimatedMissCount float64
	ScoreBasedEstimatedMissCount *float64
	AimEstimatedSliderBreaks     float64
	SpeedEstimatedSliderBreaks   float64
}

type TaikoPerformanceAttributes struct {
	PerformanceAttributes
	Difficulty            float64
	Accuracy              float64
	EstimatedUnstableRate *float64
}

type CatchPerformanceAttributes struct {
	PerformanceAttributes
}

type ManiaPerformanceAttributes struct {
	PerformanceAttributes
	Difficulty float64
}
