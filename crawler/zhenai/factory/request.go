package factory

import (
	"fmt"
	"io"
	"net/http"
)

func NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	req.Header.Set("Cookie", fmt.Sprintf("%s=%s;%s=%s;",
		"FSSBBIl1UgzbN7NO",
		"5_A7ueMiKN05arKvmpPpEIkz2liWZlfOydjMH2FAVZBHPWwF5pjWcaiegVO3iMn7N19YPQtM.Qj_XsN2bDFvqvG",
		"FSSBBIl1UgzbN7NP",
		"53rtrDKq1Gp3qqqm61lfpLGI6ujYYQPDoXaVBlBjVq8by0wmmBAy4S8t65JcrC9kOg8EMGzcowg7MIgZorHOpNl4Id9cH7JR6mQLD3bHYvgJgT1ooTLHWg8OuktPk9aKXpqyhkutcsF9PgxxXf7PsRJRfBKMZWWh9Egsg6nJtfeZA1z4DyDpnT1pKkv34h264ot.haxBAji6Rva83r9Lff5heMxeN2ykLzW__YpEN8jHFj3NuXyeqU3INYh1WO9MrOP51abfgHYpGI841DqHaW0"))

	return req, nil
}
