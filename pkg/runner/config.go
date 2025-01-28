package runner

import (
	"github.com/abklabs/svmkit/pkg/runner/deb"
)

const (
	aptLockTimeout = 120
)

type Config struct {
	PackageConfig  *deb.PackageConfig `pulumi:"packageConfig,optional"`
	AptLockTimeout *int               `pulumi:"aptLockTimeout,optional"`
}

func (c *Config) SetDefaults() {
	if c.AptLockTimeout == nil {
		temp := aptLockTimeout
		c.AptLockTimeout = &temp
	}
}

func (c *Config) UpdatePackageGroup(grp *deb.PackageGroup) error {
	if c.PackageConfig == nil {
		return nil
	}

	return c.PackageConfig.UpdatePackageGroup(grp)
}
