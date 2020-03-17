package commands

import (
	"github.com/chikin-4x/saml2aws/helper/credentials"
	"github.com/chikin-4x/saml2aws/helper/wincred"
)

func init() {
	credentials.CurrentHelper = &wincred.Wincred{}
}
