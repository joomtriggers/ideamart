<?php

namespace Ideamart\USSD\Ideamart\Sender;

use Ideamart\USSD\Ideamart\Sender\Contracts\FormatterInterface;
use Inqurtime\System\Session\Contracts\SessionHandler;

class RequestFormatter implements FormatterInterface {

	protected $ussdOperation;

	protected $session;

	protected $sessionId;

	public function __construct(SessionHandler $sessionHandler) {
		$this->session = $sessionHandler;
	}
	public function resolveJsonStream($message, $provider) {
		$this->ussdOperation = $this->session->getOperation();
		$this->sessionId = $this->session->getSessionId();
		$this->version = "1.0";
		$this->encoding = "440";
		$messageDetails = array(
			"message" => $message,
			"destinationAddress" => $this->session->getSourceAddress(),
		);
		$messageDetails = $this->mapToArray(
			[
				'deliveryStatusRequest',
				'binaryHeader',
				'version',
				'encoding',
				'ussdOperation',
				'sessionId',
			],
			$messageDetails
		);

		$applicationDetails = array(
			'applicationId' => $provider->app_id,
			'password' => $provider->password,
		);

		return json_encode($applicationDetails + $messageDetails);
	}

	private function mapToArray(array $array, array $details) {
		foreach ($array as $element) {
			if (isset($this->{$element})) {
				$details = array_merge($details, array($element => $this->{$element}));
			}
		}

		return $details;
	}
}