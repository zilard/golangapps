package manual

import "github.com/zilard/golangapps/jujucloudmanager/environs"

const (
	providerType = "manual"
)


func init() {
	p := ManualProvider{}
	environs.RegisterProvider(providerType, p, "null")
}



