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
     * @dev deleteGlobalVirtualGroup defines a method for sp delete a global virtual group.
     */
    function deleteGlobalVirtualGroup(
        uint32 globalVirtualGroupId
    ) external returns (bool success);

    /**
     * @dev globalVirtualGroupFamily queries the global virtual group family by family id.
     */
    function globalVirtualGroupFamily(
        uint32 familyId
    ) external view returns (GlobalVirtualGroupFamily memory gvgfamily);

    /**
     * @dev globalVirtualGroupFamilies queries all the global virtual group family.
     */
    function globalVirtualGroupFamilies(
        PageRequest calldata pagination
    )
        external
        view
        returns (
            GlobalVirtualGroupFamily[] memory gvgFamilies,
            PageResponse calldata pageResponse
        );

    /**
     * @dev swapOut defines a method for sp to swap out.
     */
    function swapOut(
        uint32 gvgFamilyId,
        uint32[] memory gvgIds,
        uint32 successorSpId,
        Approval memory successorSpApproval
    ) external returns (bool success);

    /**
     * @dev completeSwapOut defines a method for sp complete to swap out.
     */
    function completeSwapOut(
        uint32 gvgFamilyId,
        uint32[] memory gvgIds
    ) external returns (bool success);

    /**
     * @dev spExit defines a method for sp complete to swap out.
     */
    function spExit() external returns (bool success);

    /**
     * @dev completeSPExit defines a method for sp complete to exit.
     */
    function completeSPExit(
        string memory storageProvider,
        string memory operator
    ) external returns (bool success);

    /**
     * @dev deposit defines a method to deposit more tokens for the objects stored on it.
     */
    function deposit(
        uint32 globalVirtualGroupId,
        Coin memory deposit
    ) external returns (bool success);

    /**
     * @dev reserveSwapIn defines a method to reserve swap in.
     */
    function reserveSwapIn(
        uint32 targetSpId,
        uint32 gvgFamilyId,
        uint32 globalVirtualGroupId
    ) external returns (bool success);

    /**
     * @dev completeSwapIn defines a method to complete swap in.
     */
    function completeSwapIn(
        uint32 gvgFamilyId,
        uint32 globalVirtualGroupId
    ) external returns (bool success);

    /**
     * @dev cancelSwapIn defines a method to cancel swap in.
     */
    function cancelSwapIn(
        uint32 gvgFamilyId,
        uint32 globalVirtualGroupId
    ) external returns (bool success);

    /**
     * @dev CreateGlobalVirtualGroup defines an Event emitted when a sp create a global virtual group.
     */
    event CreateGlobalVirtualGroup(
        address indexed storageProvider,
        uint256 familyId
    );

    /**
     * @dev DeleteGlobalVirtualGroup defines an Event emitted when a sp delete a global virtual group.
     */
    event DeleteGlobalVirtualGroup(address indexed storageProvider);

    /**
     * @dev SwapOut defines an Event emitted when a sp swap out.
     */
    event SwapOut(address indexed storageProvider, uint256 familyId);

    /**
     * @dev SPExit defines an Event emitted when a sp exit.
     */
    event SPExit(address indexed storageProvider);

    /**
     * @dev CompleteSPExit defines an Event emitted when a sp complete to exit.
     */
    event CompleteSPExit(
        address indexed storageProvider,
        address indexed operator
    );

    /**
     * @dev Deposit defines an Event emitted when a sp to deposit more tokens for the objects stored on it.
     */
    event Deposit(address indexed storageProvider);

    /**
     * @dev ReserveSwapIn defines an Event emitted when a sp to reserve swap in.
     */
    event ReserveSwapIn(address indexed storageProvider);

    /**
     * @dev CompleteSwapIn defines an Event emitted when a sp to complete swap in.
     */
    event CompleteSwapIn(address indexed storageProvider);

    /**
     * @dev CancelSwapIn defines an Event emitted when a sp to cancel swap in.
     */
    event CancelSwapIn(address indexed storageProvider);
}
