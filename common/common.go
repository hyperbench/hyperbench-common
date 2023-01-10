package common

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
 * @brief define the common struct
 * @file common.go
 * @author: wangxiaohui
 * @date 2022-01-18
 */

// Invoke define need filed for invoke contract.
type Invoke struct {
	Func string        `mapstructure:"func"`
	Args []interface{} `mapstructure:"args"`
}

// Transfer define need filed for transfer.
type Transfer struct {
	From   string `mapstructure:"from"`
	To     string `mapstructure:"to"`
	Amount int64  `mapstructure:"amount"`
	Extra  string `mapstructure:"extra"`
}

// Query define need filed for query info.
type Query struct {
	Func string        `mapstructure:"func"`
	Args []interface{} `mapstructure:"args"`
}

// Option for receive options.
type Option map[string]interface{}

// Context the context in vm.
type Context string

// Statistic contains statistic time.
type Statistic struct {
	From *ChainInfo `mapstructure:"from"`
	To   *ChainInfo `mapstructure:"to"`
}

// ChainInfo contains txCount and blockHeight.
type ChainInfo struct {
	TxCount     uint64 `mapstructure:"txCount"`
	BlockHeight int64  `mapstructure:"blockHeight"`
	TimeStamp   int64  `mapstructure:"timeStamp"`
}
