<?php

namespace Ideamart\USSD\Ideamart\Maker;


use Joomtriggers\Ideamart\Contracts\USSD\SessionHandlerInterface;

class MessageTranslator {
    protected $session;
    protected $configLoader;
    protected $serviceBroker;

    public function __construct(
        SessionHandlerInterface $sessionHandler,
        ConfigLoader $configLoader,
        ServiceBroker $serviceBroker
    ) {
        $this->session = $sessionHandler;
        $this->configLoader = $configLoader;
        $this->serviceBroker = $serviceBroker;
    }

    public function translate(array $request, Parser $parser) {
        $message = $request['message'];
        $this->session->setParameters($request);
        $this->session->setAppId($applicationRepository->searchApplication('ideamart', $request['applicationId'])['_id']);
        $this->configLoader->setApplication();
        $this->updateService();
        $menu = $parser->getMenuPlain();
        $message = $this->messageConverter($message);
        $this->messageIterator($menu['options'], $message, $parser);
    }

    private function messageIterator(array $menu, array $message, Parser $parser) {

        if (empty($message)) {
            if ($parser->getMenuPlain()['type'] == 'master_menu') {
                $this->session->setOperation("mt-cont");
            }
            return "";

        }
        foreach ($message as $key) {
            if ($this->session->isAction()) {
                $option = array_pop($message);
                $this->session->setOption($option);
                $this->session->setOperation("mt-cont");
            } //check key, go in if set, otherwise just move forward
            elseif (isset($menu[$key]) && !$this->session->isAction()) {
                if ($menu[$key]['type'] == "sub_menu") {
                    $operation = $menu[$key]['response'];
                    $this->session->setOperation($operation);
                    $this->session->setMenuPath($this->session->getMenuPath() . '.options.' . $key . '.sub_menu');
                    array_shift($message);
                    return $this->messageIterator($menu[$key]['sub_menu']['options'], $message, $parser);
                } elseif ($menu[$key]['type'] == "message") {
                    $operation = $menu[$key]['response'];
                    $this->session->setOption($key);
                    $this->session->setOperation($operation);
                    $this->session->setMenuPath($this->session->getMenuPath() . '.options.' . $key);
                } else {
                    $option = array_pop($message);
                    $this->session->setOption($option);
                    if (isset($menu[$key])) {
                        $operation = $menu[$key]['response'];
                        $this->session->setOperation($operation);
                    }
                }
            }
        }

    }

    public function getSession() {
        return $this->session;
    }

    private function messageConverter($message = "") {

        $code = $this->config->get('system.' . $this->session->getAppId() . '.ussd.code');
        $count = 1;
        $message = str_replace($code, "", $message, $count);
        $message = explode("*", $message);
        $message[sizeof($message) - 1] = trim($message[sizeof($message) - 1], '#');

        return array_filter($message);
    }

    private function updateService() {

        $app_id = $this->session->getAppId();

        $provider = $this->config->get('system.' . $app_id . '.ussd.provider');
        $config = $this->config->get('system.' . $app_id . '.ussd.providers.' . $provider);

        $this->serviceBroker->app_id = $config['app_id'];
        $this->serviceBroker->server_url = $config['server'];
        $this->serviceBroker->password = $config['app_secret'];
    }
}
