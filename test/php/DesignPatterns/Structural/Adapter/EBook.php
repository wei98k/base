<?php declare(strict_types=1);


namespace DesignPatterns\Structural\Adapter;


interface EBook
{
    public function unlock();

    public function pressNext();

    /**
     * @return int[]
     */
    public function getPage(): array;
}