// Copyright 2025 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mocks

//go:generate $MOCKGEN -typed -source=../database/database.go -destination=dbclient.go -package mocks github.com/Azure/ARO-HCP/internal/database DBClient
//go:generate $MOCKGEN -typed -source=../database/lock.go -destination=lock.go -package mocks github.com/Azure/ARO-HCP/internal/database LockClientInterface
//go:generate $MOCKGEN -typed -source=../database/transaction.go -destination=dbtransaction.go -package mocks github.com/Azure/ARO-HCP/internal/database DBTransaction DBTransactionResult
//go:generate $MOCKGEN -typed -source=../ocm/ocm.go -destination=ocm.go -package mocks github.com/Azure/ARO-HCP/internal/ocm ClusterServiceClientSpec
