// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

interface IBank {
    /**
     * @dev send defines a method for sending coins from one account to another account.
     */
    function send(
        address toAddress,
        Coin[] memory amount
    ) external returns (bool success);

    /**
     * @dev totalSupply queries the total supply of all coins.
     */
    function totalSupply() external view returns (Coin[] memory supply);

    /**
     * @dev Send defines an Event emitted when a given amount of tokens send fromAddress to toAddress
     */
    event Send(
        address indexed fromAddress,
        address indexed toAddress,
        string amount
    );
}
