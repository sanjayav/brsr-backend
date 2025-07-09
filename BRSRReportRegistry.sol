// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract BRSRReportRegistry {
    address public owner;
    uint256 public fineRatePerDay = 1 ether;
    uint256 public deadline;

    struct Report {
        string company;
        string ipfsHash;
        uint256 timestamp;
        uint256 finePaid;
        bool submitted;
    }

    mapping(address => Report) public reports;
    mapping(address => bool) public admins;
    address[] public submitters;

    event ReportSubmitted(address indexed submitter, string ipfsHash, uint256 timestamp, uint256 finePaid);
    event ReportUpdated(address indexed submitter, string ipfsHash, uint256 timestamp);
    event ReportDeleted(address indexed submitter);
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

    function submitReport(string memory _company, string memory _ipfsHash) external payable {
        require(!reports[msg.sender].submitted, "Already submitted");

        uint256 currentTime = block.timestamp;
        uint256 fine = 0;
        if (currentTime > deadline) {
            uint256 daysLate = (currentTime - deadline) / 1 days;
            fine = daysLate * fineRatePerDay;
            require(msg.value >= fine, "Insufficient fine payment");
        }

        reports[msg.sender] = Report({
            company: _company,
            ipfsHash: _ipfsHash,
            timestamp: currentTime,
            finePaid: fine,
            submitted: true
        });

        submitters.push(msg.sender);
        emit ReportSubmitted(msg.sender, _ipfsHash, currentTime, fine);
    }

    function updateReport(string memory _ipfsHash) external onlyAdmin {
        require(reports[msg.sender].submitted, "Report not submitted");
        reports[msg.sender].ipfsHash = _ipfsHash;
        reports[msg.sender].timestamp = block.timestamp;
        emit ReportUpdated(msg.sender, _ipfsHash, block.timestamp);
    }

    function deleteReport(address _user) external onlyAdmin {
        require(reports[_user].submitted, "No report to delete");
        delete reports[_user];
        emit ReportDeleted(_user);
    }

    function getFine(address _user) external view returns (uint256) {
        if (reports[_user].submitted) return 0;
        if (block.timestamp <= deadline) return 0;
        uint256 daysLate = (block.timestamp - deadline) / 1 days;
        return daysLate * fineRatePerDay;
    }

    function viewAllReports() external view onlyAdmin returns (Report[] memory) {
        uint256 count = submitters.length;
        Report[] memory allReports = new Report[](count);
        for (uint256 i = 0; i < count; i++) {
            allReports[i] = reports[submitters[i]];
        }
        return allReports;
    }

    function withdraw() external onlyOwner {
        payable(owner).transfer(address(this).balance);
    }
}
