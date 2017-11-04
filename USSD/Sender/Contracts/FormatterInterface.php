<?php

namespace Ideamart\USSD\Ideamart\Sender\Contracts;

interface FormatterInterface {

	public function resolveJsonStream($product, $provider);
}