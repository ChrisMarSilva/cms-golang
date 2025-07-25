package adapters

import (
	"context"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/rinha-backend-2025/internal/dtos"
	"github.com/chrismarsilva/rinha-backend-2025/internal/utils"
	"github.com/redis/go-redis/v9"
)

type HealthCheckService struct {
	Config      *utils.Config
	RedisClient *redis.Client
	Healthy     atomic.Value
}

func NewHealthCheckService(config *utils.Config, redisClient *redis.Client) *HealthCheckService {
	h := &HealthCheckService{
		Config:      config,
		RedisClient: redisClient,
	}

	h.Healthy.Store(dtos.ProcessorStatusDto{Service: "default", URL: h.Config.UrlDefault})
	return h
}

func (h *HealthCheckService) Start() {
	go func() {
		ticker := time.NewTicker(6 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if h.acquireLock() {
				h.updateHealthyProcessor()
				h.releaseLock()
			} else {
				h.readHealthyFromRedis()
			}
		}
	}()
}

func (h *HealthCheckService) acquireLock() bool {
	// SetNX: Definido key para conter uma string value se key não existir.
	ok, _ := h.RedisClient.SetNX(context.Background(), "health_check_lock", "locked", 10*time.Second).Result()
	return ok
}

func (h *HealthCheckService) releaseLock() {
	// Del: Remove as chaves especificadas.
	h.RedisClient.Del(context.Background(), "health_check_lock")
}

func (h *HealthCheckService) updateHealthyProcessor() {
	if h.isHealthy(h.Config.UrlDefault+"/payments/service-health", "default") {
		h.Healthy.Store(dtos.ProcessorStatusDto{URL: h.Config.UrlDefault + "/payments", Service: "default"})
		h.saveHealthyToRedis("default")
		return
	}

	if h.isHealthy(h.Config.UrlFallback+"/payments/service-health", "fallback") {
		h.Healthy.Store(dtos.ProcessorStatusDto{URL: h.Config.UrlFallback + "/payments", Service: "fallback"})
		h.saveHealthyToRedis("fallback")
		return
	}

	h.Healthy.Store(dtos.ProcessorStatusDto{URL: "", Service: "out"})
	h.saveHealthyToRedis("out")
}

func (h *HealthCheckService) isHealthy(url, service string) bool {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error checking health for %s: %v", url, err)
		return false
	}
	defer resp.Body.Close()

	var response dtos.HealthCheckResponseDto
	//_ = json.NewDecoder(resp.Body).Decode(&response)
	_ = sonic.ConfigFastest.NewDecoder(resp.Body).Decode(&response)

	return !response.Failing && response.MinResponseTime <= 1000
}

func (h *HealthCheckService) saveHealthyToRedis(service string) {
	healthService := map[string]interface{}{"service": service, "timestamp": time.Now().Unix()}
	// HMSet: Define os campos especificados com seus respectivos valores no hash armazenado em key.
	h.RedisClient.HMSet(context.Background(), "healthy_processor_status", healthService)
}

func (h *HealthCheckService) readHealthyFromRedis() {
	healthService, _ := h.RedisClient.HGetAll(context.Background(), "healthy_processor_status").Result()

	switch healthService["service"] {
	case "default":
		h.Healthy.Store(dtos.ProcessorStatusDto{URL: h.Config.UrlDefault + "/payments", Service: "default"})
	case "fallback":
		h.Healthy.Store(dtos.ProcessorStatusDto{URL: h.Config.UrlFallback + "/payments", Service: "fallback"})
	default:
		h.Healthy.Store(dtos.ProcessorStatusDto{URL: "", Service: "out"})
	}
}

func (h *HealthCheckService) GetCurrent() dtos.ProcessorStatusDto {
	return h.Healthy.Load().(dtos.ProcessorStatusDto)
}
