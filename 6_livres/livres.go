package main

import (
	"fmt"
)

type Livre struct {
	ID          int
	Titre       string
	Auteur      string
	Description string
}

func NouveauLivre(id int, titre, auteur, description string) Livre {
	return Livre{
		ID:          id,
		Titre:       titre,
		Auteur:      auteur,
		Description: description,
	}
}

func AfficherDetails(livre Livre) {
	fmt.Printf("ID: %d\nTitre: %s\nAuteur: %s\nDescription: %s\n", livre.ID, livre.Titre, livre.Auteur, livre.Description)
}

type Bibliotheque struct {
	Livres []Livre
}

func (b *Bibliotheque) AjouterLivre(livre Livre) {
	b.Livres = append(b.Livres, livre)
}

func (b *Bibliotheque) AfficherListeLivres() {
	for _, livre := range b.Livres {
		AfficherDetails(livre)
	}
}

func (b *Bibliotheque) RechercherLivreParID(id int) *Livre {
	for _, livre := range b.Livres {
		if livre.ID == id {
			return &livre
		}
	}
	return nil
}

func main() {
	livre1 := NouveauLivre(1, "Le Petit Prince", "Antoine de Saint-Exupéry", "Un conte poétique et philosophique.")
	fmt.Println("----- Détails du livre:")
	AfficherDetails(livre1)

	livre2 := NouveauLivre(2, "1984", "George Orwell", "Un roman d'anticipation dystopique.")
	bibliotheque := Bibliotheque{}
	bibliotheque.AjouterLivre(livre1)
	bibliotheque.AjouterLivre(livre2)
	fmt.Println("----- Liste des livres:")
	bibliotheque.AfficherListeLivres()

	livreRecherche := bibliotheque.RechercherLivreParID(1)
	if livreRecherche != nil {
		fmt.Println("----- Livre trouvé:")
		AfficherDetails(*livreRecherche)
	} else {
		fmt.Println("Livre non trouvé.")
	}
}
