package grifts

import (
	"supertests/models"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {
	grift.Desc("seed", "Database seeds")
	grift.Add("seed", func(c *grift.Context) error {
		return models.DB.Transaction(func(tx *pop.Connection) error {
			tests := []models.Test{
				{
					Title:       "Quais Famosos estão te xavecando no whatsapp?",
					Cover:       "o-que-os-famosos-estao-falando-sobre-voce/whatsappcapa.jpg",
					Slug:        "quais-famosos-estao-te-xavecando-no-whatsapp",
					Description: "Faça o teste e veja quais famosos estão te enviando mensagens no WhatsAPP",
					Message:     "[[nome]], olha o que os famosos estão falando sobre você, gostou?",
					Class:       "App\\Testes\\famosos",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					Title:       "Qual deveria ser seu salario?",
					Cover:       "qual-deveria-ser-seu-salario/capa.jpg",
					Slug:        "qual-deveria-ser-seu-salario",
					Description: "Faça o teste e descubra quanto deveria ser realmente seu salário",
					Message:     "[[nome]], é isto que você deveria estar ganhando. Gostou?",
					Class:       "App\\Testes\\Salario",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					Title:       "O que o coelhinho da Páscoa vai trazer pra você?",
					Cover:       "o-que-o-coelhinho-da-pascoa-vai-trazer-para-voce/coelinho230317.jpg",
					Slug:        "o-que-o-coelhinho-da-pascoa-vai-trazer-para-voce",
					Description: "Faça o teste e veja o que o coelhinho da Páscoa vai trazer para você!",
					Message:     "[[nome]], olha o que o coelho trouxe para você, gostou?",
					Class:       "App\\Testes\\CoelhoPascoa",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					Title:       "Como seria seus filhos Gêmeos?",
					Cover:       "como-seria-seus-filhos-gemeos/gemeoscapa.jpg",
					Slug:        "como-seria-seus-filhos-gemeos",
					Description: "Faça o teste e veja como seria seus filhos gêmeos",
					Message:     "[[nome]], olha como seria seus filhos, gostou?",
					Class:       "App\\Testes\\gemeos",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					Title:       "Qual celebridade se parece com você?",
					Cover:       "qual-celebridade-se-parece-com-voce/celebridadecapa.jpg",
					Slug:        "qual-celebridade-se-parece-com-voce",
					Description: "Descubra agora qual celebridade se parece mais com você!",
					Message:     "[[nome]], esta é a celebridade que você se parece. Gostou?",
					Class:       "App\\Testes\\Celebridade",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					Title:       "Qual sua verdadeira nacionalidade?",
					Cover:       "qual-sua-verdadeira-nacionalidade/nacionalidadecapa.jpg",
					Slug:        "qual-sua-verdadeira-nacionalidade",
					Description: "Faça o teste e veja qual é a sua nacionalidade",
					Message:     "[[nome]], Essa é a sua verdadeira nacionalidade. Gostou?",
					Class:       "App\\Testes\\nacionalidade",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			}

			if err := tx.TruncateAll(); err != nil {
				return err
			}

			for _, test := range tests {
				if err := tx.Create(&test); err != nil {
					return err
				}
			}

			return nil
		})
	})
})
