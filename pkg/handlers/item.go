package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/xumoyun005/todo-app"
	"net/http"
	"strconv"
)

// @Summary Create todo item
// @Security ApiKeyAuth
// @Tags items
// @Description Create an item inside a list
// @ID create-item
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Param input body todo.TodoItem true "Item info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Update todo item
// @Security ApiKeyAuth
// @Tags items
// @Description Update item by item ID
// @ID update-item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param input body todo.UpdateItemInput true "Updated item info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"success"})
}

// @Summary Delete todo item
// @Security ApiKeyAuth
// @Tags items
// @Description Delete item by item ID
// @ID delete-item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "success",
	})
}

// @Summary Get all items from list
// @Security ApiKeyAuth
// @Tags items
// @Description Get all items by list ID
// @ID get-all-items
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Success 200 {array} todo.TodoItem
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id}/items [get]
func (h *Handler) getItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get item by ID
// @Security ApiKeyAuth
// @Tags items
// @Description Get single item by item ID
// @ID get-item-by-id
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} todo.TodoItem
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}
