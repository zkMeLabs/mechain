// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

/**
 * @dev GrantData gives permissions to execute the provide method with expiration time.
 */
struct GrantData {
    string authorization;
    int64 expiration;
}

/**
 * @dev GrantAuthorization extends a grant with both the addresses of the grantee and granter.
 * It is used in genesis.proto and query.proto
 */
struct GrantAuthorization {
    address granter;
    address grantee;
    string authorization;
    int64 expiration;
}

interface IAuthz {
    /**
     * @dev grant implements the MsgServer.Grant method to create a new grant.
     */
    function grant(
        address grantee,
        string memory authzType,
        string memory authorization,
        Coin[] memory limit,
        int64 expiration
    ) external returns (bool success);

    /**
     * @dev exec attempts to execute the provided messages using
     * authorizations granted to the grantee. Each message should have only
     * one signer corresponding to the granter of the authorization.
     */
    function revoke(
        address grantee,
        string memory msgTypeUrl
    ) external returns (bool success);

    /**
     * @dev exec revokes any authorization corresponding to the provided method name on the
     * granter's account that has been granted to the grantee.
     */
    function exec(string[] memory msgs) external returns (bool success);

    /**
     * @dev grants returns list of `Authorization`, granted to the grantee by the granter.
     */
    function grants(
        address granter,
        address grantee,
        string memory msgTypeUrl,
        PageRequest calldata pagination
    )
        external
        view
        returns (
            GrantData[] calldata grants,
            PageResponse calldata pageResponse
        );

    /**
     * @dev granterGrants returns list of `GrantAuthorization`, granted by granter.
     */
    function granterGrants(
        address granter,
        PageRequest calldata pagination
    )
        external
        view
        returns (
            GrantAuthorization[] calldata grants,
            PageResponse calldata pageResponse
        );

    /**
     * @dev granteeGrants returns a list of `GrantAuthorization` by grantee.
     */
    function granteeGrants(
        address grantee,
        PageRequest calldata pagination
    )
        external
        view
        returns (
            GrantAuthorization[] calldata grants,
            PageResponse calldata pageResponse
        );

    /**
     * @dev Grant defines an Event emitted when create a new grant
     */
    event Grant(
        address indexed granter,
        address indexed grantee,
        string authzType
    );

    /**
     * @dev Revoke defines an Event emitted when create a granter revoke authorization for grantee
     */
    event Revoke(
        address indexed granter,
        address indexed grantee,
        string msgTypeUrl
    );

    /**
     * @dev Exec defines an Event emitted when execute the provided messages using
     * authorizations granted to the grantee
     */
    event Exec(address indexed grantee);
}
