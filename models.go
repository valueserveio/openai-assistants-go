package openaiassistantsgo

type AssistantObject struct {
  ID           string `json:"id"`
  Object       string `json:"object"`
  CreatedAt    int    `json:"created_at"`
  Name         string `json:"name"`
  Description  any    `json:"description"`
  Model        string `json:"model"`
  Instructions string `json:"instructions"`
  Tools        []struct {
    Type string `json:"type"`
  } `json:"tools"`
  Metadata struct {
  } `json:"metadata"`
  TopP           float64 `json:"top_p"`
  Temperature    float64 `json:"temperature"`
  ResponseFormat string  `json:"response_format"`
}
