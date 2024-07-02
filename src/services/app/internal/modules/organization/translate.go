package organization

import (
	"github.com/luminancetech/varso/src/services/app/client/models"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build"
)

func TranslateOrganizationsToMap(orgs []*build.Organization) map[string]*models.Organization {
	translated := map[string]*models.Organization{}

	for _, org := range orgs {
		translated[org.ID.String()] = TranslateOrganization(org)
	}

	return translated
}

func TranslateOrganizations(orgs []*build.Organization) []*models.Organization {
	var translated []*models.Organization

	for _, org := range orgs {
		translated = append(translated, TranslateOrganization(org))
	}

	return translated
}

func TranslateOrganization(org *build.Organization) *models.Organization {
	return &models.Organization{
		Uuid:         org.ID.String(),
		UniqueName:   org.UniqueName,
		Description:  org.Description,
		Name:         org.Name,
		WebsiteUrl:   org.WebsiteURL,
		LogoImageUrl: org.LogoImageURL,
	}
}
