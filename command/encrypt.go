package command

import (
	"crypto/tls"
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

type Encrypt struct {
	Command
	Recipients []string          `short:"r" long:"recipient" description:"The recipient of the CMS encoded secret message." required:"yes"`
	LDAP       LDAPConfiguration `short:"c" long:"ldap-configuration" description:"The configuration to use to connect to the LDAP server." required:"yes"`
}

func (cmd *Encrypt) Execute(args []string) error {
	logger := cmd.InitLogger(true)

	logger.Debugf("LDAP connection info: %v", cmd.LDAP)

	options := []ldap.DialOpt{}
	logger.Infof("connecting to LDAP server %s", cmd.LDAP.Endpoint.String())
	if cmd.LDAP.Endpoint.Scheme == "ldaps" && cmd.LDAP.Insecure {
		options = append(options, ldap.DialWithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}))
	}
	client, err := ldap.DialURL(cmd.LDAP.Endpoint.String(), options...)
	if err != nil {
		logger.Errorf("error dialling LDAP server: %v", err)
		return err
	}
	defer client.Close()
	if cmd.LDAP.Username != "" && cmd.LDAP.Password != "" {
		logger.Infof("binding to LDAP server with user %s", cmd.LDAP.Username)
		if err = client.Bind(cmd.LDAP.Username, cmd.LDAP.Password); err != nil {
			logger.Errorf("error binding to LDAP server with user %s: %v", cmd.LDAP.Username, err)
			return err
		}
	}
	for _, recipient := range cmd.Recipients {
		logger.Debugf("retrieving certificate for user %s", recipient)
		request := ldap.NewSearchRequest(
			cmd.LDAP.BaseDN, // the base DN to search
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			"(&(objectClass=organizationalPerson))", // the filter to apply
			[]string{"dn", "cn", "userCertificate"}, // a list attributes to retrieve
			nil,
		)

		sr, err := client.Search(request)
		if err != nil {
			logger.Errorf("error performing search for user  %s: %v", recipient, err)
			continue
		}

		for _, entry := range sr.Entries {
			fmt.Printf("%s:\n%v\n", entry.DN, entry.GetAttributeValue("userCertificate"))
		}
	}
	return nil
}
