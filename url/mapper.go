package url

// ToURL parses DTO to URL Struct
func ToURL(dto DTO) URL {
	return URL{Code: dto.Code, OriginalURL: dto.OriginalURL}
}

// ToDTO parses URL to DTO Struct
func ToDTO(url URL) DTO {
	return DTO{ID: url.ID, Code: url.Code, OriginalURL: url.OriginalURL}
}
