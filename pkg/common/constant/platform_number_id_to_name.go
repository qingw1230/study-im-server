package constant

const (
	WindowsPlatformId = 1
	LinuxPlatformId   = 2

	WindowsPlatformStr = "Windows"
	LinuxPlatformStr   = "Linux"
)

var platformId2Name = map[int32]string{
	WindowsPlatformId: WindowsPlatformStr,
	LinuxPlatformId:   LinuxPlatformStr,
}

var platformName2Id = map[string]int32{
	WindowsPlatformStr: WindowsPlatformId,
	LinuxPlatformStr:   LinuxPlatformId,
}

func PlatformIdToName(num int32) string {
	return platformId2Name[num]
}

func PlatformNameToId(name string) int32 {
	return platformName2Id[name]
}
