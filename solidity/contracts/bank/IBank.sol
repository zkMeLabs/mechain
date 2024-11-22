// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

/**
 * @dev Output defines the send coins info.
 */
struct Output {
    address toAddress;
    Coin[] amount;
}

/**
 * @dev SendEnabled maps coin denom to a send_enabled status (whether a denom is spendable).
 */
struct SendEnabled {
    string denom;
    bool enabled;
}

/**
 * @dev Params defines the parameters for the bank module.
 */
struct Params {
    SendEnabled[] sendEnabled;
    bool defaultSendEnabled;
}
/**
 * @dev DenomUnit represents a struct that describes a given
 * denomination unit of the basic token.
 */
struct DenomUnit {
    // denom represents the string name of the given denom unit (e.g uatom).
    string denom;
    // exponent represents power of 10 exponent that one must
    // raise the base_denom to in order to equal the given DenomUnit's denom
    // 1 denom = 1^exponent base_denom
    // (e.g. with a base_denom of uatom, one can create a DenomUnit of 'atom' with
    // exponent = 6, thus: 1 atom = 10^6 uatom).
    uint32 exponent;
    // aliases is a list of string aliases for the given denom
    string[] aliases;
}

/**
 * @dev Metadata represents a struct that describes a basic token.
 */
struct Metadata {
    string description;
    // denomUnits represents the list of DenomUnit's for a given coin
    DenomUnit[] denomUnits;
    // base represents the base denom (should be the DenomUnit with exponent = 0).
    string base;
    // display indicates the suggested denom that should be
    // displayed in clients.
    string display;
    // name defines the name of the token (eg: Cosmos Atom)
    string name;
    // symbol is the token symbol usually shown on exchanges (eg: ATOM). This can
    // be the same as the display.
    string symbol;
    // URI to a document (on or off-chain) that contains additional information. Optional.
    string uri;
    // URIHash is a sha256 hash of a document pointed by URI. It's used to verify that
    // the document didn't change. Optional.
    string uriHash;
}

/**
 * @dev DenomOwner defines structure representing an account that owns or holds a
 * particular denominated token. It contains the account address and account
 * balance of the denominated token.
 */
struct DenomOwner {
    address accountAddress;
    Coin balance;
}

interface IBank {
    /**
     * @dev send defines a method for sending coins from one account to another account.
     */
    function send(
        address toAddress,
        Coin[] memory amount
    ) external returns (bool success);

    /**
     * @dev multiSend defines a method for sending coins from some accounts to other accounts.
     */
    function multiSend(Output[] memory outputs) external returns (bool success);

    /**
     * @dev balance queries the balance of a single coin for a single account.
     */
    function balance(
        address accountAddress,
        string memory denom
    ) external view returns (Coin memory balance);

    /**
     * @dev allBalances queries the balance of all coins for a single account.
     */
    function allBalances(
        address accountAddress,
        PageRequest memory pageRequest
    )
        external
        view
        returns (Coin[] memory balances, PageResponse memory pageResponse);

    /**
     * @dev totalSupply queries the total supply of all coins.
     */
    function totalSupply(
        PageRequest memory pageRequest
    )
        external
        view
        returns (Coin[] memory supply, PageResponse memory pageResponse);

    /**
     * @dev spendableBalances queries the spenable balance of all coins for a single account.
     */
    function spendableBalances(
        address accountAddress,
        PageRequest memory pageRequest
    )
        external
        view
        returns (Coin[] memory balances, PageResponse memory pageResponse);

    /**
     * @dev spendableBalanceByDenom queries an account's spendable balance for a specific denom.
     */
    function spendableBalanceByDenom(
        address accountAddress,
        string memory denom
    ) external view returns (Coin memory balance);

    /**
     * @dev supplyOf queries the supply of a single coin.
     */
    function supplyOf(
        string memory denom
    ) external view returns (Coin memory amount);

    /**
     * @dev params queries the parameters of x/bank module.
     */
    function params() external view returns (Params memory params);

    /**
     * @dev denomMetadata queries the client metadata of a given coin denomination.
     */
    function denomMetadata(
        string memory denom
    ) external view returns (Metadata memory metadata);

    /**
     * @dev denomsMetadata queries the client metadata for all registered coin denominations.
     */
    function denomsMetadata(
        PageRequest memory pageRequest
    )
        external
        view
        returns (Metadata[] memory metadatas, PageResponse memory pageResponse);

    /**
     * @dev DenomOwners queries for all account addresses that own a particular token denomination.
     */
    function denomOwners(
        string memory denom,
        PageRequest memory pageRequest
    )
        external
        view
        returns (
            DenomOwner[] memory denomOwners,
            PageResponse memory pageResponse
        );

    /**
     * @dev sendEnabled queries for SendEnabled entries.
     */
    function sendEnabled(
        string[] memory denoms,
        PageRequest memory pageRequest
    )
        external
        view
        returns (
            SendEnabled[] memory sendEnableds,
            PageResponse memory pageResponse
        );

    /**
     * @dev Send defines an Event emitted when a given amount of tokens send fromAddress to toAddress
     */
    event Send(
        address indexed fromAddress,
        address indexed toAddress,
        string amount
    );

    /**
     * @dev MultiSend defines an Event emitted when a given amount of tokens send from a address to some other addresses
     */
    event MultiSend(address indexed fromAddress, string amount);
}
