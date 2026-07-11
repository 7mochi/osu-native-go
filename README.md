# osu-native-go

[![Go Version](https://img.shields.io/github/go-mod/go-version/7mochi/osu-native-go)](https://github.com/7mochi/osu-native-go)
[![License](https://img.shields.io/badge/license-mit%20license-brightgreen.svg)][license]

Go wrapper for [osu-native], providing difficulty and performance calculation for all [osu!] modes.

## Example

### Calculating performance

```go
package main

import (
	"fmt"
	"log"

	"github.com/7mochi/osu-native-go"
)

func main() {
	beatmap, err := osunative.NewBeatmapFromFile("/path/to/file.osu")
	if err != nil {
		log.Fatal(err)
	}
	defer beatmap.Close()

	ruleset, err := osunative.NewRulesetFromID(0)
	if err != nil {
		log.Fatal(err)
	}
	defer ruleset.Close()

	mods, err := osunative.NewModsCollection()
	if err != nil {
		log.Fatal(err)
	}
	defer mods.Close()

	for _, name := range []string{"DT", "CL"} {
		mod, err := osunative.NewMod(name)
		if err != nil {
			log.Fatal(err)
		}
		mods.Add(mod)
	}

	score := &osunative.ScoreInfo{
		Accuracy:   0.94,
		MaxCombo:   116,
		CountGreat: 65,
		CountOk:    6,
		CountMiss:  0,
	}

	diffCalc, err := osunative.NewOsuDifficultyCalculator(ruleset, beatmap)
	if err != nil {
		log.Fatal(err)
	}
	defer diffCalc.Close()

	da, err := diffCalc.Calculate(mods)
	if err != nil {
		log.Fatal(err)
	}
	diffAttrs := da.(*osunative.OsuDifficultyAttributes)

	perfCalc, err := osunative.NewOsuPerformanceCalculator()
	if err != nil {
		log.Fatal(err)
	}
	defer perfCalc.Close()

	pa, err := perfCalc.Calculate(ruleset, beatmap, mods, score, diffAttrs)
	if err != nil {
		log.Fatal(err)
	}
	perfAttrs := pa.(*osunative.OsuPerformanceAttributes)

	fmt.Println(perfAttrs.Total)
}
```

## Installation

```bash
go get github.com/7mochi/osu-native-go
```

## Supported Platforms

- **Windows**: x64
- **Linux**: x64, ARM64
- **macOS**: ARM64 (Apple Silicon)

## Thanks to

- [minisbett](https://github.com/minisbett) for maintaining [osu-native].

[osu!]: https://osu.ppy.sh/
[osu-native]: https://github.com/minisbett/osu-native
[license]: https://github.com/7mochi/osu-native-go/blob/master/LICENSE
