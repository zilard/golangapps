
package manual

import (
	"bytes"
	stdcontext "context"
	"fmt"
	"strings"

	"github.com/zilard/golangapps/cloudmanager/errors"
	"github.com/zilard/golangapps/cloudmanager/jsonschema"
	"github.com/zilard/golangapps/cloudmanager/utils/v3/ssh"

	"github.com/zilard/golangapps/cloudmanager/cloud"
	"github.com/zilard/golangapps/cloudmanager/environs"
	environscloudspec "github.com/zilard/golangapps/cloudmanager/environs/cloudspec"
	"github.com/zilard/golangapps/cloudmanager/environs/config"
	"github.com/zilard/golangapps/cloudmanager/environs/context"
	"github.com/zilard/golangapps/cloudmanager/environs/manual/sshprovisioner"
)

// ManualProvider contains the logic for using a random ubuntu machine as a 
// controller, connected via SSH
type ManualProvider struct {
	environProviderCredentials
	ping func(endpoint string) error
}

// Verify that we conform to the interface
var _ environs.EnvironProvider = (*ManualProvider)(nil)


var initUbuntuUser = sshprovisioner.InitUbuntuUser


