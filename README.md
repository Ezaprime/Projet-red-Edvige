Cyberynov ⊂(◉‿◉)つ

Bienvenue sur Cyberynov — un RPG console en Go (Golang) dans un univers cyberpunk.
Crée ton héros, gère ta sacoche (inventaire), combats au tour par tour, achète au marchand, fabrique chez le forgeron et termine le mode Histoire. Le boss final ne tombe qu’avec l’arme Puff (très chère, niveau 3+).

1) Projet Golang 🧩

Exécutable CLI écrit en Go (≥ 1.20).

Architecture simple : un module, plusieurs fichiers, un point d’entrée.

Données de gameplay (prix, dégâts, recettes, coûts de mana) centralisées dans des maps pour faciliter l’équilibrage.

Construction incrémentale via Git (petites branches/PR, main compilable en continu).

2) Fonctionnalités 🕹️

Sacoche (inventaire) : Stimpak, batteries de mana, puces de sorts, armes, armures, upgrades de capacité.

Combat au tour par tour : initiative, attaque basique, sorts (Coup de poing / Boule de feu), arme équipée, IA avec pics de dégâts.

Marchand & Forgeron : achats (dont la Puff, chère et gated niv. 3+) ; craft d’armures via recettes.

Mode Histoire : enchaînement de chapitres jusqu’au boss final (immunisé aux attaques/sorts classiques → Puff requise).

Entraînement : combats sûrs, récompenses ÷4 (XP + crédits) par rapport à l’histoire.

Progression : XP → level up (+PV base, +Attaque, +Mana). Boule de feu se débloque auto au niveau 2.

UX Console : bannière animée au démarrage, couleurs ANSI, barres PV/Mana dégradées, récap de fin de combat.

3) Technologies & Langages 🛠️

Langage : Go (Golang).

I/O console : fmt.Scanln + sorties ANSI (couleurs, effets).

Structures :

Character, Monster, Equipment, Chapter (structs),

Slices pour la sacoche,

Maps pour prix/dégâts/coûts/recettes.

Passage par pointeur (*Character) pour modifier l’état en place.

4) Utilisation 🚀
Prérequis

Go installé : go version (≥ 1.20)

Git

Installation

Cloner le dépôt :

git clone https://github.com/<org>/<repo>.git


Entrer dans le projet :

cd <repo>


Synchroniser les dépendances (si besoin) :

go mod tidy

Démarrer

Lancer (dev) :

go run .


Compiler puis lancer :

go build -o cyberynov
./cyberynov


Au démarrage : bannière animée “CYBERYNOV” → saisie du nom du joueur → menu principal.

5) Démo du jeu / Présentation 🎬

Menu principal

1) Infos : stats, classe, chapitre, crédits, barres PV/Mana.

2) Sacoche : utiliser Stimpak/batteries, apprendre Boule de feu via puce, équiper armes/armures, augmenter la capacité.

3) Marchand : acheter consommables et armes (Pistolet/SMG/Katana), Puff disponible niv. 3+.

4) Forgeron : craft d’armures avec ressources (recettes simples).

5) Entraînement : combat d’essai, gains ÷4 (XP + crédits).

6) Mode Histoire : progression par chapitres jusqu’au boss final.

0) Quitter.

Combat

Jet d’initiative → alternance joueur/ennemi.

Actions : attaque basique, sorts, arme équipée, sacoche.

IA : coup renforcé périodique, feedbacks colorés/animés.

Objectif

Atteindre niveau ≥ 3 + assez de crédits → acheter/équiper Puff → battre le boss (seule action efficace sur lui).

Version 💾

Go : 1.22.x (compatible ≥ 1.20)

OS : Linux / macOS / Windows (terminal ANSI recommandé)

Authors ✍️

Kerem  — gameplay, menus, combats

Taali — sacoche, marchand/forge, histoire

Collaboration via GitHub : petites branches, PR régulières, merge sur main.