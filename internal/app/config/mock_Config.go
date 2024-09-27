package config

type MockConfig struct {
	container map[string]any
}

func NewMockConfig() *MockConfig {
	return &MockConfig{
		container: make(map[string]any),
	}
}

func (c *MockConfig) GetString(key *Item[string]) string {
	return c.container[key.key].(string)
}

func (c *MockConfig) GetInt(key *Item[int]) int {
	return c.container[key.key].(int)
}

func (c *MockConfig) GetBool(key *Item[bool]) bool {
	return c.container[key.key].(bool)
}

func (c *MockConfig) GetFloat64(key *Item[float64]) float64 {
	return c.container[key.key].(float64)
}

func (c *MockConfig) SetString(key *Item[string], value string) {
	c.container[key.key] = value
}

func (c *MockConfig) SetInt(key *Item[int], value int) {
	c.container[key.key] = value
}

func (c *MockConfig) SetBool(key *Item[bool], value bool) {
	c.container[key.key] = value
}

func (c *MockConfig) Save() {
}
