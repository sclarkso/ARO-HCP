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

//go:build E2Etests

package e2e

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ARO-HCP E2E Tests")
}

var _ = BeforeSuite(func() {
	if err := setup(context.Background()); err != nil {
		panic(err)
	}
})

var _ = AfterSuite(func() {
	// Cleanup is done by Resource Group DeferCleanup
	dumpMaestroResources()
})

// dumpMaestroResources dumps the Maestro DB resources table to the artifact
// directory for post-mortem analysis of status feedback delivery issues.
func dumpMaestroResources() {
	artifactDir := os.Getenv("ARTIFACT_DIR")
	if artifactDir == "" {
		return
	}
	cmd := exec.Command("kubectl", "exec", "-n", "maestro",
		"deployment/maestro-db", "-c", "postgresql", "--",
		"bash", "-c",
		`psql -U "$POSTGRES_USER" "$POSTGRES_DB" -c "COPY (SELECT id, name, consumer_name, source, status, created_at, updated_at FROM resources) TO STDOUT WITH CSV HEADER;"`,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "maestro db dump failed (non-fatal): %v\n", err)
		return
	}
	outPath := filepath.Join(artifactDir, "maestro-resources.csv")
	if err := os.WriteFile(outPath, output, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write maestro db dump: %v\n", err)
	}
}
