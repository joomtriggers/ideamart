package sms;

type Configuration interface {
	Configure(map[string]string)
	GetConfigurations() map[string]string
	GetConfiguration(string) string
	SetConfiguration(string,string) string
	GetSecret()
	GetApplication()
	GetServer()


}
//<?php
//
//namespace Joomtriggers\Ideamart\Contracts;
//
//interface ConfigurationInterface {
//    public function configure(array $config);
//
//    public function getConfigurations();
//
//    public function getConfiguration($key);
//
//    public function setConfiguration($key,$value);
//
//    public function getSecret();
//
//    public function getApplication();
//
//    public function getServer();
//}
//
