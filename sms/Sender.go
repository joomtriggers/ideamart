<?php

namespace Joomtriggers\Ideamart\SMS;

use Joomtriggers\Ideamart\Contracts\AddressBrokerInterface;
use Joomtriggers\Ideamart\Contracts\ConfigurationInterface;
use Joomtriggers\Ideamart\Contracts\MessageBrokerInterface;
use Joomtriggers\Ideamart\Contracts\SenderInterface;
use Joomtriggers\Ideamart\Core;

class Sender implements SenderInterface
{
    use Core;
    protected $response = [];

    public function send(
        MessageBrokerInterface $messageBrokerInterface,
        AddressBrokerInterface $addressBrokerInterface,
        ConfigurationInterface $configurationInterface)
    {
        //"message":"Hello" "destinationAddresses":["tel:94777123456"], "password":"password", "applicationId":"APP_999999"
        $this->response['message'] = $messageBrokerInterface->getMessage();
        $this->response['destinationAddresses'] = $addressBrokerInterface->getSubscribers();
        $this->response['password'] = $configurationInterface->getSecret();
        $this->response['applicationId'] = $configurationInterface->getApplication();
        return $this->sendRequest(json_encode($this->response), $configurationInterface->getServer());
    }
}