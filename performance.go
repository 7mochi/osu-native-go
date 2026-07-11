package osunative

/*
#include "native/bin/cabinet.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

type PerformanceCalculator interface {
	Calculate(ruleset *Ruleset, beatmap *Beatmap, mods *ModsCollection, score *ScoreInfo, diffAttrs any) (any, error)
	Close()
}

func CreatePerformanceCalculator(ruleset *Ruleset) (PerformanceCalculator, error) {
	switch ruleset.ID() {
	case 0:
		return NewOsuPerformanceCalculator()
	case 1:
		return NewTaikoPerformanceCalculator()
	case 2:
		return NewCatchPerformanceCalculator()
	case 3:
		return NewManiaPerformanceCalculator()
	default:
		return nil, UnexpectedRuleset
	}
}

type OsuPerformanceCalculator struct {
	handle C.ManagedObjectHandle
	closed bool
}

func NewOsuPerformanceCalculator() (*OsuPerformanceCalculator, error) {
	var native C.NativeOsuPerformanceCalculator
	result := C.OsuPerformanceCalculator_Create(&native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	c := &OsuPerformanceCalculator{handle: native.handle}
	runtime.SetFinalizer(c, (*OsuPerformanceCalculator).Close)
	return c, nil
}

func (c *OsuPerformanceCalculator) Calculate(
	ruleset *Ruleset, beatmap *Beatmap, mods *ModsCollection,
	score *ScoreInfo, diffAttrs any,
) (any, error) {
	da, ok := diffAttrs.(*OsuDifficultyAttributes)
	if !ok {
		return nil, UnexpectedRuleset
	}

	var ns C.NativeScoreInfo
	ns.rulesetHandle = ruleset.handle
	ns.beatmapHandle = beatmap.handle
	ns.modsHandle = mods.handle
	ns.maxCombo = C.int(score.MaxCombo)
	ns.accuracy = C.double(score.Accuracy)
	ns.countMiss = C.int(score.CountMiss)
	ns.countMeh = C.int(score.CountMeh)
	ns.countOk = C.int(score.CountOk)
	ns.countGood = C.int(score.CountGood)
	ns.countGreat = C.int(score.CountGreat)
	ns.countPerfect = C.int(score.CountPerfect)
	ns.countSmallTickMiss = C.int(score.CountSmallTickMiss)
	ns.countSmallTickHit = C.int(score.CountSmallTickHit)
	ns.countLargeTickMiss = C.int(score.CountLargeTickMiss)
	ns.countLargeTickHit = C.int(score.CountLargeTickHit)
	ns.countSliderTailHit = C.int(score.CountSliderTailHit)
	if score.LegacyTotalScore != nil {
		setNullableInt64(unsafe.Pointer(&ns.legacyTotalScore), *score.LegacyTotalScore)
	}

	var nd C.NativeOsuDifficultyAttributes
	nd.starRating = C.double(da.StarRating)
	nd.maxCombo = C.int(da.MaxCombo)
	nd.aimDifficulty = C.double(da.AimDifficulty)
	nd.aimDifficultSliderCount = C.double(da.AimDifficultSliderCount)
	nd.speedDifficulty = C.double(da.SpeedDifficulty)
	nd.speedNoteCount = C.double(da.SpeedNoteCount)
	nd.flashlightDifficulty = C.double(da.FlashlightDifficulty)
	nd.sliderFactor = C.double(da.SliderFactor)
	nd.aimTopWeightedSliderFactor = C.double(da.AimTopWeightedSliderFactor)
	nd.speedTopWeightedSliderFactor = C.double(da.SpeedTopWeightedSliderFactor)
	nd.aimDifficultStrainCount = C.double(da.AimDifficultStrainCount)
	nd.speedDifficultStrainCount = C.double(da.SpeedDifficultStrainCount)
	nd.nestedScorePerObject = C.double(da.NestedScorePerObject)
	nd.legacyScoreBaseMultiplier = C.double(da.LegacyScoreBaseMultiplier)
	nd.maximumLegacyComboScore = C.double(da.MaximumLegacyComboScore)
	nd.drainRate = C.double(da.DrainRate)
	nd.hitCircleCount = C.int(da.HitCircleCount)
	nd.sliderCount = C.int(da.SliderCount)
	nd.spinnerCount = C.int(da.SpinnerCount)

	var np C.NativeOsuPerformanceAttributes
	result := C.OsuPerformanceCalculator_Calculate(c.handle, ns, nd, &np)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}

	return &OsuPerformanceAttributes{
		PerformanceAttributes:        PerformanceAttributes{Total: float64(np.total)},
		Aim:                          float64(np.aim),
		Speed:                        float64(np.speed),
		Accuracy:                     float64(np.accuracy),
		Flashlight:                   float64(np.flashlight),
		EffectiveMissCount:           float64(np.effectiveMissCount),
		SpeedDeviation:               readNullableDouble(unsafe.Pointer(&np.speedDeviation)),
		ComboBasedEstimatedMissCount: float64(np.comboBasedEstimatedMissCount),
		ScoreBasedEstimatedMissCount: readNullableDouble(unsafe.Pointer(&np.scoreBasedEstimatedMissCount)),
		AimEstimatedSliderBreaks:     float64(np.aimEstimatedSliderBreaks),
		SpeedEstimatedSliderBreaks:   float64(np.speedEstimatedSliderBreaks),
	}, nil
}

func (c *OsuPerformanceCalculator) Close() {
	if !c.closed {
		C.OsuPerformanceCalculator_Destroy(c.handle)
		c.closed = true
		runtime.SetFinalizer(c, nil)
	}
}

type TaikoPerformanceCalculator struct {
	handle C.ManagedObjectHandle
	closed bool
}

func NewTaikoPerformanceCalculator() (*TaikoPerformanceCalculator, error) {
	var native C.NativeTaikoPerformanceCalculator
	result := C.TaikoPerformanceCalculator_Create(&native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	c := &TaikoPerformanceCalculator{handle: native.handle}
	runtime.SetFinalizer(c, (*TaikoPerformanceCalculator).Close)
	return c, nil
}

func (c *TaikoPerformanceCalculator) Calculate(
	ruleset *Ruleset, beatmap *Beatmap, mods *ModsCollection,
	score *ScoreInfo, diffAttrs any,
) (any, error) {
	da, ok := diffAttrs.(*TaikoDifficultyAttributes)
	if !ok {
		return nil, UnexpectedRuleset
	}

	var ns C.NativeScoreInfo
	ns.rulesetHandle = ruleset.handle
	ns.beatmapHandle = beatmap.handle
	ns.modsHandle = mods.handle
	ns.maxCombo = C.int(score.MaxCombo)
	ns.accuracy = C.double(score.Accuracy)
	ns.countMiss = C.int(score.CountMiss)
	ns.countMeh = C.int(score.CountMeh)
	ns.countOk = C.int(score.CountOk)
	ns.countGreat = C.int(score.CountGreat)
	ns.countSliderTailHit = C.int(score.CountSliderTailHit)

	var nd C.NativeTaikoDifficultyAttributes
	nd.starRating = C.double(da.StarRating)
	nd.maxCombo = C.int(da.MaxCombo)
	nd.mechanicalDifficulty = C.double(da.MechanicalDifficulty)
	nd.rhythmDifficulty = C.double(da.RhythmDifficulty)
	nd.readingDifficulty = C.double(da.ReadingDifficulty)
	nd.colourDifficulty = C.double(da.ColourDifficulty)
	nd.staminaDifficulty = C.double(da.StaminaDifficulty)
	nd.monoStaminaFactor = C.double(da.MonoStaminaFactor)
	nd.consistencyFactor = C.double(da.ConsistencyFactor)
	nd.staminaTopStrains = C.double(da.StaminaTopStrains)

	var np C.NativeTaikoPerformanceAttributes
	result := C.TaikoPerformanceCalculator_Calculate(c.handle, ns, nd, &np)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}

	return &TaikoPerformanceAttributes{
		PerformanceAttributes: PerformanceAttributes{Total: float64(np.total)},
		Difficulty:            float64(np.difficulty),
		Accuracy:              float64(np.accuracy),
		EstimatedUnstableRate: readNullableDouble(unsafe.Pointer(&np.estimatedUnstableRate)),
	}, nil
}

func (c *TaikoPerformanceCalculator) Close() {
	if !c.closed {
		C.TaikoPerformanceCalculator_Destroy(c.handle)
		c.closed = true
		runtime.SetFinalizer(c, nil)
	}
}

type CatchPerformanceCalculator struct {
	handle C.ManagedObjectHandle
	closed bool
}

func NewCatchPerformanceCalculator() (*CatchPerformanceCalculator, error) {
	var native C.NativeCatchPerformanceCalculator
	result := C.CatchPerformanceCalculator_Create(&native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	c := &CatchPerformanceCalculator{handle: native.handle}
	runtime.SetFinalizer(c, (*CatchPerformanceCalculator).Close)
	return c, nil
}

func (c *CatchPerformanceCalculator) Calculate(
	ruleset *Ruleset, beatmap *Beatmap, mods *ModsCollection,
	score *ScoreInfo, diffAttrs any,
) (any, error) {
	da, ok := diffAttrs.(*CatchDifficultyAttributes)
	if !ok {
		return nil, UnexpectedRuleset
	}

	var ns C.NativeScoreInfo
	ns.rulesetHandle = ruleset.handle
	ns.beatmapHandle = beatmap.handle
	ns.modsHandle = mods.handle
	ns.maxCombo = C.int(score.MaxCombo)
	ns.accuracy = C.double(score.Accuracy)
	ns.countGreat = C.int(score.CountGreat)
	ns.countSmallTickMiss = C.int(score.CountSmallTickMiss)
	ns.countSmallTickHit = C.int(score.CountSmallTickHit)
	ns.countLargeTickMiss = C.int(score.CountLargeTickMiss)
	ns.countLargeTickHit = C.int(score.CountLargeTickHit)

	var nd C.NativeCatchDifficultyAttributes
	nd.starRating = C.double(da.StarRating)
	nd.maxCombo = C.int(da.MaxCombo)

	var np C.NativeCatchPerformanceAttributes
	result := C.CatchPerformanceCalculator_Calculate(c.handle, ns, nd, &np)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}

	return &CatchPerformanceAttributes{
		PerformanceAttributes: PerformanceAttributes{Total: float64(np.total)},
	}, nil
}

func (c *CatchPerformanceCalculator) Close() {
	if !c.closed {
		C.CatchPerformanceCalculator_Destroy(c.handle)
		c.closed = true
		runtime.SetFinalizer(c, nil)
	}
}

type ManiaPerformanceCalculator struct {
	handle C.ManagedObjectHandle
	closed bool
}

func NewManiaPerformanceCalculator() (*ManiaPerformanceCalculator, error) {
	var native C.NativeManiaPerformanceCalculator
	result := C.ManiaPerformanceCalculator_Create(&native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	c := &ManiaPerformanceCalculator{handle: native.handle}
	runtime.SetFinalizer(c, (*ManiaPerformanceCalculator).Close)
	return c, nil
}

func (c *ManiaPerformanceCalculator) Calculate(
	ruleset *Ruleset, beatmap *Beatmap, mods *ModsCollection,
	score *ScoreInfo, diffAttrs any,
) (any, error) {
	da, ok := diffAttrs.(*ManiaDifficultyAttributes)
	if !ok {
		return nil, UnexpectedRuleset
	}

	var ns C.NativeScoreInfo
	ns.rulesetHandle = ruleset.handle
	ns.beatmapHandle = beatmap.handle
	ns.modsHandle = mods.handle
	ns.maxCombo = C.int(score.MaxCombo)
	ns.accuracy = C.double(score.Accuracy)
	ns.countMiss = C.int(score.CountMiss)
	ns.countMeh = C.int(score.CountMeh)
	ns.countOk = C.int(score.CountOk)
	ns.countGood = C.int(score.CountGood)
	ns.countGreat = C.int(score.CountGreat)
	ns.countPerfect = C.int(score.CountPerfect)

	var nd C.NativeManiaDifficultyAttributes
	nd.starRating = C.double(da.StarRating)
	nd.maxCombo = C.int(da.MaxCombo)

	var np C.NativeManiaPerformanceAttributes
	result := C.ManiaPerformanceCalculator_Calculate(c.handle, ns, nd, &np)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}

	return &ManiaPerformanceAttributes{
		PerformanceAttributes: PerformanceAttributes{Total: float64(np.total)},
		Difficulty:            float64(np.difficulty),
	}, nil
}

func (c *ManiaPerformanceCalculator) Close() {
	if !c.closed {
		C.ManiaPerformanceCalculator_Destroy(c.handle)
		c.closed = true
		runtime.SetFinalizer(c, nil)
	}
}
