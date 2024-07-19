package models

type MegaMenus struct {
	Data []MMData `json:"data,omitempty"`
}

type MegaMenu struct {
	Data MMData `json:"data,omitempty"`
}
type MMData struct {
	ID         int          `json:"id,omitempty"`
	Attributes MMAttributes `json:"attributes,omitempty"`
}
type MMItems struct {
	Data []MMData `json:"data,omitempty"`
}
type MMAttributes struct {
	Title    string  `json:"title,omitempty"`
	Slug     string  `json:"slug,omitempty"`
	Url      string  `json:"url,omitempty"`
	Items    MMItems `json:"items,omitempty"`
	Children MMItems `json:"children,omitempty"`
}

func (m *MegaMenu) FindParentURLs(targetURL string) []string {
	var parentURLs []string
	for _, item := range m.Data.Attributes.Items.Data {
		if findParentURLs(&parentURLs, item, targetURL) {
			parentURLs = append(parentURLs, m.Data.Attributes.Url)
		}
	}
	return parentURLs
}

func findParentURLs(parentURLs *[]string, data MMData, targetURL string) bool {
	if data.Attributes.Url == targetURL {
		return true
	}

	for _, child := range data.Attributes.Children.Data {
		if findParentURLs(parentURLs, child, targetURL) {
			*parentURLs = append(*parentURLs, data.Attributes.Url)
			return true
		}
	}

	return false
}

func (m *MegaMenu) FindBreadcrumbs(targetURL string) []Breadcrumb {
	var breadcrumbs []Breadcrumb
	for _, item := range m.Data.Attributes.Items.Data {
		findBreadcrumbs(&breadcrumbs, item, targetURL)
	}
	return breadcrumbs
}

func findBreadcrumbs(breadcrumbs *[]Breadcrumb, data MMData, targetURL string) bool {
	if data.Attributes.Url == targetURL {
		*breadcrumbs = append(*breadcrumbs, Breadcrumb{Name: data.Attributes.Title, Slug: data.Attributes.Url})
		return true
	}

	for _, child := range data.Attributes.Children.Data {
		if findBreadcrumbs(breadcrumbs, child, targetURL) {
			*breadcrumbs = append(*breadcrumbs, Breadcrumb{Name: data.Attributes.Title, Slug: data.Attributes.Url})
			return true
		}
	}

	return false
}
