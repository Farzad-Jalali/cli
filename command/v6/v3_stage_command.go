package v6

import (
	"code.cloudfoundry.org/cli/actor/loggingaction"
	"context"
	"strings"
	"time"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/v6/shared"
)

//go:generate counterfeiter . V3StageActor

type V3StageActor interface {
	GetStreamingLogsForApplicationByNameAndSpace(appName string, spaceGUID string, client loggingaction.LogCacheClient) (<-chan loggingaction.LogMessage, <-chan error, v3action.Warnings, error, context.CancelFunc)
	StagePackage(packageGUID string, appName string) (<-chan v3action.Droplet, <-chan v3action.Warnings, <-chan error)
}

type V3StageCommand struct {
	RequiredArgs        flag.AppName `positional-args:"yes"`
	PackageGUID         string       `long:"package-guid" description:"The guid of the package to stage" required:"true"`
	usage               interface{}  `usage:"CF_NAME v3-stage APP_NAME --package-guid PACKAGE_GUID"`
	envCFStagingTimeout interface{}  `environmentName:"CF_STAGING_TIMEOUT" environmentDescription:"Max wait time for buildpack staging, in minutes" environmentDefault:"15"`

	UI             command.UI
	Config         command.Config
	LogCacheClient loggingaction.LogCacheClient
	SharedActor    command.SharedActor
	Actor          V3StageActor
}

func (cmd *V3StageCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	cmd.SharedActor = sharedaction.NewActor(config)

	ccClient, _, err := shared.NewV3BasedClients(config, ui, true, "")
	if err != nil {
		return err
	}

	cmd.Actor = v3action.NewActor(ccClient, config, nil, nil)
	cmd.LogCacheClient = shared.NewLogCacheClient(ccClient.LogCacheEndpoint(), config)

	return nil
}

func (cmd V3StageCommand) Execute(args []string) error {
	cmd.UI.DisplayWarning(command.ExperimentalWarning)

	err := cmd.SharedActor.CheckTarget(true, true)
	if err != nil {
		return err
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	cmd.UI.DisplayTextWithFlavor("Staging package for {{.AppName}} in org {{.OrgName}} / space {{.SpaceName}} as {{.Username}}...", map[string]interface{}{
		"AppName":   cmd.RequiredArgs.AppName,
		"OrgName":   cmd.Config.TargetedOrganization().Name,
		"SpaceName": cmd.Config.TargetedSpace().Name,
		"Username":  user.Name,
	})

	logStream, logErrStream, logWarnings, logErr, stopStreaming := cmd.Actor.GetStreamingLogsForApplicationByNameAndSpace(cmd.RequiredArgs.AppName, cmd.Config.TargetedSpace().GUID, cmd.LogCacheClient)
	cmd.UI.DisplayWarnings(logWarnings)
	if logErr != nil {
		return logErr
	}
	defer stopStreaming()

	dropletStream, warningsStream, errStream := cmd.Actor.StagePackage(cmd.PackageGUID, cmd.RequiredArgs.AppName)
	var droplet v3action.Droplet
	droplet, err = shared.PollStage(dropletStream, warningsStream, errStream, logStream, logErrStream, cmd.UI)
	if err != nil {
		return err
	}
	stopStreaming()

	cmd.UI.DisplayNewline()
	cmd.UI.DisplayText("Package staged")

	t, err := time.Parse(time.RFC3339, droplet.CreatedAt)
	if err != nil {
		return err
	}

	table := [][]string{
		{cmd.UI.TranslateText("droplet guid:"), droplet.GUID},
		{cmd.UI.TranslateText("state:"), strings.ToLower(string(droplet.State))},
		{cmd.UI.TranslateText("created:"), cmd.UI.UserFriendlyDate(t)},
	}

	cmd.UI.DisplayKeyValueTable("", table, 3)
	return nil
}
