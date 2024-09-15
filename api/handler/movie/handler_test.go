package moviehandler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/labstack/echo/v4"
	"org.idev.koala/backend/app"
	"org.idev.koala/backend/component/kafka"
)

// MockKafkaConsumer is a mock implementation of the kafka.Consumer interface
type MockKafkaConsumer struct {
	consumeFunc func(ctx context.Context, handler func([]byte) error) error
}

func (m *MockKafkaConsumer) Consume(ctx context.Context, handler func([]byte) error) error {
	return m.consumeFunc(ctx, handler)
}

func TestSendVoteMovieMessage(t *testing.T) {
	mockAppCtx := &app.AppContext{
		Ctx: context.TODO(),
		Config: &app.Config{
			KafkaHost: "localhost",
			KafkaPort: 9092,
		},
	}
	mockKafkaProducer, err := kafka.NewProducer(mockAppCtx.Config.KafkaHost, mockAppCtx.Config.KafkaPort)
	if err != nil {
		t.Fatalf("Failed to create mock Kafka producer: %v", err)
	}
	mockAppCtx.KafkaProducer = mockKafkaProducer

	handler := NewMovieHandler(mockAppCtx)

	err = handler.sendVoteMovieMessage(MovieVotedMessage{
		MovieID:  "123",
		UserID:   "user123",
		VoteType: "up",
	})
	if err != nil {
		for _, err := range err.(sarama.ProducerErrors) {
			t.Errorf("Failed to send vote movie message: %v", err)
		}
		t.Fatalf("Failed to send vote movie message: %v", err)
	}

	consumer, err := kafka.NewConsumer(mockAppCtx.Config.KafkaHost, mockAppCtx.Config.KafkaPort, "movie_voted")
	if err != nil {
		t.Fatalf("Failed to create mock Kafka consumer: %v", err)
	}

	resultChan := make(chan []byte)
	go consumer.Consume(context.TODO(), func(msg []byte) error {
		t.Log(string(msg))
		resultChan <- msg
		return nil
	})

	result := <-resultChan
	consumer.Close()

	msg := MovieVotedMessage{}
	err = json.Unmarshal(result, &msg)
	if err != nil {
		t.Fatalf("Failed to unmarshal message: %v", err)
	}

	if msg.MovieID != "123" || msg.UserID != "user123" || msg.VoteType != "up" {
		t.Fatalf("Expected message %q, got %q", "{\"movie_id\":\"123\",\"user_id\":\"user123\",\"vote_type\":\"up\"}", string(result))
	}
}

func TestStreamMovieVotes(t *testing.T) {
	// Create a new echo instance
	e := echo.New()

	// Create a mock AppContext
	mockAppCtx := &app.AppContext{
		Ctx: context.TODO(),
		Config: &app.Config{
			KafkaHost: "localhost",
			KafkaPort: 9092,
		},
	}
	mockKafkaProducer, err := kafka.NewProducer(mockAppCtx.Config.KafkaHost, mockAppCtx.Config.KafkaPort)
	if err != nil {
		t.Fatalf("Failed to create mock Kafka producer: %v", err)
	}
	mockAppCtx.KafkaProducer = mockKafkaProducer

	// Create a new MovieHandler with the mock AppContext
	handler := NewMovieHandler(mockAppCtx)

	// Create a new request
	req := httptest.NewRequest(http.MethodGet, "/movies/123/votes/stream", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/movies/:id/votes/stream")
	c.SetParamNames("id")
	c.SetParamValues("123")

	go handler.StreamMovieVotes()(c)

	time.Sleep(2 * time.Second)

	handler.sendVoteMovieMessage(MovieVotedMessage{
		MovieID:  "123",
		UserID:   "user123",
		VoteType: "up",
	})

	time.Sleep(5 * time.Second)

	expectedBody := "data:{\"movie_id\":\"123\",\"user_id\":\"user123\",\"vote_type\":\"up\"}\n\n"
	body := rec.Body.String()
	t.Logf("body: %s", body)
	if body != expectedBody {
		t.Fatalf("Expected body %q, got %q", expectedBody, body)
	}
}
