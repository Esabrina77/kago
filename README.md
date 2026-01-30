# 🚀 kaGO - Go Project Initializer

## 🇫🇷 Français

### Présentation

**kaGO** est un outil en ligne de commande (CLI) conçu pour automatiser la création de nouveaux projets en Go. Il génère instantanément une structure de dossiers professionnelle et initialise le module Go pour vous faire gagner du temps.

### Fonctionnalités

- **Initialisation Rapide** : Créez un projet en une seule commande.
- **Deux Modes** :
  - `simple` : Pour les scripts et petits outils (un seul `main.go`).
  - `web` : Une architecture professionnelle (`cmd`, `internal`, `pkg`, `api`).
- **Automatisation** : Lance automatiquement `go mod init`.
- **Embarqué** : Utilise `go:embed` pour inclure les templates directement dans l'exécutable.

### Utilisation

**Bash**

```
go run main.go -type=web mon-projet-api
```

---

# 🇺🇸 English

### Overview

**kaGO** is a command-line interface (CLI) tool designed to automate the bootstrapping of new Go projects. It instantly generates a professional directory structure and initializes the Go module to save you development time.

### Features

- **Fast Bootstrapping** : Create a project with a single command.
- **Two Modes** :
- `simple`: For scripts and small utilities (single `main.go`).
- `web`: A professional architecture (`cmd`, `internal`, `pkg`, `api`).
- **Automation** : Automatically runs `go mod init` for you.
- **Self-contained** : Uses `go:embed` to bundle templates within the binary.

### Usage

**Bash**

```
go run main.go -type=web my-api-project
```

---

## 🛠 Technique / Technical Details

### Architecture Decision Records (ADR)

- **Language** : Go 1.21+
- **Standard Library** : Built using `flag`, `os`, and `os/exec` to ensure zero external dependencies.
- **Embedding** : Templates are managed via the `embed` package to allow single-binary distribution.
- **Safety** : Includes directory existence checks to prevent accidental overwrites.
