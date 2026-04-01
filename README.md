# Go Auth Backend — UpcycleConnect

API d'authentification développée en Go pour le projet UpcycleConnect.

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=UpcycleConnect-gpr3_go_auth_backend&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=UpcycleConnect-gpr3_go_auth_backend)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=UpcycleConnect-gpr3_go_auth_backend&metric=bugs)](https://sonarcloud.io/summary/new_code?id=UpcycleConnect-gpr3_go_auth_backend)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=UpcycleConnect-gpr3_go_auth_backend&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=UpcycleConnect-gpr3_go_auth_backend)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=UpcycleConnect-gpr3_go_auth_backend&metric=coverage)](https://sonarcloud.io/summary/new_code?id=UpcycleConnect-gpr3_go_auth_backend)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=UpcycleConnect-gpr3_go_auth_backend&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=UpcycleConnect-gpr3_go_auth_backend)

---

## Prérequis

- [Go](https://golang.org/dl/) **1.21+**
- [Docker](https://www.docker.com/) & **Docker Compose** (pour la base de données)

Vérifiez votre installation :

```bash
go version
docker -v
```

---

## Démarrage

### 1. Cloner le dépôt

```bash
git clone https://github.com/UpcycleConnect-gpr3/go_auth_backend.git
cd go_auth_backend
```

### 2. Configurer les variables d'environnement

```bash
cp example.env .env.development
```

Éditez le fichier `.env` avec vos valeurs.

### 3. Lancer la base de données

```bash
docker-compose up -d
```

### 4. Installer les dépendances Go

```bash
go mod tidy
```

### 5. Démarrer le serveur

```bash
go run main.go serve
```
---

---

## Commandes disponibles

### Via Go

| Commande | Description |
|---|---|
| `go run main.go serve` | Démarre le serveur API |
| `go mod tidy` | Installe/nettoie les dépendances |
| `go build -o api .` | Compile le binaire |
| `go test ./...` | Lance les tests |


### Via Docker Compose

```bash
docker-compose up -d      # Démarre les services (DB...)
docker-compose down       # Arrête les services
docker-compose logs -f    # Affiche les logs
```
---

## Variables d'environnement

Copiez `example.env` vers `.env` et renseignez les valeurs. Exemple de variables attendues :



> Consultez `example.env` pour la liste complète des variables requises.

---

## Contribution

1. Créez une branche : `git checkout -b feat/ma-feature`
2. Committez : `git commit -m "feat: description"`
3. Pushez : `git push origin feat/ma-feature`
4. Ouvrez une Pull Request


openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048
openssl pkey -in private_key.pem -pubout -out public_key.pem

## Template

### Actions

```go
package <model>_actions

import (
    "authentication_backend/utils/rules"
)

func <action>Validate<Model>(<model>Dto <model>_models.<ModelUse>) []rules.ValidationError {
	var errs []rules.ValidationError

	rules.<Rule>(<model>Dto.<Attribute>, 5, "<attribute>", &errs)

	return errs
}

func <Action><Model>(<model>Dto <model>_models.<ModelUse>) []rules.ValidationError {

	validationError := <action>Validate<Model>(<model>Dto)

	if len(validationError) > 0 {
		return validationError
	}

	<model>_models.<Action><Model>(<model>Dto)

	return nil
}

```
