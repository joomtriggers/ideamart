<?php

namespace Joomtriggers\Ideamart\SMS\Brokers;
use Joomtriggers\Ideamart\Contracts\MessageBrokerInterface;

/**
 * Class: MessageBroker
 *
 * @category Category
 * @package  Package
 * @author   Gnanakeethan Balasubramaniam <gnana@keethan.me>
 * @license  MIT http://opensource.org/licenses/MIT/
 * @link     http://link/
 *
 * @see MessageBrokerInterface
 */
class MessageBroker implements MessageBrokerInterface
{
    /**
     * Construct
     *
     * @return void
     */
    public function __construct()
    {

    }

    protected $message = "";

    public function setMessage($message)
    {
        $this->message = $message;

    }

    public function getMessage()
    {
        return $this->message;
    }

}

