# Proposition de Système de Mise à Jour Automatique

## 1. Introduction

Ce document propose une approche pour implémenter un système de mise à jour automatique pour l'exécutable `deepl-cli`, similaire à celui utilisé par des outils comme OhMyZsh. L'objectif est de permettre à l'application de vérifier et d'installer les nouvelles versions disponibles au démarrage, améliorant ainsi l'expérience utilisateur et la maintenance.

## 2. Principes de Fonctionnement

Le processus de mise à jour automatique suivra les étapes générales suivantes :

1.  **Vérification de la Version Actuelle :** Au lancement, l'application déterminera sa propre version (qui sera intégrée lors de la compilation).
2.  **Interrogation d'une Source Distante :** L'application contactera une source distante (par exemple, les [GitHub Releases](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository)) pour vérifier si une version plus récente est disponible.
3.  **Téléchargement de la Nouvelle Version :** Si une nouvelle version est détectée, l'exécutable correspondant à l'OS et à l'architecture de l'utilisateur sera téléchargé dans un emplacement temporaire.
4.  **Remplacement de l'Exécutable :** L'application tentera de se remplacer elle-même par le nouvel exécutable téléchargé. C'est l'étape la plus délicate, car un processus en cours d'exécution ne peut pas directement écraser son propre fichier. Cela implique généralement :
    *   Renommer l'ancien exécutable (par exemple, `deepl.old`).
    *   Déplacer le nouvel exécutable vers le chemin d'origine (`deepl`).
    *   Gérer les erreurs potentielles (permissions, téléchargement corrompu, etc.).
    *   Nettoyer l'ancien exécutable lors d'un démarrage ultérieur réussi.

## 3. Implémentation Technique (Go)

Pour faciliter l'implémentation de ce système en Go, il est fortement recommandé d'utiliser une bibliothèque existante qui gère les complexités liées au remplacement d'un exécutable en cours d'exécution et à la gestion des versions. La bibliothèque `github.com/rhysd/go-github-selfupdate` est un excellent candidat pour les projets hébergés sur GitHub.

### Étapes d'Implémentation :

1.  **Intégration de la Bibliothèque :**
    *   Ajouter `github.com/rhysd/go-github-selfupdate` aux dépendances du projet (`go get github.com/rhysd/go-github-selfupdate`).

2.  **Définition de la Version :**
    *   La version de l'application doit être définie de manière à pouvoir être lue par le code (par exemple, via une variable `const` ou une variable injectée au moment de la compilation via les `ldflags`).
    *   Exemple : `go build -ldflags "-X main.version=1.0.0"`

3.  **Logique de Vérification et de Mise à Jour :**
    *   Au début de la fonction `main` (ou dans une fonction dédiée appelée par `main`), implémenter la logique de vérification.
    *   Utiliser la fonction `selfupdate.UpdateSelf` ou `selfupdate.UpdateSelfFromSource` de la bibliothèque.
    *   Cette fonction prendra en charge la comparaison des versions, le téléchargement et le remplacement.

    ```go
    package main

    import (
    	"fmt"
    	"os"
    	"time"

    	"github.com/rhysd/go-github-selfupdate/selfupdate"
    )

    var version = "0.0.1" // Cette valeur sera remplacée par ldflags lors de la compilation

    func doSelfUpdate() {
    	latest, err := selfupdate.UpdateSelf(version, "Jsamson33/deepl-cli")
    	if err != nil {
    		fmt.Println("Self-update failed:", err)
    		return
    	}
    	if latest.Version.String() != version {
    		fmt.Println("Successfully updated to version", latest.Version)
    		fmt.Println("Release Note:\n", latest.ReleaseNotes)
    	} else {
    		fmt.Println("Current version is the latest")
    	}
    }

    func main() {
    	// Optionnel: Vérifier la mise à jour seulement après un certain délai ou avec un flag spécifique
    	// if time.Now().Sub(lastUpdateTime) > 24 * time.Hour {
    	//     doSelfUpdate()
    	// }

    	doSelfUpdate()

    	// Reste de la logique de l'application
    	// ...
    }
    ```

4.  **Gestion des Erreurs et de l'Expérience Utilisateur :**
    *   Fournir des messages clairs à l'utilisateur concernant l'état de la mise à jour (nouvelle version disponible, mise à jour réussie, échec).
    *   Gérer les cas où l'utilisateur n'a pas les permissions nécessaires pour la mise à jour.
    *   Considérer l'ajout d'un flag CLI (`--no-self-update`) pour désactiver la vérification automatique.

## 4. Considérations Importantes

*   **Permissions :** Le système de mise à jour nécessite des permissions d'écriture sur le répertoire où l'exécutable est installé. Cela peut être un problème sur certains systèmes (ex: `/usr/bin`). Une solution est de demander à l'utilisateur d'installer l'exécutable dans un répertoire où il a les droits (ex: `/usr/local/bin` ou un répertoire personnel).
*   **Signatures Binaires :** Pour une sécurité accrue, il est fortement recommandé de signer les binaires des releases et de vérifier ces signatures lors de la mise à jour. La bibliothèque `go-github-selfupdate` supporte la vérification des signatures GPG.
*   **Fréquence de Vérification :** Éviter de vérifier les mises à jour à chaque lancement pour ne pas surcharger GitHub API ou ralentir le démarrage de l'application. Une vérification quotidienne ou hebdomadaire est généralement suffisante.
*   **Rollback :** En cas d'échec de la mise à jour, l'ancien exécutable est généralement conservé (renommé). Cela permet un rollback manuel si nécessaire.
*   **Compatibilité :** S'assurer que les nouvelles versions sont compatibles avec les anciennes configurations et données utilisateur.

## 5. Étapes Suivantes

1.  Discuter de l'opportunité d'implémenter cette fonctionnalité.
2.  Si approuvé, intégrer la bibliothèque `go-github-selfupdate`.
3.  Mettre en place le processus de release sur GitHub Actions pour générer des binaires avec des versions et potentiellement des signatures.
4.  Implémenter la logique de mise à jour dans le code.
5.  Tester rigoureusement le système de mise à jour sur différentes plateformes.