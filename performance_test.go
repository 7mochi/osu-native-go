package osunative

import (
	"math"
	"testing"
)

func TestOsuPerformance(t *testing.T) {
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

	diffCalc, err := NewOsuDifficultyCalculator(ruleset, beatmap)
	if err != nil {
		t.Fatal(err)
	}
	defer diffCalc.Close()

	da, err := diffCalc.Calculate(mods)
	if err != nil {
		t.Fatal(err)
	}
	d := da.(*OsuDifficultyAttributes)

	perfCalc, err := NewOsuPerformanceCalculator()
	if err != nil {
		t.Fatal(err)
	}
	defer perfCalc.Close()

	score := &ScoreInfo{
		MaxCombo:   116,
		Accuracy:   0.94,
		CountGreat: 65,
		CountOk:    6,
		CountMiss:  0,
	}

	pa, err := perfCalc.Calculate(ruleset, beatmap, mods, score, d)
	if err != nil {
		t.Fatal(err)
	}
	p := pa.(*OsuPerformanceAttributes)

	if math.Abs(p.Total-200.0) > 1e-6 {
		t.Errorf("expected total pp ~200.0, got %f", p.Total)
	}
}

func TestTaikoPerformance(t *testing.T) {
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

	diffCalc, err := NewTaikoDifficultyCalculator(ruleset, beatmap)
	if err != nil {
		t.Fatal(err)
	}
	defer diffCalc.Close()

	da, err := diffCalc.Calculate(mods)
	if err != nil {
		t.Fatal(err)
	}
	d := da.(*TaikoDifficultyAttributes)

	perfCalc, err := NewTaikoPerformanceCalculator()
	if err != nil {
		t.Fatal(err)
	}
	defer perfCalc.Close()

	score := &ScoreInfo{
		MaxCombo:   500,
		Accuracy:   0.96,
		CountGreat: 400,
	}

	pa, err := perfCalc.Calculate(ruleset, beatmap, mods, score, d)
	if err != nil {
		t.Fatal(err)
	}
	p := pa.(*TaikoPerformanceAttributes)

	if p.Total <= 0 {
		t.Errorf("expected positive total pp, got %f", p.Total)
	}
}

func TestCatchPerformance(t *testing.T) {
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

	diffCalc, err := NewCatchDifficultyCalculator(ruleset, beatmap)
	if err != nil {
		t.Fatal(err)
	}
	defer diffCalc.Close()

	da, err := diffCalc.Calculate(mods)
	if err != nil {
		t.Fatal(err)
	}
	d := da.(*CatchDifficultyAttributes)

	perfCalc, err := NewCatchPerformanceCalculator()
	if err != nil {
		t.Fatal(err)
	}
	defer perfCalc.Close()

	score := &ScoreInfo{
		MaxCombo:   500,
		Accuracy:   0.98,
		CountGreat: 500,
	}

	pa, err := perfCalc.Calculate(ruleset, beatmap, mods, score, d)
	if err != nil {
		t.Fatal(err)
	}
	p := pa.(*CatchPerformanceAttributes)

	if p.Total <= 0 {
		t.Errorf("expected positive total pp, got %f", p.Total)
	}
}

func TestManiaPerformance(t *testing.T) {
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

	diffCalc, err := NewManiaDifficultyCalculator(ruleset, beatmap)
	if err != nil {
		t.Fatal(err)
	}
	defer diffCalc.Close()

	da, err := diffCalc.Calculate(mods)
	if err != nil {
		t.Fatal(err)
	}
	d := da.(*ManiaDifficultyAttributes)

	perfCalc, err := NewManiaPerformanceCalculator()
	if err != nil {
		t.Fatal(err)
	}
	defer perfCalc.Close()

	score := &ScoreInfo{
		MaxCombo:     1000,
		Accuracy:     0.97,
		CountPerfect: 500,
		CountGreat:   50,
		CountGood:    5,
		CountOk:      2,
		CountMeh:     1,
		CountMiss:    0,
	}

	pa, err := perfCalc.Calculate(ruleset, beatmap, mods, score, d)
	if err != nil {
		t.Fatal(err)
	}
	p := pa.(*ManiaPerformanceAttributes)

	if p.Total <= 0 {
		t.Errorf("expected positive total pp, got %f", p.Total)
	}
}
