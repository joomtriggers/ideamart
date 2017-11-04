<?php

namespace Ideamart\USSD\Ideamart\Receiver;

use Illuminate\Contracts\Foundation\Application;

class RequestValidator {

	protected $rules;

	public function __construct() {
	}

	public function validate(array $request,array $rules) {
        return $request;
	}
}
