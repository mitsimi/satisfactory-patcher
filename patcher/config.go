package patcher

import (
	"errors"

	"gopkg.in/ini.v1"
)

type Config struct {
	Engine      EngineConfig
	Game        GameConfig
	Scalability ScalabilityConfig
}

type Player struct {
	ConfiguredInternetSpeed int
	ConfiguredLanSpeed      int
}

type IpNetDriver struct {
	MaxClientRate         int
	MaxInternetClientRate int
}

type EpicNetDriver struct {
	MaxClientRate         int
	MaxInternetClientRate int
}

type EngineConfig struct {
	Player        Player        `ini:"/Script/Engine.Player"`
	IpNetDriver   IpNetDriver   `ini:"/Script/OnlineSubsystemUtils.IpNetDriver"`
	EpicNetDriver EpicNetDriver `ini:"/Script/SocketSubsystemEpic.EpicNetDriver"`
}

type GameNetworkManager struct {
	TotalNetBandwidth   int
	MaxDynamicBandwidth int
	MinDynamicBandwidth int
}

type GameConfig struct {
	GameNetworkManager GameNetworkManager `ini:"/Script/Engine.GameNetworkManager"`
}

type NetworkQuality struct {
	TotalNetBandwidth   int
	MaxDynamicBandwidth int
	MinDynamicBandwidth int
}

type ScalabilityConfig struct {
	NetworkQuality NetworkQuality `ini:"NetworkQuality@3"`
}

const (
	EngineIni      = "Engine.ini"
	GameIni        = "Game.ini"
	ScalabilityIni = "Scalability.ini"
)

var INIs = []string{EngineIni, GameIni, ScalabilityIni}

func applyPatch(cfg *ini.File, path string) error {
	var err error = nil
	switch path {
	case EngineIni:
		data := patchedEngine()
		err = cfg.ReflectFrom(data)
		break
	case GameIni:
		data := patchedGame()
		err = cfg.ReflectFrom(data)
		break
	case ScalabilityIni:
		data := patchedScalability()
		err = cfg.ReflectFrom(data)
		break
	default:
		return errors.New("could not apply a patch for unknown file")
	}

	return err
}

func patchedEngine() *EngineConfig {
	v := 104857600
	return &EngineConfig{
		Player: Player{
			ConfiguredInternetSpeed: v,
			ConfiguredLanSpeed:      v,
		},
		IpNetDriver: IpNetDriver{
			MaxClientRate:         v,
			MaxInternetClientRate: v,
		},
		EpicNetDriver: EpicNetDriver{
			MaxClientRate:         v,
			MaxInternetClientRate: v,
		},
	}
}

func patchedGame() *GameConfig {
	v := 104857600
	return &GameConfig{
		GameNetworkManager: GameNetworkManager{
			TotalNetBandwidth:   v,
			MaxDynamicBandwidth: v,
			MinDynamicBandwidth: v,
		},
	}
}

func patchedScalability() *ScalabilityConfig {
	v := 104857600
	return &ScalabilityConfig{
		NetworkQuality: NetworkQuality{
			TotalNetBandwidth:   v,
			MaxDynamicBandwidth: v,
			MinDynamicBandwidth: v,
		},
	}
}
