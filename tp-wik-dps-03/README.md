# WIK-DPS-TP03

## Description
Docker Compose avec 4 réplicas de l'API Go et un reverse-proxy nginx
qui load balance les requêtes entre les conteneurs.

## Architecture

```
┌──────────────────────────────────────────────────────┐
│                   Docker Network                     │
│                                                      │
│  ┌─────────────┐     ┌──────────────────────────┐    │
│  │             │     │   Service "ping" x4      │    │
│  │    nginx    │──── │  ┌────┐ ┌────┐ ┌────┐    │    │
│  │  :80 (int.) │     │  │ C1 │ │ C2 │ │ C3 │... │    │
│  │             │     │  └────┘ └────┘ └────┘    │    │
│  └──────┬──────┘     └──────────────────────────┘    │
│         │                                            │
└─────────│────────────────────────────────────────────┘
          │ :8080
     [navigateur / curl]
```

## Lancer le projet

```bash
docker compose up --build
```

## Tester

Puis ouvrir un autre terminal pour observer les logs en temps réel. 
```bash
docker compose logs -f ping
```
Lancer dans un second terminal
```bash
for ($i = 1; $i -le 8; $i++) { curl http://localhost:8080/ping }
```
Résultat attendu :
```bash
ping-2  | [PING] Requête reçue sur le conteneur : ea4e8be0419a
ping-1  | [PING] Requête reçue sur le conteneur : 474ff5a8aa22
ping-4  | [PING] Requête reçue sur le conteneur : 60fb66c1affe
ping-3  | [PING] Requête reçue sur le conteneur : fa5d99949ab9
ping-2  | [PING] Requête reçue sur le conteneur : ea4e8be0419a
ping-1  | [PING] Requête reçue sur le conteneur : 474ff5a8aa22
ping-4  | [PING] Requête reçue sur le conteneur : 60fb66c1affe
ping-3  | [PING] Requête reçue sur le conteneur : fa5d99949ab9
```

Vérification que l'API n'est pas accessible directement (seulement via nginx) :
```bash
# Doit fonctionner
curl http://localhost:8080/ping

# Ne doit pas fonctionner
curl http://localhost:9090/ping
```
## Observer le load balancing

Dans les logs Docker, le hostname change à chaque requête :
chaque conteneur (ping-1, ping-2, ping-3, ping-4) répond à tour de rôle.

## Arrêter

```bash
docker compose down
```