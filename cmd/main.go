package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/glorfischi/kswp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Options struct {
	kswp.Kswp
}

func main() {
	o, err := getOptions()
	if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot parse configuration: \n\n%s\n", err.Error())
    os.Exit(1);
	}
  // Main command to execute
	cmd := cobra.Command{
		Use:  "kswp [workspace]",
		Short: "kswp allows for simple management of kubeconfig files",
    DisableFlagsInUseLine: true,
    Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.swp(cmd, args); err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), "%s\n", err.Error())
			}
		},
    ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
      if len(args) != 0 {
        return nil, cobra.ShellCompDirectiveNoFileComp
      }
      var nouns []string
      for _, c := range o.Configs {
        nouns = append(nouns, c.Name)
      }
      return nouns, cobra.ShellCompDirectiveNoFileComp
    },
	}

  // Completion comand. Will not show up in any help screen or auto-completion
	var completionCmd = &cobra.Command{
		Use:   "completion",
		Short: "Generates bash completion scripts",
    DisableFlagsInUseLine: true,
    ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
    Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
    switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, false)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
		},
    Hidden: true,
	}

	cmd.AddCommand(completionCmd)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func (o Options) swp(cmd *cobra.Command, args []string) error {
	return o.Swap(args[0])
}

const kubeConfigKey = "kubeconfig"
const configsKey = "configs"


// Read config 
func getOptions() (Options, error) {
	o := Options{}

	v := viper.New()
  v.SetEnvPrefix("kswp") // prefix kswp for envVars
	v.SetConfigName("kswp")

	home, err := os.UserHomeDir()
	if err == nil {
		v.AddConfigPath(filepath.Join(home, ".kube"))
		v.AddConfigPath(filepath.Join(home, ".config", "kswp"))
	}
	v.AddConfigPath(filepath.Join("etc", "kswp"))
	v.AddConfigPath(".")

	v.SetDefault(kubeConfigKey, filepath.Join(home, ".kube", "config"))
	v.AutomaticEnv()

	err = v.ReadInConfig()
  if err != nil {
		return o, err
	}

	o.KubeConf = v.GetString(kubeConfigKey)
	err = v.UnmarshalKey(configsKey, &o.Configs)
	return o, nil
}


