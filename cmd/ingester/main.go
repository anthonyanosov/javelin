package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type githubRepo struct {
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	Description     string   `json:"description"`
	HTMLURL         string   `json:"html_url"`
	StargazersCount int      `json:"stargazers_count"`
	ForksCount      int      `json:"forks_count"`
	Language        string   `json:"language"`
	Topics          []string `json:"topics"`
	License         *struct {
		SPDXID string `json:"spdx_id"`
	} `json:"license"`
	UpdatedAt time.Time `json:"updated_at"`
}

type githubReadme struct {
	Content string `json:"content"`
}

type pluginDocument struct {
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	RepoURL     string    `json:"repo_url"`
	Description string    `json:"description"`
	Readme      string    `json:"readme"`
	Stars       int       `json:"stars"`
	Forks       int       `json:"forks"`
	Language    string    `json:"language"`
	Topics      []string  `json:"topics"`
	License     string    `json:"license"`
	UpdatedAt   time.Time `json:"updated_at"`
	IndexedAt   time.Time `json:"indexed_at"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	repo := "nvim-telescope/telescope.nvim"
	esURL := getEnv("ELASTICSEARCH_URL", "http://localhost:9200")
	index := getEnv("ELASTICSEARCH_INDEX", "plugins")

	doc, err := fetchPluginDocument(ctx, repo)
	if err != nil {
		exitWithError(fmt.Errorf("fetch plugin metadata: %w", err))
	}

	if err := indexDocument(ctx, esURL, index, repo, doc); err != nil {
		exitWithError(fmt.Errorf("index plugin document: %w", err))
	}

	fmt.Printf("indexed %s into %s/%s\n", repo, esURL, index)
}

func fetchPluginDocument(ctx context.Context, repo string) (*pluginDocument, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s", repo)

	repoReq, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, err
	}
	repoReq.Header.Set("Accept", "application/vnd.github+json")
	repoReq.Header.Set("User-Agent", "javelin-ingester")

	repoRes, err := http.DefaultClient.Do(repoReq)
	if err != nil {
		return nil, err
	}
	defer repoRes.Body.Close()

	if repoRes.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(repoRes.Body)
		return nil, fmt.Errorf("repo request failed: status=%d body=%s", repoRes.StatusCode, string(body))
	}

	var gr githubRepo
	if err := json.NewDecoder(repoRes.Body).Decode(&gr); err != nil {
		return nil, err
	}

	readme, err := fetchReadme(ctx, repo)
	if err != nil {
		return nil, err
	}

	license := ""
	if gr.License != nil {
		license = gr.License.SPDXID
	}

	return &pluginDocument{
		Name:        gr.Name,
		FullName:    gr.FullName,
		RepoURL:     gr.HTMLURL,
		Description: gr.Description,
		Readme:      readme,
		Stars:       gr.StargazersCount,
		Forks:       gr.ForksCount,
		Language:    gr.Language,
		Topics:      gr.Topics,
		License:     license,
		UpdatedAt:   gr.UpdatedAt,
		IndexedAt:   time.Now().UTC(),
	}, nil
}

func fetchReadme(ctx context.Context, repo string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/readme", repo)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "javelin-ingester")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return "", fmt.Errorf("readme request failed: status=%d body=%s", res.StatusCode, string(body))
	}

	var payload githubReadme
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return "", err
	}

	decoded, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(payload.Content, "\n", ""))
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

func indexDocument(ctx context.Context, esURL, index, repo string, doc *pluginDocument) error {
	payload, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	docID := strings.ReplaceAll(repo, "/", "_")
	url := fmt.Sprintf("%s/%s/_doc/%s", strings.TrimRight(esURL, "/"), index, docID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("index request failed: status=%d body=%s", res.StatusCode, string(body))
	}

	return nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}
