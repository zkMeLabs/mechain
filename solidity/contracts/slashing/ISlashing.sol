// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

/**
 * @dev Params represents the parameters used for by the slashing module.
 */
struct Params {
    int64 signedBlocksWindow;
    string minSignedPerWindow;
    int64 downtimeJailDuration;
    string slashFractionDoubleSign;
    string slashFractionDowntime;
}

/**
 * @dev ValidatorSigningInfo defines a validator's signing info for monitoring their
 * liveness activity.
 */
struct ValidatorSigningInfo {
    address consAddress;
    // Height at which validator was first a candidate OR was unjailed
    int64 startHeight;
    // Index which is incremented each time the validator was a bonded
    // in a block and may have signed a precommit or not. This in conjunction with the
    // `SignedBlocksWindow` param determines the index in the `MissedBlocksBitArray`.
    int64 indexOffset;
    // Timestamp until which the validator is jailed due to liveness downtime.
    int64 jailedUntil;
    // Whether or not a validator has been tombstoned (killed out of validator set). It is set
    // once the validator commits an equivocation or for any other configured misbehiavor.
    bool tombstoned;
    // A counter kept to avoid unnecessary array reads.
    // Note that `Sum(MissedBlocksBitArray)` always equals `MissedBlocksCounter`.
    int64 missedBlocksCounter;
}

interface ISlashing {
    /**
     * @dev unjail defines a method for unjailing a jailed validator, thus returning
     * them into the bonded validator set, so they can begin receiving provisions
     * and rewards again.
     */
    function unjail() external returns (bool success);

    /**
     * @dev Params queries the parameters of slashing module
     */
    function params() external view returns (Params memory params);

    /**
     * @dev signingInfo queries the signing info of given cons address
     */
    function signingInfo(address consAddress) external view returns (ValidatorSigningInfo memory valSigningInfo);

    /**
     * @dev signingInfos queries signing info of all validators
     */
    function signingInfos(PageRequest calldata pagination) external view returns (ValidatorSigningInfo[] memory infos, PageResponse memory pageResponse);

    /**
     * @dev Unjail defines an Event emitted when a validator unjail
     */
    event Unjail(
        address indexed validator
    );
}
