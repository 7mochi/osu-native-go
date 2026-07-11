package osunative

import (
	"math"
	"testing"
)

func TestOsuDifficulty(t *testing.T) {
	t.Skip("requires osu-native native library")

	ruleset, err := NewRulesetFromID(0)
	if err != nil {
		t.Fatal(err)
	}
	defer ruleset.Close()

	beatmap, err := NewBeatmapFromFile("tests/5438072.osu")
	if err != nil {
		t.Fatal(err)
	}
	defer beatmap.Close()

	modDT, err := NewMod("DT")
	if err != nil {
		t.Fatal(err)
	}
	defer modDT.Close()

	mods, err := NewModsCollection()
	if err != nil {
		t.Fatal(err)
	}
	defer mods.Close()
	mods.Add(modDT)

	calc, err := NewOsuDifficultyCalculator(ruleset, beatmap)
	if err != nil {
		t.Fatal(err)
	}
	defer calc.Close()

	da, err := calc.Calculate(mods)
	if err != nil {
		t.Fatal(err)
	}
	d := da.(*OsuDifficultyAttributes)

	if math.Abs(d.StarRating-7.60) > 1e-6 {
		t.Errorf("expected star rating ~7.60, got %f", d.StarRating)
	}
}

func TestTaikoDifficulty(t *testing.T) {
	t.Skip("requires osu-native native library")

	ruleset, err := NewRulesetFromID(1)
	if err != nil {
		t.Fatal(err)
	}
	defer ruleset.Close()

	beatmap, err := NewBeatmapFromFile("tests/221923.osu")
	if err != nil {
		t.Fatal(err)
	}
	defer beatmap.Close()

	modDT, err := NewMod("DT")
	if err != nil {
		t.Fatal(err)
	}
	defer modDT.Close()

	mods, err := NewModsCollection()
	if err != nil {
		t.Fatal(err)
	}
	defer mods.Close()
	mods.Add(modDT)

	calc, err := NewTaikoDifficultyCalculator(ruleset, beatmap)
	if err != nil {
		t.Fatal(err)
	}
	defer calc.Close()

	da, err := calc.Calculate(mods)
	if err != nil {
		t.Fatal(err)
	}
	d := da.(*TaikoDifficultyAttributes)

	if math.Abs(d.StarRating-5.0) > 1e-6 {
		t.Errorf("expected star rating ~5.0, got %f", d.StarRating)
	}
}

func TestCatchDifficulty(t *testing.T) {
	t.Skip("requires osu-native native library")

	ruleset, err := NewRulesetFromID(2)
	if err != nil {
		t.Fatal(err)
	}
	defer ruleset.Close()

	beatmap, err := NewBeatmapFromFile("tests/4289411.osu")
	if err != nil {
		t.Fatal(err)
	}
	defer beatmap.Close()

	mods, err := NewModsCollection()
	if err != nil {
		t.Fatal(err)
	}
	defer mods.Close()

	calc, err := NewCatchDifficultyCalculator(ruleset, beatmap)
	if err != nil {
		t.Fatal(err)
	}
	defer calc.Close()

	da, err := calc.Calculate(mods)
	if err != nil {
		t.Fatal(err)
	}
	d := da.(*CatchDifficultyAttributes)

	if math.Abs(d.StarRating-3.0) > 1e-6 {
		t.Errorf("expected star rating ~3.0, got %f", d.StarRating)
	}
}

func TestManiaDifficulty(t *testing.T) {
	t.Skip("requires osu-native native library")

	ruleset, err := NewRulesetFromID(3)
	if err != nil {
		t.Fatal(err)
	}
	defer ruleset.Close()

	beatmap, err := NewBeatmapFromFile("tests/5107047.osu")
	if err != nil {
		t.Fatal(err)
	}
	defer beatmap.Close()

	mods, err := NewModsCollection()
	if err != nil {
		t.Fatal(err)
	}
	defer mods.Close()

	calc, err := NewManiaDifficultyCalculator(ruleset, beatmap)
	if err != nil {
		t.Fatal(err)
	}
	defer calc.Close()

	da, err := calc.Calculate(mods)
	if err != nil {
		t.Fatal(err)
	}
	d := da.(*ManiaDifficultyAttributes)

	if math.Abs(d.StarRating-4.0) > 1e-6 {
		t.Errorf("expected star rating ~4.0, got %f", d.StarRating)
	}
}
