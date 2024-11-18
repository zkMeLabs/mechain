// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

/**
 * @dev DelegationDelegatorReward represents the properties of a delegator's delegation reward
 */
struct DelegationDelegatorReward {
    address validatorAddress;
    DecCoin[] rewards;
}


/**
 * @dev Params defines the set of params for the distribution module.
 */
struct Params {
    uint256 communityTax;
    uint256 baseProposerReward;
    uint256 bonusProposerReward;
    bool withdrawAddrEnabled;
}


/**
 * @dev QueryValidatorSlashesRequest is the request type for the
 * Query/ValidatorSlashes RPC method
 */
struct QueryValidatorSlashesRequest {
    string validatorAddress;
    uint64 startingHeight;
    uint64 endingHeight;
    PageRequest pagination;
}

/**
 * @dev ValidatorSlashEvent represents a validator slash event.
 * Height is implicit within the store key.
 * This is needed to calculate appropriate amount of staking tokens
 * for delegations which are withdrawn after a slash has occurred.
 */
struct ValidatorSlashEvent {
    uint64 validatorPeriod;
    uint256 fraction;
}

interface IDistribution {
    /**
     * @dev setWithdrawAddress SetWithdrawAddress defines a method to change the withdraw address
     * for a delegator (or validator self-delegation).
     */
    function setWithdrawAddress(
        address withdrawAddress
    ) external returns (bool success);

    /**
     * @dev withdrawDelegatorReward defines a method to withdraw rewards of delegator
     * from a single validator.
     */
    function withdrawDelegatorReward(
        address validatorAddress
    ) external returns (Coin[] memory amount);

    /**
     * @dev withdrawValidatorCommission defines a method to withdraw the
     * full commission to the validator address.
     */
    function withdrawValidatorCommission() external returns (Coin[] memory amount);

    /**
     * @dev fundCommunityPool defines a method to allow an account to directly
     * fund the community pool.
     */
    function fundCommunityPool(
        Coin[] memory amount
    ) external returns (bool success);

    /**
     * @dev validatorDistributionInfo queries validator commision and self-delegation rewards for validator
     */
    function validatorDistributionInfo(
        address validatorAddress
    ) external view returns (address operatorAddress, DecCoin[] memory selfBondRewards, DecCoin[] memory commission);

    /**
     * @dev validatorOutstandingRewards queries rewards of a validator address.
     */
    function validatorOutstandingRewards(
        address validatorAddress
    ) external view returns (DecCoin[] memory rewards);

    /**
     * @dev validatorCommission queries accumulated commission for a validator.
     */
    function validatorCommission(
        address validatorAddress
    ) external view returns (DecCoin[] memory commission);

    /**
     * @dev delegationRewards queries the total rewards accrued by a delegation.
     */
    function delegationRewards(
        address delegatorAddress,
        address validatorAddress
    ) external view returns (DecCoin[] memory rewards);

    /**
     * @dev delegationTotalRewards queries the total rewards accrued by a each validator.
     */
    function delegationTotalRewards(
        address delegatorAddress
    ) external view returns (DelegationDelegatorReward[] memory rewards, DecCoin[] memory total);

    /**
     * @dev communityPool queries the community pool coins.
     */
    function communityPool() external view returns (DecCoin[] memory pool);

    /**
     * @dev params queries params of the distribution module.
     */
    function params() external view returns (Params memory params);

    /**
     * @dev validatorSlashes queries slash events of a validator.
     */
    function validatorSlashes(
        address validatorAddress,
        uint64 startingHeight,
        uint64 endingHeight,
        PageRequest calldata pagination
    ) external view returns (ValidatorSlashEvent[] memory validatorSlashEvents, PageResponse memory pageResponse);

    /**
     * @dev delegatorValidators queries the validators of a delegator.
     */
    function delegatorValidators(
        address delegatorAddress
    ) external view returns (address[] memory validators);

    /**
     * @dev delegatorWithdrawAddress queries withdraw address of a delegator.
     */
    function delegatorWithdrawAddress(
        address delegatorAddress
    ) external view returns (address withdrawAddress);

    /**
     * @dev SetWithdrawAddress defines an Event emitted when a user change the withdraw address
     */
    event SetWithdrawAddress(
        address indexed delegatorAddress,
        address indexed withdrawAddress
    );

    /**
     * @dev WithdrawDelegatorReward defines an Event emitted when withdraw rewards by delegator
     */
    event WithdrawDelegatorReward(
        address indexed delegatorAddress,
        address indexed withdrawAddress,
        string amount
    );

    /**
     * @dev WithdrawValidatorCommission defines an Event emitted when withdraw commission by validator
     */
    event WithdrawValidatorCommission(
        address indexed validatorAddress,
        string amount
    );

    /**
     * @dev FundCommunityPool defines an Event emitted when a user fund community pool
     */
    event FundCommunityPool(
        address indexed depositor,
        string amount
    );
}
