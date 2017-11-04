<?php

namespace Joomtriggers\Ideamart\Contracts;

use Joomtriggers\Ideamart\SMS\Handler;

interface ReceiverInterface{
    public function receive($request,Handler $handler);
}


