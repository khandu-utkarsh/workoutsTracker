package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	models "workout_app_backend/internal/models"
)

// LLMClient handles communication with the LLM service
type LLMClient struct {
	baseURL    string
	httpClient *http.Client
}

// ChatMessage represents a message in the conversation
type ChatMessage struct {
	Role    string `json:"role"` // "user", "assistant", "system"
	Content string `json:"content"`
}

// ChatRequest represents the request to the LLM service
type ChatRequest struct {
	Messages       []ChatMessage          `json:"messages"`
	UserID         string                 `json:"user_id"`
	ConversationID string                 `json:"conversation_id"`
	Context        map[string]interface{} `json:"context,omitempty"`
}

// ChatResponse represents the response from the LLM service
type ChatResponse struct {
	Message MessageContent `json:"message"`
}

// MessageContent represents the content of the message from the LLM service
type MessageContent struct {
	Content          string                 `json:"content"`
	AdditionalKwargs map[string]interface{} `json:"additional_kwargs"`
	ResponseMetadata map[string]interface{} `json:"response_metadata"`
	Type             string                 `json:"type"`
	Name             *string                `json:"name"`
	ID               string                 `json:"id"`
}

// NewLLMClient creates a new LLM client
func NewLLMClient(baseURL string) *LLMClient {
	return &LLMClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// ProcessChatMessage sends a chat message to the LLM service and returns the response
func (c *LLMClient) ProcessChatMessage(ctx context.Context, messages []models.Message, userID, conversationID int64, context map[string]interface{}) (*ChatResponse, error) {
	// Convert messages to the format expected by the LLM service
	chatMessages := make([]ChatMessage, len(messages))
	for i, msg := range messages {
		chatMessages[i] = ChatMessage{
			Role:    string(msg.MessageType),
			Content: msg.Content,
		}
	}

	// Create the request
	request := ChatRequest{
		Messages:       chatMessages,
		UserID:         fmt.Sprintf("%d", userID),
		ConversationID: fmt.Sprintf("%d", conversationID),
		Context:        context,
	}

	// Marshal the request
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	url := fmt.Sprintf("%s/api/v1/process_messages", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("LLM service returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Print response for debugging
	fmt.Println("Response body:", string(bodyBytes))

	// Parse the response
	var chatResponse ChatResponse
	if err := json.Unmarshal(bodyBytes, &chatResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	fmt.Println("Chat response:", chatResponse)
	return &chatResponse, nil
}

// HealthCheck checks if the LLM service is healthy
func (c *LLMClient) HealthCheck(ctx context.Context) error {
	url := fmt.Sprintf("%s/api/v1/chat/health", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create health check request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send health check request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("LLM service health check failed with status %d", resp.StatusCode)
	}

	return nil
}
