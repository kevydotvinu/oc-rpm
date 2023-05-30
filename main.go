package main

import (
	"fmt"
	"github.com/openshift/oc/pkg/cli/admin/release"
	imagemanifest "github.com/openshift/oc/pkg/cli/image/manifest"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/templates"
	"os"
)

type CommandOptions struct {
	genericclioptions.IOStreams
	genericclioptions.KubeTemplatePrintFlags

	ParallelOptions imagemanifest.ParallelOptions
	SecurityOptions imagemanifest.SecurityOptions
	FilterOptions   imagemanifest.FilterOptions
}

func RootCommandOptions(streams genericclioptions.IOStreams) *CommandOptions {
	return &CommandOptions{
		IOStreams:              streams,
		KubeTemplatePrintFlags: *genericclioptions.NewKubeTemplatePrintFlags(),
		ParallelOptions:        imagemanifest.ParallelOptions{MaxPerRegistry: 4},
	}
}
func RootCommand(f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	rootOptions := release.NewPackageOptions(streams)
	rootCmd := &cobra.Command{
		Use:   "rpm [RELEASE]",
		Short: "List RPMs used in an OpenShift node. Not for production use.",
		Long:  "Display the list of RPM packages in an OpenShift release. Not for production use.",
		Example: templates.Examples(`
                        # Show the list of RPMs in the cluster's current release
			oc rpm

			# Show the list of RPMs in the release
			oc rpm 4.13.0
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
				return err
			}
			kcmdutil.CheckErr(rootOptions.Complete(f, cmd, args))
			kcmdutil.CheckErr(rootOptions.Validate())
			kcmdutil.CheckErr(rootOptions.Run())
			return nil
		},
	}
	flags := rootCmd.Flags()
	rootOptions.SecurityOptions.Bind(flags)
	rootOptions.FilterOptions.Bind(flags)
	rootOptions.ParallelOptions.Bind(flags)
	return rootCmd
}

func main() {
	kubeConfigFlags := genericclioptions.NewConfigFlags(true).WithDiscoveryBurst(350).WithDiscoveryQPS(50.0)
	f := kcmdutil.NewFactory(kubeConfigFlags)
	rootCmd := RootCommand(f, genericclioptions.IOStreams{})

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
