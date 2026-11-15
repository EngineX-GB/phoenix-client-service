package config

import "os"

func GetListOfConfigs() map[string]string {
	var configMap = make(map[string]string)
	configMap["ENV"] = GetEnvironment()
	configMap["DATASOURCE_HOSTNAME"] = GetDataSourceHostname()
	configMap["DATASOURCE_USERNAME"] = GetDataSourceUsername()
	configMap["DATASOURCE_NAME"] = GetDataSourceName()
	configMap["MARKET_CONNECTOR_ENDPOINT"] = GetMarketConnectorEndpoint()
	configMap["INGESTION_SERVICE_ENDPOINT"] = GetIngestionServiceEndpoint()
	return configMap
}

func getEnvProperty(envVariable string, defaultValue string) string {
	env := os.Getenv(envVariable)
	if env == "" {
		return defaultValue
	}
	return env
}

func GetEnvironment() string {
	return getEnvProperty("ENV", "PROD")
}

func GetDataSourceUsername() string {
	return getEnvProperty("DATASOURCE_USERNAME", "root")
}

func GetDataSourcePassword() string {
	return getEnvProperty("DATASOURCE_PASSWORD", "root")
}

func GetDataSourceName() string {
	return getEnvProperty("DATASOURCE_NAME", "db_phoenix")
}

func GetDataSourceHostname() string {
	return getEnvProperty("DATASOURCE_HOST", "localhost:3306")
}

func GetMarketConnectorEndpoint() string {
	return getEnvProperty("MARKET_CONNECTOR_ENDPOINT", "http://localhost:8001")
}

func GetIngestionServiceEndpoint() string {
	return getEnvProperty("INGESTION_SERVICE_ENDPOINT", "http://localhost:8000")
}

func GetServicePort() string {
	return getEnvProperty("SERVICE_PORT", "8081")
}
