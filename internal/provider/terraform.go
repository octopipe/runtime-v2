package provider

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
	commonv1alpha1 "github.com/octopipe/cloudx/apis/common/v1alpha1"
)

type terraformProvider struct {
	execPath string
}

func NewTerraformProvider() (Provider, error) {
	err := os.MkdirAll("/tmp/terraform-ins", os.ModePerm)
	if err != nil {
		return nil, err
	}

	installer := &releases.ExactVersion{
		Product:    product.Terraform,
		Version:    version.Must(version.NewVersion("1.0.6")),
		InstallDir: "/tmp/terraform-ins",
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		return nil, err
	}

	return terraformProvider{
		execPath: execPath,
	}, nil
}

func (p terraformProvider) Apply(workdirPath string, input map[string]interface{}) ([]commonv1alpha1.StackSetPluginOutput, string, error) {
	tf, err := tfexec.NewTerraform(workdirPath, p.execPath)
	if err != nil {
		return nil, "", err
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		return nil, "", err
	}

	execVarsFile := filepath.Join(workdirPath, "exec.tfvars")
	f, err := os.Create(execVarsFile)
	if err != nil {
		return nil, "", err
	}

	for key, val := range input {
		f.WriteString(fmt.Sprintf("%s = \"%s\"\n", key, val))
	}

	err = tf.Apply(context.Background(), tfexec.VarFile(execVarsFile))
	if err != nil {
		return nil, "", err
	}

	out, err := tf.Output(context.Background())
	if err != nil {
		return nil, "", err
	}

	outputs := []commonv1alpha1.StackSetPluginOutput{}

	for key, value := range out {
		outputs = append(outputs, commonv1alpha1.StackSetPluginOutput{
			Key:   key,
			Value: string(value.Value),
		})
	}

	stateFile, err := os.ReadFile(fmt.Sprintf("%s/terraform.tfstate", workdirPath))
	if err != nil {
		return nil, "", err
	}

	return outputs, string(stateFile), nil
}

func (p terraformProvider) Destroy(workdirPath string, input map[string]interface{}) error {
	tf, err := tfexec.NewTerraform(workdirPath, p.execPath)
	if err != nil {
		return err
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		return err
	}

	execVarsFile := filepath.Join(workdirPath, "exec.tfvars")
	f, err := os.Create(execVarsFile)
	if err != nil {
		return err
	}

	for key, val := range input {
		f.WriteString(fmt.Sprintf("%s = \"%s\"\n", key, val))
	}

	err = tf.Destroy(context.Background(), tfexec.VarFile(execVarsFile))
	if err != nil {
		return err
	}

	return nil
}
