<?php
/**
 * Class: Producer
 *
 * PHP Version: 7.0.0
 *
 * @category Category
 * @package  Package
 * @author   Gnanakeethan Balasubramaniam <gnana@keethan.me>
 * @license  MIT http://opensource.org/licenses/MIT/
 * @link     http://link/
 */

namespace Joomtriggers\Ideamart\USSD\Maker;

/**
 * Class: Producer
 *
 * @category Category
 * @package  Package
 * @author   Gnanakeethan Balasubramaniam <gnana@keethan.me>
 * @license  MIT http://opensource.org/licenses/MIT/
 * @link     http://link/
 */
class Producer
{
    protected $parser;
    protected $translator;
    protected $responder;
    protected $request;

    public function __construct(
        Parser $parser,
        MessageTranslator $translator,
        Responder $responder
    ) {
        $this->parser = $parser;
        $this->translator = $translator;
        $this->responder = $responder;
    }
    public function setRequest(array $request) {
        $this->request = $request;
        return $this;
    }
    public function loadSetup() {
        return $this;
    }
    public function translateMessage() {
        $this->translator->translate($this->request, $this->parser);
        return $this;
    }
    public function makeMessage() {
        return $this->responder->produceResponse($this->parser, $this->translator);
    }
    public function __toString() {
        return $this->responder->getMessage();
    }
}
