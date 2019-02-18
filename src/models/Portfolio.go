package Models

type Portfolio struct
{
	Id string `json:"id"`
	Name string `json:"name"`
	DisplayImage string `json:"dspImg"`
	Description []string `json:"desc"`
	Image []string `json:"images"`
	Video []string `json:"videos"`
	Year string `json:"year"`
}