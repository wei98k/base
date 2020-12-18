<?php declare(strict_types=1);


namespace DesignPatterns\Structural\Adapter;


class Kindle implements EBook
{

    private int $page = 1;
    private int $totalPages = 100;

    public function unlock()
    {
    }

    public function pressNext()
    {
        $this->page++;
    }

    /**
     * @return int[]
     */
    public function getPage(): array
    {
        return [$this->page, $this->totalPages];
    }
}