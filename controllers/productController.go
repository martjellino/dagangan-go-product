package controller

import (
	"go-product/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func ReadProducts(c *gin.Context) {
	var product []models.Product
	models.DB.Find(&product)
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})

}

func ReadProductByID(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "It is failed. The data is not correct...",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "the product can't be updated...",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "it's already updated!",

		"Name":     product.Name,
		"Merchant": product.Merchant,
		"Desc":     product.Desc,
		"Stock":    product.Stock,
		"Price":    product.Price,
	})
}

func RemoveProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.Debug().Where("id = ?", id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "It is failed. The data is not correct...",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	if err := models.DB.Where("id = ?", id).Delete(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "the product can't be deleted...",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "it's already deleted!",
	})
}

func MerchantFilter(c *gin.Context) {
	merchantFilter := c.Query("merchant")

	var product []models.Product
	var count int64

	models.DB.Model(&models.Product{}).Where("merchant = ?", merchantFilter).Count(&count)
	models.DB.Where("merchant = ?", merchantFilter).Find(&product)

	response := make(map[string]interface{})
	response["totalData"] = count
	response["data"] = product

	c.JSON(http.StatusOK, response)
}

func ReadProductByPage(c *gin.Context) {
	pageSize := 3 // 3 product per page
	pageNum, _ := strconv.Atoi(c.Param("pageNum"))
	offset := (pageNum - 1) * pageSize

	var product []models.Product
	models.DB.Offset(offset).Limit(pageSize).Find(&product)
	c.JSON(http.StatusOK, product)
}
