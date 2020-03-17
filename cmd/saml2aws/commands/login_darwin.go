package commands

import (
	"github.com/chikin-4x/saml2aws/helper/credentials"
	"github.com/chikin-4x/saml2aws/helper/osxkeychain"
)

func init() {
	credentials.CurrentHelper = &osxkeychain.Osxkeychain{}
}
