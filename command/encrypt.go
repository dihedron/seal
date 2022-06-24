package command

import (
	"crypto/tls"

	"github.com/go-ldap/ldap/v3"
)

type Encrypt struct {
	Command
	Recipients []string `short:"r" long:"recipient" description:"The recipient of the CMS encoded secret message." required:"yes"`
	Address    URL      `short:"l" long:"ldap-address" description:"The address of the LDAP sever." required:"yes"`
	Insecure   bool     `short:"i" long:"insecure-skip-verify" description:"Whether to skip TLS certificate verification." optional:"yes"`
}

func (cmd *Encrypt) Execute(args []string) error {
	logger := cmd.InitLogger(true)
	client, err := ldap.DialURL("ldap://ldap.example.com:389", ldap.DialWithTLSConfig(&tls.Config{
		InsecureSkipVerify: true,
	}))
	if err != nil {
		logger.Errorf("error dialling LDAP server: %v", err)
		return err
	}
	defer client.Close()
	for _, recipient := range cmd.Recipients {
		logger.Debugf("retrieving certificate for user %s", recipient)

	}
	return nil
}

/*

type Check struct {
	Command
	Rules        []Rule `short:"r" long:"rule" description:"The rules(s) to apply." optional:"yes"`
	Verbose      bool   `short:"v" long:"verbose" description:"Print a verbose report including the change contents." optional:"yes"`
	ShowMessages bool   `short:"m" long:"show-messages" description:"Whether to report changes that have associated informational messages." optional:"yes"`
	ShowWarnings bool   `short:"w" long:"show-warnings" description:"Whether to report changes that have associated warning messages." optional:"yes"`
	ShowFailures bool   `short:"f" long:"show-failures" description:"Whether to report changes that have associated failure messages." optional:"yes"`
	DryRun       bool   `short:"d" long:"dry-run" description:"Whether the check should not exit with an error value in case of failed checks (i.e. \"dry run\")." optional:"yes"`
	Format       string `short:"o" long:"format" description:"The output format (JSON, YAML or plain text)." optional:"yes" choice:"json" choice:"yaml" choice:"human" default:"human"`
}

func (cmd *Check) Execute(args []string) error {
	var (
		err  error
		file *os.File
		data []byte
	)
	logger := cmd.InitLogger(true)

	switch len(args) {
	case 0:
		logger.Debug("reading plan from standard input")
		if data, err = ioutil.ReadAll(os.Stdin); err != nil {
			return fmt.Errorf("error reading from standard input: %w", err)
		}
	case 1:
		logger.Debugf("reading plan from file %s", args[0])
		if file, err = os.Open(args[0]); err != nil {
			logger.Errorf("error opening file %s: %v", args[0], err)
			os.Exit(1)
		}
		defer file.Close()
		if data, err = ioutil.ReadAll(file); err != nil {
			return fmt.Errorf("error reading file %s: %w", args[0], err)
		}
	default:
		logger.Errorf("too many files: %v", args)
		return fmt.Errorf("too many file: %v", args)
	}

	logger.Debug("data read from input, size: %d", len(data))

	plan := Plan{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err = json.Unmarshal(data, &plan); err != nil {
		return fmt.Errorf("error unmarshalling input data: %w", err)
	}
	report := &Report{
		Records: []Record{},
	}
	for _, change := range plan.ResourceChanges {
		record := Record{}
		if cmd.Verbose {
			record.Change = &change
		}
		for _, rule := range cmd.Rules {
			outcome := Outcome{
				Rule: rule.code,
			}
			outcome.Result, outcome.Messages, outcome.Warnings, outcome.Failures, err = rule.Apply(change)
			if err != nil {
				logger.Warnf("error applying rule %s: %v\n", rule.code, err)
				continue
			}
			record.Outcomes = append(record.Outcomes, outcome)
		}
		if (!cmd.ShowMessages && !cmd.ShowWarnings && !cmd.ShowFailures) ||
			((cmd.ShowMessages && record.HasMessages()) || (cmd.ShowWarnings && record.HasWarnings()) || (cmd.ShowFailures && record.HasFailures())) {
			report.Records = append(report.Records, record)
		}
	}

	switch cmd.Format {
	case "json", "j":
		fmt.Println(ToJSON(report, true))
	case "yaml", "y":
		fmt.Printf("%s", ToYAML(report))
	case "human", "h":
		for _, record := range report.Records {
			fmt.Printf("----------------------------------------------------------------\n")
			fmt.Printf("change:\n")
			if cmd.Verbose {
				fmt.Printf("  details:\n")
				fmt.Printf("%s\n", ToJSON(record.Change, true))
			}
			if len(record.Outcomes) > 0 {
				fmt.Printf("  rules:\n")
				for _, outcome := range record.Outcomes {
					fmt.Printf("  - rule     : %v\n", outcome.Rule)
					fmt.Printf("    outcome  : %v\n", outcome.Result)
					if len(outcome.Messages) > 0 {
						fmt.Printf("    messages :\n")
						for _, m := range outcome.Messages {
							if isatty.IsTerminal(os.Stdout.Fd()) {
								fmt.Printf("    - message: %s\n", green(m))
							} else {
								fmt.Printf("    - message: %s\n", m)
							}
						}
					}
					if len(outcome.Warnings) > 0 {
						fmt.Printf("  warnings:\n")
						for _, w := range outcome.Warnings {
							if isatty.IsTerminal(os.Stdout.Fd()) {
								fmt.Printf("    - warning: %s\n", yellow(w))
							} else {
								fmt.Printf("    - warning: %s\n", w)
							}
						}
					}
					if len(outcome.Failures) > 0 {
						fmt.Printf("    failures :\n")
						for _, f := range outcome.Failures {
							if isatty.IsTerminal(os.Stdout.Fd()) {
								fmt.Printf("    - failure: %s\n", red(f))
							} else {
								fmt.Printf("    - failure: %s\n", f)
							}
						}
					}
				}
			}
		}
		fmt.Printf("----------------------------------------------------------------\n")
	}
	logger.Debugf("total: %d", len(report.Records))
	if !cmd.DryRun {
		if report.HasFailures() {
			os.Exit(1)
		}
	}
	return nil
}
*/
