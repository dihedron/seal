package command

type Commands struct {
	// Check applies a set of rules to the plan and produces a report and an action.
	Encrypt Encrypt `command:"encrypt" alias:"e" description:"Encrypt a messa in a CMS-encoded file."`
}
