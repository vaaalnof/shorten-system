package avatar

import (
	"net/url"
	"strings"
)

func DefaultAvatar(
	firstName string,
	lastName string,
) *string {

	initial := ""

	if firstName != "" {
		initial += strings.ToUpper(firstName[:1])
	}

	if lastName != "" {
		initial += strings.ToUpper(lastName[:1])
	}

	avatar := "https://api.dicebear.com/9.x/initials/svg?seed=" +
		url.QueryEscape(initial)

	return &avatar
}
