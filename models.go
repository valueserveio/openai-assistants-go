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

type ThreadObject struct {
  ID        string `json:"id"`
  Object    string `json:"object"`
  CreatedAt int    `json:"created_at"`
  Metadata  struct {
  } `json:"metadata"`
}

type MessageObject struct {
  ID        string `json:"id"`
  Object    string `json:"object"`
  CreatedAt int    `json:"created_at"`
  ThreadID  string `json:"thread_id"`
  Role      string `json:"role"`
  Content   []struct {
    Type string `json:"type"`
    Text struct {
      Value       string `json:"value"`
      Annotations []any  `json:"annotations"`
    } `json:"text"`
  } `json:"content"`
  AssistantID string `json:"assistant_id"`
  RunID       string `json:"run_id"`
  Attachments []any  `json:"attachments"`
  Metadata    struct {
  } `json:"metadata"`
}

type RunObject struct {
  ID           string `json:"id"`
  Object       string `json:"object"`
  CreatedAt    int    `json:"created_at"`
  AssistantID  string `json:"assistant_id"`
  ThreadID     string `json:"thread_id"`
  Status       string `json:"status"`
  StartedAt    int    `json:"started_at"`
  ExpiresAt    any    `json:"expires_at"`
  CancelledAt  any    `json:"cancelled_at"`
  FailedAt     any    `json:"failed_at"`
  CompletedAt  int    `json:"completed_at"`
  LastError    any    `json:"last_error"`
  Model        string `json:"model"`
  Instructions any    `json:"instructions"`
  Tools        []struct {
    Type string `json:"type"`
  } `json:"tools"`
  Metadata struct {
  } `json:"metadata"`
  IncompleteDetails any `json:"incomplete_details"`
  Usage             struct {
    PromptTokens     int `json:"prompt_tokens"`
    CompletionTokens int `json:"completion_tokens"`
    TotalTokens      int `json:"total_tokens"`
  } `json:"usage"`
  Temperature         float64 `json:"temperature"`
  TopP                float64 `json:"top_p"`
  MaxPromptTokens     int     `json:"max_prompt_tokens"`
  MaxCompletionTokens int     `json:"max_completion_tokens"`
  TruncationStrategy  struct {
    Type         string `json:"type"`
    LastMessages any    `json:"last_messages"`
  } `json:"truncation_strategy"`
  ResponseFormat    string `json:"response_format"`
  ToolChoice        string `json:"tool_choice"`
  ParallelToolCalls bool   `json:"parallel_tool_calls"`
}
