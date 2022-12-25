// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

contract SaleStock {
    struct Stock {
        string hash;
        string goodsId;
        string style;
        uint32 weight; //产品的重量
        string pichash;
        uint16 price;
        string shifu;
        uint time;
    }
    address public _sender;
    address public _txOrigin;

    constructor() {
        _sender = msg.sender;
        _txOrigin = tx.origin;
    }

    modifier onlySender() {
        require(msg.sender == _sender, "Not admin");
        _;
    }

    mapping(string => Stock) private _stock; //use value hash similar like uuid to storage EvidenceData

    function setData(
        string memory _hash,
        string memory _goodsId,
        string memory _style,
        uint32 _weight,
        string memory _pichash,
        uint16 _price,
        string memory _shifu
    ) public onlySender {
        _stock[_hash].hash = _hash;
        _stock[_hash].goodsId = _goodsId;
        _stock[_hash].style = _style;
        _stock[_hash].weight = _weight;
        _stock[_hash].pichash = _pichash;
        _stock[_hash].price = _price;
        _stock[_hash].shifu = _shifu;
        _stock[_hash].time = block.timestamp;

    }

    function getData(string memory _hash)
        public
        view
        returns (
            string memory,
            string memory,
            string memory,
            uint32,
            string memory,
            uint16,
            string memory,
            uint
        )
    {
        Stock storage stock = _stock[_hash];
        return (
            stock.hash,
            stock.goodsId,
            stock.style,
            stock.weight,
            stock.pichash,
            stock.price,
            stock.shifu,
            stock.time
        );
    }
}

