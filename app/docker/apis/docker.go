package apis

import (
	"docker-api/app/docker/service"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ListContainer 容器列表接口
// @Summary Docker 容器列表接口
// @Description 查询容器列表
// @Tags 容器操作
// @Accept application/json
// @Produce application/json
// @Param all query string false "是否查询全部容器"
// @Param size query string false "是否容器显示大小"
// @Param latest query string false "是否仅显示最新创建容器"
// @Param quiet query string false "是否只显示容器id"
// @Param limit query int false  "显示数量"
// @Success 200 {object} gin.H
// @Router /api/v1/app/dockers [get]
func ListContainer(c *gin.Context) {
	option := types.ContainerListOptions{}
	if c.Query("all") == "true" {
		option.All = true
	}
	if c.Query("size") == "true" {
		option.Size = true
	}
	if c.Query("latest") == "true" {
		option.Latest = true
	}
	if c.Query("quiet") == "true" {
		option.Quiet = true
	}
	if c.Query("limit") != "" {
		if limit, err := strconv.ParseInt(c.Query("limit"), 10, 64); err == nil {
			option.Limit = int(limit)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "result": map[string]interface{}{"containers": service.List(option)}})
}

// StopContainer 停止容器接口
// @Summary Docker 停止容器接口
// @Description 停止容器
// @Tags 容器操作
// @Accept application/json
// @Produce application/json
// @Param containerId path string true "容器id"
// @Success 200 {object} gin.H
// @Router /api/v1/app/docker/{containerId}/stop [post]
func StopContainer(c *gin.Context) {
	containerId := c.Param("containerId")
	service.Stop(containerId)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// RestartContainer 重启容器接口
// @Summary Docker 重启容器接口
// @Description 重启容器
// @Tags 容器操作
// @Accept application/json
// @Produce application/json
// @Param containerId path string true "容器id"
// @Success 200 {object} gin.H
// @Router /api/v1/app/docker/{containerId}/restart [post]
func RestartContainer(c *gin.Context) {
	containerId := c.Param("containerId")
	service.Restart(containerId)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
