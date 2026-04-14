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

## Future Ideas

- Plugin ranking based on usage and popularity
- Alternative suggestions ("plugins like telescope.nvim")
- Curated plugin collections
- CLI search tool
- Web UI for browsing plugins

