<?php declare(strict_types=1);


namespace DesignPatterns\Creational\AbstractFactory;


class UnixCsvWriter implements CsvWriter
{

    public function writer(array $line): string
    {
        return join(',', $line) . "\n";
    }
}