// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"github.com/google/gce-tcb-verifier/gcetcbendorsement"
	epb "github.com/google/gce-tcb-verifier/proto/endorsement"
	"github.com/google/gce-tcb-verifier/verify"
	"github.com/spf13/cobra"
)

var errNoGetter = errors.New("getter was nil")

type verifyCommand struct {
	root        string
	show        bool
	endorsement *epb.VMLaunchEndorsement
}

type verifyKeyType struct{}

var verifyKey verifyKeyType

func (c *verifyCommand) persistentPreRunE(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("verify expects exactly one argument, got %d", len(args))
	}
	if !c.show {
		c.endorsement = &epb.VMLaunchEndorsement{}
		return ReadProto(cmd.Context(), args[0], c.endorsement)
	}
	return nil
}

func rootOfTrust(ctx context.Context, root string) (*x509.CertPool, error) {
	backend, err := backendFrom(ctx)
	if err != nil {
		return nil, err
	}
	rot := x509.NewCertPool()
	var data []byte
	if root != "" {
		data, err = backend.IO.ReadFile(root)
	} else {
		if backend.Getter == nil {
			return nil, errNoGetter
		}
		data, err = backend.Getter.Get(gcetcbendorsement.DefaultRootURL)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get root certificate: %w", err)
	}
	// Certificate may be PEM, but also may be DER.
	if !rot.AppendCertsFromPEM(data) {
		rootCert, err := x509.ParseCertificate(data)
		if err != nil {
			return nil, fmt.Errorf("failed to parse root certificate as PEM or DER")
		}
		rot.AddCert(rootCert)
	}
	return rot, nil
}

func (c *verifyCommand) runE(cmd *cobra.Command, args []string) error {
	backend, err := backendFrom(cmd.Context())
	if err != nil {
		return err
	}
	if c.show {
		if c.root == "" {
			c.root = gcetcbendorsement.DefaultRootCmd
		}
		out, done, err := backend.IO.Create("-")
		if err != nil {
			return err
		}
		defer done()
		_, err = fmt.Fprintf(out, "%s \\\n&& \\\n%s\n",
			gcetcbendorsement.OpensslVerifyCertShellCmd(os.Args[0], args[0], c.root),
			gcetcbendorsement.OpensslVerifyShellCmd(os.Args[0], args[0]),
		)
		return err
	}
	rot, err := rootOfTrust(cmd.Context(), c.root)
	if err != nil {
		return err
	}

	return verify.EndorsementProto(c.endorsement, &verify.Options{
		Now:          backend.Now,
		Getter:       backend.Getter,
		RootsOfTrust: rot,
	})
}

func makeVerify(ctx0 context.Context) *cobra.Command {
	v := &verifyCommand{}
	cmd := &cobra.Command{
		Use: "verify PATH [options]",
		Long: `Outputs the result of verifying the endorsement.

If the root certificate is not provided, the command will attempt to download
it directly from https://pki.goog.
`,
		PersistentPreRunE: v.persistentPreRunE,
		RunE:              v.runE,
	}
	ctx := context.WithValue(ctx0, verifyKey, v)
	cmd.PersistentFlags().StringVar(&v.root, "root_cert", "", "Path to the root certificate.")
	cmd.PersistentFlags().BoolVar(&v.show, "show", false,
		"Output the openssl shell command to perform the equivalent logic.")
	cmd.SetContext(ctx)
	return cmd
}
