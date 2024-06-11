// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

// VisibilityType is the resources public status.
enum VisibilityType {
    UnSpecified,
    PublicRead,
    Private,
    // If the bucket Visibility is inherit, it's finally set to private. If the object Visibility is inherit, it's the same as bucket.
    Inherit
}

// SourceType represents the source of resource creation, which can
// from Greenfield native or from a cross-chain transfer from BSC
enum SourceType {
    Origin,
    MirrorPending,
    BscCrossChain,
    OpCrossChain
}

// BucketStatus represents the status of a bucket. After a user successfully
// sends a CreateBucket transaction onto the chain, the status is set to 'Created'.
// When a Discontinue Object transaction is received on chain, the status is set to 'Discontinued'.
enum BucketStatus {
    Created,
    Discontinued,
    Migrating
}

struct Tag {
    string key;
    string value;
}

// Approval is the signature information returned by the Primary Storage Provider (SP) to the user
// after allowing them to create a bucket or object, which is then used for verification on the chain
// to ensure agreement between the Primary SP and the user.
struct Approval {
    // expiredHeight is the block height at which the signature expires.
    uint64 expiredHeight;
    // globalVirtualGroupFamilyId is the family id that stored.
    uint32 globalVirtualGroupFamilyId;
    // The signature needs to conform to the EIP 712 specification.
    bytes sig;
}

struct BucketInfo {
    // owner is the account address of bucket creator, it is also the bucket owner.
    address owner;
    // bucket_name is a globally unique name of bucket
    string bucketName;
    // visibility defines the highest permissions for bucket. When a bucket is public, everyone can get storage objects in it.
    VisibilityType visibility;
    // id is the unique identification for bucket.
    uint256 id;
    // source_type defines which chain the user should send the bucket management transactions to
    SourceType sourceType;
    // create_at define the block timestamp when the bucket created.
    int64 createAt;
    // payment_address is the address of the payment account
    address paymentAddress;
    // global_virtual_group_family_id defines the unique id of gvg family
    uint32 globalVirtualGroupFamilyId;
    // charged_read_quota defines the traffic quota for read in bytes per month.
    // The available read data for each user is the sum of the free read data provided by SP and
    // the ChargeReadQuota specified here.
    uint64 chargedReadQuota;
    // bucket_status define the status of the bucket.
    BucketStatus bucketStatus;
    // tags defines a list of tags the bucket has
    Tag[] tags;
    // sp_as_delegated_agent_disabled indicates that whether bucket owner disable SP as the upload agent.
    // when a bucket is created, by default, this is false, means SP is allowed to create object for delegator
    bool spAsDelegatedAgentDisabled;
}

interface IStorage {
    /**
     * @dev createBucket defines a method for create a bucket.
     */
    function createBucket(
        string memory bucketName,
        VisibilityType visibility,
        address paymentAddress,
        address primarySpAddress,
        Approval memory primarySpApproval,
        uint64 chargedReadQuota
    ) external returns (bool success);

    /**
     * @dev listBuckets queries all the buckets.
     */
    function listBuckets(
        PageRequest calldata pagination
    ) external view returns (BucketInfo[] memory bucketInfos, PageResponse calldata pageResponse);

    /**
     * @dev Send defines an Event emitted when a given amount of tokens send fromAddress to toAddress
     */
    event CreateBucket(
        address indexed creator,
        address indexed paymentAddress,
        address indexed primarySpAddress,
        uint256 id
    );
}