// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract BRSRReportRegistry {
    address public owner;
    uint256 public fineRatePerDay = 1 ether; // Adjustable fine rate
    uint256 public deadline; // UNIX timestamp

    struct Report {
        string company;
        string ipfsHash;
        uint256 timestamp;
        uint256 finePaid;
        bool submitted;
    }

    mapping(address => Report) public reports;
    mapping(address => bool) public admins;

    event ReportSubmitted(address indexed submitter, string ipfsHash, uint256 timestamp, uint256 finePaid);
    event FineAdjusted(address indexed submitter, uint256 newFine);
    event AdminUpdated(address indexed admin, bool status);
    event DeadlineUpdated(uint256 newDeadline);

    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    modifier onlyAdmin() {
        require(admins[msg.sender] || msg.sender == owner, "Not admin");
        _;
    }

    constructor(uint256 _deadline) {
        owner = msg.sender;
        deadline = _deadline;
        admins[owner] = true;
    }

    function updateDeadline(uint256 _newDeadline) external onlyOwner {
        deadline = _newDeadline;
        emit DeadlineUpdated(_newDeadline);
    }

    function setFineRate(uint256 _newRate) external onlyOwner {
        fineRatePerDay = _newRate;
    }

    function setAdmin(address _admin, bool _status) external onlyOwner {
        admins[_admin] = _status;
        emit AdminUpdated(_admin, _status);
    }

    function submitReport(string calldata _company, string calldata _ipfsHash) external payable {
        require(!reports[msg.sender].submitted, "Already submitted");

        uint256 timeNow = block.timestamp;
        uint256 fine = 0;

        if (timeNow > deadline) {
            uint256 delayDays = (timeNow - deadline) / 1 days;
            fine = delayDays * fineRatePerDay;
            require(msg.value >= fine, "Insufficient fine payment");
        }

        reports[msg.sender] = Report({
            company: _company,
            ipfsHash: _ipfsHash,
            timestamp: timeNow,
            finePaid: fine,
            submitted: true
        });

        emit ReportSubmitted(msg.sender, _ipfsHash, timeNow, fine);
    }

    function adjustFine(address _reporter, uint256 _newFine) external onlyAdmin {
        require(reports[_reporter].submitted, "No report submitted");
        reports[_reporter].finePaid = _newFine;
        emit FineAdjusted(_reporter, _newFine);
    }

    function withdraw() external onlyOwner {
        payable(owner).transfer(address(this).balance);
    }

    function getReport(address _reporter) external view returns (Report memory) {
        return reports[_reporter];
    }
}
