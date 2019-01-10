pragma solidity ^0.4.24;

contract SimpleElection {

    mapping (address => string) public OwnerToCandidate;

    function elect(string _candidateName) public {
        OwnerToCandidate[msg.sender] = _candidateName;
    }

}