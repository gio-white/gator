package config

func (c *Config) SetUser(user string) error {
    c.CurrentUserName = user
    
    return Write(*c)
}