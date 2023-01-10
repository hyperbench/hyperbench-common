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
 * @brief Define the structure used in the vm
 * @file result.go
 * @author: wangxiaohui
 * @date 2022-01-18
 */

import (
	"context"
)

// TxContext is the context of transaction (`vm.Run`)
type TxContext struct {
	context.Context
	TxIndex
}

// TxIndex is the unique index of a transaction (`vm.Run`)
type TxIndex struct {
	EngineIdx int64 // EngineIdx is the index of engine (maybe support multiple test stage in future)
	TxIdx     int64 // TxIdx is the index of sent transaction
	MissIdx   int64 // MissIdx is the index of missed transaction
}

// VMContext is the context of vm
type VMContext struct {
	WorkerIdx int64 `mapstructure:"worker"` // WorkerIdx is the index of worker
	VMIdx     int64 `mapstructure:"vm"`     // VMIdx is the index of vm
}
