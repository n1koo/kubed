package cmds

import (
	"os"

	"github.com/appscode/go/flags"
	"github.com/appscode/kubed/pkg/backup"
	"github.com/spf13/cobra"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewCmdBackup() *cobra.Command {
	var opt backup.Options

	cmd := &cobra.Command{
		Use:     "backup",
		Short:   "Takes backup of YAML files of cluster",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			flags.EnsureRequiredFlags(cmd, "context", "backup-dir")

			err := os.MkdirAll(opt.BackupDir, 0777)
			if err != nil {
				return err
			}
			restConfig, err := createKubeConfig(opt.Context)
			if err != nil {
				return err
			}
			return backup.Backup(restConfig, opt)
		},
	}
	cmd.Flags().BoolVar(&opt.Sanitize, "sanitize", false, " Sanitize fields in YAML")
	cmd.Flags().StringVar(&opt.BackupDir, "backup-dir", "", "Directory where YAML files will be saved")
	cmd.Flags().StringVar(&opt.Context, "context", "", "The name of the kubeconfig context to use")
	return cmd
}

func createKubeConfig(ctx string) (*rest.Config, error) {
	apiConfig, err := clientcmd.NewDefaultPathOptions().GetStartingConfig()
	if err != nil {
		return nil, err
	}
	overrides := &clientcmd.ConfigOverrides{CurrentContext: ctx}
	return clientcmd.NewDefaultClientConfig(*apiConfig, overrides).ClientConfig()
}
