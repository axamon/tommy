package tommy

// https://tommyweb.telecomitalia.local/businesspublisher/layouts/default/reportOptions.jsp?IDSConversationID=R_cc&modelId=1351994954604915072&viewKey=
// https://tommygrc.telecomitalia.local:8443/mashzone/login.jsp

import (
	"regexp"
)

var identificativo = regexp.MustCompile(`(?m)\d{19}`)

// GetIdentificativoPDF recupera l'identificativo univoco dei documenti Tommy.
func GetIdentificativoPDF(text string) (string, error) {

	return identificativo.FindString(text), nil
}
