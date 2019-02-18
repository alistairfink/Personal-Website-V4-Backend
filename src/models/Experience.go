package Models

type Experience struct
{
	Id string `json:"id"`
	Position string `json"position"`
	Start string `json"start"`
	End string `json"end"`
	Company string `json"comp"`
	Location string `json"location"`
	Data []string `json"data"`
	Image string `json"img"`
}