pragma solidity >=0.4.24 <0.6.0;

contract SimpleElection {

    mapping (address => string) public OwnerToCandidate;

    function elect(string memory _candidateName) public {
        OwnerToCandidate[msg.sender] = _candidateName;
    }

}
