package services

import (
	"fmt"

	azparams "github.com/redhat-developer/mapt/cmd/mapt/cmd/azure/constants"
	params "github.com/redhat-developer/mapt/cmd/mapt/cmd/constants"
	maptContext "github.com/redhat-developer/mapt/pkg/manager/context"
	azureAKS "github.com/redhat-developer/mapt/pkg/provider/azure/action/aks"
	spotAzure "github.com/redhat-developer/mapt/pkg/spot/azure"
	"github.com/redhat-developer/mapt/pkg/util/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	cmdAKS     = "aks"
	cmdAKSDesc = "aks operations"

	paramVersion            = "version"
	paramVersionDesc        = "AKS K8s cluster version"
	paramVMSizeDesc         = "VMSize to be used on the user pool. Typically this is used to provision spot node pools"
	defaultVersion          = "1.30"
	paramOnlySystemPool     = "only-system-pool"
	paramOnlySystemPoolDesc = "if we do not need bunch of resources we can run only the systempool. More info https://learn.microsoft.com/es-es/azure/aks/use-system-pools?tabs=azure-cli#system-and-user-node-pools"
)

func GetAKSCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   cmdAKS,
		Short: cmdAKSDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
	}
	c.AddCommand(getCreateAKS(), getDestroyAKS())
	return c
}

func getCreateAKS() *cobra.Command {
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

			// ParseEvictionRate
			var spotToleranceValue = spotAzure.DefaultEvictionRate
			if viper.IsSet(azparams.ParamSpotTolerance) {
				var ok bool
				spotToleranceValue, ok = spotAzure.ParseEvictionRate(
					viper.GetString(azparams.ParamSpotTolerance))
				if !ok {
					return fmt.Errorf("%s is not a valid spot tolerance value", viper.GetString(azparams.ParamSpotTolerance))
				}
			}

			if err := azureAKS.Create(
				&azureAKS.AKSRequest{
					Prefix:            viper.GetString(params.ProjectName),
					Location:          viper.GetString(azparams.ParamLocation),
					KubernetesVersion: viper.GetString(paramVersion),
					OnlySystemPool:    viper.IsSet(paramOnlySystemPool),
					VMSize:            viper.GetString(azparams.ParamVMSize),
					Spot:              viper.IsSet(azparams.ParamSpot),
					SpotTolerance:     spotToleranceValue}); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
	flagSet := pflag.NewFlagSet(params.CreateCmdName, pflag.ExitOnError)
	flagSet.StringP(params.ConnectionDetailsOutput, "", "", params.ConnectionDetailsOutputDesc)
	flagSet.StringToStringP(params.Tags, "", nil, params.TagsDesc)
	flagSet.StringP(azparams.ParamLocation, "", azparams.DefaultLocation, azparams.ParamLocationDesc)
	flagSet.StringP(azparams.ParamVMSize, "", azparams.DefaultVMSize, paramVMSizeDesc)
	flagSet.StringP(paramVersion, "", defaultVersion, paramVersionDesc)
	flagSet.Bool(azparams.ParamSpot, false, azparams.ParamSpotDesc)
	flagSet.Bool(paramOnlySystemPool, false, paramOnlySystemPoolDesc)
	flagSet.StringP(azparams.ParamSpotTolerance, "", azparams.DefaultSpotTolerance, azparams.ParamSpotToleranceDesc)
	c.PersistentFlags().AddFlagSet(flagSet)
	return c
}

func getDestroyAKS() *cobra.Command {
	return &cobra.Command{
		Use:   params.DestroyCmdName,
		Short: params.DestroyCmdName,
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
			if err := azureAKS.Destroy(); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
}