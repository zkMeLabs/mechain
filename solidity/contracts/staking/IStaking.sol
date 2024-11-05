// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

/**
 * @dev Description defines a validator description.
 */
struct Description {
    // moniker defines a human-readable name for the validator.
    string moniker;
    // identity defines an optional identity signature (ex. UPort or Keybase).
    string identity;
    // website defines an optional website link.
    string website;
    // securityContact defines an optional email for security contact.
    string securityContact;
    // details define other optional details.
    string details;
}

/**
 * @dev CommissionRates defines the initial commission rates to be used for creating a validator.
 */
struct CommissionRates {
    // rate defines the maximum commission rate which validator can ever charge, as a fraction.
    uint256 rate;
    // maxRate defines the maximum commission rate which validator can ever charge, as a fraction.
    uint256 maxRate;
    // maxChangeRate defines the maximum daily increase of the validator commission, as a fraction.
    uint256 maxChangeRate;
}

/**
 * @dev Commission defines commission parameters for a given validator.
 */
struct Commission {
    // commissionRates defines the initial commission rates to be used for creating a validator
    CommissionRates commissionRates;
    // updateTime defines the validator update commissionRates time
    int64 updateTime;
}

/**
 * @dev BondStatus is the status of a validator.
 */
enum BondStatus {
    // Unspecified defines an invalid validator status.
    Unspecified,
    // Unbonded defines a validator that is not bonded.
    Unbonded,
    // Unbonding defines a validator that is unbonding.
    Unbonding,
    // Bonded defines a validator that is bonded.
    Bonded
}

/**
 * @dev Validator defines a validator, together with the total amount of the
 * Validator's bond shares and their exchange rate to coins. Slashing results in
 * a decrease in the exchange rate, allowing correct calculation of future
 * undelegations without iterating over delegators. When coins are delegated to
 * this validator, the validator is credited with a delegation whose number of
 * bond shares is based on the amount of coins delegated divided by the current
 * exchange rate. Voting power can be calculated as total bonded shares
 * multiplied by exchange rate.
 */
struct Validator {
    // operatorAddress defines the address of the validator's operator
    address operatorAddress;
    // consensusPubkey is the consensus public key of the validator
    string consensusPubkey;
    // jailed defined whether the validator has been jailed from bonded status or not.
    bool jailed;
    // status is the validator status (bonded/unbonding/unbonded).
    BondStatus status;
    // tokens define the delegated tokens (incl. self-delegation).
    uint256 tokens;
    // delegatorShares defines total shares issued to a validator's delegators.
    uint256 delegatorShares;
    // description defines the description terms for the validator.
    Description description;
    // unbonding_height defines, if unbonding, the height at which this validator has begun unbonding.
    int64 unbondingHeight;
    // unbonding_time defines, if unbonding, the min time for the validator to complete unbonding.
    int64 unbondingTime;
    // commission defines the commission parameters.
    Commission commission;
    // minSelfDelegation is the validator's self declared minimum self delegation.
    uint256 minSelfDelegation;
    // unbondingOnHoldRefCount strictly positive if this validator's unbonding has been stopped by external modules
    int64 unbondingOnHoldRefCount;
    // unbondingIds list of unbonding ids, each uniquely identifying an unbonding of this validator
    uint64[] unbondingIds;
    // selfDelAddress defines the address of the validator for self delegation.
    string selfDelAddress;
    // relayerAddress defines the address of the validator's authorized relayer.
    string relayerAddress;
    // challengerAddress defines the address of the validator's authorized challenger.
    string challengerAddress;
    // blsKey defines the bls pubkey of the validator's authorized relayer/challenger/operator.
    string blsKey;
}

/**
 * @dev Delegation represents the bond with tokens held by an account. It is
 * owned by one delegator, and is associated with the voting power of one
 * validator.
 */
struct Delegation {
    // delegatorAddress is the address of the delegator.
    address delegatorAddress;
    // validatorAddress is the address of the validator.
    address validatorAddress;
    // shares define the delegation shares received.
    Dec shares;
}

/**
 * @dev DelegationResponse is equivalent to Delegation except that it contains a
 * balance in addition to shares which is more suitable for client responses.
 */
struct DelegationResponse {
    Delegation delegation;
    Coin balance;
}

/**
 * @dev UnbondingDelegationEntry defines an unbonding object with relevant metadata.
 */
struct UnbondingDelegationEntry {
    // creationHeight is the height which the unbonding took place.
    int64 creationHeight;
    // completionTime is the unix time for unbonding completion.
    int64 completionTime;
    // initialBalance defines the tokens initially scheduled to receive at completion.
    uint256 initialBalance;
    // balance defines the tokens to receive at completion.
    uint256 balance;
}

/**
 * @dev RedelegationEntry defines a redelegation object with relevant metadata.
 */
struct RedelegationEntry {
    // creationHeight defines the height which the redelegation took place.
    int64 creationHeight;
    // completionTime defines the unix time for redelegation completion.
    int64 completionTime;
    // initialBalance defines the initial balance when redelegation started.
    uint256 initialBalance;
    // shareDst is the amount of destination-validator shares created by redelegation.
    uint256 shareDst;
}

/**
 * @dev RedelegationEntryResponse is equivalent to a RedelegationEntry except that it
 * contains a balance in addition to shares which is more suitable for client responses.
 */
struct RedelegationEntryResponse {
    RedelegationEntry redelegationEntry;
    uint256 balance;
}

/**
 * @dev UnbondingDelegation stores all of a single delegator's unbonding bonds
 * for a single validator in an time-ordered list.
 */
struct UnbondingDelegation {
    address delegatorAddress;
    address validatorAddress;
    UnbondingDelegationEntry[] entries;
}

/**
 * @dev RedelegationResponse is equivalent to a Redelegation except that its entries
 * contain a balance in addition to shares which is more suitable for client responses.
 */
struct RedelegationResponse {
    Redelegation redelegation;
    RedelegationEntryResponse[] entries;
}

/**
 * @dev Redelegation contains the list of a particular delegator's redelegating bonds
 * from a particular source validator to a particular destination validator.
 */
struct Redelegation {
    // delegatorAddress is the bech32-encoded address of the delegator.
    address delegatorAddress;
    // validatorSrcAddress is the validator redelegation source operator address.
    address validatorSrcAddress;
    // validatorDstAddress is the validator redelegation destination operator address.
    address validatorDstAddress;
    // entries are the redelegation entries.
    RedelegationEntry[] entries;
}

/**
 * @dev HistoricalInfo contains header and validator information for a given block.
 * It is stored as part of staking module's state, which persists the `n` most
 * recent HistoricalInfo
 * (`n` is set by the staking module's `historical_entries` parameter).
 */
struct HistoricalInfo {
    Header header;
    Validator[] valset;
}

/**
 * @dev Pool is used for tracking bonded and not-bonded token supply of the bond
 * denomination.
 */
struct Pool {
    uint256 notBondedTokens;
    uint256 bondedTokens;
}

/**
 * @dev Params defines the parameters for the staking module.
 */
struct Params {
    // unbondingTime is the time duration of unbonding.
    int64 unbondingTime;
    // maxValidators is the maximum number of validators.
    uint32 maxValidators;
    // maxEntries is the max entries for either unbonding delegation or redelegation (per pair/trio).
    uint32 maxEntries;
    // historicalEntries is the number of historical entries to persist.
    uint32 historicalEntries;
    // bondDenom defines the bondable coin denomination.
    string bondDenom;
    // min_commission_rate is the chain-wide minimum commission rate that a validator can charge their delegators
    uint256 minCommissionRate;
}

/**
 * @dev Header defines the structure of a Tendermint block header.
 */
struct Header {
    // basic block info
    Consensus version;
    string chainId;
    int64 height;
    int64 time;

    // prev block info
    BlockID lastBlockId;

    // hashes of block data
    string lastCommitHash;     // commit from validators from the last block
    string dataHash;           // transactions

    // hashes from the app output from the prev block
    string validatorsHash;     // validators for the current block
    string nextValidatorsHash; // validators for the next block
    string consensusHash;      // consensus params for current block
    string appHash;            // state after txs from the previous block
    string lastResultsHash;    // root hash of all results from the txs from the previous block

    // consensus info
    string evidenceHash;       // evidence included in the block
    string proposerAddress;    // original proposer of the block
}

/**
 * @dev Consensus captures the consensus rules for processing a block in the blockchain,
 * including all blockchain data structures and the rules of the application's
 * state transition machine.
 */
struct Consensus {
    uint64 block;
    uint64 app;
}

/**
 * @dev BlockID
 */
struct BlockID {
    string hash;
    PartSetHeader partSetHeader;
}

/**
 * @dev PartsetHeader
 */
struct PartSetHeader {
    uint32 total;
    string hash;
}

interface IStaking {
    /**
     * @dev editValidator defines a method for editing an existing validator.
     */
    function editValidator(
        Description calldata description,
        int256 commissionRate,
        int256 minSelfDelegation,
        address relayerAddress,
        address challengerAddress,
        string memory blsKey,
        string memory blsProof
    ) external returns (bool success);

    /**
     * @dev delegate defines a method for performing a delegation of coins
     * from a delegator to a validator.
     */
    function delegate(
        address validatorAddress,
        uint256 amount
    ) external returns (bool success);

    /**
     * @dev undelegate defines a method for performing an undelegation from a
     * delegate and a validator.
     */
    function undelegate(
        address validatorAddress,
        uint256 amount
    ) external returns (uint256 completionTime);

    /**
     * @dev redelegate defines a method for performing a redelegation
     * of coins from a delegator and source validator to a destination validator.
     */
    function redelegate(
        address validatorSrcAddress,
        address validatorDstAddress,
        uint256 amount
    ) external returns (uint256 completionTime);

    /**
     * @dev cancelUnbondingDelegation defines a method for performing canceling the unbonding delegation
     * and delegate back to previous validator.
     */
    function cancelUnbondingDelegation(
        address validatorAddress,
        uint256 amount,
        uint256 creationHeight
    ) external returns (bool success);

    /**
     * @dev validators queries all validators that match the given status.
     */
    function validators(
        BondStatus status,
        PageRequest calldata pagination
    ) external view returns (Validator[] calldata validators, PageResponse calldata pageResponse);

    /**
     * @dev validator queries validator info for given validator address.
     */
    function validator(
        address validatorAddr
    ) external view returns (Validator calldata validator);

    /**
     * @dev validatorDelegations queries delegate info for given validator.
     */
    function validatorDelegations(
        address validatorAddr,
        PageRequest calldata pagination
    ) external view returns (DelegationResponse[] calldata response, PageResponse calldata pageResponse);

    /**
     * @dev validatorUnbondingDelegations queries unbonding delegations of a validator.
     */
    function validatorUnbondingDelegations(
        address validatorAddr,
        PageRequest calldata pagination
    ) external view returns (UnbondingDelegation[] calldata response, PageResponse calldata pageResponse);

    /**
     * @dev delegation queries delegate info for given validator delegator pair.
     */
    function delegation(
        address delegatorAddr,
        address validatorAddr
    ) external view returns (DelegationResponse calldata response);

    /**
     * @dev unbondingDelegation queries unbonding info for given validator delegator pair.
     */
    function unbondingDelegation(
        address delegatorAddr,
        address validatorAddr
    ) external view returns (UnbondingDelegation calldata response);

    /**
     * @dev delegatorDelegations queries all delegations of a given delegator address.
     */
    function delegatorDelegations(
        address delegatorAddr,
        PageRequest calldata pagination
    ) external view returns (DelegationResponse[] calldata response, PageResponse calldata pageResponse);

    /**
     * @dev delegatorUnbondingDelegations queries all unbonding delegations of a given delegator address.
     */
    function delegatorUnbondingDelegations(
        address delegatorAddr,
        PageRequest calldata pagination
    ) external view returns (UnbondingDelegation[] calldata response, PageResponse calldata pageResponse);

    /**
     * @dev redelegations queries redelegations of given address.
     */
    function redelegations(
        address delegatorAddr,
        address srcValidatorAddr,
        address dstValidatorAddr,
        PageRequest calldata pagination
    ) external view returns (RedelegationResponse[] calldata redelegationResponses, PageResponse calldata pageResponse);

    /**
     * @dev delegatorValidators queries all validators info for given delegator address.
     */
    function delegatorValidators(
        address delegatorAddr,
        PageRequest calldata pagination
    ) external view returns (Validator[] calldata validators, PageResponse calldata pageResponse);

    /**
     * @dev delegatorValidator queries validator info for given delegator validator pair.
     */
    function delegatorValidator(
        address delegatorAddr,
        address validatorAddr
    ) external view returns (Validator calldata validator);

    /**
     * @dev historicalInfo queries the historical info for given height.
     */
    function historicalInfo(
        int64 height
    ) external view returns (HistoricalInfo calldata historicalInfo);

    /**
     * @dev pool queries the pool info.
     */
    function pool() external view returns (Pool calldata pool);

    /**
     * @dev params queries the staking params.
     */
    function params() external view returns (Params calldata params);

    /**
     * @dev EditValidator defines an Event emitted when a validator edited.
     */
    event EditValidator(
        address indexed validator,
        int256 commissionRate,
        int256 minSelfDelegation
    );

    /**
     * @dev Delegate defines an Event emitted when a given amount of tokens are delegated from the
     * delegator address to the validator address.
     */
    event Delegate(
        address indexed delegator,
        address indexed validator,
        uint256 amount
    );

    /**
     * @dev Undelegate defines an Event emitted when a given amount of tokens are undelegate by delegator
     */
    event Undelegate(
        address indexed delegatorAddress,
        address indexed validatorAddress,
        uint256 amount,
        uint256 completionTime
    );

    /**
     * @dev Redelegate defines an Event emitted when a given amount of tokens are redelegated from
     * the source validator address to the destination validator address.
     */
    event Redelegate(
        address indexed delegatorAddress,
        address indexed validatorSrcAddress,
        address indexed validatorDstAddress,
        uint256 amount,
        uint256 completionTime
    );

    /**
     * @dev CancelUnbondingDelegation defines an Event emitted when a given amount of tokens are cancel
     * bond delegate from the validator by the delegator
     */
    event CancelUnbondingDelegation(
        address indexed delegatorAddress,
        address indexed validatorAddress,
        uint256 amount,
        uint256 creationHeight
    );
}
