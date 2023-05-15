package get_file

type Response struct {
	ID          string `json:"id"`
	FileName    string `json:"file_name"`
	FileContent []byte `json:"file_content"`
}
