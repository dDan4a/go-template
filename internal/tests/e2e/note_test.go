//go:build e2e

package e2e

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/dDan4a/go-template/internal/cmd/server"
	"github.com/dDan4a/go-template/internal/controller/http/converter"
	"github.com/dDan4a/go-template/internal/tests"
)

func TestNotes(t *testing.T) {
	srv := httptest.NewServer(server.NewHandler(tests.Postgres(t)))
	t.Cleanup(srv.Close)

	created := request[converter.NoteResponse](t, http.MethodPost, srv.URL+"/notes", `{"text":"buy milk"}`, http.StatusCreated)
	if created.Text != "buy milk" {
		t.Errorf("created text = %q, want %q", created.Text, "buy milk")
	}

	got := request[converter.NoteResponse](t, http.MethodGet, srv.URL+"/notes/"+strconv.FormatInt(created.ID, 10), "", http.StatusOK)
	if got.ID != created.ID || got.Text != created.Text {
		t.Errorf("got note %+v, want %+v", got, created)
	}

	request[map[string]string](t, http.MethodPost, srv.URL+"/notes", `{"text":"  "}`, http.StatusBadRequest)
	request[map[string]string](t, http.MethodGet, srv.URL+"/notes/0", "", http.StatusNotFound)
}

func request[T any](t *testing.T, method, url, body string, wantStatus int) T {
	t.Helper()

	req, err := http.NewRequestWithContext(t.Context(), method, url, strings.NewReader(body))
	if err != nil {
		t.Fatalf("build request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("%s %s: %v", method, url, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != wantStatus {
		t.Fatalf("%s %s: status %d, want %d", method, url, resp.StatusCode, wantStatus)
	}

	var out T
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatalf("decode %s %s: %v", method, url, err)
	}

	return out
}
