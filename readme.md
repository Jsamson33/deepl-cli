# 🌍 deepl-cli : Votre Traducteur DeepL en Ligne de Commande

`deepl-cli` est un outil en ligne de commande léger et puissant écrit en Go, conçu pour vous permettre d'accéder à la qualité de traduction exceptionnelle de DeepL directement depuis votre terminal. Traduisez rapidement du texte sans quitter votre environnement de travail !

## ✨ Fonctionnalités

* Traduction rapide de texte via l'API DeepL.
* Détection automatique de la langue source.
* Configuration facile de votre clé API et des langues par défaut via un fichier `~/.deepl/.deepl.toml`.
* Support des langues source et cible directement en ligne de commande.

## 🚀 Installation

### Pré-requis

* Une clé API DeepL. Vous pouvez en obtenir une en vous inscrivant au [plan DeepL API Free](https://www.deepl.com/pro/developer).
* Go 1.16+ installé sur votre machine.

### Depuis les Binaires Pré-compilés (Recommandé)

Des binaires pré-compilés pour Linux, macOS et Windows sont disponibles dans la section [Releases de ce dépôt GitHub](https://github.com/votre-utilisateur/deepl-cli/releases).
1.  Téléchargez l'archive correspondant à votre système d'exploitation et architecture.
2.  Décompressez l'archive.
3.  Placez l'exécutable `deepl` (ou `deepl.exe` sur Windows) dans un répertoire présent dans votre `PATH` (par exemple, `/usr/local/bin` sur Linux/macOS, ou un dossier personnalisé ajouté à votre `PATH` sur Windows).

### Depuis les Sources

```bash
git clone [https://github.com/votre-utilisateur/deepl-cli.git](https://github.com/votre-utilisateur/deepl-cli.git)
cd deepl-cli
go build -o deepl main.go
# Déplacez l'exécutable vers un répertoire de votre PATH
# mv deepl /usr/local/bin/deepl
⚙️ Configuration
Avant d'utiliser deepl-cli, vous devez configurer votre clé API DeepL et éventuellement vos langues par défaut.

Créez un répertoire .deepl dans votre répertoire personnel :
mkdir -p ~/.deepl
Créez un fichier nommé .deepl.toml à l'intérieur de ce répertoire (~/.deepl/.deepl.toml) avec le contenu suivant, en remplaçant VOTRE_CLE_API_DEEPL par votre véritable clé :
api_key = "VOTRE_CLE_API_DEEPL"
target_lang = "fr" # Langue cible par défaut (ex: "fr" pour Français, "en" pour Anglais)
# source_lang = "en" # Optionnel : langue source par défaut (peut être auto-détectée)
Note sur les langues : Utilisez les codes ISO 639-1 (par exemple en, fr, de). Pour une liste complète, consultez la documentation de l'API DeepL.

💻 Utilisation
Traduire du texte (langues par défaut)
deepl "Hello world, how are you?"
Le texte sera traduit vers la target_lang configurée dans votre fichier .deepl.toml. La langue source sera détectée automatiquement.

Spécifier les langues source et cible
Utilisez les drapeaux -s (source) et -t (target) :

deepl -s en -t fr "Hello world, how are you?"
deepl -s fr -t en "Bonjour le monde, comment allez-vous ?"
Aide
Pour voir toutes les options disponibles :

deepl --help
