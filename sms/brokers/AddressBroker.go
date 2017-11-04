<?php

namespace Joomtriggers\Ideamart\SMS\Brokers;

use Joomtriggers\Ideamart\Contracts\AddressBrokerInterface;

class AddressBroker implements AddressBrokerInterface
{

    protected $numbers;

    /**
     * Constructor
     *
     * @return void
     */
    public function __construct()
    {
    }

    /**
     * Return All Current Subscriber List
     *
     * @return array
     */
    public function getSubscribers()
    {
        return $this->numbers;
    }

    /**
     * Adding Subscriber
     *
     * @param string $number Number of Subscribers
     *
     * @return void
     */
    public function addSubscriber($number)
    {
        $this->numbers[] = $number;
        $this->numbers = array_unique(array_filter($this->numbers));
    }

    /**
     * Remove Certain Subscribers
     *
     * @return void
     */
    public function removeSubscriber($number)
    {
        if (($key = array_search($number, $this->numbers)) !== false) {
            unset($this->numbers[$key]);
        }
    }


}
