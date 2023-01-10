package base

/**
 *  Copyright (C) 2021 HyperBench.
 *  SPDX-License-Identifier: Apache-2.0
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 * @brief the base implementation of blockChain
 * @file base.go
 * @author: wangxiaohui
 * @date 2022-01-18
 */

import (
	"github.com/hyperbench/hyperbench-common/common"
	"github.com/op/go-logging"
)

// BlockchainBase the base implementation of blockChain.
type BlockchainBase struct {
	ClientConfig
	Logger *logging.Logger
}

// ClientConfig define the filed for client config.
type ClientConfig struct {
	// client
	ClientType string `mapstructure:"type"`
	ConfigPath string `mapstructure:"config"`

	// contract
	ContractPath string        `mapstructure:"contract"`
	Args         []interface{} `mapstructure:"args"`

	// options
	Options map[string]interface{} `mapstructure:"options"`

	// index
	VmID     int `mapstructure:"vmID"`
	WorkerID int `mapstructure:"workerID"`
}

// Confirm confirm invoke result for tx.
func (b *BlockchainBase) Confirm(result *common.Result, ops ...common.Option) *common.Result {
	return result
}

// DeployContract send tx for deploy contract.
func (b *BlockchainBase) DeployContract() error {
	return nil
}

// Invoke send tx for invoke contract.
func (b *BlockchainBase) Invoke(common.Invoke, ...common.Option) *common.Result {
	return &common.Result{}
}

// Transfer send tx for transfer.
func (b *BlockchainBase) Transfer(common.Transfer, ...common.Option) *common.Result {
	return &common.Result{}
}

// Query query info.
func (b *BlockchainBase) Query(common.Query, ...common.Option) interface{} {
	return nil
}

// Option receive some options.
func (b *BlockchainBase) Option(common.Option) error {
	return nil
}

// GetContext get context for execute tx in vm.
func (b *BlockchainBase) GetContext() (string, error) {
	return "", nil
}

// SetContext set context for execute tx in vm.
func (b *BlockchainBase) SetContext(ctx string) error {
	return nil
}

// ResetContext reset context.
func (b *BlockchainBase) ResetContext() error {
	return nil
}

// Statistic statistic remote execute result.
func (b *BlockchainBase) Statistic(statistic common.Statistic) (*common.RemoteStatistic, error) {
	return &common.RemoteStatistic{}, nil
}

// LogStatus records blockheight and time
func (b *BlockchainBase) LogStatus() (int64, error) {
	return 0, nil
}

// NewBlockchainBase new blockchain base.
func NewBlockchainBase(clientConfig ClientConfig) *BlockchainBase {
	return &BlockchainBase{
		ClientConfig: clientConfig,
		Logger:       common.GetLogger("client"),
	}
}
