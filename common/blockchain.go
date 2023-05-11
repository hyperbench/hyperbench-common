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
 * @brief define the service need provided in blockchain
 * @file blockchain.go
 * @author: linguopeng
 * @date 2022-01-28
 */

// Blockchain define the service need provided in blockchain.
type Blockchain interface {

	// DeployContract should deploy contract with config file
	DeployContract() error

	// Invoke just invoke the contract
	Invoke(Invoke, ...Option) *Result

	// Transfer a amount of money from a account to the other one
	Transfer(Transfer, ...Option) *Result

	// Confirm check the result of `Invoke` or `Transfer`
	Confirm(*Result, ...Option) *Result

	// Verify check the relative time of transaction
	Verify(*Result, ...Option) *Result

	// InvokeBatch just batch invoke the contract
	InvokeBatch(...Invoke) []*Result

	// TransferBatch send batch transfer txs, which a amount of money from a account to the other one
	TransferBatch(...Transfer) []*Result

	// ConfirmBatch check the result of `Invoke` or `Transfer`
	ConfirmBatch(...*Result) []*Result

	// VerifyBatch check the relative time of transactions
	VerifyBatch(...*Result) []*Result

	// Query do some query
	Query(Query, ...Option) interface{}

	// Option pass the options to affect the action of client
	Option(Option) error

	// GetContext Generate TxContext based on New/Init/DeployContract
	// GetContext will only be run in master
	// return the information how to invoke the contract, maybe include
	// contract address, abi or so.
	// the return value will be send to worker to tell them how to invoke the contract
	GetContext() (string, error)

	// SetContext set test context into go client
	// SetContext will be run once per worker
	SetContext(ctx string) error

	// ResetContext reset test group context in go client
	ResetContext() error

	// Statistic query the statistics information in the time interval defined by
	// nanosecond-level timestamps `from` and `to`
	Statistic(statistic Statistic) (*RemoteStatistic, error)

	// LogStatus records blockheight and time
	LogStatus() (*ChainInfo, error)

	Close()
}
