<img src="/assets/javelin_icon.png" alt="javelin" width="200" />

# Javelin

Javelin is a search and discovery engine for Neovim plugins. It aims to make it easy to explore, compare, and find plugins across the ecosystem with fast, relevant search powered by Elasticsearch.

At its core, Javelin indexes plugin metadata from GitHub repositories and exposes a search-focused API and (eventually) a web interface for discovery.

---

## Goals

- Provide fast and relevant search over Neovim plugins
- Improve discoverability beyond fragmented GitHub lists and dotfiles
- Support filtering by tags, features, and usage intent
- Offer a clean API for building clients (web, CLI, etc.)

---

## Architecture Overview

Javelin is split into a few logical components:

- **Ingester**: Collects plugin data from GitHub and other sources
- **Search Engine**: Uses Elasticsearch to index and query plugin data
- **API Server**: Exposes search and plugin endpoints for clients
- **Storage (future)**: Persistent database for canonical plugin metadata

---

## Tech Stack

- Go (backend services and tooling)
- Elasticsearch (search and indexing layer)
- Docker (local development infrastructure)

---

## Project Structure

- `cmd/` - Application entrypoints (API server, ingester)
- `internal/` - Core business logic and integrations
- `pkg/` - Shared models and types
- `configs/` - Configuration and index mappings
- `docker/` - Local infrastructure setup
- `scripts/` - Utility scripts for development

---

## Status

This project is in early development. Core functionality is not yet implemented.

---

## Quickstart: Index One Plugin into Elasticsearch

This is a minimal end-to-end flow that:

- starts Elasticsearch with Docker
- creates a `plugins` index with mappings
- scrapes one plugin (`nvim-telescope/telescope.nvim`) from GitHub
- indexes it into Elasticsearch

### 1) Start Elasticsearch

```bash
docker compose -f docker/docker-compose.yml up -d
```

### 2) Create the index

```bash
curl -X PUT "http://localhost:9200/plugins" \
  -H "Content-Type: application/json" \
  -d @configs/elasticsearch/plugins-index.json
```

If the index already exists, delete and recreate it:

```bash
curl -X DELETE "http://localhost:9200/plugins"
curl -X PUT "http://localhost:9200/plugins" \
  -H "Content-Type: application/json" \
  -d @configs/elasticsearch/plugins-index.json
```

### 3) Run the ingester

```bash
go run ./cmd/ingester
```

Optional environment variables:

- `ELASTICSEARCH_URL` (default: `http://localhost:9200`)
- `ELASTICSEARCH_INDEX` (default: `plugins`)

### 4) Verify it was indexed

```bash
curl -X GET "http://localhost:9200/plugins/_search?q=telescope&pretty"
```

You should see a document with id `nvim-telescope_telescope.nvim`.

### What this currently scrapes

The ingester pulls from GitHub API:

- repo metadata (name, stars, forks, topics, language, license, updated time)
- README content (decoded from Base64)

This is intentionally small so you can expand it incrementally (multi-plugin crawl, ranking, API server, Neovim UI).

---

## Future Ideas

- Plugin ranking based on usage and popularity
- Alternative suggestions ("plugins like telescope.nvim")
- Curated plugin collections
- CLI search tool
- Web UI for browsing plugins

