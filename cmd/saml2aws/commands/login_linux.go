package commands

import (
	"github.com/chikin-4x/saml2aws/helper/credentials"
	"github.com/chikin-4x/saml2aws/helper/linuxkeyring"
)

func init() {
	if keyringHelper, err := linuxkeyring.NewKeyringHelper(); err == nil {
		credentials.CurrentHelper = keyringHelper
	}
}
