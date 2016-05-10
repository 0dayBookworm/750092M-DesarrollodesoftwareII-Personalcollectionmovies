package util

var (
	Ip          string
	Port        string
	ServicePath string

	PublicName string
	PublicPath string

	TemplatesPath string
	Home          string
)

func InitServiceInformation(pIP, pPort, pServicePath string) {
	Ip = pIP
	Port = pPort
	ServicePath = pServicePath
}

func InitPublic(pPublicName, pPublicPath string) {
	PublicName = pPublicName
	PublicPath = pPublicPath
}

func InitTemplates(pTemplatesPath, pHome string) {
	TemplatesPath = pTemplatesPath
	Home = pHome
}
