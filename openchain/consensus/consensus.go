/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package consensus

import (
	pb "github.com/openblockchain/obc-peer/protos"

	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// =============================================================================
// Init.
// =============================================================================

// Package-level logger.
var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("consensus")
}

// =============================================================================
// Interface definitions go here.
// =============================================================================

// Consenter should be implemented by every consensus algorithm implementation (plugin).
type Consenter interface {
	RecvMsg(msg *pb.OpenchainMessage) error // Called by the helper's `HandleMsg()`. This is where the message processing happens.
}

// CPI (Consensus Programming Interface)
type CPI interface {
	SetConsenter(c Consenter)                                             // Is called by the controller package, links `helper` with the plugin package.
	HandleMsg(msg *pb.OpenchainMessage) error                             // Routes to the Consenter's `RecvMsg()`.
	Broadcast(msgPayload []byte) error                                    // May be called by the Consenter's `RecvMsg()` after the processing is done.
	ExecTXs(ctx context.Context, txs []*pb.Transaction) ([]byte, []error) // Is called by the Consenter's `RecvMsg()` during processing.
	Unicast(msgPayload []byte, receiver string) error                     // May be called by the Consenter's `RecvMsg()` after the processing is done.
}
