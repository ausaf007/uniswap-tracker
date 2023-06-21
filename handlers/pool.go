package handlers

import (
	"github.com/ausaf007/uniswap-tracker/services"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type PoolHandler struct {
	service *services.TrackingService
}

type PoolDataResponse struct {
	Token0Balance string `json:"token0Balance"`
	Token1Balance string `json:"token1Balance"`
	Tick          string `json:"tick"`
}

type HistoricPoolDataResponse struct {
	Token0Balance string `json:"token0Balance"`
	Token0Delta   string `json:"token0Delta"`
	Token1Balance string `json:"token1Balance"`
	Token1Delta   string `json:"token1Delta"`
	BlockNumber   int64  `json:"blockNumber"`
}

type HistoricPoolDataResponseSlice []HistoricPoolDataResponse

func NewPoolHandler(service *services.TrackingService) *PoolHandler {
	return &PoolHandler{service: service}
}

func (h *PoolHandler) PoolDataHandler(c *fiber.Ctx) error {
	poolID, err := strconv.ParseUint(c.Params("pool_id"), 10, 64)
	if err != nil {
		log.Error("Invalid Pool ID: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid pool_id"})
	}
	block := c.Query("block", "latest")

	poolData, err := h.service.GetPoolData(uint(poolID), block)
	if err != nil {
		log.Error("Unable to fetch Pool Data: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	responseData := PoolDataResponse{
		Token0Balance: poolData.Token0Balance,
		Token1Balance: poolData.Token1Balance,
		Tick:          poolData.Tick,
	}

	return c.JSON(responseData)
}

func (h *PoolHandler) HistoricPoolDataHandler(c *fiber.Ctx) error {
	poolID, err := strconv.ParseUint(c.Params("pool_id"), 10, 64)
	if err != nil {
		log.Error("Invalid Pool ID: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid pool_id"})
	}

	poolData, err := h.service.GetHistoricPoolData(uint(poolID))
	if err != nil {
		log.Error("Unable to fetch Historic Pool Data: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var responseData HistoricPoolDataResponseSlice
	for _, pd := range poolData {
		responseData = append(responseData, HistoricPoolDataResponse{
			Token0Balance: pd.Token0Balance,
			Token0Delta:   pd.Token0Delta,
			Token1Balance: pd.Token1Balance,
			Token1Delta:   pd.Token1Delta,
			BlockNumber:   pd.BlockNumber,
		})
	}

	return c.JSON(responseData)
}

func (h *PoolHandler) PoolMappingHandler(c *fiber.Ctx) error {
	poolMap, err := h.service.GetPoolMapping()
	if err != nil {
		log.Error("Unable to fetch pool table data: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(poolMap)
}
