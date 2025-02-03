package patcher

type Engine struct {
	Player        EnginePlayer        `ini:"/Script/Engine.Player"`
	IpNetDriver   EngineIpNetDriver   `ini:"/Script/OnlineSubsystemUtils.IpNetDriver"`
	EpicNetDriver EngineEpicNetDriver `ini:"/Script/SocketSubsystemEpic.EpicNetDriver"`
}

type EnginePlayer struct {
	ConfiguredInternetSpeed int
	ConfiguredLanSpeed      int
}

type EngineIpNetDriver struct {
	MaxClientRate         int
	MaxInternetClientRate int
}

type EngineEpicNetDriver struct {
	MaxClientRate         int
	MaxInternetClientRate int
}

type Game struct {
	GameNetworkManager NetworkManager `ini:"/Script/Engine.GameNetworkManager"`
}

type NetworkManager struct {
	TotalNetBandwidth   int
	MaxDynamicBandwidth int
	MinDynamicBandwidth int
}

type Scalability struct {
	NetworkQuality ScalabilityNetworkQuality `ini:"NetworkQuality@3"`
}

type ScalabilityNetworkQuality struct {
	TotalNetBandwidth   int
	MaxDynamicBandwidth int
	MinDynamicBandwidth int
}

const (
	EngineIni      = "Engine.ini"
	GameIni        = "Game.ini"
	ScalabilityIni = "Scalability.ini"
)

func PatchedEngine() interface{} {
	v := 104857600
	return Engine{
		Player: EnginePlayer{
			ConfiguredInternetSpeed: v,
			ConfiguredLanSpeed:      v,
		},
		IpNetDriver: EngineIpNetDriver{
			MaxClientRate:         v,
			MaxInternetClientRate: v,
		},
		EpicNetDriver: EngineEpicNetDriver{
			MaxClientRate:         v,
			MaxInternetClientRate: v,
		},
	}
}

func PatchedGame() interface{} {
	v := 104857600
	return Game{
		GameNetworkManager: NetworkManager{
			TotalNetBandwidth:   v,
			MaxDynamicBandwidth: v,
			MinDynamicBandwidth: v,
		},
	}
}

func PatchedScalability() interface{} {
	v := 104857600
	return Scalability{
		NetworkQuality: ScalabilityNetworkQuality{
			TotalNetBandwidth:   v,
			MaxDynamicBandwidth: v,
			MinDynamicBandwidth: v,
		},
	}
}
