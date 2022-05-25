package main

type TasmotaInfo struct {
	Warning string `json:"WARNING"`
	Status  struct {
		Module       int      `json:"Module"`
		DeviceName   string   `json:"DeviceName"`
		FriendlyName []string `json:"FriendlyName"`
		Topic        string   `json:"Topic"`
		ButtonTopic  string   `json:"ButtonTopic"`
		Power        int      `json:"Power"`
		PowerOnState int      `json:"PowerOnState"`
		LedState     int      `json:"LedState"`
		LedMask      string   `json:"LedMask"`
		SaveData     int      `json:"SaveData"`
		SaveState    int      `json:"SaveState"`
		SwitchTopic  string   `json:"SwitchTopic"`
		SwitchMode   []int    `json:"SwitchMode"`
		ButtonRetain int      `json:"ButtonRetain"`
		SwitchRetain int      `json:"SwitchRetain"`
		SensorRetain int      `json:"SensorRetain"`
		PowerRetain  int      `json:"PowerRetain"`
		InfoRetain   int      `json:"InfoRetain"`
		StateRetain  int      `json:"StateRetain"`
	} `json:"Status"`
	StatusPRM struct {
		Baudrate      int    `json:"Baudrate"`
		SerialConfig  string `json:"SerialConfig"`
		GroupTopic    string `json:"GroupTopic"`
		OtaURL        string `json:"OtaUrl"`
		RestartReason string `json:"RestartReason"`
		Uptime        string `json:"Uptime"`
		StartupUTC    string `json:"StartupUTC"`
		Sleep         int    `json:"Sleep"`
		CfgHolder     int    `json:"CfgHolder"`
		BootCount     int    `json:"BootCount"`
		BCResetTime   string `json:"BCResetTime"`
		SaveCount     int    `json:"SaveCount"`
		SaveAddress   string `json:"SaveAddress"`
	} `json:"StatusPRM"`
	StatusFWR struct {
		Version       string `json:"Version"`
		BuildDateTime string `json:"BuildDateTime"`
		Boot          int    `json:"Boot"`
		Core          string `json:"Core"`
		Sdk           string `json:"SDK"`
		CPUFrequency  int    `json:"CpuFrequency"`
		Hardware      string `json:"Hardware"`
		Cr            string `json:"CR"`
	} `json:"StatusFWR"`
	StatusLOG struct {
		SerialLog  int      `json:"SerialLog"`
		WebLog     int      `json:"WebLog"`
		MqttLog    int      `json:"MqttLog"`
		SysLog     int      `json:"SysLog"`
		LogHost    string   `json:"LogHost"`
		LogPort    int      `json:"LogPort"`
		SSID       []string `json:"SSId"`
		TelePeriod int      `json:"TelePeriod"`
		Resolution string   `json:"Resolution"`
		SetOption  []string `json:"SetOption"`
	} `json:"StatusLOG"`
	StatusMEM struct {
		ProgramSize      int      `json:"ProgramSize"`
		Free             int      `json:"Free"`
		Heap             int      `json:"Heap"`
		ProgramFlashSize int      `json:"ProgramFlashSize"`
		FlashSize        int      `json:"FlashSize"`
		FlashChipID      string   `json:"FlashChipId"`
		FlashFrequency   int      `json:"FlashFrequency"`
		FlashMode        int      `json:"FlashMode"`
		Features         []string `json:"Features"`
		Drivers          string   `json:"Drivers"`
		Sensors          string   `json:"Sensors"`
	} `json:"StatusMEM"`
	StatusNET struct {
		Hostname   string  `json:"Hostname"`
		IPAddress  string  `json:"IPAddress"`
		Gateway    string  `json:"Gateway"`
		Subnetmask string  `json:"Subnetmask"`
		DNSServer1 string  `json:"DNSServer1"`
		DNSServer2 string  `json:"DNSServer2"`
		Mac        string  `json:"Mac"`
		Webserver  int     `json:"Webserver"`
		HTTPAPI    int     `json:"HTTP_API"`
		WifiConfig int     `json:"WifiConfig"`
		WifiPower  float64 `json:"WifiPower"`
	} `json:"StatusNET"`
	StatusMQT struct {
		MqttHost       string `json:"MqttHost"`
		MqttPort       int    `json:"MqttPort"`
		MqttClientMask string `json:"MqttClientMask"`
		MqttClient     string `json:"MqttClient"`
		MqttUser       string `json:"MqttUser"`
		MqttCount      int    `json:"MqttCount"`
		MaxPacketSize  int    `json:"MAX_PACKET_SIZE"`
		Keepalive      int    `json:"KEEPALIVE"`
		SocketTimeout  int    `json:"SOCKET_TIMEOUT"`
	} `json:"StatusMQT"`
	StatusTIM struct {
		Utc      string `json:"UTC"`
		Local    string `json:"Local"`
		StartDST string `json:"StartDST"`
		EndDST   string `json:"EndDST"`
		Timezone string `json:"Timezone"`
		Sunrise  string `json:"Sunrise"`
		Sunset   string `json:"Sunset"`
	} `json:"StatusTIM"`
	StatusSNS struct {
		Time string `json:"Time"`
	} `json:"StatusSNS"`
	StatusSTS struct {
		Time      string `json:"Time"`
		Uptime    string `json:"Uptime"`
		UptimeSec int    `json:"UptimeSec"`
		Heap      int    `json:"Heap"`
		SleepMode string `json:"SleepMode"`
		Sleep     int    `json:"Sleep"`
		LoadAvg   int    `json:"LoadAvg"`
		MqttCount int    `json:"MqttCount"`
		Power     string `json:"POWER"`
		Wifi      struct {
			Ap        int    `json:"AP"`
			SSID      string `json:"SSId"`
			BSSID     string `json:"BSSId"`
			Channel   int    `json:"Channel"`
			Mode      string `json:"Mode"`
			Rssi      int    `json:"RSSI"`
			Signal    int    `json:"Signal"`
			LinkCount int    `json:"LinkCount"`
			Downtime  string `json:"Downtime"`
		} `json:"Wifi"`
	} `json:"StatusSTS"`
}
