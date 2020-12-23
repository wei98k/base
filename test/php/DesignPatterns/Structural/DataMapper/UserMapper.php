<?php declare(strict_types=1);


namespace DesignPatterns\Structural\DataMapper;


use InvalidArgumentException;

class UserMapper
{
    private StorageAdapter $adapter;

    public function __construct(StorageAdapter $storage)
    {
        $this->adapter = $storage;
    }

    public function findById(int $id): User
    {
        $result = $this->adapter->find($id);

        if($result === null) {
            throw new InvalidArgumentException("User #$id not found");
        }

        return $this->mapRowToUser($result);
    }

    public function mapRowToUser(array $row): User
    {
        return User::fromState($row);
    }
}