<?php

namespace Ideamart\USSD\Ideamart\Maker;

use Inqurtime\System\Session\Contracts\SessionHandler;

/**
 * Class Responder
 * @package Inqurtime\System\USSD\Ideamart\Maker
 */
class Responder {
    protected $session;
    protected $config;
    protected $menu_path;
    protected $option;
    protected $actionPath;
    protected $variables;
    protected $actionTrigger;
    protected $actionMessage;
    protected $message;

    public function __construct(SessionHandler $sessionHandler) {
        $this->session = $sessionHandler;
    }

    public function produceResponse(Parser $parser, MessageTranslator $translator,array $request) {

        $session = $translator->getSession();
        $this->config = $request;
        $this->menu_path = $session->getMenuPath();
        $this->option = $session->getOption();
        if ($this->option) {
            $menu = $parser->getMenuPlain($session);
            if ($menu['type'] == "message") {
                $this->message = $menu['message'];
            }
        } elseif ($this->option == null) {
            $this->message = $parser->makeMenu($session);
        }

        return $this->message;
    }


    public function getMessage() {
        return $this->message;
    }

}
