// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

// Global virtual group family serve as a means of grouping global virtual groups.
// Each bucket must be associated with a unique global virtual group family and cannot cross families.
struct GlobalVirtualGroupFamily {
    // id is the identifier of the global virtual group family.
    uint32 id;
    // primarySpId
    uint32 primarySpId;
    // globalVirtualGroupIds is a list of identifiers of the global virtual groups associated with the family.
    uint32[] globalVirtualGroupIds;
    // virtualPaymentAddress is the payment address associated with the global virtual group family.
    address virtualPaymentAddress;
}

interface IVirtualGroup {
    /**
     * @dev createGlobalVirtualGroup defines a method for sp create a global virtual group.
     */
    function createGlobalVirtualGroup(
        uint32 familyId,
        uint32[] memory secondarySpIds,
        Coin memory deposit
    ) external returns (bool success);

    /**
     * @dev globalVirtualGroupFamilies queries all the global virtual group family.
     */
    function globalVirtualGroupFamilies(
        PageRequest calldata pagination
    ) external view returns (GlobalVirtualGroupFamily[] memory gvgFamilies, PageResponse calldata pageResponse);

    /**
     * @dev CreateBucket defines an Event emitted when a sp create a global virtual group.
     */
    event CreateGlobalVirtualGroup(
        address indexed storageProvider,
        uint256 familyId
    );
}