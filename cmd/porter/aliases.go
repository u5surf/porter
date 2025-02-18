package main

import (
	"strings"

	"github.com/deislabs/porter/pkg/porter"
	"github.com/spf13/cobra"
)

func buildAliasCommands(p *porter.Porter) []*cobra.Command {
	return []*cobra.Command{
		buildCreateAlias(p),
		buildBuildAlias(p),
		buildInstallAlias(p),
		buildUpgradeAlias(p),
		buildUninstallAlias(p),
		buildInvokeAlias(p),
		buildPublishAlias(p),
		buildListAlias(p),
		buildShowAlias(p),
		buildArchiveAlias(p),
	}
}

func buildCreateAlias(p *porter.Porter) *cobra.Command {
	cmd := buildBundleCreateCommand(p)
	cmd.Example = strings.Replace(cmd.Example, "porter bundle create", "porter create", -1)
	cmd.Annotations = map[string]string{
		"group": "alias",
	}
	return cmd
}

func buildBuildAlias(p *porter.Porter) *cobra.Command {
	cmd := buildBundleBuildCommand(p)
	cmd.Example = strings.Replace(cmd.Example, "porter bundle build", "porter build", -1)
	cmd.Annotations = map[string]string{
		"group": "alias",
	}
	return cmd
}

func buildInstallAlias(p *porter.Porter) *cobra.Command {
	cmd := buildBundleInstallCommand(p)
	cmd.Example = strings.Replace(cmd.Example, "porter bundle install", "porter install", -1)
	cmd.Annotations = map[string]string{
		"group": "alias",
	}
	return cmd
}

func buildUpgradeAlias(p *porter.Porter) *cobra.Command {
	cmd := buildBundleUpgradeCommand(p)
	cmd.Example = strings.Replace(cmd.Example, "porter bundle upgrade", "porter upgrade", -1)
	cmd.Annotations = map[string]string{
		"group": "alias",
	}
	return cmd
}

func buildInvokeAlias(p *porter.Porter) *cobra.Command {
	cmd := buildBundleInvokeCommand(p)
	cmd.Example = strings.Replace(cmd.Example, "porter bundle invoke", "porter invoke", -1)
	cmd.Annotations = map[string]string{
		"group": "alias",
	}
	return cmd
}

func buildUninstallAlias(p *porter.Porter) *cobra.Command {
	cmd := buildBundleUninstallCommand(p)
	cmd.Example = strings.Replace(cmd.Example, "porter bundle uninstall", "porter uninstall", -1)
	cmd.Annotations = map[string]string{
		"group": "alias",
	}
	return cmd
}

func buildPublishAlias(p *porter.Porter) *cobra.Command {
	cmd := buildBundlePublishCommand(p)
	cmd.Example = strings.Replace(cmd.Example, "porter bundle publish", "porter publish", -1)
	cmd.Annotations = map[string]string{
		"group": "alias",
	}
	return cmd
}

func buildShowAlias(p *porter.Porter) *cobra.Command {
	cmd := buildInstanceShowCommand(p)
	cmd.Example = strings.Replace(cmd.Example, "porter bundle instance show", "porter show", -1)
	cmd.Annotations = map[string]string{
		"group": "alias",
	}
	return cmd
}

func buildListAlias(p *porter.Porter) *cobra.Command {
	cmd := buildInstancesListCommand(p)
	cmd.Example = strings.Replace(cmd.Example, "porter bundle instances list", "porter list", -1)
	cmd.Annotations = map[string]string{
		"group": "alias",
	}
	return cmd
}

func buildArchiveAlias(p *porter.Porter) *cobra.Command {
	cmd := buildBundleArchiveCommand(p)
	cmd.Example = strings.Replace(cmd.Example, "porter bundle archive", "porter archive", -1)
	cmd.Annotations = map[string]string{
		"group": "alias",
	}
	return cmd
}
