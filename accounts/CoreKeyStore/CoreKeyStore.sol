pragma solidity ^0.4.13;



contract CoreKeyStores {



    // map of cluster -> account -> keyjson

    mapping(address => mapping(address => string)) keys;



    function keyExists(address cluster, address account)

        internal

        returns (bool)

    {

        if(bytes(keys[cluster][account]).length != 0) return true;

        return false;

    }



    function storeKey (

        address cluster,

        address account,

        string keyjson

    )

        public

    {

        require(cluster != account);

        require(!keyExists(cluster, account));

        keys[cluster][account] = keyjson;

    }



    function getKey(address cluster, address account)

        public

        constant

        returns (string)

    {

        require(cluster != account);

        require(keyExists(cluster, account));

        return keys[cluster][account];

    }

}
