package app

import (
	"errors"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
	"path/filepath"
)

// wel instance
type Instance struct {
	Id       string            `yaml:"id" json:"id"`
	Name     string            `yaml:"name" json:"name"`
	Options  map[string]string `yaml:"options" json:"options"`
	Interval int               `yaml:"interval" json:"interval"` // seconds

	Operations []*Operation `yaml:"operations" json:"operations"`

	Thresholds []*Threshold   `yaml:"thresholds" json:"thresholds"`
	Dashboards []*Dashboard   `yaml:"dashboards" json:"dashboards"`
	Charts     []*ChartCanvas `yaml:"charts" json:"charts"`
}

// load instance from file
func LoadInstanceFromPath(path string) (*Instance, error) {
	if len(path) == 0 {
		return nil, errors.New("path should not be empty")
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	instance := new(Instance)
	err = yaml.Unmarshal(data, instance)
	if err != nil {
		return nil, err
	}

	return instance, err
}

// create a new instance
func NewInstance() *Instance {
	return &Instance{
		Options:    map[string]string{},
		Operations: []*Operation{},
		Thresholds: []*Threshold{},
		Dashboards: []*Dashboard{},
		Charts:     []*ChartCanvas{},
	}
}

func (this *Instance) FindDashboard(dashboardId string) *Dashboard {
	for _, dashboard := range this.Dashboards {
		if dashboard.Id == dashboardId {
			return dashboard
		}
	}
	return nil
}

// write to a file
func (this *Instance) Write(file string) error {
	data, err := yaml.Marshal(this)
	if err != nil {
		return err
	}
	dir := filepath.Dir(file)
	_, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0777)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return ioutil.WriteFile(file, data, 0666)
}
