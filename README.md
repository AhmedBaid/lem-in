# 🐜 Lem-in: Simulation d'une Fourmilière

Lem-in est un projet de simulation où un groupe de fourmis doit se déplacer d'une salle de départ vers une salle d'arrivée en empruntant un réseau de tunnels. L'objectif est d'optimiser leurs déplacements pour minimiser le nombre de tours nécessaires.

---

## 📌 Fonctionnalités
✅ Lecture et analyse des données d'entrée (nombre de fourmis, salles, liens).  
✅ Recherche des chemins possibles entre la salle de départ et la salle d'arrivée.  
✅ Répartition intelligente des fourmis pour optimiser leur parcours.  
✅ Affichage du mouvement des fourmis à chaque tour.  

---

## 📁 Structure du projet
```
lem-in/
│── cmd/
│   └── main.go         # Point d'entrée du programme
│── bfs/
│   └── bfs.go          # Algorithme de recherche de chemins (BFS)
│── parsing/
│   └── parsing.go      # Lecture et analyse du fichier d'entrée
│── printage/
│   └── printage.go     # Affichage du déplacement des fourmis
│── utils/
│   └── utils.go        # Fonctions utilitaires et structures de données
│── README.md           # Documentation
```

---

## 📥 Installation et exécution
### 1️⃣ Prérequis
- **Go** installé sur votre machine (`go version` pour vérifier).

### 2️⃣ Cloner le projet
```sh
git clone https://github.com/votre-repo/lem-in.git
cd lem-in
```

### 3️⃣ Compiler et exécuter
```sh
go run . test0.txt
```
Remplace `test0.txt` par le fichier contenant la description de la fourmilière.

---

## 📜 Format du fichier d'entrée
Le fichier d'entrée doit respecter le format suivant :
```
3               # Nombre de fourmis
##start         # Début de la définition des salles
A 1 2          # Salle de départ
B 3 4
C 5 6
##end           # Début de la salle d'arrivée
D 7 8          # Salle de fin
A-B            # Liens entre les salles
A-C
B-D
C-D
```

---

## 🔎 Explication du fonctionnement
1. **Lecture du fichier d'entrée** (`parsing.go`) :
   - Extraction du nombre de fourmis, des salles et des tunnels.
2. **Recherche des chemins optimaux** (`graph.go`) :
   - Algorithme BFS pour trouver les meilleurs chemins.
3. **Simulation du déplacement des fourmis** (`printage.go`) :
   - Répartition des fourmis sur les chemins et affichage des déplacements.

---

## 📖 Exemple d'exécution
**Entrée (test0.txt) :**
```
3
##start
A 1 2
B 3 4
C 5 6
##end
D 7 8
A-B
A-C
B-D
C-D
```

**Sortie attendue :**
```
L1-B L2-C
L1-D L2-D L3-B
L3-D
```

---

## 🤝 Contribution
Les contributions sont les bienvenues ! Pour proposer des améliorations :
1. **Fork** le projet.
2. Crée une branche (`git checkout -b feature-nouvelle-fonctionnalité`).
3. Fais tes modifications et commit (`git commit -m "Ajout d'une nouvelle fonctionnalité"`).
4. Pousse tes changements (`git push origin feature-nouvelle-fonctionnalité`).
5. Crée une Pull Request.

---

## 🛠 Technologies utilisées
- **Langage :** Go
- **Algorithme principal :** BFS (Breadth-First Search)
- **Gestion des fichiers :** `os` et `strings`

---

## 📄 Licence
Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.