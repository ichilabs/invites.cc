FROM golang:1.19

# Définir le répertoire de travail
WORKDIR /app

RUN go mod init invites.cc

RUN go get github.com/lib/pq

# Copier le code source de l'application
COPY . .

# Exposer le port sur lequel l'application écoute
EXPOSE 8080

# Définir la commande pour exécuter l'application
CMD ["go", "run", "main.go"]
