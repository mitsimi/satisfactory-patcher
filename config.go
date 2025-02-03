package main

type Engine struct {
	Player struct {
		ConfiguredInternetSpeed int
		ConfiguredLanSpeed      int
	} `ini:"/Script/Engine.Player"`
	IpNetDriver struct {
		MaxClientRate         int
		MaxInternetClientRate int
	} `ini:"/Script/OnlineSubsystemUtils.IpNetDriver"`
	EpicNetDriver struct {
		MaxClientRate         int
		MaxInternetClientRate int
	} `ini:"/Script/SocketSubsystemEpic.EpicNetDriver"`
}

type Game struct {
	GameNetworkManager struct {
		TotalNetBandwidth   int
		MaxDynamicBandwidth int
		MinDynamicBandwidth int
	} `ini:"/Script/Engine.GameNetworkManager"`
}

type Scalability struct {
	NetworkQuality struct {
		TotalNetBandwidth   int
		MaxDynamicBandwidth int
		MinDynamicBandwidth int
	} `ini:"NetworkQuality@3"`
}
