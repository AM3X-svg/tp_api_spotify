# API Spotify - Projet Go

![Spotify Banner](https://upload.wikimedia.org/wikipedia/commons/1/19/Spotify_logo_without_text.svg)

## Description

Ce projet est une application web en **Go** qui consomme l’**API Spotify** pour afficher :

* Les albums de **Damso**
* La musique de **Laylow**

La page d’accueil présente les deux artistes côte à côte avec un design moderne **Liquid Glass**, des boutons interactifs et un fond animé. Les utilisateurs peuvent cliquer sur un artiste pour accéder à ses albums ou morceaux dans une page dédiée.

---

## Fonctionnalités

* **Page d’accueil** : Damso et Laylow présentés avec images et boutons cliquables
* **Page Damso** : Affiche tous les albums sous forme de boxes avec image, titre et bouton “Écouter”
* **Page Laylow** : Affiche les morceaux avec même design
* **Design moderne** : Fond animé, effet Liquid Glass sur cartes et boutons, responsive
* **Interaction API Spotify** : Utilisation de `client_credentials` pour récupérer les albums et tracks

---

## Architecture du projet

```
apispotify/
├─ src/
│  ├─ main.go              # Point d'entrée du serveur
│  ├─ go.mod               # Modules Go
│  ├─ controller/
│  │  └─ controller.go     # Gestion des routes et récupération API
│  ├─ pages/
│  │  ├─ index.html        # Page d'accueil
│  │  ├─ damso.html        # Page des albums de Damso
│  │  ├─ laylow.html       # Page des tracks de Laylow
│  │  └─ initTemps.go      # Initialisation des templates
│  ├─ router/
│  │  └─ router.go         # Définition des routes HTTP
│  ├─ static/
│  │  ├─ css/
│  │  │  └─ styles.css     # Styles CSS (Liquid Glass, animations, responsive)
│  │  └─ images/           # Images statiques (profil des artistes, albums)
│  └─ struct/
│     └─ struct.go         # Structures Go pour l’API Spotify
```

---

## Installation

### Prérequis

* [Go 1.25+](https://go.dev/dl/) installé sur votre machine
* Accès Internet pour récupérer les données Spotify

### Étapes

1. Clonez le projet :

```bash
git clone <repo_url>
cd apispotify/src
```

2. Installer les dépendances Go (si nécessaire) :

```bash
go mod tidy
```

3. Lancer le serveur local sur le port 8080 :

```bash
go run main.go
```

4. Ouvrir le navigateur :

```
http://localhost:8080
```

---

## Utilisation

* **Page d’accueil** : Cliquez sur les cartes Damso ou Laylow
* **Page Damso** : Cliquez sur “Écouter” pour accéder à l’album sur Spotify
* **Page Laylow** : Idem pour les morceaux

---

## Technologies

* **Go** pour le backend et le serveur HTTP
* **HTML / CSS** pour le frontend
* **API Spotify** pour récupérer albums et tracks
* **Glassmorphism / Liquid Glass** pour le design
* **Flexbox et animations CSS** pour responsive et interactions

---

## Structure des templates

* `index.html` : Page principale avec cartes artistes
* `damso.html` : Albums de Damso en boxes
* `laylow.html` : Tracks de Laylow en boxes
* `initTemps.go` : Initialisation des templates Go avec `html/template`

---

## Style & UI

* Effet **Glassmorphism** sur cartes et boutons
* **Fond animé** avec dégradé vert Spotify → noir
* Hover sur les images et cartes pour effet interactif
* Responsive design sur mobile et tablette

---

## Contributions

Ce projet est open-source et peut être amélioré par :

* Ajout d’autres artistes
* Affichage des playlists ou top tracks
* Amélioration du design UI / UX
* Gestion des tokens Spotify de manière sécurisée

---

## Licence

MIT License - libre d’utilisation et modification
