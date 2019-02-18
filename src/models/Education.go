package Models

type Education struct
{
	Id string `json:"id"`
	School string `json"school"`
	Start string `json"start"`
	End string `json"end"`
	Scholarship []string `json"scholarships"`
	Award []string `json"awards"`
	Location string `json"location"`
	Title string `json"title"`
	NotableProject []string `json"notableProj"`
	ExtraCurcicular []string `json"extraCuric"`
	Image string `json"img"`	
}