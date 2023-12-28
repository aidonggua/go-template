package configuration

type Profile string

const (
	ProfileDev     Profile = "dev"
	ProfileProd    Profile = "prod"
	ProfileTest    Profile = "test"
	ProfileDefault Profile = "default"
)
