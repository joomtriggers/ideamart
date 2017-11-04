package sms

type ConfigurationInterface interface {
	SetServer(string)
	SetApplication(string)
	SetPassword(string)
	GetPassword() string
	GetApplication() string
	GetServer() string
}
type Configuration struct {
	Server        string `json:"server"`
	Password      string `json:"password"`
	ApplicationId string `json:"applicationId"`
}

func (c *Configuration) SetServer(server string) *Configuration {
	c.Server = server;
	return c
}
func (c *Configuration) SetApplication(appId string) *Configuration {
	c.ApplicationId = appId;
	return c
}
func (c *Configuration) SetPassword(password string) *Configuration {
	c.Password = password;
	return c;
}
func (c *Configuration) GetServer() string {
	return c.Server
}
func (c *Configuration) GetApplication() string {
	return c.ApplicationId
}
func (c *Configuration) GetPassword() string {
	return c.Password
}
