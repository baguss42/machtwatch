package factory

import (
	"github.com/baguss42/machtwatch/entity"
)

type BrandFactory struct {
	entity.Brand
}

func (f *BrandFactory) Build() entity.Brand {
	return entity.Brand{
		Name:        "Seiko",
		Description: "Seiko is one of the few fully integrated watch manufactures",
		Level:       "large",
		Logo:        "https://ae01.alicdn.com/kf/He340df489bae489ca212adbc6e2557221/Seiko-Watch-Pria-5-Automatic-Watch-Top-Mewah-Merek-Tahan-Air-Olahraga-Jam-Tangan-Jam-Tangan.jpg",
	}
}
