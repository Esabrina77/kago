# kaGO 🚀

**kaGO** est un outil en ligne de commande (CLI) conçu pour initialiser rapidement des environnements de travail en Go. L'objectif est de passer de l'idée au code en une seule commande, avec des structures de dossiers éprouvées.

## 📌 Vision
L'idée est de proposer un outil similaire à `npx create-react-app` ou aux générateurs de projets Node.js, mais adapté à l'écosystème Go. 

kaGO permet de générer deux types de structures :
1.  **Simple** : Pour les scripts rapides ou les exercices (main.go, go.mod).
2.  **Web API** : Une structure complète (Architecture, Middleware, Gitignore, Tests) prête à être lancée sur le port 8080.

## 🛠️ Installation
*À venir (via go install)*

## 🚀 Utilisation
```bash
# Pour un projet simple
kago mon-projet --type=simple

# Pour une API Web complète
kago mon-api --type=web
