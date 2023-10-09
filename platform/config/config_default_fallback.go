package config

func (c *DefaultConfig) GetStringDefault(name string, defaultValue string) (result string) {
	value, found := c.GetString(name)

	if found {
		result = value
	}
	return
}

func (c *DefaultConfig) GetIntDefault(name string, defaultValue int) (result int) {
	value, found := c.GetInt(name)

	if found {
		result = value
	}
	return
}

func (c *DefaultConfig) GetBoolDefault(name string, defaultValue bool) (result bool) {
	value, found := c.GetBool(name)

	if found {
		result = value
	}
	return
}

func (c *DefaultConfig) GetFloatDefault(name string, defaultValue float64) (result float64) {
	value, found := c.GetFloat(name)

	if found {
		result = value
	}
	return
}
