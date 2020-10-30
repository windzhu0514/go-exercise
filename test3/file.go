package main

type name struct {
	Airlines []struct {
		DeptCity struct {
			IsDomestic string `json:"isDomestic"`
			JcName     string `json:"jcName"`
		} `json:"deptCity"`
		DestCities []struct {
			AirlineFlag string `json:"airlineFlag"`
			IsDomestic  string `json:"isDomestic"`
		} `json:"destCities"`
	} `json:"airlines"`
	DeptHotCities []struct {
		IsDomestic string `json:"isDomestic"`
		JcName     string `json:"jcName"`
		Name       string `json:"name"`
		SpellName  string `json:"spellName"`
		Tag        string `json:"tag"`
		ThreeCode  string `json:"threeCode"`
	} `json:"deptHotCities"`
	DestHotCities []struct {
		IsDomestic string `json:"isDomestic"`
		JcName     string `json:"jcName"`
		Name       string `json:"name"`
		SpellName  string `json:"spellName"`
		Tag        string `json:"tag"`
		ThreeCode  string `json:"threeCode"`
	} `json:"destHotCities"`
	IfSuccess string `json:"ifSuccess"`
	UsedTime  int64  `json:"usedTime"`
}
