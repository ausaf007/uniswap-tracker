package handlers

import (
	"github.com/ausaf007/uniswap-tracker/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

//type PoolHandler struct {
//	service *services.TrackingService
//}
//
//func NewPoolHandler(service *services.TrackingService) *PoolHandler {
//	return &PoolHandler{service: service}
//}
//
//func (h *PoolHandler) PoolDataHandler(c *fiber.Ctx) error {
//	poolAddress := c.Params("pool_id")
//	block := c.Query("block", "latest")
//
//	poolData, err := h.service.GetPoolData(poolAddress, block)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
//	}
//
//	return c.JSON(poolData)
//}
//
//func (h *PoolHandler) HistoricPoolDataHandler(c *fiber.Ctx) error {
//	poolAddress := c.Params("pool_id")
//
//	poolData, err := h.service.GetHistoricPoolData(poolAddress)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
//	}
//
//	return c.JSON(poolData)
//}

type PoolHandler struct {
	service *services.TrackingService
}

func NewPoolHandler(service *services.TrackingService) *PoolHandler {
	return &PoolHandler{service: service}
}

func (h *PoolHandler) PoolDataHandler(c *fiber.Ctx) error {
	poolID, err := strconv.ParseUint(c.Params("pool_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid pool_id"})
	}
	block := c.Query("block", "latest")

	poolData, err := h.service.GetPoolData(uint(poolID), block)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(poolData)
}

func (h *PoolHandler) HistoricPoolDataHandler(c *fiber.Ctx) error {
	poolID, err := strconv.ParseUint(c.Params("pool_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid pool_id"})
	}

	poolData, err := h.service.GetHistoricPoolData(uint(poolID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(poolData)
}
