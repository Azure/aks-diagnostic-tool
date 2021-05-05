package collector

import (
	"path/filepath"

	"github.com/Azure/aks-periscope/pkg/interfaces"
	"github.com/Azure/aks-periscope/pkg/utils"
)

// HelmCollector defines a Helm Collector struct
type HelmCollector struct {
	BaseCollector
}

var _ interfaces.Collector = &IPTablesCollector{}

// NewHelmCollector is a constructor
func NewHelmCollector(exporter interfaces.Exporter) *HelmCollector {
	return &HelmCollector{
		BaseCollector: BaseCollector{
			collectorType: Helm,
			exporter:      exporter,
		},
	}
}

// Collect implements the interface method
func (collector *HelmCollector) Collect() error {
	rootPath, err := utils.CreateCollectorDir(collector.GetName())
	if err != nil {
		return err
	}
	helmListFile := filepath.Join(rootPath, collector.GetName())
	output, err := utils.RunCommandOnHost("helm", "list", "--all-namespaces")
	if err != nil {
		return err
	}
	err = utils.WriteToFile(helmListFile, output)
	if err != nil {
		return err
	}

	collector.AddToCollectorFiles(helmListFile)

	helmHistoryFile := filepath.Join(rootPath, collector.GetName())
	output, err = utils.RunCommandOnHost("helm", "history", "azure-arc")
	if err != nil {
		return err
	}
	err = utils.WriteToFile(helmHistoryFile, output)
	if err != nil {
		return err
	}

	collector.AddToCollectorFiles(helmHistoryFile)
	return nil
}