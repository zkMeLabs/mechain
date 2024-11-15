// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

struct PaymentAccount {
    string addr;
    string owner;
    bool refundable;
}

struct VersionedParams {
    uint64 reserveTime;
    uint256 validatorTaxRate;
}

struct Params {
    VersionedParams versionedParams;
    uint64 paymentAccountCountLimit;
    uint64 forcedSettleTime;
    uint64 maxAutoSettleFlowCount;
    uint64 maxAutoResumeFlowCount;
    string feeDenom;
    uint256 withdrawTimeLockThreshold;
    uint64 withdrawTimeLockDuration;
}

struct OutFlow {
    string toAddress;
    uint256 rate;
    int32 status;
}

struct StreamRecord {
    string account;
    int64 crudTimestamp;
    uint256 netflowRate;
    uint256 staticBalance;
    uint256 bufferBalance;
    uint256 lockBalance;
    int32 status;
    int64 settleTimestamp;
    uint64 outFlowCount;
    uint256 frozenNetflowRate;
}

struct PaymentAccountCount {
    string owner;
    uint64 count;
}

struct DynamicBalance {
    uint256 dynamicBalance;
    StreamRecord streamRecord;
    int64 currentTimestamp;
    uint256 bankBalance;
    uint256 availableBalance;
    uint256 lockedFee;
    uint256 changeRate;
}

struct AutoSettleRecord {
    int64 timestamp;
    string addr;
}

struct DelayedWithdrawalRecord {
    string addr;
    uint256 amount;
    string from;
    int64 unlockTimestamp;
}

interface IPayment {
    /**
     * @dev createPaymentAccount defines a method for create a payment account.
     */
    function createPaymentAccount() external returns (bool success);

    /**
     * @dev deposit defines a method for deposit.
     */
    function deposit(
        string memory to,
        uint256 amount
    ) external returns (bool success);

    /**
     * @dev disableRefund defines a method for disable refund.
     */
    function disableRefund(string memory addr) external returns (bool success);

    /**
     * @dev withdraw defines a method for withdraw.
     */
    function withdraw(
        string memory from,
        uint256 amount
    ) external returns (bool success);

    /**
     * @dev updateParams defines a method for update params of modular payment.
     */
    function updateParams(
        string memory authority,
        Params memory params
    ) external returns (bool success);

    /**
     * @dev paymentAccountsByOwner defines a method for queries all payment accounts by a owner.
     */
    function paymentAccountsByOwner(
        string memory owner
    ) external view returns (string[] memory accounts);

    /**
     * @dev paymentAccount defines a method for queries a payment account by payment account address.
     */
    function paymentAccount(
        string memory addr
    ) external view returns (PaymentAccount calldata paymentAccount);

    /**
     * @dev paymentAccounts defines a method for queries all payment accounts.
     */
    function paymentAccounts(
        PageRequest calldata pagination
    )
        external
        view
        returns (
            PaymentAccount[] memory paymentAccounts,
            PageResponse calldata pageResponse
        );

    /**
     * @dev paymentAccountCount defines a method for queries the count of payment account by owner.
     */
    function paymentAccountCount(
        string memory owner
    ) external view returns (PaymentAccountCount calldata paymentAccountCount);

    /**
     * @dev paymentAccountCounts defines a method for queries all counts of payment account for all owners.
     */
    function paymentAccountCounts(
        PageRequest calldata pagination
    )
        external
        view
        returns (
            PaymentAccountCount[] memory paymentAccountCounts,
            PageResponse calldata pageResponse
        );

    /**
     * @dev params defines a method for queries the parameters of the module payment.
     */
    function params() external view returns (Params calldata params);

    /**
     * @dev paramsByTimestamp defines a method for queries the parameters of the module payment by timestamp.
     */
    function paramsByTimestamp(
        int64 timestamp
    ) external view returns (Params calldata params);

    /**
     * @dev outFlows defines a method for queries out flows by account.
     */
    function outFlows(
        string memory account
    ) external view returns (OutFlow[] memory outFlows);

    /**
     * @dev streamRecord defines a method for queries a stream record by account.
     */
    function streamRecord(
        string memory account
    ) external view returns (StreamRecord calldata streamRecord);

    /**
     * @dev streamRecords defines a method for queries all stream records.
     */
    function streamRecords(
        PageRequest calldata pagination
    )
        external
        view
        returns (
            StreamRecord[] memory streamRecords,
            PageResponse calldata pageResponse
        );

    /**
     * @dev dynamicBalance defines a method for queries dynamic balance of a payment account.
     */
    function dynamicBalance(
        string memory account
    ) external view returns (DynamicBalance calldata dynamicBalance);

    /**
     * @dev autoSettleRecords defines a method for queries all auto settle records.
     */
    function autoSettleRecords(
        PageRequest calldata pagination
    )
        external
        view
        returns (
            AutoSettleRecord[] memory autoSettleRecords,
            PageResponse calldata pageResponse
        );

    /**
     * @dev delayedWithdrawal defines a method for queries delayed withdrawal of a account.
     */
    function delayedWithdrawal(
        string memory account
    )
        external
        view
        returns (DelayedWithdrawalRecord calldata delayedWithdrawal);

    /**
     * @dev CreatePaymentAccount defines an Event emitted when a user create a payment account
     */
    event CreatePaymentAccount(address indexed creator);

    /**
     * @dev Deposit defines an Event emitted when a user deposit
     */
    event Deposit(address indexed operator);

    /**
     * @dev DisableRefund defines an Event emitted when a user disable refund
     */
    event DisableRefund(address indexed owner);

    /**
     * @dev Withdraw defines an Event emitted when a user withdraw
     */
    event Withdraw(address indexed creator);

    /**
     * @dev UpdateParams defines an Event emitted when a user update params of modular payment
     */
    event UpdateParams(address indexed creator);
}
