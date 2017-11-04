<?php

/**
 * Class: Handler
 *
 * PHP Version 5.6
 *
 * @category SMS
 * @package  Ideamart
 * @author   Gnanakeethan Balasubramaniam <gnana@keethan.me>
 * @license  MIT http://opensource.org/licenses/MIT/
 * @link     http://github.com/joomtriggers/ideamart-php/
 */

namespace Joomtriggers\Ideamart\SMS;


use Joomtriggers\Ideamart\Contracts\AddressBrokerInterface;
use Joomtriggers\Ideamart\Contracts\MessageBrokerInterface;
use Joomtriggers\Ideamart\Contracts\ReceiverInterface;
use Joomtriggers\Ideamart\Contracts\ServiceBrokerInterface;
use Joomtriggers\Ideamart\Contracts\SenderInterface;
use Joomtriggers\Ideamart\Contracts\ConfigurationInterface;

/**
 * Class: Handler
 *
 * @category SMS
 * @package  Ideamart
 * @author   Gnanakeethan Balasubramaniam <gnana@keethan.me>
 * @license  MIT http://opensource.org/licenses/MIT/
 * @link     http://github.com/joomtriggers/ideamart-php/
 */
class Handler
{

    private $_mode = "sending";

    public function __construct(
        AddressBrokerInterface $addressBroker,
        MessageBrokerInterface $messageBroker,
        ConfigurationInterface $configurationBroker,
        SenderInterface $sender,
        ReceiverInterface $receiver
    ) {
        $this->addressBroker = $addressBroker;
        $this->messageBroker = $messageBroker;
        $this->configurationBroker = $configurationBroker;
        $this->sender = $sender;
        $this->receiver = $receiver;
    }


    /**
     * Add Subscriber
     *
     * @param $number String Number of Subscriber
     *
     * @return Handler
     */
    public function addSubscriber($number)
    {
        $this->addressBroker->addSubscriber($number);

        return $this;
    }

    /**
     * Remove Subscriber
     *
     * @return Handler
     */
    public function removeSubscriber($number)
    {
        $this->addressBroker->removeSubscriber($number);

        return $this;
    }

    /**
     * Getting Subscribers
     *
     * @return AddressBrokerInterface
     */
    public function getSubscribers()
    {
        return $this->addressBroker->getSubscribers();
    }


    /**
     * Setting the message to send
     *
     * @return Handler
     */
    public function setMessage($message)
    {
        $this->messageBroker->setMessage($message);

        return $this;
    }

    /**
     * Get the message
     *
     * @return MessageBrokerInterface
     *
     */
    public function getMessage()
    {
        if($this->_mode = "receiving"){
            return $this->receiver->getMessage();
        }
        return $this->messageBroker->getMessage();
    }

    public function configure(Array $configuration)
    {
        $this->configurationBroker->configure($configuration);

        return $this;
    }
    public function getConfigurations(){
        return $this->configurationBroker->getConfigurations();
    }

    public function getConfiguration($key){
        return $this->configurationBroker->getConfiguration($key);
    }

    public function setConfiguration($key,$value){
        $this->configurationBroker->setConfiguration($key,$value);
        return $this;
    }

    public function setApplication($app){
        $this->configurationBroker->setConfiguration('APP_ID',$app);
        return $this;
    }
    public function setServer($server){
        $this->configurationBroker->setConfiguration('SERVER_URL',$server);
        return $this;
    }
    public function setSecret($secret){
        $this->configurationBroker->setConfiguration('APP_SECRET',$secret);
        return $this;
    }


    /**
     * @return mixed
     */
    public function send()
    {
        return $this->sender->send(
            $this->messageBroker,
            $this->addressBroker,
            $this->configurationBroker);
    }

    /**
     *
     * Receiving the Request
     *
     * @return Handler
     */
    public function receive($request=null)
    {
        $this->_mode = "receiving";
        if (!isset($request)) {
            $post = file_get_contents('php://input');
            $request = json_decode($post);
        }

        return $this->receiver->receive($request, $this);
    }
     /**
      * Returns the Source Address
      *
      * @return void
      */
    public function getSender()
    {
        return $this->receiver->getSender();
    }

    /**
     * Returns the Encoding
     *
     * @return encoding
     */
    public function getEncoding()
    {
        return $this->receiver->getEncoding();
    }
    /**
     * Returns Application ID of Request
     *
     * @return app_id
     */
    public function getApplication()
    {
        return $this->receiver->getApplication();
    }


}


