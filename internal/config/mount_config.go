// Copyright 2021 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

const (
	// Default log rotation config values.
	defaultMaxFileSizeMB   = 512
	defaultBackupFileCount = 10
	defaultCompress        = true
)

type WriteConfig struct {
	CreateEmptyFile bool `yaml:"create-empty-file"`
}

type LogConfig struct {
	Severity        LogSeverity     `yaml:"severity"`
	Format          string          `yaml:"format"`
	FilePath        string          `yaml:"file-path"`
	LogRotateConfig LogRotateConfig `yaml:"log-rotate"`
}

type MountConfig struct {
	WriteConfig `yaml:"write"`
	LogConfig   `yaml:"logging"`
}

// LogRotateConfig defines the parameters for log rotation. It consists of three
// configuration options:
// 1. max-file-size-mb: specifies the maximum size in megabytes that a log file
// can reach before it is rotated. The default value is 512 megabytes.
// 2. backup-file-count: determines the maximum number of backup log files to
// retain after they have been rotated. The default value is 10. When value is
// set to 0, all backup files are retained.
// 3. compress: indicates whether the rotated log files should be compressed
// using gzip. The default value is False.
type LogRotateConfig struct {
	MaxFileSizeMB   int  `yaml:"max-file-size-mb"`
	BackupFileCount int  `yaml:"backup-file-count"`
	Compress        bool `yaml:"compress"`
}

func DefaultLogRotateConfig() LogRotateConfig {
	return LogRotateConfig{
		MaxFileSizeMB:   defaultMaxFileSizeMB,
		BackupFileCount: defaultBackupFileCount,
		Compress:        defaultCompress,
	}
}

func NewMountConfig() *MountConfig {
	mountConfig := &MountConfig{}
	mountConfig.LogConfig = LogConfig{
		// Making the default severity as INFO.
		Severity: INFO,
		// Setting default values of log rotate config.
		LogRotateConfig: DefaultLogRotateConfig(),
	}
	return mountConfig
}
