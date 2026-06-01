# tp-wik-dps-01

## Description

Petit serveur HTTP en Go qui expose une route `/ping`.
Cette route renvoie les en-têtes de la requête dans un format JSON.

## Prérequis

- Go installé (version 1.20+)

## Lancer le projet

1. Ouvrir un terminal.
2. Se placer dans le dossier du projet :
   ```bash
   cd ./tp-wik-dps-01
   ```
3. Lancer le serveur :
   ```bash
   go run .
   ```

Le serveur écoute par défaut sur le port `9090`.

## Accéder à l'API

- `GET http://localhost:9090/ping`

## Configuration

Vous pouvez changer le port en définissant la variable d'environnement `PING_LISTEN_PORT` avant de lancer l'application.

Par exemple :

```bash
set PING_LISTEN_PORT=8080
go run .
```
