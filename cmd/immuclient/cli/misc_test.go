/*
Copyright 2021 CodeNotary, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cli

import (
	"github.com/codenotary/immudb/pkg/client/tokenservice"
	"os"
	"strings"
	"testing"

	"github.com/codenotary/immudb/pkg/client"

	test "github.com/codenotary/immudb/cmd/immuclient/immuclienttest"
	"github.com/codenotary/immudb/pkg/server"
	"github.com/codenotary/immudb/pkg/server/servertest"
)

func TestHealthCheck(t *testing.T) {
	options := server.DefaultOptions().WithAuth(true)
	bs := servertest.NewBufconnServer(options)

	bs.Start()
	defer bs.Stop()

	defer os.RemoveAll(options.Dir)
	defer os.Remove(".state-")

	ts := tokenservice.NewFileTokenService().WithTokenFileName("testTokenFile").WithHds(&test.HomedirServiceMock{})
	ic := test.NewClientTest(&test.PasswordReader{
		Pass: []string{"immudb"},
	}, ts).WithOptions(client.DefaultOptions())
	ic.Connect(bs.Dialer)
	ic.Login("immudb")

	cli := new(cli)
	cli.immucl = ic.Imc
	msg, err := cli.healthCheck([]string{})
	if err != nil {
		t.Fatal("HealthCheck fail", err)
	}
	if !strings.Contains(msg, "Health check OK") {
		t.Fatal("HealthCheck fail")
	}
}

func TestHistory(t *testing.T) {
	options := server.DefaultOptions().WithAuth(true)
	bs := servertest.NewBufconnServer(options)

	bs.Start()
	defer bs.Stop()

	defer os.RemoveAll(options.Dir)
	defer os.Remove(".state-")

	ts := tokenservice.NewFileTokenService().WithTokenFileName("testTokenFile").WithHds(&test.HomedirServiceMock{})
	ic := test.NewClientTest(&test.PasswordReader{
		Pass: []string{"immudb"},
	}, ts).WithOptions(client.DefaultOptions())
	ic.Connect(bs.Dialer)
	ic.Login("immudb")

	cli := new(cli)

	cli.immucl = ic.Imc
	msg, err := cli.history([]string{"key"})
	if err != nil {
		t.Fatal("History fail", err)
	}
	if !strings.Contains(msg, "key not found") {
		t.Fatalf("History fail %s", msg)
	}

	_, err = cli.set([]string{"key", "value"})
	if err != nil {
		t.Fatal("History fail", err)
	}

	msg, err = cli.history([]string{"key"})
	if err != nil {
		t.Fatal("History fail", err)
	}
	if !strings.Contains(msg, "hash") {
		t.Fatalf("History fail %s", msg)
	}
}
func TestVersion(t *testing.T) {
	options := server.DefaultOptions().WithAuth(true)
	bs := servertest.NewBufconnServer(options)

	bs.Start()
	defer bs.Stop()

	defer os.RemoveAll(options.Dir)
	defer os.Remove(".state-")

	ts := tokenservice.NewFileTokenService().WithTokenFileName("testTokenFile").WithHds(&test.HomedirServiceMock{})
	ic := test.NewClientTest(&test.PasswordReader{
		Pass: []string{"immudb"},
	}, ts).WithOptions(client.DefaultOptions())
	ic.Connect(bs.Dialer)
	ic.Login("immudb")

	cli := new(cli)
	cli.immucl = ic.Imc
	msg, err := cli.version([]string{"key"})
	if err != nil {
		t.Fatal("version fail", err)
	}
	if !strings.Contains(msg, "no version info available") {
		t.Fatalf("version fail %s", msg)
	}
}
