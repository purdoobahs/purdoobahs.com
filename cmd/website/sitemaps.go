package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/purdoobahs/purdoobahs.com/internal/purdoobahs"
	"github.com/purdoobahs/purdoobahs.com/internal/sitemap"
	"github.com/purdoobahs/purdoobahs.com/internal/traditions"
)

func (app *application) generateSitemaps() error {
	imageEntryGeoLocation := "West Lafayette, Indiana USA"
	imageEntryLicense := "https://creativecommons.org/licenses/by-nc-nd/4.0/"

	// get all purdoobahs
	allPurdoobahs, err := app.purdoobahService.All()
	if err != nil {
		return err
	}

	// get all section years
	allSectionYears, err := app.purdoobahService.AllSectionYears()
	if err != nil {
		return err
	}

	// get all traditions
	allTraditions, err := app.traditionService.All()
	if err != nil {
		return err
	}

	// index
	err = app.generateIndexSitemap()
	if err != nil {
		return err
	}

	// root
	err = app.generateRootSitemap()
	if err != nil {
		return err
	}

	// profiles
	err = app.generateProfilesSitemap(allPurdoobahs, imageEntryGeoLocation, imageEntryLicense)
	if err != nil {
		return err
	}

	// sections
	err = app.generateSectionsSitemap(allSectionYears, imageEntryGeoLocation, imageEntryLicense)
	if err != nil {
		return err
	}

	// traditions
	err = app.generateTraditionsSitemap(allTraditions, imageEntryGeoLocation, imageEntryLicense)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) generateIndexSitemap() error {
	homeUrl := "https://www.purdoobahs.com"

	// use today's date to generate Last Modified
	lastModified := time.Now().Format(time.RFC3339)

	// generate index sitemap
	indexSitemap := sitemap.NewIndexFile([]sitemap.Entry{
		sitemap.NewSitemapEntry(fmt.Sprintf("%s%s", homeUrl, "/sitemap-root.xml"), lastModified),
		sitemap.NewSitemapEntry(fmt.Sprintf("%s%s", homeUrl, "/purdoobah/sitemap.xml"), lastModified),
		sitemap.NewSitemapEntry(fmt.Sprintf("%s%s", homeUrl, "/section/sitemap.xml"), lastModified),
		sitemap.NewSitemapEntry(fmt.Sprintf("%s%s", homeUrl, "/tradition/sitemap.xml"), lastModified),
	})

	// write index sitemap to disk
	indexSitemapFilepath := "/static/file/sitemap-index.xml"
	err := indexSitemap.WriteToFile(fmt.Sprintf(".%s", indexSitemapFilepath))
	if err != nil {
		return err
	}

	// add to CacheBuster
	err = app.cacheBuster.Add(indexSitemapFilepath)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) generateRootSitemap() error {
	homeUrl := "https://www.purdoobahs.com"

	// use today's date to generate Last Modified
	lastModified := time.Now().Format(time.RFC3339)

	// generate root sitemap
	rootSitemap := sitemap.NewFile([]sitemap.UrlEntry{})

	// add new UrlEntry for each root route
	routes := []string{"", "alumni", "tradition", "cravers-hall-of-fame"}
	for _, route := range routes {
		var images []sitemap.ImageEntry
		urlEntry, err := sitemap.NewUrlEntry(
			fmt.Sprintf("%s/%s", homeUrl, route),
			lastModified,
			sitemap.Weekly,
			0.5,
			images,
		)
		if err != nil {
			return err
		}
		rootSitemap.AddUrl(urlEntry)
	}

	// write root sitemap to disk
	rootSitemapFilepath := "/static/file/sitemap-root.xml"
	err := rootSitemap.WriteToFile(fmt.Sprintf(".%s", rootSitemapFilepath))
	if err != nil {
		return err
	}

	// add to CacheBuster
	err = app.cacheBuster.Add(rootSitemapFilepath)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) generateProfilesSitemap(allPurdoobahs []*purdoobahs.Purdoobah, imageEntryGeoLocation, imageEntryLicense string) error {
	homeUrl := "https://www.purdoobahs.com/purdoobah"
	baseImageUrl := "https://www.purdoobahs.com"

	// use today's date to generate Last Modified
	lastModified := time.Now().Format(time.RFC3339)

	// init profiles sitemap
	profilesSitemap := sitemap.NewFile([]sitemap.UrlEntry{})

	// add new UrlEntry for every Purdoobah (by ID, not Name/BirthCertificateName)
	for _, purdoobah := range allPurdoobahs {
		// generate profile image entry
		var images []sitemap.ImageEntry
		if !strings.HasSuffix(purdoobah.Metadata.Image.File, "_unknown.webp") {
			// only add profile image if they have one
			profileImage, err := sitemap.NewImageEntry(
				fmt.Sprintf("%s%s", baseImageUrl, purdoobah.Metadata.Image.File),
				purdoobah.Name,
				purdoobah.Metadata.Image.Alt,
				imageEntryGeoLocation,
				imageEntryLicense,
			)
			if err != nil {
				return err
			}
			images = append(images, profileImage)
		}

		// generate URL entry
		urlEntry, err := sitemap.NewUrlEntry(
			fmt.Sprintf("%s/%s", homeUrl, purdoobah.ID),
			lastModified,
			sitemap.Weekly,
			0.5,
			images,
		)
		if err != nil {
			return err
		}
		profilesSitemap.AddUrl(urlEntry)
	}

	// write profiles sitemap to disk
	profilesSitemapFilepath := "/static/file/sitemap-profiles.xml"
	err := profilesSitemap.WriteToFile(fmt.Sprintf(".%s", profilesSitemapFilepath))
	if err != nil {
		return err
	}

	// add to CacheBuster
	err = app.cacheBuster.Add(profilesSitemapFilepath)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) generateSectionsSitemap(allSectionYears []int, imageEntryGeoLocation, imageEntryLicense string) error {
	homeUrl := "https://www.purdoobahs.com/section"
	baseImageUrl := "https://www.purdoobahs.com/static/image/section"

	// use today's date to generate Last Modified
	lastModified := time.Now().Format(time.RFC3339)

	// init profiles sitemap
	sectionsSitemap := sitemap.NewFile([]sitemap.UrlEntry{})

	// add new UrlEntry for every unique section year
	for _, uniqueYear := range allSectionYears {
		var images []sitemap.ImageEntry

		if uniqueYear == -1 {
			urlEntry, err := sitemap.NewUrlEntry(
				fmt.Sprintf("%s/unknown", homeUrl),
				lastModified,
				sitemap.Weekly,
				0.5,
				images,
			)
			if err != nil {
				return err
			}
			sectionsSitemap.AddUrl(urlEntry)
			continue
		}

		// generate section header image entry
		if app.doesSectionHaveSocialImage(uniqueYear) {
			// generate list of all purdoobah names from this year for image caption
			sectionByYear, err := app.purdoobahService.SectionByYear(uniqueYear)
			if err != nil {
				return err
			}
			var purdoobahNames []string
			for _, purdoobah := range sectionByYear {
				purdoobahNames = append(purdoobahNames, purdoobah.Name)
			}
			commaDelimitedPurdoobahNames := strings.Join(purdoobahNames, ", ")

			// only add section header image if it has one
			sectionImage, err := sitemap.NewImageEntry(
				fmt.Sprintf("%s/%d.webp", baseImageUrl, uniqueYear),
				fmt.Sprintf("The Section of %d", uniqueYear),
				fmt.Sprintf("The Section of %d: %s", uniqueYear, commaDelimitedPurdoobahNames),
				imageEntryGeoLocation,
				imageEntryLicense,
			)
			if err != nil {
				return err
			}
			images = append(images, sectionImage)
		}

		// generate URL entry
		urlEntry, err := sitemap.NewUrlEntry(
			fmt.Sprintf("%s/%d", homeUrl, uniqueYear),
			lastModified,
			sitemap.Weekly,
			0.5,
			images,
		)
		if err != nil {
			return err
		}
		sectionsSitemap.AddUrl(urlEntry)
	}

	// write sections sitemap to disk
	sectionsSitemapFilepath := "/static/file/sitemap-sections.xml"
	err := sectionsSitemap.WriteToFile(fmt.Sprintf(".%s", sectionsSitemapFilepath))
	if err != nil {
		return err
	}

	// add to CacheBuster
	err = app.cacheBuster.Add(sectionsSitemapFilepath)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) generateTraditionsSitemap(allTraditions []*traditions.Tradition, imageEntryGeoLocation, imageEntryLicense string) error {
	homeUrl := "https://www.purdoobahs.com/tradition"
	baseImageUrl := "https://www.purdoobahs.com"

	// use today's date to generate Last Modified
	lastModified := time.Now().Format(time.RFC3339)

	// init profiles sitemap
	traditionsSitemap := sitemap.NewFile([]sitemap.UrlEntry{})

	// add new UrlEntry for every Purdoobah (by ID, not Name/BirthCertificateName)
	for _, tradition := range allTraditions {
		// generate tradition image entry
		var images []sitemap.ImageEntry
		if !strings.HasSuffix(tradition.Metadata.Image.File, "_unknown.webp") {
			// only add tradition image if it has one
			traditionImage, err := sitemap.NewImageEntry(
				fmt.Sprintf("%s%s", baseImageUrl, tradition.Metadata.Image.File),
				tradition.Name,
				tradition.Metadata.Image.Alt,
				imageEntryGeoLocation,
				imageEntryLicense,
			)
			if err != nil {
				return err
			}
			images = append(images, traditionImage)
		}

		// generate URL entry
		urlEntry, err := sitemap.NewUrlEntry(
			fmt.Sprintf("%s/%s", homeUrl, tradition.ID),
			lastModified,
			sitemap.Weekly,
			0.5,
			images,
		)
		if err != nil {
			return err
		}
		traditionsSitemap.AddUrl(urlEntry)
	}

	// write traditions sitemap to disk
	traditionsSitemapFilepath := "/static/file/sitemap-traditions.xml"
	err := traditionsSitemap.WriteToFile(fmt.Sprintf(".%s", traditionsSitemapFilepath))
	if err != nil {
		return err
	}

	// add to CacheBuster
	err = app.cacheBuster.Add(traditionsSitemapFilepath)
	if err != nil {
		return err
	}

	return nil
}
