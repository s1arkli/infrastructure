package deploy

import (
	_ "github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
	_ "github.com/redis/go-redis/v9"
	_ "github.com/segmentio/kafka-go"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/viper"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "go.uber.org/zap"
	_ "gorm.io/gen"
	_ "gorm.io/gorm"

	_ "github.com/prometheus/client_golang/prometheus"
)

//swag程序，扫描项目生成doc文档
//go install github.com/swaggo/swag/cmd/swag@latest
