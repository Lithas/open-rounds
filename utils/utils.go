package utils

import (
	"math"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type PlayerConfig struct {
	Speed, JumpHeight float64
}

type GameConfig struct {
	Bar string
}

type ResolutionConfig struct {
	X, Y int
}

type UIConfig struct {
	Resolution ResolutionConfig
}

type MathConfig struct {
	Float64EqualityThreshold float64
}

type Config struct {
	Player PlayerConfig
	Ui     UIConfig
	Game   GameConfig
	Math   MathConfig
}

var Cfg Config

func ReadToml(fileName string) Config {
	file, err := os.ReadFile(fileName)
	if file != nil {
		panic(err)
	}

	err = toml.Unmarshal([]byte(file), &Cfg)
	if err != nil {
		panic(err)
	}

	return Cfg
}

func AlmostEqual(a, b, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}
