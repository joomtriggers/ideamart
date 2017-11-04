<?php
/**
 * Class: Receiver
 *
 * PHP Version: 7.0.0
 *
 * @category Category
 * @package  Package
 * @author   Gnanakeethan Balasubramaniam <gnana@keethan.me>
 * @license  MIT http://opensource.org/licenses/MIT/
 * @link     http://link/
 */

namespace Ideamart\USSD\Ideamart\Receiver;

/**
 * Class: Receiver
 *
 * @category Category
 * @package  Package
 * @author   Gnanakeethan Balasubramaniam <gnana@keethan.me>
 * @license  MIT http://opensource.org/licenses/MIT/
 * @link     http://link/
 */
class Receiver
{

    protected $request = [];

    protected $validator;

    protected $formatter;

    /**
     * Constructor
     *
     * @param RequestFormatter $formatter Formatter
     * @param RequestValidator $validator Validator
     */
    public function __construct(RequestFormatter $formatter, RequestValidator $validator)
    {
        $this->validator = $validator;
        $this->formatter = $formatter;

    }

    /**
     * Receiver
     *
     * @param array $request Request Array
     *
     * @return Object
     */
    public function receive(array $request)
    {
        $request = $this->validator->validate($request);
        $this->request = $this->formatter->format($request);

        return $this->request;
    }

}
