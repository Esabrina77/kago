# 🚀 kaGO - Go Project Initializer

## 🇫🇷 Français

### Présentation

**kaGO** est un outil CLI qui automatise la création de vos projets Go. Plus besoin de créer manuellement vos dossiers ou de taper `go mod init` à chaque fois.

### 📥 Installation

Pour installer l'outil de manière globale sur votre machine :

**Bash**

```
go install github.com/Esabrina77/kago@latest
```

_(Vérifiez que votre répertoire `$GOPATH/bin` est bien présent dans votre PATH système)_ .

### 🛠 Utilisation

Une fois installé, utilisez simplement la commande `kago` n'importe où :

**Bash**

```
# Projet simple (main.go unique)
kago -type=simple mon-projet

# Structure Web professionnelle
kago -type=web ma-super-api
```

---

## 🇺🇸 English

### Overview

**kaGO** is a CLI tool that automates Go project bootstrapping. Stop manually creating folders and running `go mod init` for every new project.

### 📥 Installation

To install the tool globally on your machine:

**Bash**

```
go install github.com/Esabrina77/kago@latest
```

_(Ensure your `$GOPATH/bin` directory is in your system PATH)_ .

### 🛠 Usage

Once installed, simply use the `kago` command anywhere:

**Bash**

```
# Simple project (single main.go)
kago -type=simple my-project

# Professional Web architecture
kago -type=web my-awesome-api
```
