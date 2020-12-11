<?php declare(strict_types=1);


namespace DesignPatterns\Creational\AbstractFactory;


interface CsvWriter
{
    public function writer(array $line): string;
}