# WIK-DPS-TP02

## Description
Dockerisation de l'API Go du TP01 avec deux stratégies de build.

## Images Docker

### Image mono-stage
```bash
docker build -t tp02:single .
docker run -p 9090:9090 tp02:single
```

### Image multi-stages
```bash
docker build -f Dockerfile.multi -t tp02:multi .
docker run -p 9090:9090 tp02:multi
```

## Test
```bash
curl http://localhost:9090/ping
```

## Résultats du scan (Trivy / Docker Scout)
```bash
docker run tp02:single whoami`
```
doit afficher : appuser

```bash
docker run tp02:multi whoami`
```
doit afficher : appuser

```bash
docker scout cves tp02:single
docker scout cves tp02:multi
```
doit afficher des vulnérabilités de ce genre : CVE-2024-5525  MEDIUM  busybox 1.36.x    Command injection

### tp02:single
- TAILLE : 492MB

### tp02:multi
- TAILLE : 26.4MB

> L'image multi-stage est plus légère et présente moins de vulnérabilités
> car elle ne contient pas le compilateur Go ni les sources.