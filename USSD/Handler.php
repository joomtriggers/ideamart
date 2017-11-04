<?php

/**
 * Class:Handler
 *
 * PHP Version: 5.6
 *
 * @category Category
 * @package  Package
 * @author   Gnanakeethan Balasubramaniam <gnana@keethan.me>
 * @license  MIT http://opensource.org/licenses/MIT/
 * @link     http://link/
 */
namespace Ideamart\USSD\Ideamart;



use Joomtriggers\Ideamart\Contracts\USSD\ProducerInterface;
use Joomtriggers\Ideamart\Contracts\USSD\ReceiverInterface;
use Joomtriggers\Ideamart\Contracts\USSD\SenderInterface;
use Joomtriggers\Ideamart\Contracts\USSD\SessionHandlerInterface;

/**
 * Class:Handler
 *
 * @category Category
 * @package  Package
 * @author   Gnanakeethan Balasubramaniam <gnana@keethan.me>
 * @license  MIT http://opensource.org/licenses/MIT/
 * @link     http://link/
 */
class Handler
{
    /**
     * @var array
     */
    protected $request = [];
    /**
     */
    protected $sessionHandler;
    /**
     */
    protected $receiver;
    /**
     */
    protected $sender;
    /**
     */
    private $producer;
    /**
     * @var
     */
    private $product;
    /**
     * @var
     */
    private $appRepo;

    /**
     * __construct
     *
     * @param SessionHandlerInterface $sessionHandler Session Handling object
     * @param ReceiverInterface       $receiver       Receive handler
     * @param SenderInterface         $sender         Send Handler
     * @param ProducerInterface       $producer       Menu Manager
     */
    public function __construct(
        SessionHandlerInterface $sessionHandler,
        ReceiverInterface $receiver,
        SenderInterface $sender,
        ProducerInterface $producer
    )
    {
        $this->sessionHandler = $sessionHandler;
        $this->sender = $sender;
        $this->receiver = $receiver;
        $this->producer = $producer;
    }

    /**
     * @param array $request
     *
     * @return $this
     */
    public function receive(array $request)
    {
        $request = $this->receiver->receive($request);
        $this->product = $this->producer
            ->setRequest($request)
            ->translateMessage()
            ->makeMessage();
        return $this;
    }

    /**
     * @param $option
     */
    public function getOption($option)
    {
    }

    /**
     * @param $option
     */
    public function setOption($option)
    {
    }

    /**
     * @param $custom_message
     */
    public function setMessage($custom_message)
    {
    }

    /**
     * @param $server_details
     */
    public function setConfiguration($server_details)
    {
    }

    /**
     * @param $CONSTANT
     */
    public function setResponse($CONSTANT)
    {
    }

    /**
     *
     */
    public function makeResponse()
    {
    }
}
