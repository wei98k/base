<?php declare(strict_types=1);


namespace DesignPatterns\Creational\AbstractFactory;


class WinCsvWriter implements CsvWriter
{

    public function writer(array $line): string
    {
        return join(',', $line) . "\r\n";
    }
}