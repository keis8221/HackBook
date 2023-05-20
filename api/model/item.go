package model

type RakutenItems struct {
	GenreInformation []any `json:"GenreInformation"`
	Items            []struct {
		Item struct {
			AffiliateURL   string `json:"affiliateUrl"`
			ArtistName     string `json:"artistName"`
			Author         string `json:"author"`
			Availability   string `json:"availability"`
			BooksGenreID   string `json:"booksGenreId"`
			ChirayomiURL   string `json:"chirayomiUrl"`
			DiscountPrice  int    `json:"discountPrice"`
			DiscountRate   int    `json:"discountRate"`
			Hardware       string `json:"hardware"`
			Isbn           string `json:"isbn"`
			ItemCaption    string `json:"itemCaption"`
			ItemPrice      int    `json:"itemPrice"`
			ItemURL        string `json:"itemUrl"`
			Jan            string `json:"jan"`
			Label          string `json:"label"`
			LargeImageURL  string `json:"largeImageUrl"`
			LimitedFlag    int    `json:"limitedFlag"`
			ListPrice      int    `json:"listPrice"`
			MediumImageURL string `json:"mediumImageUrl"`
			Os             string `json:"os"`
			PostageFlag    int    `json:"postageFlag"`
			PublisherName  string `json:"publisherName"`
			ReviewAverage  string `json:"reviewAverage"`
			ReviewCount    int    `json:"reviewCount"`
			SalesDate      string `json:"salesDate"`
			SmallImageURL  string `json:"smallImageUrl"`
			Title          string `json:"title"`
		} `json:"Item"`
	} `json:"Items"`
	Carrier   int `json:"carrier"`
	Count     int `json:"count"`
	First     int `json:"first"`
	Hits      int `json:"hits"`
	Last      int `json:"last"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}
