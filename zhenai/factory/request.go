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

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	req.Header.Set("Cookie", fmt.Sprintf("%s=%s;%s=%s;", "FSSBBIl1UgzbN7NO", "5_A7ueMiKN05arKvmpPpEIkz2liWZlfOydjMH2FAVZBHPWwF5pjWcaiegVO3iMn7N19YPQtM.Qj_XsN2bDFvqvG", "FSSBBIl1UgzbN7NP", "53rkFfKqACLWqqqm6AwKGtA31c_G4tUv5OiEvq_squ13KTaFdS5LGfoFUKzQoKD.GTW0T6piW4YOSYA9BY50IrYmm8MzdhOFx48FnPhZi8s4UpYPUcudAKpw2wYq8mUPrHMjgn6ZYNEigtjONk3NlLHXwHyfUn0CGuX.A6Pc._mcrUXZrAMWvdPb3oy7HrtkNsFlCiU0nDF_KYW3_.mVEk.v5ndLYQpkAruzWxAHeq_iyXiSF0pSpZPMqK8WpNsKO3"))

	return req, nil
}
