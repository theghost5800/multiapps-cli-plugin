package commands

import (
	"context"
	"os/exec"
	"time"

	"github.com/cloudfoundry-incubator/multiapps-cli-plugin/log"
	pb "github.com/cloudfoundry-incubator/multiapps-cli-plugin/test-proto"
	"github.com/cloudfoundry-incubator/multiapps-cli-plugin/ui"
	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/plugin"
	"google.golang.org/grpc"
)

// MtasLocalCommand is a command for listing all deployed MTAs
type MtasLocalCommand struct {
	BaseCommand
}

// GetPluginCommand returns the plugin command details
func (c *MtasLocalCommand) GetPluginCommand() plugin.Command {
	return plugin.Command{
		Name:     "mtas-local",
		HelpText: "List all multi-target apps from local rpc server",
		UsageDetails: plugin.Usage{
			Usage: "cf mtas-local [-u URL]",
			Options: map[string]string{
				"u": "Deploy service URL, by default 'deploy-service.<system-domain>'",
			},
		},
	}
}

// Execute executes the command
func (c *MtasLocalCommand) Execute(args []string) ExecutionStatus {
	log.Tracef("Executing command '"+c.name+"': args: '%v'\n", args)

	localServer := exec.Command("java", "-jar", "path-to-your-built-jar")
	err := localServer.Start()
	if err != nil {
		ui.Failed("Fail to run local rpc server: %s", err)
		return Failure
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		ui.Failed("did not connect: %v", err)
		return Failure
	}

	client := pb.NewMtasClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var host string

	// Parse command arguments and check for required options
	flags, err := c.CreateFlags(&host, args)
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}
	parser := NewCommandFlagsParser(flags, NewDefaultCommandFlagsParser(nil), NewDefaultCommandFlagsValidator(nil))
	err = parser.Parse(args)
	if err != nil {
		c.Usage(err.Error())
		return Failure
	}

	commandContext, err := c.GetContext()
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}

	token, err := c.cliConnection.AccessToken()
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}

	apiEndpoint, err := c.cliConnection.ApiEndpoint()
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}

	ui.Say("Getting multi-target apps in org %s / space %s as %s...",
		terminal.EntityNameColor(commandContext.Org), terminal.EntityNameColor(commandContext.Space), terminal.EntityNameColor(commandContext.Username))

	r2, err2 := client.GetMtas(ctx, &pb.MtasRequest{OrganizationName: commandContext.Org, SpaceName: commandContext.Space, ApiUrl: apiEndpoint, Token: token})
	if err2 != nil {
		ui.Failed("could not greet: %v", err2.Error())
		errKill := localServer.Process.Kill()
		if errKill != nil {
			ui.Say("Not able to kill child process:%s", errKill.Error())
		}
		return Failure
	}

	mtas := r2.Mtas
	if len(mtas) > 0 {
		table := ui.Table([]string{"mta id", "version"})
		for _, mta := range mtas {
			table.Add(mta.Id, mta.Version)
		}
		table.Print()
	} else {
		ui.Say("No multi-target apps found")
	}

	errKill := localServer.Process.Kill()
	if errKill != nil {
		ui.Say("Not able to kill child process:%s", errKill.Error())
	}

	// ui.Say("LocalServer exit code %d and PID:%d", localServer.ProcessState.ExitCode(), localServer.ProcessState.Pid)

	return Success
}
