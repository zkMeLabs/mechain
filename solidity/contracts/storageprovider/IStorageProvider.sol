// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

// Status is the status of a storage provider.
enum Status {
    STATUS_IN_SERVICE,
    STATUS_IN_JAILED,
    STATUS_GRACEFUL_EXITING,
    STATUS_IN_MAINTENANCE,
    STATUS_FORCED_EXITING
}

// Description defines a storage provider description.
struct Description {
    // moniker defines a human-readable name for the storage provider
    string moniker;
    // identity defines an optional identity signature (ex. UPort or Keybase).
    string identity;
    // website defines an optional website link.
    string website;
    // security_contact defines an optional email for security contact.
    string security_contact;
    // details define other optional details.
    string details;
}

// StorageProvider defines the meta info of storage provider
struct StorageProvider {
    // id is the identifier of the storage provider, used in virtual group
    uint32 id;
    // operator_address defines the account address of the storage provider's operator; It also is the unique index key of sp.
    string operator_address;
    // funding_address defines one of the storage provider's accounts which is used to deposit and reward.
    string funding_address;
    // seal_address defines one of the storage provider's accounts which is used to SealObject
    string seal_address;
    // approval_address defines one of the storage provider's accounts which is used to approve use's createBucket/createObject request
    string approval_address;
    // gc_address defines one of the storage provider's accounts which is used for gc purpose.
    string gc_address;
    // maintenance_address defines one of the storage provider's accounts which is used for testing while in maintenance mode
    string maintenance_address;
    // total_deposit defines the number of tokens deposited by this storage provider for staking.
    uint256 total_deposit;
    // status defines the current service status of this storage provider
    Status status;
    // endpoint define the storage provider's network service address
    string endpoint;
    // description defines the description terms for the storage provider.
    Description description;
    // bls_key defines the bls pub key of the Storage provider for sealing object and completing migration
    string bls_key;
}

struct SpStoragePrice {
    // Storage Provider ID
    uint32 sp_id;
    // Update time as Unix timestamp in seconds
    uint256 update_time_sec;
    // Read price in wei per byte
    uint256 read_price;
    // Free read quota in bytes
    uint64 free_read_quota;
    // Store price in wei per byte
    uint256 store_price;
}

interface IStorageProvider {
    /**
     * @dev updateSPPrice defines a method for sp update storage-provider price info.
     */
    function updateSPPrice(
        uint256 readPrice,
        uint64 freeReadQuota,
        uint256 storePrice
    ) external returns (bool success);

    /**
     * @dev storageProvider queries a storage provider with specify id.
     */
    function storageProvider(
        uint32 id
    ) external view returns (StorageProvider calldata storageProvider);

    /**
     * @dev storageProviders queries a list of GetStorageProviders items.
     */
    function storageProviders(
        PageRequest calldata pagination
    )
        external
        view
        returns (
            StorageProvider[] calldata storageProviders,
            PageResponse calldata pageResponse
        );

    /**
     * @dev storageProviderByOperatorAddress queries a StorageProvider by specify operator address.
     */
    function storageProviderByOperatorAddress(
        address operatorAddress
    ) external view returns (StorageProvider calldata storageProvider);

    /**
     * @dev storageProviderPrice get the latest storage price of specific sp.
     */
    function storageProviderPrice(
        address operatorAddress
    ) external view returns (SpStoragePrice calldata spStoragePrice);

    /**
     * @dev UpdateSPPrice defines an Event emitted when a sp update storage-provider price info.
     */
    event UpdateSPPrice(address indexed storageProvider);
}
