package hosts

import (
	params "github.com/redhat-developer/mapt/cmd/mapt/cmd/constants"
	maptContext "github.com/redhat-developer/mapt/pkg/manager/context"
	"github.com/redhat-developer/mapt/pkg/provider/aws/action/rhel"
	"github.com/redhat-developer/mapt/pkg/provider/util/instancetypes"
	"github.com/redhat-developer/mapt/pkg/util"
	"github.com/redhat-developer/mapt/pkg/util/ghactions"
	"github.com/redhat-developer/mapt/pkg/util/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	cmdRHEL     = "rhel"
	cmdRHELDesc = "manage rhel dedicated host"

	rhelVersion        string = "version"
	rhelVersionDesc    string = "version for the RHEL OS"
	rhelVersionDefault string = "9.4"
	subsUsername       string = "rh-subscription-username"
	subsUsernameDesc   string = "username to register the subscription"
	subsUserpass       string = "rh-subscription-password"
	subsUserpassDesc   string = "password to register the subscription"
	profileSNC         string = "snc"
	profileSNCDesc     string = "if this flag is set the RHEL will be setup with SNC profile. Setting up all requirements to run https://github.com/crc-org/snc"
)

func GetRHELCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   cmdRHEL,
		Short: cmdRHELDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
	}
	c.AddCommand(getRHELCreate(), getRHELDestroy())
	return c
}

func getRHELCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   params.CreateCmdName,
		Short: params.CreateCmdName,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			// Initialize context
			maptContext.Init(
				viper.GetString(params.ProjectName),
				viper.GetString(params.BackedURL),
				viper.GetString(params.ConnectionDetailsOutput),
				viper.GetStringMapString(params.Tags))

			// Initialize gh actions runner if needed
			if viper.IsSet(params.InstallGHActionsRunner) {
				err := ghactions.InitGHRunnerArgs(viper.GetString(params.GHActionsRunnerToken),
					viper.GetString(params.GHActionsRunnerName),
					viper.GetString(params.GHActionsRunnerRepo),
					viper.GetStringSlice(params.GHActionsRunnerLabels))
				if err != nil {
					logging.Fatal(err)
				}
			}

			instanceRequest := &instancetypes.AwsInstanceRequest{
				CPUs:       viper.GetInt32(params.CPUs),
				MemoryGib:  viper.GetInt32(params.Memory),
				Arch:       util.If(viper.GetString(params.LinuxArch) == "arm64", instancetypes.Arm64, instancetypes.Amd64),
				NestedVirt: viper.GetBool(profileSNC) || viper.GetBool(params.NestedVirt),
			}

			// Run create
			if err := rhel.Create(
				&rhel.Request{
					Prefix:               "main",
					Version:              viper.GetString(rhelVersion),
					Arch:                 viper.GetString(params.LinuxArch),
					InstanceRequest:      instanceRequest,
					VMType:               viper.GetStringSlice(vmTypes),
					SubsUsername:         viper.GetString(subsUsername),
					SubsUserpass:         viper.GetString(subsUserpass),
					ProfileSNC:           viper.IsSet(profileSNC),
					Spot:                 viper.IsSet(spot),
					Airgap:               viper.IsSet(airgap),
					SetupGHActionsRunner: viper.IsSet(params.InstallGHActionsRunner),
				}); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
	flagSet := pflag.NewFlagSet(params.CreateCmdName, pflag.ExitOnError)
	flagSet.StringP(params.ConnectionDetailsOutput, "", "", params.ConnectionDetailsOutputDesc)
	flagSet.StringToStringP(params.Tags, "", nil, params.TagsDesc)
	flagSet.StringP(rhelVersion, "", rhelVersionDefault, rhelVersionDesc)
	flagSet.StringP(params.LinuxArch, "", params.LinuxArchDefault, params.LinuxArchDesc)
	flagSet.StringSliceP(vmTypes, "", []string{}, vmTypesDescription)
	flagSet.StringP(subsUsername, "", "", subsUsernameDesc)
	flagSet.StringP(subsUserpass, "", "", subsUserpassDesc)
	flagSet.Bool(airgap, false, airgapDesc)
	flagSet.Bool(spot, false, spotDesc)
	flagSet.Bool(profileSNC, false, profileSNCDesc)
	flagSet.AddFlagSet(params.GetGHActionsFlagset())
	flagSet.AddFlagSet(params.GetCpusAndMemoryFlagset())
	c.PersistentFlags().AddFlagSet(flagSet)
	// if err := c.MarkFlagRequired(subsUsername); err != nil {
	// 	logging.Error(err)
	// 	return nil
	// }
	// if err := c.MarkFlagRequired(subsUserpass); err != nil {
	// 	logging.Error(err)
	// 	return nil
	// }
	return c
}

func getRHELDestroy() *cobra.Command {
	c := &cobra.Command{
		Use:   params.DestroyCmdName,
		Short: params.DestroyCmdName,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			maptContext.InitBase(
				viper.GetString(params.ProjectName),
				viper.GetString(params.BackedURL))

			if err := rhel.Destroy(); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
	return c
}
