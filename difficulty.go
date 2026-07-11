package osunative

/*
#include "build/generated/cabinet.h"
*/
import "C"
import (
	"runtime"
)

type DifficultyCalculator interface {
	Calculate(mods *ModsCollection) (any, error)
	Close()
}

func CreateDifficultyCalculator(ruleset *Ruleset, beatmap *Beatmap) (DifficultyCalculator, error) {
	switch ruleset.ID() {
	case 0:
		return NewOsuDifficultyCalculator(ruleset, beatmap)
	case 1:
		return NewTaikoDifficultyCalculator(ruleset, beatmap)
	case 2:
		return NewCatchDifficultyCalculator(ruleset, beatmap)
	case 3:
		return NewManiaDifficultyCalculator(ruleset, beatmap)
	default:
		return nil, UnexpectedRuleset
	}
}

type OsuDifficultyCalculator struct {
	handle C.ManagedObjectHandle
	closed bool
}

func NewOsuDifficultyCalculator(ruleset *Ruleset, beatmap *Beatmap) (*OsuDifficultyCalculator, error) {
	var native C.NativeOsuDifficultyCalculator
	result := C.OsuDifficultyCalculator_Create(ruleset.handle, beatmap.handle, &native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	c := &OsuDifficultyCalculator{handle: native.handle}
	runtime.SetFinalizer(c, (*OsuDifficultyCalculator).Close)
	return c, nil
}

func (c *OsuDifficultyCalculator) Calculate(mods *ModsCollection) (any, error) {
	var attrs C.NativeOsuDifficultyAttributes
	result := C.OsuDifficultyCalculator_Calculate(c.handle, mods.handle, &attrs)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	return &OsuDifficultyAttributes{
		DifficultyAttributes: DifficultyAttributes{
			StarRating: float64(attrs.starRating),
			MaxCombo:   int32(attrs.maxCombo),
		},
		AimDifficulty:                float64(attrs.aimDifficulty),
		AimDifficultSliderCount:      float64(attrs.aimDifficultSliderCount),
		SpeedDifficulty:              float64(attrs.speedDifficulty),
		SpeedNoteCount:               float64(attrs.speedNoteCount),
		FlashlightDifficulty:         float64(attrs.flashlightDifficulty),
		SliderFactor:                 float64(attrs.sliderFactor),
		AimTopWeightedSliderFactor:   float64(attrs.aimTopWeightedSliderFactor),
		SpeedTopWeightedSliderFactor: float64(attrs.speedTopWeightedSliderFactor),
		AimDifficultStrainCount:      float64(attrs.aimDifficultStrainCount),
		SpeedDifficultStrainCount:    float64(attrs.speedDifficultStrainCount),
		NestedScorePerObject:         float64(attrs.nestedScorePerObject),
		LegacyScoreBaseMultiplier:    float64(attrs.legacyScoreBaseMultiplier),
		MaximumLegacyComboScore:      float64(attrs.maximumLegacyComboScore),
		DrainRate:                    float64(attrs.drainRate),
		HitCircleCount:               int32(attrs.hitCircleCount),
		SliderCount:                  int32(attrs.sliderCount),
		SpinnerCount:                 int32(attrs.spinnerCount),
	}, nil
}

func (c *OsuDifficultyCalculator) Close() {
	if !c.closed {
		C.OsuDifficultyCalculator_Destroy(c.handle)
		c.closed = true
		runtime.SetFinalizer(c, nil)
	}
}

type TaikoDifficultyCalculator struct {
	handle C.ManagedObjectHandle
	closed bool
}

func NewTaikoDifficultyCalculator(ruleset *Ruleset, beatmap *Beatmap) (*TaikoDifficultyCalculator, error) {
	var native C.NativeTaikoDifficultyCalculator
	result := C.TaikoDifficultyCalculator_Create(ruleset.handle, beatmap.handle, &native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	c := &TaikoDifficultyCalculator{handle: native.handle}
	runtime.SetFinalizer(c, (*TaikoDifficultyCalculator).Close)
	return c, nil
}

func (c *TaikoDifficultyCalculator) Calculate(mods *ModsCollection) (any, error) {
	var attrs C.NativeTaikoDifficultyAttributes
	result := C.TaikoDifficultyCalculator_Calculate(c.handle, mods.handle, &attrs)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	return &TaikoDifficultyAttributes{
		DifficultyAttributes: DifficultyAttributes{
			StarRating: float64(attrs.starRating),
			MaxCombo:   int32(attrs.maxCombo),
		},
		MechanicalDifficulty: float64(attrs.mechanicalDifficulty),
		RhythmDifficulty:     float64(attrs.rhythmDifficulty),
		ReadingDifficulty:    float64(attrs.readingDifficulty),
		ColourDifficulty:     float64(attrs.colourDifficulty),
		StaminaDifficulty:    float64(attrs.staminaDifficulty),
		MonoStaminaFactor:    float64(attrs.monoStaminaFactor),
		ConsistencyFactor:    float64(attrs.consistencyFactor),
		StaminaTopStrains:    float64(attrs.staminaTopStrains),
	}, nil
}

func (c *TaikoDifficultyCalculator) Close() {
	if !c.closed {
		C.TaikoDifficultyCalculator_Destroy(c.handle)
		c.closed = true
		runtime.SetFinalizer(c, nil)
	}
}

type CatchDifficultyCalculator struct {
	handle C.ManagedObjectHandle
	closed bool
}

func NewCatchDifficultyCalculator(ruleset *Ruleset, beatmap *Beatmap) (*CatchDifficultyCalculator, error) {
	var native C.NativeCatchDifficultyCalculator
	result := C.CatchDifficultyCalculator_Create(ruleset.handle, beatmap.handle, &native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	c := &CatchDifficultyCalculator{handle: native.handle}
	runtime.SetFinalizer(c, (*CatchDifficultyCalculator).Close)
	return c, nil
}

func (c *CatchDifficultyCalculator) Calculate(mods *ModsCollection) (any, error) {
	var attrs C.NativeCatchDifficultyAttributes
	result := C.CatchDifficultyCalculator_Calculate(c.handle, mods.handle, &attrs)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	return &CatchDifficultyAttributes{
		DifficultyAttributes: DifficultyAttributes{
			StarRating: float64(attrs.starRating),
			MaxCombo:   int32(attrs.maxCombo),
		},
	}, nil
}

func (c *CatchDifficultyCalculator) Close() {
	if !c.closed {
		C.CatchDifficultyCalculator_Destroy(c.handle)
		c.closed = true
		runtime.SetFinalizer(c, nil)
	}
}

type ManiaDifficultyCalculator struct {
	handle C.ManagedObjectHandle
	closed bool
}

func NewManiaDifficultyCalculator(ruleset *Ruleset, beatmap *Beatmap) (*ManiaDifficultyCalculator, error) {
	var native C.NativeManiaDifficultyCalculator
	result := C.ManiaDifficultyCalculator_Create(ruleset.handle, beatmap.handle, &native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	c := &ManiaDifficultyCalculator{handle: native.handle}
	runtime.SetFinalizer(c, (*ManiaDifficultyCalculator).Close)
	return c, nil
}

func (c *ManiaDifficultyCalculator) Calculate(mods *ModsCollection) (any, error) {
	var attrs C.NativeManiaDifficultyAttributes
	result := C.ManiaDifficultyCalculator_Calculate(c.handle, mods.handle, &attrs)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	return &ManiaDifficultyAttributes{
		DifficultyAttributes: DifficultyAttributes{
			StarRating: float64(attrs.starRating),
			MaxCombo:   int32(attrs.maxCombo),
		},
	}, nil
}

func (c *ManiaDifficultyCalculator) Close() {
	if !c.closed {
		C.ManiaDifficultyCalculator_Destroy(c.handle)
		c.closed = true
		runtime.SetFinalizer(c, nil)
	}
}
