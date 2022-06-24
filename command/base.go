package command

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
	"strings"

	"github.com/dihedron/seal/logging"
	"github.com/dihedron/seal/logging/console"
	"github.com/dihedron/seal/logging/file"
	"github.com/dihedron/seal/logging/noop"
	"github.com/dihedron/seal/logging/uber"
)

type Command struct {
	// LogLevel sets the debugging level of the application.
	LogLevel string `short:"D" long:"log-level" description:"The debug level of the application." optional:"yes" choice:"off" choice:"trace" choice:"debug" choice:"info" choice:"warn" choice:"error" default:"info" env:"SEAL_LOG_LEVEL"`
	// LogStream is the type of logger to use.
	LogStream string `short:"L" long:"log-stream" description:"The logger to use." optional:"yes" choice:"zap" choice:"stdout" choice:"stderr" choice:"file" choice:"log" choice:"none" default:"stderr" env:"SEAL_LOG_STREAM"`
	// CPUProfile sets the (optional) path of the file for CPU profiling info.
	CPUProfile string `short:"C" long:"cpu-profile" description:"The (optional) path where the CPU profiler will store its data." optional:"yes"`
	// MemProfile sets the (optional) path of the file for memory profiling info.
	MemProfile string `short:"M" long:"mem-profile" description:"The (optional) path where the memory profiler will store its data." optional:"yes"`
	// AutomationFriendly enables automation-friendly JSON output.
	AutomationFriendly bool `short:"A" long:"automation-friendly" description:"Whether to output in automation friendly JSON format." optional:"yes"`
	// Parameters are a set of <key>:<value> or <key>=<value> pairs, that ca be used for substitution in inputs.
	Parameters []Parameter `short:"P" long:"parameter" description:"A set of parameters, in <key>:<value> format." optional:"yes"`
}

// InitLogger initialises the logger.
func (cmd *Command) InitLogger(global bool) logging.Logger {
	switch cmd.LogLevel {
	case "trace":
		logging.SetLevel(logging.LevelTrace)
	case "debug":
		logging.SetLevel(logging.LevelDebug)
	case "info":
		logging.SetLevel(logging.LevelInfo)
	case "warn":
		logging.SetLevel(logging.LevelWarn)
	case "error":
		logging.SetLevel(logging.LevelError)
	case "off":
		logging.SetLevel(logging.LevelOff)
	}
	var logger logging.Logger = &noop.Logger{}
	switch cmd.LogStream {
	case "none":
		logger = &noop.Logger{}
	case "stdout":
		logger = console.NewLogger(console.StdOut)
	case "stderr":
		logger = console.NewLogger(console.StdErr)
	case "zap":
		logger, _ = uber.NewLogger()
	case "file":
		exe, _ := os.Executable()
		log := fmt.Sprintf("%s-%d.log", strings.Replace(exe, ".exe", "", -1), os.Getpid())
		logger = file.NewLogger(log)
	}
	if global {
		logging.SetLogger(logger)
	}
	return logger
}

func (cmd *Command) ProfileCPU() *Closer {
	logger := logging.GetLogger()
	var f *os.File
	if cmd.CPUProfile != "" {
		var err error
		f, err = os.Create(cmd.CPUProfile)
		if err != nil {
			logger.Errorf("could not create CPU profile at %s: %v", cmd.CPUProfile, err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			logger.Errorf("could not start CPU profiler: %v", err)
		}
	}
	return &Closer{
		file: f,
	}
}

func (cmd *Command) ProfileMemory() {
	logger := logging.GetLogger()
	if cmd.MemProfile != "" {
		f, err := os.Create(cmd.MemProfile)
		if err != nil {
			logger.Errorf("could not create memory profile at %s: %v", cmd.MemProfile, err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			logger.Errorf("could not write memory profile: %v", err)
		}
	}
}

// BindVariables provides a way to bind variables in strings with
// parameters provided either in the environment, or on the command
// line as '--parameter' pairs. In order to bind all variables in a
// string do as follows:
//   cmd.Bind("{env:SOME_ENV_VAR}={cli:key1}+{cli:key2}")
// and it will be bound according to the following rules:
// * variables of the form {env:<ENVVAR>} are bound to the corresponding
//   environment variable
// * variables of the forma {cli:<PARAM>} are bound to the corresponding
//   parameter as provided on the command line via the command line switch.
func (cmd *Command) BindVariables(value string) string {
	logger := logging.GetLogger()
	result := value
	re := regexp.MustCompile(`(?:\{(env|cli)\:([a-zA-Z0-9-_@\.]+)\})`)
	groups := re.FindAllStringSubmatch(value, -1)
	if groups == nil {
		// no match, no need to bind
		return ""
	}
	for _, group := range groups {
		logger.Debugf("variable match: '%s' '%s' '%s'\n", group[0], group[1], group[2])
		switch group[1] {
		case "cli":
			for _, parameter := range cmd.Parameters {
				if group[2] == parameter.Key {
					logger.Debugf("replacing variable '%s' with value '%s' in '%s'\n", group[0], parameter.Value, result)
					result = strings.ReplaceAll(result, group[0], parameter.Value)
					continue
				}
			}
		case "env":
			e := os.Getenv(group[2])
			logger.Debugf("replacing variable '%s' with value '%s' in '%s'\n", group[0], e, result)
			result = strings.ReplaceAll(result, group[0], e)
		}
	}
	logger.Debugf("'%s'", result)
	return result
}

// Parameter represents a parameter that can be used to
// values in configuration files.
type Parameter struct {
	Key   string
	Value string
}

func (p *Parameter) UnmarshalFlag(value string) error {
	re := regexp.MustCompile(`^([a-zA-Z0-9-_@\.]+)(?:\:|=)(.*)$`)
	matches := re.FindStringSubmatch(value)
	if matches == nil {
		return fmt.Errorf("invalid format for parameter '%s'", value)
	}
	p.Key = matches[1]
	p.Value = matches[2]
	return nil
}

type Closer struct {
	file *os.File
}

func (c *Closer) Close() {
	if c.file != nil {
		pprof.StopCPUProfile()
		c.file.Close()
	}
}
