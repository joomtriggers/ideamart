<?php

namespace Joomtriggers\Ideamart\Contracts;

interface MessageBrokerInterface {

    public function setMessage($message);
    public function getMessage();
}


