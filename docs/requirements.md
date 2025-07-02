# Requirements pour l'outil de traduction Deepl CLI

Ce document décrit les exigences fonctionnelles et techniques pour le développement d'un outil de traduction en ligne de commande nommé `deepl`, utilisant l'API de DeepL.

## 1. Description du projet

L'objectif est de créer un utilitaire en ligne de commande simple et efficace qui permet aux utilisateurs de traduire du texte en utilisant l'API DeepL. L'outil sera configurable et facile à utiliser pour des traductions rapides depuis le terminal.

## 2. Fonctionnalités

### 2.1 Fonctionnalités principales

* **Traduction de texte :** Traduire un morceau de texte fourni en argument.
* **Détection automatique de la langue source :** Si la langue source n'est pas spécifiée, l'outil tentera de la détecter automatiquement.
* **Configuration flexible :** La clé API DeepL ainsi que les langues source et cible par défaut devront être configurées via un fichier de configuration dédié.
* **Affichage du résultat :** Le texte traduit sera affiché directement dans le terminal.

### 2.2 Fonctionnalités optionnelles (pour les futures versions)

* **Traduction de fichiers :** Possibilité de traduire le contenu d'un fichier et d'écrire la traduction dans un nouveau fichier.
* **Support des glossaires :** Utilisation des glossaires DeepL pour des traductions spécifiques.
* **Mode interactif :** Un mode où l'utilisateur peut taper du texte ligne par ligne pour la traduction.
* **Affichage des langues supportées :** Une commande pour lister toutes les langues supportées par DeepL.

## 3. Spécifications techniques

### 3.1 Langage de développement

* Le projet sera entièrement codé en **Go (Golang)**.

### 3.2 Utilisation en ligne de commande

* L'exécutable final devra être utilisable via la commande `deepl`.
* Exemples d'utilisation :
    * `deepl "Bonjour le monde"` (traduit en langue cible par défaut)
    * `deepl -s fr -t en "Bonjour le monde"` (traduit du français vers l'anglais)

### 3.3 Fichier de configuration

* Un fichier de configuration nommé `.deepl.toml` (ou un autre format comme YAML/JSON si préféré, mais TOML est courant pour Go) sera créé dans le répertoire `~/.deepl/` (ou `$HOME/.deepl/` sur Linux/macOS, `%USERPROFILE%\.deepl\` sur Windows).
* Ce fichier contiendra au minimum :
    * `api_key`: La clé API DeepL.
    * `source_lang`: La langue source par défaut (optionnel, peut être auto-détectée).
    * `target_lang`: La langue cible par défaut (obligatoire).

### 3.4 API DeepL

* Utilisation de l'API REST de DeepL. Une bibliothèque Go pour l'API DeepL pourrait être utilisée si disponible et maintenue, sinon des requêtes HTTP directes devront être implémentées.
* Le plan DeepL API Free sera utilisé pour le développement et les tests.

## 4. Étapes de réalisation

### 4.1 Phase 1 : Initialisation et Configuration

1.  **Initialisation du projet Go :** Créer un nouveau module Go.
2.  **Structure des répertoires :** Mettre en place la structure de base du projet (voir section 7).
3.  **Gestion de la configuration :**
    * Implémenter la logique pour lire le fichier `~/.deepl/.deepl.toml`.
    * Gérer les erreurs si le fichier ou la clé API est manquante.
    * Utiliser une bibliothèque de parsing de configuration (ex: `viper`, `gopkg.in/toml.v2`).

### 4.2 Phase 2 : Logique de Traduction

1.  **Parsing des arguments CLI :**
    * Utiliser une bibliothèque comme `cobra` ou `flag` pour gérer les arguments de la ligne de commande (`-s`, `-t`, texte à traduire).
    * Prioriser les arguments CLI sur les valeurs par défaut du fichier de configuration.
2.  **Appel à l'API DeepL :**
    * Implémenter la logique pour construire la requête HTTP vers l'API DeepL.
    * Gérer les réponses de l'API (succès, erreurs, codes HTTP).
    * Afficher la traduction ou les messages d'erreur de manière claire.

### 4.3 Phase 3 : Build et Distribution

1.  **Script de compilation :** Préparer un script (ou une cible Makefile) pour compiler l'exécutable pour différentes plateformes (Linux, macOS, Windows).
2.  **GitHub Actions pour les Releases :**
    * Mettre en place un workflow GitHub Actions pour compiler l'outil pour différentes architectures et systèmes d'exploitation.
    * Ce workflow devra créer des *releases* GitHub avec les binaires attachés lors d'un push de tag.

### 4.4 Phase 4 : Documentation et tests

1.  **Tests unitaires et d'intégration :** Écrire des tests pour les composants clés (parsing de configuration, appels API, etc.).
2.  **Documentation :**
    * Créer un fichier `README.md` détaillé.
    * Créer une documentation plus approfondie dans le répertoire `doc/`.

