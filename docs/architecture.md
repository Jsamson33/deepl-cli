# Architecture Technique du Projet DeepL CLI

## 1. Vue d'ensemble du Projet

Le projet DeepL CLI est un utilitaire en ligne de commande développé en Go (Golang) qui permet aux utilisateurs de traduire du texte en exploitant l'API de DeepL. L'objectif principal est de fournir un outil simple, rapide et configurable pour des traductions textuelles directement depuis le terminal.

## 2. Architecture Générale

L'architecture du projet est conçue pour être modulaire et facile à maintenir, en séparant les préoccupations entre la gestion de la ligne de commande, la configuration et l'interaction avec l'API externe.

```mermaid
graph TD
    A[Utilisateur] -->|Exécute deepl commande| B(CLI - cmd/deepl/main.go)
    B -->|Parse arguments (cobra)| C{Root Command - cmd/deepl/cmd/root.go}
    C -->|Charge configuration| D[Configuration - internal/config/config.go]
    C -->|Priorise arguments sur config|
    C -->|Appelle API DeepL| E(DeepL API)
    E -->|Retourne traduction| C
    C -->|Affiche résultat| A
    D -->|.deepl.toml (fichier de config)|
```

## 3. Composants Clés

### 3.1. `cmd/deepl/main.go`

Point d'entrée principal de l'application. Il initialise et exécute la commande racine définie par Cobra.

### 3.2. `cmd/deepl/cmd/root.go`

Contient la définition de la commande racine (`deepl`) et gère le parsing des arguments de la ligne de commande (`-s`, `-t`, texte à traduire). Il orchestre également le chargement de la configuration et l'appel à la fonction de traduction.

### 3.3. `internal/config/config.go`

Gère le chargement et la validation du fichier de configuration `.deepl.toml`. Ce fichier est situé dans le répertoire `~/.deepl/` de l'utilisateur et contient la clé API DeepL, ainsi que les langues source et cible par défaut.

### 3.4. Logique de Traduction (`translateText` fonction dans `root.go`)

Cette fonction est responsable de la construction de la requête HTTP vers l'API DeepL, de l'envoi de la requête, de la gestion des réponses (succès ou erreurs) et de l'extraction du texte traduit.

## 4. Flux de Données

1.  L'utilisateur exécute la commande `deepl` avec le texte à traduire et des options facultatives (langue source, langue cible).
2.  Le module `main.go` lance la commande racine.
3.  La commande racine (`root.go`) utilise Cobra pour parser les arguments CLI.
4.  Elle charge la configuration depuis `.deepl.toml` via `config.go`.
5.  Les arguments CLI priment sur les valeurs par défaut du fichier de configuration.
6.  La fonction `translateText` est appelée avec le texte, les langues (source et cible) et la clé API.
7.  `translateText` construit une requête HTTP POST vers l'API DeepL, incluant le texte, les langues et la clé API (via l'en-tête `Authorization`).
8.  L'API DeepL traite la requête et renvoie la traduction.
9.  La réponse de l'API est lue, désérialisée, et le texte traduit est extrait.
10. Le texte traduit est affiché sur la console de l'utilisateur.

## 5. Configuration

Le fichier de configuration `.deepl.toml` doit être créé manuellement par l'utilisateur dans `~/.deepl/`. Il doit contenir au minimum la clé API DeepL et la langue cible par défaut. Exemple :

```toml
api_key = "YOUR_DEEPL_API_KEY"
target_lang = "en"
source_lang = "fr" # Optionnel
```

## 6. Utilisation de l'API DeepL

L'outil utilise l'API REST de DeepL. Les requêtes sont effectuées vers `https://api-free.deepl.com/v2/translate` en utilisant la méthode POST avec un corps de requête JSON. L'authentification se fait via l'en-tête `Authorization` avec le format `DeepL-Auth-Key YOUR_API_KEY`.
