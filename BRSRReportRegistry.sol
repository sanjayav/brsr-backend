// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract BRSRReportRegistry {
    struct Report {
        address submittedBy;
        string ipfsHash;
        string company;
        string financialYear;
        uint256 timestamp;
    }

    Report[] public reports;

    event ReportSubmitted(address indexed submitter, string ipfsHash, string company, string financialYear);

    function submitReport(string memory ipfsHash, string memory company, string memory financialYear) public {
        reports.push(Report({
            submittedBy: msg.sender,
            ipfsHash: ipfsHash,
            company: company,
            financialYear: financialYear,
            timestamp: block.timestamp
        }));

        emit ReportSubmitted(msg.sender, ipfsHash, company, financialYear);
    }

    function getReport(uint256 index) public view returns (
        address submittedBy, string memory ipfsHash, string memory company, string memory financialYear, uint256 timestamp
    ) {
        Report memory r = reports[index];
        return (r.submittedBy, r.ipfsHash, r.company, r.financialYear, r.timestamp);
    }

    function getTotalReports() public view returns (uint256) {
        return reports.length;
    }
}
