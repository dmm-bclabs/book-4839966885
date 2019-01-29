pragma solidity >=0.4.24 <0.6.0;

contract Strcut {

    struct Author {
        address addr;
        string name;
    }  
    
    Author[] public authors;
     
    constructor() public {
    }

    function getAuthor() public returns (address, string memory) {
        address authorAddr = 0x0cE446255506E92DF41614C46F1d6df9Cc969183;
        authors.push(Author(authorAddr, 'blockchain_lab'));
        return (authors[0].addr, authors[0].name);
    }
}
  
