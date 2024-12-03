// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

struct Params {
    uint64 maximumStatementsNum;
    uint64 maximumGroupNum;
    uint64 maximumRemoveExpiredPoliciesIteration;
}

interface IPermission {
    /**
     * @dev updateParams defines a method for update params.
     */
    function updateParams(
        string memory authority,
        Params memory params
    ) external returns (bool success);

    /**
     * @dev params defines a method for queries the parameters of the module.
     */
    function params() external view returns (Params calldata params);

    /**
     * @dev UpdateParams defines an Event emitted when a user update params
     */
    event UpdateParams(address indexed creator);
}
