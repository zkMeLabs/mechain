// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

/**
 * @dev Coin defines a token with a denomination and an amount.
 */
struct Coin {
    string denom;
    uint256 amount;
}