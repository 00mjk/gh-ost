/*
   Copyright 2016 GitHub Inc.
	 See https://github.com/github/gh-ost/blob/master/LICENSE
*/

package mysql

import (
	"fmt"
	"net"
)

// ConnectionConfig is the minimal configuration required to connect to a MySQL server
type ConnectionConfig struct {
	Key        InstanceKey
	User       string
	Password   string
	ImpliedKey *InstanceKey
}

func NewConnectionConfig() *ConnectionConfig {
	config := &ConnectionConfig{
		Key: InstanceKey{},
	}
	config.ImpliedKey = &config.Key
	return config
}

func (this *ConnectionConfig) Duplicate() *ConnectionConfig {
	config := &ConnectionConfig{
		Key: InstanceKey{
			Hostname: this.Key.Hostname,
			Port:     this.Key.Port,
		},
		User:     this.User,
		Password: this.Password,
	}
	config.ImpliedKey = &config.Key
	return config
}

func (this *ConnectionConfig) String() string {
	return fmt.Sprintf("%s, user=%s", this.Key.DisplayString(), this.User)
}

func (this *ConnectionConfig) Equals(other *ConnectionConfig) bool {
	return this.Key.Equals(&other.Key) || this.ImpliedKey.Equals(other.ImpliedKey)
}

func (this *ConnectionConfig) GetDBUri(databaseName string) string {
	var ip = net.ParseIP(this.Key.Hostname)
	if (ip != nil) && (ip.To4() == nil) {
		// Wrap IPv6 literals in square brackets
		return fmt.Sprintf("%s:%s@tcp([%s]:%d)/%s", this.User, this.Password, this.Key.Hostname, this.Key.Port, databaseName)
	} else {
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", this.User, this.Password, this.Key.Hostname, this.Key.Port, databaseName)
	}
}
