// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

contract Mock {
    function latestRoundData() external view
        returns (
            uint80 roundId,
            int answer,
            uint startedAt,
            uint updatedAt,
            uint80 answeredInRound
        ) {
            // 984.3445 usdt
        return (0, 98434450000, 0, 0, 0);
    }
}