<?php


namespace Joomtriggers\Ideamart\SMS;

use Joomtriggers\Ideamart\Contracts\ReceiverInterface;

class Receiver implements ReceiverInterface
{
    protected $sender = "";
    protected $encoding = 0;
    protected $app_id="";
    protected $message="";
    /**
     * Receives the request
     *
     * @param mixed   $request Array
     * @param Handler $handler Handler
     *
     * @return $handler
     */
    public function receive($request,Handler $handler)
    {

        $this->sender = $request->sourceAddress;
        $this->app_id = $request->requestId;
        $this->encoding = $request->encoding;
        $this->message = $request->message;


        return $handler;
    }

    /**
     * undocumented function
     *
     */
    public function getSender()
    {
        return $this->sender;
    }
    /**
     * undocumented function
     *
     */
    public function getEncoding()
    {
        return $this->encoding;
    }

    public function getApplication(){
        return $this->app_id;
    }

    /**
     * Message
     *
     */
    public function getMessage()
    {
        return $this->message;
    }



}
