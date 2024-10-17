// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

interface IStorageProvider {
    /**
     * @dev updateSPPrice defines a method for sp update storage-provider price info.
     */
    function updateSPPrice(
        int256 readPrice,
        uint64 freeReadQuota,
        int256 storePrice
    ) external returns (bool success);

    /**
     * @dev UpdateSPPrice defines an Event emitted when a sp update storage-provider price info.
     */
    event UpdateSPPrice(address indexed storageProvider);
}
