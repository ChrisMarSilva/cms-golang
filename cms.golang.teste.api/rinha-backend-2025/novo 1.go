


,
func PaymentsHandler(w http.ResponseWriter, r *http.Request) {
	if err := database.Rdb.LPush(database.RedisCtx, "payments_pending", data).Err(); err != nil {


func PaymentsSummaryHandler(w http.ResponseWriter, r *http.Request) {
paymentsData, err := database.Rdb.HGetAll(database.RedisCtx, "payments").Result()
	resp := summarizePayments(paymentsData, from, to, useFilter)

		if payment.Processor == "DEFAULT" && payment.Status == "PROCESSED_DEFAULT" {
		} else if payment.Processor == "FALLBACK" && payment.Status == "PROCESSED_FALLBACK" {

			TotalAmount:   core.RoundedFloat(defaultSum),
type RoundedFloat float64
	TotalAmount   RoundedFloat `json:"totalAmount"`


func (r RoundedFloat) MarshalJSON() ([]byte, error) {
	rounded := math.Round(float64(r)*10) / 10
	return json.Marshal(rounded)
}


	mux.HandleFunc("/purge-payments", PurgePaymentsHandler)
func PurgePaymentsHandler(w http.ResponseWriter, r *http.Request) {
	err := database.Rdb.Del(database.RedisCtx, "payments").Err()




	instanceID := os.Getenv("INSTANCE_ID")
			acquired, err := database.Rdb.SetNX(database.RedisCtx, "rinha-leader-lock", instanceID, 10*time.Second).Result()


processorURLs := []string{
		os.Getenv("PROCESSOR_DEFAULT_URL"),
		os.Getenv("PROCESSOR_FALLBACK_URL"),
	}

	redisKeys := []string{"health:default", "health:fallback"}

for i, url := range processorURLs {
if err := database.Rdb.Set(redisCtx, redisKeys[i], data, 0).Err(); err != nil {


func RetrieveHealthStates(ctx context.Context) (*core.HealthManager, error) {

	defaultKey := "health:default"
	fallbackKey := "health:fallback"
	defaultVal, err := database.Rdb.Get(ctx, defaultKey).Result()
	fallbackVal, err := database.Rdb.Get(ctx, fallbackKey).Result()


func StartWorker() {
				res, err := database.Rdb.RPopLPush(context.Background(), "payments_pending", processingQueue).Result()

if err := processPayment(context.Background(), payment); err != nil {
database.Rdb.LPush(database.RedisCtx, "payments_pending", res)

	if health.DefaultProcessor.Failing || health.DefaultProcessor.MinResponseTime > health.FallBackProcessor.MinResponseTime+50 {
requestedAt := time.Now().UTC().Format(time.RFC3339Nano)

	err = database.Rdb.HSet(database.RedisCtx, "payments", payment.CorrelationID, paymentData).Err()


----

	s.app.Get("/payments-summary", s.handleGetSummary)
	s.app.Get("/internal/summary", s.handleGetInternalSummary)
	s.app.Post("/purge-payments", s.handlePurgeAllData)
	s.app.Post("/internal/purge", s.handleInternalPurge)

	Status        string          `json:"status"`
	CreatedAt     string          `json:"createdAt"`
	RequestedAt   string          `json:"requestedAt"`

	
import "github.com/shopspring/decimal"
	Amount        decimal.Decimal `json:"amount"`

		payments: make([]Payment, 0, 50000),
	snapshot := make([]Payment, len(s.payments))
	copy(snapshot, s.payments)

	func (q *Queue) retry(job Job) {
	if job.Retries < common.MaxRetries {
		job.Retries++

		delay := common.RetryDelay * time.Duration(1<<(job.Retries-1))
		time.AfterFunc(delay, func() {
			q.jobs <- job
		})
	} else {
		log.Printf(
			"Pagamento %s descartado após %d tentativas.",
			job.Payment.CorrelationID,
			common.MaxRetries,
		)
	}
}



--------------

payload, _ := json.Marshal(p)

		redisQueue.ClearStream(context.Background())
	err := q.redisClient.Del(ctx, PaymentStream).Err()

	TotalAmount   float64 `json:"totalAmount"`
	s.Default.TotalAmount = float64(int64(s.Default.TotalAmount*100+0.5)) / 100
		

	package internal

import (
	"fmt"
	"time"
)

func ParseDateTime(dateStr string) time.Time {
	formats := []string{
		"2006-01-02T15:04:05",
		"2006-01-02",
		time.RFC3339Nano,
		"2006-01-02",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z07:00",
	}
	for _, f := range formats {
		t, err := time.Parse(f, dateStr)
		if err == nil {
			return t
		}
	}
	fmt.Printf("Failed to parse date %s with formats %v\n", dateStr, formats)
	return time.Time{} // Return zero value if no format matches
}

--------------




	

--------------








------------------------------------------------------------------------------------------

func CreateRouter(mux *http.ServeMux, config Config) {

store := &store.Store{RedisClient: redisClient}
newProcessor := distributor.NewPaymentProcessor(config.Workers, store, healthCheckService)
handler := &Handler{paymentProcessor: newProcessor}



func (h *Handler) HandlePayments(w http.ResponseWriter, r *http.Request) {
Amount:        int64(math.Round(incoming.Amount * 100)),
	if err := h.paymentProcessor.Store.RedisClient.LPush(r.Context(), "payments:queue", payload).Err(); err != nil {


func (h *Handler) HandlePaymentsSummary(w http.ResponseWriter, r *http.Request) {
	from, to, err := parseTimeRange(fromStr, toStr)

	var payments []models.Payment
	if !from.IsZero() && !to.IsZero() {
		payments, err = h.paymentProcessor.Store.GetPaymentsByTime(r.Context(), from, to)
	} else {
		// If no time range, get all payments
		payments, err = h.paymentProcessor.Store.GetPaymentsByTime(r.Context(), time.Unix(0, 0), time.Now().UTC())
	}

summary := PaymentsToSummary(payments, from, to)
TotalAmount:   float64(summary.Default.TotalAmount) / 100.0,


func (h *Handler) HandlePurgePayments(w http.ResponseWriter, r *http.Request) {
err := h.paymentProcessor.Store.PurgeAllData(ctx)

	response := map[string]string{
		"status":  "success",
		"message": "Payment data purged successfully",
	}
	
	
Workers:  10,

package health

var (
	DEFAULT_PAYMENT_PROCESSOR_URL  = "http://payment-processor-default:8080/payments"
	FALLBACK_PAYMENT_PROCESSOR_URL = "http://payment-processor-fallback:8080/payments"
	DEFAULT_HEALTH_URL             = "http://payment-processor-default:8080/payments/service-health"
	FALLBACK_HEALTH_URL            = "http://payment-processor-fallback:8080/payments/service-health"
)



package internal

import (
	"fmt"
	"rinha-backend-arthur/internal/models"
	"time"
)

func buildSummary(payments []models.Payment) models.PaymentSummary {
	summary := models.PaymentSummary{
		Default:  models.Summary{},
		Fallback: models.Summary{},
	}

	for _, payment := range payments {
		switch payment.Service {
		case "default":
			summary.Default.TotalRequests++
			summary.Default.TotalAmount += payment.Amount
		case "fallback":
			summary.Fallback.TotalRequests++
			summary.Fallback.TotalAmount += payment.Amount
		}
	}

	// Round to 2 decimal places to match payment processor behavior
	// summary.Default.TotalAmount = math.Round(summary.Default.TotalAmount*100) / 100
	// summary.Fallback.TotalAmount = math.Round(summary.Fallback.TotalAmount*100) / 100

	return summary
}
func parseTimeRange(fromStr, toStr string) (from, to time.Time, err error) {
	if fromStr == "" && toStr == "" {
		return // Both zero values means no filtering
	}

	if fromStr == "" || toStr == "" {
		return // Return zero values to indicate no filtering
	}

	from, err = ParseFlexibleTime(fromStr)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid 'from' time format: %w", err)
	}

	to, err = ParseFlexibleTime(toStr)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid 'to' time format: %w", err)
	}

	if from.After(to) {
		return time.Time{}, time.Time{}, fmt.Errorf("'from' must be before or equal to 'to'")
	}

	return from, to, nil
}

func PaymentsToSummary(payments []models.Payment, from, to time.Time) models.PaymentSummary {
	validPayments := []models.Payment{}

	isTimeRangeSet := !from.IsZero() && !to.IsZero()

	for _, payment := range payments {
		include := true

		if isTimeRangeSet && (payment.RequestedAt.Before(from) || payment.RequestedAt.After(to)) {
			include = false
		}

		if include {
			validPayments = append(validPayments, payment)
		}
	}

	summary := buildSummary(validPayments)
	return summary
}

func ParseFlexibleTime(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, nil
	}

	// Formatos ISO 8601 UTC conforme especificação
	formats := []string{
		time.RFC3339,               // "2006-01-02T15:04:05Z07:00"
		"2006-01-02T15:04:05.000Z", // "2020-07-10T12:34:56.000Z"
		"2006-01-02T15:04:05Z",     // "2020-07-10T12:34:56Z"
		"2006-01-02T15:04:05",      // "2000-01-01T00:00:00" (without timezone)
		"2006-01-02",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			if t.Location() == time.UTC || format == "2006-01-02T15:04:05" || format == "2006-01-02" {
				return t.UTC(), nil
			}
			return t.UTC(), nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid ISO UTC date format '%s' (expected: 2020-07-10T12:34:56.000Z)", timeStr)
}




package store

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"rinha-backend-arthur/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Store struct {
	RedisClient *redis.Client
}

func (s *Store) StorePayment(ctx context.Context, payment models.Payment) error {
	paymentData := map[string]any{
		"correlationId":    payment.CorrelationId.String(),
		"amount":           payment.Amount,
		"paymentProcessor": payment.Service,
		"requestedAt":      payment.RequestedAt.Format(time.RFC3339Nano),
	}
	paymentJSON, err := json.Marshal(paymentData)
	if err != nil {
		return fmt.Errorf("failed to marshal payment data: %w", err)
	}
	// Use nanosecond precision for score
	score := float64(payment.RequestedAt.UnixNano())
	err = s.RedisClient.ZAdd(ctx, "payments", redis.Z{
		Score:  score,
		Member: paymentJSON,
	}).Err()
	if err != nil {
		return fmt.Errorf("failed to store payment: %w", err)
	}
	return nil
}

func (s *Store) GetPaymentsByTime(ctx context.Context, from, to time.Time) ([]models.Payment, error) {
	minScore := float64(from.UnixNano())
	maxScore := float64(to.UnixNano())
	results, err := s.RedisClient.ZRangeByScore(ctx, "payments", &redis.ZRangeBy{
		Min: fmt.Sprintf("%f", minScore),
		Max: fmt.Sprintf("%f", maxScore),
	}).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve payments: %w", err)
	}
	var payments []models.Payment
	for _, paymentDataJSON := range results {
		var paymentData map[string]interface{}
		if err := json.Unmarshal([]byte(paymentDataJSON), &paymentData); err != nil {
			continue // Skip malformed data
		}
		payment, err := s.parsePaymentFromData(paymentData)
		if err != nil {
			continue // Skip invalid data
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (s *Store) GetAllPayments(ctx context.Context) ([]models.Payment, error) {
	// Get all payments from the single hash
	paymentsData, err := s.RedisClient.HGetAll(ctx, "payments").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve payments: %w", err)
	}

	if len(paymentsData) == 0 {
		return []models.Payment{}, nil
	}

	var payments []models.Payment
	for _, paymentDataJSON := range paymentsData {
		var paymentData map[string]interface{}
		if err := json.Unmarshal([]byte(paymentDataJSON), &paymentData); err != nil {
			continue // Skip malformed data
		}

		payment, err := s.parsePaymentFromData(paymentData)
		if err != nil {
			continue // Skip invalid data
		}

		payments = append(payments, payment)
	}

	return payments, nil
}

func (s *Store) parsePaymentFromData(data map[string]interface{}) (models.Payment, error) {
	correlationIdStr, ok := data["correlationId"].(string)
	if !ok {
		return models.Payment{}, fmt.Errorf("invalid correlationId")
	}

	var amount int64
	switch amountVal := data["amount"].(type) {
	case float64:
		amount = int64(math.Round(amountVal)) // Convert existing float to int
	case json.Number:
		floatVal, _ := amountVal.Float64()
		amount = int64(math.Round(floatVal))
	case int64:
		amount = amountVal
	case int:
		amount = int64(amountVal)
	default:
		return models.Payment{}, fmt.Errorf("invalid amount type: %T", data["amount"])
	}

	service, ok := data["paymentProcessor"].(string)
	if !ok {
		return models.Payment{}, fmt.Errorf("invalid paymentProcessor")
	}

	requestedAtStr, ok := data["requestedAt"].(string)
	if !ok {
		return models.Payment{}, fmt.Errorf("invalid requestedAt")
	}

	correlationId, err := uuid.Parse(correlationIdStr)
	if err != nil {
		return models.Payment{}, fmt.Errorf("failed to parse correlationId: %w", err)
	}

	requestedAt, err := time.Parse(time.RFC3339Nano, requestedAtStr)
	if err != nil {
		return models.Payment{}, fmt.Errorf("failed to parse requestedAt: %w", err)
	}

	payment := models.Payment{
		PaymentRequest: models.PaymentRequest{
			CorrelationId: correlationId,
			Amount:        amount,
			RequestedAt:   requestedAt.UTC(),
		},
		Service: service,
	}

	return payment, nil
}

func (s *Store) PurgeAllData(ctx context.Context) error {
	err := s.RedisClient.Del(ctx, "payments").Err()
	if err != nil {
		return fmt.Errorf("failed to delete payments: %w", err)
	}

	keys, err := s.RedisClient.Keys(ctx, "payments:processing:*").Result()
	if err == nil && len(keys) > 0 {
		s.RedisClient.Del(ctx, keys...)
	}

	return nil
}


package distributor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"rinha-backend-arthur/internal/health"
	"rinha-backend-arthur/internal/models"
	"rinha-backend-arthur/internal/store"
	"time"

	"github.com/google/uuid"
)

type PaymentProcessor struct {
	Store   *store.Store
	workers int
	client  *http.Client
	health  *health.HealthCheckService
}

func NewPaymentProcessor(workers int, store *store.Store, healthCheckServie *health.HealthCheckService) *PaymentProcessor {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 5,
			IdleConnTimeout:     30 * time.Second,
			DisableKeepAlives:   false,
		},
	}

	processor := &PaymentProcessor{
		workers: workers,
		Store:   store,
		client:  httpClient,
		health:  healthCheckServie,
	}

	// Start health check with ticker
	go processor.health.StartHealthCheckLoop()

	for i := range workers {
		go processor.distributePayment(i)
	}

	return processor
}

func (p *PaymentProcessor) distributePayment(workerNum int) {

	if p.Store == nil || p.Store.RedisClient == nil {
		// log.Fatal("PaymentProcessor store or RedisClient is nil!")
	}

	ctx := context.Background()
	processingQueue := fmt.Sprintf("payments:processing:%d", workerNum)

	for {
		result, err := p.Store.RedisClient.RPopLPush(ctx, "payments:queue", processingQueue).Result()
		if err != nil {
			if err.Error() == "redis: nil" {
				time.Sleep(1 * time.Second)
				continue
			}
			// log.Printf("[Worker %d] Redis error: %v", workerNum, err)
			time.Sleep(1 * time.Second)
			continue
		}

		var payment models.PaymentRequest
		if err := json.Unmarshal([]byte(result), &payment); err != nil {
			fmt.Printf("[Worker %v-%d] Failed to unmarshal payment: %v\n", workerNum, workerNum, err)
			p.Store.RedisClient.LRem(ctx, processingQueue, 1, result)
			continue
		}

		if err := p.ProcessPayments(payment); err != nil {
			fmt.Printf("[Worker %v] Failed to process payment: %v\n", workerNum, err)
			p.Store.RedisClient.LPush(ctx, "payments:queue", result) // Requeue the payment
		} else {
			// Successfully processed - remove from processing queue
			p.Store.RedisClient.LRem(ctx, processingQueue, 1, result)
		}

	}
}

func (p *PaymentProcessor) ProcessPayments(paymentRequest models.PaymentRequest) error {
	// evita que o health checker mude no meio
	currentProcessor := p.health.HealthyProcessor
	if currentProcessor == nil {
		return fmt.Errorf("no healthy processor available")
	}

	paymentRequestForProcessor := struct {
		CorrelationId uuid.UUID `json:"correlationId"`
		Amount        float64   `json:"amount"`
		RequestedAt   time.Time `json:"requestedAt"`
	}{
		CorrelationId: paymentRequest.CorrelationId,
		Amount:        float64(paymentRequest.Amount) / 100.0,
		RequestedAt:   paymentRequest.RequestedAt,
	}

	requestBody, err := json.Marshal(paymentRequestForProcessor)
	if err != nil {
		return err
	}

	resp, err := p.client.Post(currentProcessor.URL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("failed to send payment request to processor %s: %w", currentProcessor.Service, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("error sending to payment processor %s: status %d", currentProcessor.Service, resp.StatusCode)
	}

	processedPayment := models.Payment{
		PaymentRequest: paymentRequest,
		Service:        currentProcessor.Service, // Use captured processor
	}

	err = p.Store.StorePayment(context.Background(), processedPayment)
	if err != nil {
		// This is critical - payment was accepted by processor but we failed to save
		// Log as error but don't return error to avoid reprocessing
		// log.Printf("CRITICAL: Payment accepted by processor but failed to save in Redis: %v", err)
	}

	return nil
}


-*---------



/internal/store/cache

type UserStore struct {
	rdb *redis.Client
}


	addr:    env.GetString("REDIS_ADDR", "localhost:6379"),
			pw:      env.GetString("REDIS_PW", ""),
			db:      env.GetInt("REDIS_DB", 0),
			enabled: env.GetBool("REDIS_ENABLED", false),
		env: env.GetString("ENV", "development"),
		
	
	func (s *UserStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	cacheKey := fmt.Sprintf("user-%d", userID)

	data, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var user store.User
	if data != "" {
		err := json.Unmarshal([]byte(data), &user)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}


func (s *UserStore) Set(ctx context.Context, user *store.User) error {
	cacheKey := fmt.Sprintf("user-%d", user.ID)

	json, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return s.rdb.SetEX(ctx, cacheKey, json, UserExpTime).Err()
}


func (s *UserStore) Delete(ctx context.Context, userID int64) {
	cacheKey := fmt.Sprintf("user-%d", userID)
	s.rdb.Del(ctx, cacheKey)
}

---------------------------------------------------------------

func IsWithInRange(checked, from, to time.Time) bool {
	return (checked.Equal(from) || checked.After(from)) && (checked.Equal(to) || checked.Before(to))
}


func (a *PaymentProcessorAdapter) EnableHealthCheck(should string) {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			resDefault, err := a.healthCheckEndpoint(a.defaultUrl + "/payments/service-health")
			resFallback, err := a.healthCheckEndpoint(a.fallbackUrl + "/payments/service-health")
			reqbody := HealthCheckStatus{ resDefault, resFallback, }
			rawBody, err := sonic.Marshal(reqbody)
			if a.db.Set(context.Background(), HealthCheckKey, rawBody, 0).Err() != nil {
		}
	}()
}

func (a *PaymentProcessorAdapter) StartWorkers() {
	for range a.workers {
		go a.retryWorkers()
	}

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()

		for range ticker.C {
			res := a.db.Get(context.Background(), HealthCheckKey)x
			resBody, err := res.Result()
			var healthCheckStatus *HealthCheckStatus
			err := sonic.ConfigFastest.Unmarshal([]byte(resBody), &healthCheckStatus);
	
			a.mu.Lock()
			a.healthStatus = healthCheckStatus
			a.mu.Unlock()
		}
	}()
}

func (a *PaymentProcessorAdapter) retryWorkers() {
	for payment := range a.retryQueue {
		time.Sleep(time.Millisecond * 500)
		a.Process(payment)
	}
}

---------------------------------------------------------------


func RunHealthCheckWorker() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		newStatus, err := fetchAndProcessHealthChecks()

		HealthCache.Mu.Lock()
		HealthCache.Status = newStatus
		HealthCache.BestPaymentProcessorUrl, HealthCache.Err = decideBestProcessor(newStatus)
		HealthCache.Mu.Unlock()

	}
}

func decideBestProcessor(m fiber.Map) (string, error) {
	var payload types.HealthStatusPayload

	jsonBytes, err := json.Marshal(m)
	json.Unmarshal(jsonBytes, &payload)

	if payload.Default.IsFailing && payload.Fallback.IsFailing {

	if payload.Default.IsFailing || payload.Default.MinResponseTime > payload.Fallback.MinResponseTime+50 {
		return config.PaymentProcessorUrlFallback, nil
	}
	return config.PaymentProcessorUrlDefault, nil
}

func fetchAndProcessHealthChecks() (fiber.Map, error) {
	var wg sync.WaitGroup
	defaultChan := make(chan types.HealthResult, 1)
	fallbackChan := make(chan types.HealthResult, 1)

	wg.Add(2)
	go fetchHealth(config.PaymentProcessorUrlDefault+"/payments/service-health", &wg, defaultChan)
	go fetchHealth(config.PaymentProcessorUrlFallback+"/payments/service-health", &wg, fallbackChan)

	wg.Wait()
	close(defaultChan)
	close(fallbackChan)

	defaultResult := <-defaultChan
	fallbackResult := <-fallbackChan

	if defaultResult.Err != nil {
		return nil, defaultResult.Err
	}
	if fallbackResult.Err != nil {
		return nil, fallbackResult.Err
	}

	response := fiber.Map{
		"default":          defaultResult.Response,
		"fallback":         fallbackResult.Response,
		"last_checked_utc": time.Now().UTC(),
	}

	return response, nil
}


	err = database.RedisClient.HSet(database.RedisCtx, config.RedisPaymentsKey, processedPayment.CorrelationID, paymentData).Err()

	func StartWorker() {
	for i := 0; i < concurrency; i++ {
		go func(workerNum int) {
			for {
				res, err := database.RedisClient.RPopLPush(context.Background(), config.RedisQueueKey, processingQueue).Result()
				var payment types.PaymentRequest
				if err := json.Unmarshal([]byte(res), &payment); err != nil {
				if err := ProccessPayments(payment); err != nil {
				database.RedisClient.LPush(database.RedisCtx, "payments_pending", res)
			
		}(i)
	}
}
func (a *PaymentProcessorAdapter) innerProcess(payment model.PaymentRequestProcessor) error {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if !a.healthStatus.Default.Failing && a.healthStatus.Default.MinResponseTime < 300 {
		if err := a.sendPayment(payment, a.defaultUrl+"/payments", 200*time.Millisecond); err == nil {
			return nil
		}
	}
	if !a.healthStatus.Fallback.Failing && a.healthStatus.Fallback.MinResponseTime < 100 {
		if err := a.sendPayment(payment, a.fallbackUrl+"/payments", 100*time.Millisecond); err == nil {
			return nil
		}
	}

	return model.ErrUnavailableProcessor
}


	_ = sonic.ConfigFastest.NewDecoder(res.Body).Decode(&hc)

	func (a *PaymentProcessorAdapter) EnableHealthCheck(should string) {
	go func() {
		

		for range ticker.C {
			def, _ := a.healthCheck(a.defaultUrl + "/payments/service-health")
			fall, _ := a.healthCheck(a.fallbackUrl + "/payments/service-health")

			status := &model.HealthCheckStatus{Default: def, Fallback: fall}
			a.mu.Lock()
			a.healthStatus = status
			a.mu.Unlock()

			raw, _ := sonic.Marshal(status)
			a.db.Set(context.Background(), model.HealthCheckKey, raw, 30*time.Second)
		}
	}()
}

---------------------------------------------------------------

"DEFAULT"
"FALLBACK"

	requestedAt := time.Now().UTC().Format(time.RFC3339Nano)


func healthChecker() {
	processorURLs := []string{os.Getenv("PROCESSOR_DEFAULT_URL"), os.Getenv("PROCESSOR_FALLBACK_URL"),}
	redisKeys := []string{"health:default", "health:fallback"}

	for {
		<-ticker.C
		for i, url := range processorURLs {
			go func(i int, url string) {
				resp, err := http.Get(url + "/payments/service-health")
		
			}(i, url)
		}
	}
}


func RetrieveHealthStates(ctx context.Context) (*core.HealthManager, error) {

	defaultVal, err := database.Rdb.Get(ctx, defaultKey).Result()
	fallbackVal, err := database.Rdb.Get(ctx, fallbackKey).Result()

	var defaultHealth core.HealthResponse
	var fallbackHealth core.HealthResponse
	if err := json.Unmarshal([]byte(defaultVal), &defaultHealth); err != nil {x
	if err := json.Unmarshal([]byte(fallbackVal), &fallbackHealth); err != nil {
	return &core.HealthManager{
		DefaultProcessor:  defaultHealth,
		FallBackProcessor: fallbackHealth,
	}, nil
}

type PaymentRedisRepository struct {
	client *redis.Client
	ctx    context.Context
}


func (r *PaymentRedisRepository) Get(correlationID string) (entities.Payment, bool) {
	key := fmt.Sprintf("payment:%s", correlationID)

	data, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return entities.Payment{}, false
	}
	if err != nil {
		return entities.Payment{}, false
	}

	var payment entities.Payment
	if err := json.Unmarshal([]byte(data), &payment); err != nil {
		return entities.Payment{}, false
	}
	return payment, true
}

func (r *PaymentRedisRepository) Purge() {
	if err := r.client.FlushDB(r.ctx).Err(); err != nil {
		fmt.Println("Error purging Redis database:", err)
	}
}



---------------------------------------------------------------