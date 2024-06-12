
# Documentation: Reverse Proxy, Docker, et Serveurs Web

Dans cette documentation, nous allons voir comment réalisé un Reverse Proxy et manipulé Docker. Avec les étapes importantes ainsi que les commandes.

## Projet

Le projet consiste à mettre en œuvre un système intégrant un "Reverse Proxy", Docker et un serveur web. Nous avons choisi de réaliser cette infrastructure dans le cadre d'un projet intitulé "Reverse Proxy + Docker + Serveur Web".

Nous avons opté pour la mise en place d'un reverse proxy, un outil permettant de diriger le trafic web vers différentes applications en fonction de certains critères. En parallèle, nous avons utilisé Docker pour créer des conteneurs isolés pour nos différentes applications.

Les applications que nous avons incluses dans ce projet sont les suivantes :

- Un jeu de Pendu (Hangman) développé en Go, un langage de programmation haut niveau connu pour sa simplicité d'utilisation et de compréhension.

- Deux "Groupie Tracker", chacun développé dans un langage différent : Go et JavaScript. Ces applications ont pour objectif d'utiliser une API externe et de restituer les résultats sur une page web.

Cette configuration nous permet de démontrer la capacité du reverse proxy à diriger le trafic vers les différentes applications hébergées sur les conteneurs Docker, tout en offrant une isolation et une gestion efficace des applications.

## Mots-clés

- Reverse Proxy : Requète client -> serveur. ([Explications](#explications))

- Docker : Docker est un logiciel permettant la mise en place de plusieurs applications différentes ayant chacune une tâche différentes. ([Explications](#explications))

- Nginx : Logiciel libre de serveur Web ainsi qu'un proxy inverse.

- API : Un ensemble de définitions et de protocoles qui facilite le développement des applications. Exemple ci-dessous.
  ![API](https://user.oc-static.com/upload/2021/06/17/16239236573998_image38.png)

## Table des Matières

- [Explications du Reverse Proxy](#Explications-du-Reverse-Proxy)
- [Explications de Docker](#explications-de-Docker)
- [Installation de NGINX](#Installation-de-NGINX)
- [Installation d'Apache](##Installation-d'Apache)
- [Configuration du Reverse Proxy](#Configuration-du-Reverse-Proxy)
## Explications du Reverse Proxy

Avant d'entamer la configuration des serveurs, du reverse proxy et de Docker, une explication approfondie des termes s'impose. Mais qu'est-ce qu'un "reverse proxy" ?

![Schéma illustratif du fonctionnement d'un reverse proxy](https://res.cloudinary.com/delbwqa4s/image/upload/v1678860211/Reverse_proxy_flow_eac1d9aa0e.png)

Comme mentionné dans la section "[Mots-clés](#mots-clés)", un reverse proxy est un serveur placé devant les serveurs web qui traite les requêtes des clients et les redirige vers les serveurs appropriés.

### Explications Détaillées du Reverse Proxy

- **Traitement des Requêtes :** Le reverse proxy reçoit les requêtes des clients, les analyse et décide vers quel serveur backend rediriger chaque requête.
- **Avantages :**
    - **Sécurité :** Le reverse proxy peut masquer l'identité des serveurs backend, augmentant ainsi la sécurité.
    - **Performance :** Il peut équilibrer la charge entre plusieurs serveurs backend, améliorant ainsi la performance et la disponibilité.
    - **Cache :** Le reverse proxy peut mettre en cache les réponses pour réduire la charge sur les serveurs backend.

Contrairement à un "proxy traditionnel", qui permet une communication du serveur vers le client, le "proxy inversé" établit une communication du client vers le serveur. Il récupère les informations (requêtes) du client dans ses instances (serveurs) déployés et renvoie celui correspondant. Contrairement au "proxy traditionnel", qui demande des informations aux clients pour avoir accès aux données présentes sur le site de base.
## Explications de Docker

Docker est une plateforme de conteneurisation, c'est-à-dire une plateforme sur laquelle on peut lancer plusieurs applications contenues dans des "containers" (conteneurs). Ces containers, issus de "Docker Containers", sont de mini-conteneurs dans lesquels on peut déployer une application quelconque à l'aide d'une image de base. Cette image de base est représentée par un fichier nommé "Dockerfile". Cependant, lancer les applications une par une serait fastidieux, surtout s'il y en a plusieurs, comme 500 ou même seulement 20. C'est là qu'intervient "docker-compose". Docker compose permet de lancer plusieurs applications en même temps.

### Explications Détaillées de Docker

- **Dockerfile :** ![Exemple de Dockerfile](https://www.data-transitionnumerique.com/wp-content/uploads/2022/05/dockerfile.png)
    - "FROM node.." signifie qu'on utilisera Node.js dans cette application.
    - "WORKDIR /app" représente le dossier dans lequel on veut copier et coller nos fichiers utilisés dans l'application.
    - "COPY" copie et colle nos fichiers.
    - "RUN" permet de lancer l'instruction suivante.
    - "EXPOSE" : Port utilisé pour lancer cette application. Par exemple, "3000" = le port 3000.
    - "CMD" lance l'application entière.

- **docker-compose.yml :** ![Exemple de docker-compose.yml](https://www.data-transitionnumerique.com/wp-content/uploads/2022/05/docker-compose-yml.png)
    - "Version".
    - "Services" qui signifie les containers Docker, donc les applications.
        - "nodeserver" est un container Docker avec un port et un build spécifiés.
      ## Installation de NGINX

NGINX est un serveur web populaire utilisé pour servir du contenu web statique, agir en tant que proxy inverse et équilibrer la charge des serveurs. Voici comment l'installer sur différentes plateformes.

### Installation sous Ubuntu / Debian

**Mise à jour des paquets :**

- Avant d'installer NGINX, assurez-vous que vos paquets sont à jour :

    ```bash 
    sudo apt update
    ```


**Installation de NGINX :**

- Installez NGINX avec la commande suivante :

  ```bash sudo apt install nginx```

**Vérification de l'installation :**

- Vérifiez si NGINX est en cours d'exécution :

    ```bash
    sudo systemctl status nginx
    ```

### Installation sous macOS

**Installation avec Homebrew :

- Installez NGINX avec Homebrew :

    ```bash
    brew install nginx```

**Vérification de l'installation :**

- Démarrez NGINX et vérifiez dans votre navigateur web en accédant à http://localhost/.

### Installation sous Windows

**Téléchargement de l'exécutable :**

- Téléchargez l'exécutable NGINX depuis le site officiel.

    - Installation : Suivez les instructions d'installation de l'exécutable téléchargé.

**Vérification de l'installation :**

- Ouvrez une invite de commande et exécutez :

    ```bash
    cmd nginx -v
    ```

Cela devrait afficher la version de NGINX.
## Installation d'Apache HTTP Server

Apache HTTP Server est un serveur web open-source largement utilisé. Voici comment l'installer sur différentes plateformes.

### Installation sous Ubuntu / Debian

**Mise à jour des paquets :**

- Avant d'installer Apache, assurez-vous que vos paquets sont à jour :

    ```bash 
    sudo apt update
    ```

**Installation d'Apache :**

- Installez Apache avec la commande suivante :

    ```bash
    sudo apt install apache2
    ```

**Vérification de l'installation :**

- Vérifiez si Apache est en cours d'exécution :

    ```bash
    sudo systemctl status apache2
    ```

### Installation sous macOS

**Installation avec Homebrew :**

- Installez Apache avec Homebrew :

    ```bash
    brew install httpd
    ```

**Démarrage et vérification :**

- Démarrez Apache avec la commande suivante et vérifiez dans votre navigateur web en accédant à http://localhost/ :

    ```bash
    sudo apachectl start
    ```

### Installation sous Windows

**Téléchargement de l'installateur :**

- Téléchargez l'installateur Apache depuis le site officiel d'Apache.

**Installation : Suivez les instructions d'installation de l'installateur téléchargé.**

**Vérification de l'installation :**

- Ouvrez une invite de commande et exécutez :

    ```bash
    cmd httpd -v
    ```

Cela devrait afficher la version d'Apache installée.

## Configuration du Reverse Proxy

Pour le Reverse proxy il faut donner dans le fichier nginx.conf les redirection vers les sites qui s'occupe en même 
temps de vérifiés la sécurité grâce au clé générés par LetsEncrypt et mise dans le serveur Web.


## Services choisis

Premièrement nous avons utilisés des sites web que nous avions fait auparavant et intégrer les bases de données. 
L'utilisation de Nginx pour le reverse proxy avec LetsEncypt.
Chaque services, c'est à dire les sites Web, le serveur Web et le reverse proxy sont dans des containers gèrer par
Docker Compose qui nous permet de lancer la totalité des services par une simple commande:

``docker compose up -d``


## Ajout d'un site 

Pour ajouter un site, il faut intégrer tout les fichiers nécessaires dans le fichier principal, en l'occurence 
InfraCompose.
A la suite de cela, il faut configurer le fichier nginx.conf qui va permettre de rediriger vers le nom de domaines que 
nous voulons pour ce nouveau site. 
La modification du fichier docker-compose.yml est aussi nécessaire pour déployer un container avec les autres containers
et permettre de gèrer le reverse-proxy en ayant gènèrer une clé pour rendre la connexion sécuriser.
Il faut pour cela faire un Dockerfile décrivant les étapes à suivre par Docker Compose pour rendre l'application 
utilisable.