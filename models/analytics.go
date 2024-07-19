package models

type RequestInfo struct {
	ID              string `json:"id,omitempty" bson:"_id,omitempty"`
	ClientIP        string `json:"client_ip,omitempty" bson:"client_ip,omitempty"`
	HttpMethod      string `json:"http_method,omitempty" bson:"http_method,omitempty"`
	URLPath         string `json:"url_path,omitempty" bson:"url_path,omitempty"`
	QueryParameters string `json:"query_parameters,omitempty" bson:"query_parameters,omitempty"`
	RequestBodySize int64  `json:"request_body_size,omitempty" bson:"request_body_size,omitempty"`
	ResponseStatus  int    `json:"response_status,omitempty" bson:"response_status,omitempty"`
	AcceptLang      string `json:"accept_language,omitempty" bson:"accept_language,omitempty"`
	Referer         string `json:"referer,omitempty" bson:"referer,omitempty"`
	Authorization   string `json:"authorization,omitempty" bson:"authorization,omitempty"`
	Cookie          string `json:"cookie,omitempty" bson:"cookie,omitempty"`
	XForwardedFor   string `json:"x_forwarded_for,omitempty" bson:"x_forwarded_for,omitempty"`
	XRequestedWith  string `json:"x_requested_with,omitempty" bson:"x_requested_with,omitempty"`
	CacheControl    string `json:"cache_control,omitempty" bson:"cache_control,omitempty"`
	ContentType     string `json:"content_type,omitempty" bson:"content_type,omitempty"`
	Accept          string `json:"accept,omitempty" bson:"accept,omitempty"`
	IfNoneMatch     string `json:"if_none_match,omitempty" bson:"if_none_match,omitempty"`
	Host            string `json:"host,omitempty" bson:"host,omitempty"`
	DNT             string `json:"dnt,omitempty" bson:"dnt,omitempty"`
	Timestamp       int64  `json:"timestamp,omitempty" bson:"timestamp,omitempty"`

	UserAgent   UserAgentInfo   `json:"user_agent,omitempty" bson:"user_agent,omitempty"`
	Geolocation GeolocationInfo `json:"geolocation,omitempty" bson:"geolocation,omitempty"`
}

type UserAgentInfo struct {
	BrowserName    string `json:"browser_name,omitempty" bson:"browser_name,omitempty"`
	BrowserVersion string `json:"browser_version,omitempty" bson:"browser_version,omitempty"`
	UserOS         string `json:"user_os,omitempty" bson:"user_os,omitempty"`
	IsMobile       bool   `json:"is_mobile,omitempty" bson:"is_mobile,omitempty"`
}

type GeolocationInfo struct {
	Country    string `json:"country,omitempty" bson:"country,omitempty"`
	Region     string `json:"region,omitempty" bson:"region,omitempty"`
	City       string `json:"city,omitempty" bson:"city,omitempty"`
	PostalCode string `json:"postal_code,omitempty" bson:"postal_code,omitempty"`
	Latitude   string `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude  string `json:"longitude,omitempty" bson:"longitude,omitempty"`
}

type QuickStat struct {
	ID             string   `json:"_id" bson:"_id"`
	Count          int      `json:"count" bson:"count"`
	UniqueIPs      []string `json:"uniqueIPs" bson:"uniqueIPs"`
	UniqueIPsCount int      `json:"uniqueIPsCount,omitempty"`
	UrlPaths       []string `json:"urlPaths" bson:"urlPaths"`
	UserOS         []string `json:"userOS" bson:"userOS"`
}

func (qs *QuickStat) CountUniqueIPs() {
	uniqueIPMap := make(map[string]bool)
	for _, ip := range qs.UniqueIPs {
		uniqueIPMap[ip] = true
	}
	qs.UniqueIPsCount = len(uniqueIPMap)
}

type QuickStats []QuickStat

type CountUniqueIPsURL struct {
	Count   int32  `json:"count,omitempty" bson:"count,omitempty"`
	URLPath string `json:"url_path,omitempty" bson:"url_path,omitempty"`
}

type CountUniqueIPsURLs []CountUniqueIPsURL
