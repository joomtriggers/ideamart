<?php

namespace Ideamart\USSD\Ideamart\Exceptions;

class ServiceException extends \Exception {

	public $code;

	public $response;

	public $statusMessage;

	public function __construct($message, $code, $response = null) {
		parent::__construct($message);
		$this->statusMessage = $message;
		$this->code = $code;
		$this->response = $response;
	}

	public function getStatusCode() {
		return $this->code;
	}

	public function getStatusMessage() {
		return $this->statusMessage;
	}

	public function getRawResponse() {
		return $this->response;
	}

}
