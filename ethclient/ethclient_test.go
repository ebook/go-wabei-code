// Copyright 2016 The go- Authors
// This file is part of the go- library.
//
// The go- library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go- library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go- library. If not, see <http://www.gnu.org/licenses/>.

package ethclient

import "github.com/wabei/go-wabei"

// Verify that Client implements the  interfaces.
var (
	_ = .ChainReader(&Client{})
	_ = .TransactionReader(&Client{})
	_ = .ChainStateReader(&Client{})
	_ = .ChainSyncReader(&Client{})
	_ = .ContractCaller(&Client{})
	_ = .GasEstimator(&Client{})
	_ = .GasPricer(&Client{})
	_ = .LogFilterer(&Client{})
	_ = .PendingStateReader(&Client{})
	// _ = .PendingStateEventer(&Client{})
	_ = .PendingContractCaller(&Client{})
)