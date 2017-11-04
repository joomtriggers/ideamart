<?php

namespace Ideamart\USSD\Ideamart\Sender;

use Ideamart\USSD\Ideamart\Sender\Broker\ServiceBroker;
use Ideamart\USSD\Ideamart\Sender\Contracts\FormatterInterface;
use Joomtriggers\Ideamart\Core;

class Sender
{

    use Core;

    protected $app;

    protected $serviceProvider;

    protected $formatter;

    public function __construct(
        ServiceBroker $broker,
        FormatterInterface $formatterInterface
    )
    {
        $this->serviceProvider = $broker;
        $this->formatter = $formatterInterface;
    }

    public function send($product)
    {
        $product = $this->formatter->resolveJsonStream($product, $this->serviceProvider);

        return $this->sendRequest($product, $this->serviceProvider->server_url);
    }
}
