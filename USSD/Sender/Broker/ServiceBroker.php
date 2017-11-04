<?php
namespace Ideamart\USSD\Ideamart\Sender\Broker;

class ServiceBroker {
	public $server_url;

	protected $app_id;

	protected $password;


	public function __construct($server, $app, $password) {
		$this->server_url = $server;
		$this->app_id = $app;
		$this->password = $password;
	}
	/**
	 * Magic Get Method to replace Native Access Methods
	 */
	public function __get($name) {
		return $this->{$name};
	}

	/**
	 * Magic Set Method to replace Native Access Methods
	 */
	public function __set($name, $value = null) {
		return $this->{$name} = $value;
	}
}
