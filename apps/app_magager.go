package apps

type AppManager struct {
	apps map[AppType]AppConfig
}

func NewAppManager() *AppManager {
	return &AppManager{
		apps: make(map[AppType]AppConfig),
	}
}

func (am *AppManager) EnableApp(
	appType AppType,
	url string,
) {
	am.apps[appType] = AppConfig{
		Type:    appType,
		Enabled: true,
		URL:     url,
	}
}

func (am *AppManager) DisableApp(
	appType AppType,
) {
	delete(am.apps, appType)
}

func (am *AppManager) IsEnabled(
	appType AppType,
) bool {
	config, exists := am.apps[appType]

	return exists && config.Enabled
}

func (am *AppManager) GetApp(
	appType AppType,
) (AppConfig, bool) {
	config, exists := am.apps[appType]

	return config, exists
}

func (am *AppManager) GetEnabledApps() []AppConfig {

	result := make([]AppConfig, 0)

	for _, config := range am.apps {
		if config.Enabled {
			result = append(result, config)
		}
	}

	return result
}
