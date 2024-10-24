// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";
import "../virtualgroup/Types.sol";
// VisibilityType is the resources public status.
enum VisibilityType {
    UnSpecified,
    PublicRead,
    Private,
    // If the bucket Visibility is inherit, it's finally set to private. If the object Visibility is inherit, it's the same as bucket.
    Inherit
}

// SourceType represents the source of resource creation, which can
// from Mechain native or from a cross-chain transfer from BSC
enum SourceType {
    Origin,
    MirrorPending,
    BscCrossChain,
    OpCrossChain,
    PolygonCrossChain,
    ScrollCrossChain,
    LineaCrossChain,
    MantleCrossChain,
    ArbitrumCrossChain,
    OptimismCrossChain
}

// BucketStatus represents the status of a bucket. After a user successfully
// sends a CreateBucket transaction onto the chain, the status is set to 'Created'.
// When a Discontinue Object transaction is received on chain, the status is set to 'Discontinued'.
enum BucketStatus {
    Created,
    Discontinued,
    Migrating
}

// RedundancyType represents the redundancy algorithm type for object data,
// which can be either multi-replica or erasure coding.
enum RedundancyType {
    EcType,
    ReplicaType
}
// ObjectStatus represents the creation status of an object. After a user successfully
// sends a CreateObject transaction onto the chain, the status is set to 'Created'.
// After the Primary Service Provider successfully sends a Seal Object transaction onto
// the chain, the status is set to 'Sealed'. When a Discontinue Object transaction is
// received on chain, the status is set to 'Discontinued'.
enum ObjectStatus {
    Created,
    Sealed,
    Discontinued
}

struct Tag {
    string key;
    string value;
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

struct BucketExtraInfo {
    bool IsRateLimited;
    uint256 FlowRateLimit;
    uint256 CurrentFlowRate;
}

struct ObjectInfo {
    // owner is the object owner
    address owner;
    // creator is the address of the uploader, it always be same as owner address
    address creator;
    // bucketName is the name of the bucket
    string bucketName;
    // objectName is the name of object
    string objectName;
    // id is the unique identifier of object
    uint256 id;
    uint32 localVirtualGroupId;
    // payloadSize is the total size of the object payload
    uint64 payloadSize;
    // visibility defines the highest permissions for object. When an object is public, everyone can access it.
    VisibilityType visibility;
    // contentType define the format of the object which should be a standard MIME type.
    string contentType;
    // createAt define the block timestamp when the object is created
    int64 createAt;
    // objectStatus define the upload status of the object.
    ObjectStatus objectStatus;
    // redundancyType define the type of the redundancy which can be multi-replication or EC.
    RedundancyType redundancyType;
    // sourceType define the source of the object.
    SourceType sourceType;
    // checksums define the root hash of the pieces which stored in a SP.
    // add omit tag to omit the field when converting to NFT metadata
    string[] checksums;
    // tags defines a list of tags the object has
    Tag[] tags;
    // isUpdating indicates whether a object is being updated.
    bool isUpdating;
    // updatedAt define the block timestamp when the object is updated. Will not be visible until object is re-sealed.
    int64 updatedAt;
    // updatedBy defined the account address of updater(if there is). Will not be visible until object is re-sealed.
    address updatedBy;
    // version define the version of object
    int64 version;
}

struct GroupInfo {
    // owner is the owner of the group. It can not changed once it created.
    address owner;
    // group_name is the name of group which is unique under an account.
    string groupName;
    // sourceType define the source of the group.
    SourceType sourceType;
    // id is the unique identifier of group
    uint256 id;
    // extra is used to store extra info for the group
    string extra;
    // tags defines a list of tags the group has
    Tag[] tags;
}

struct GroupMember {
    // id is an unique u256 sequence for each group member. It also be used as NFT tokenID
    uint256 id;
    // group_id is the unique id of the group
    uint256 groupId;
    // member is the account address of the member
    address member;
    // expiration_time defines the timestamp(UNIX) of the member expiration
    int64 expirationTime;
}

struct GVGMapping {
    uint32 srcGlobalVirtualGroupId;
    uint32 dstGlobalVirtualGroupId;
    bytes secondarySpBlsSignature;
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
     * @dev updateBucketInfo defines a method for update a bucket.
     */
    function updateBucketInfo(
        string memory bucketName,
        VisibilityType visibility,
        address paymentAddress,
        int128 chargedReadQuota
    ) external returns (bool success);

    /**
     * @dev deleteBucket defines a method for delete a bucket.
     */
    function deleteBucket(
        string memory bucketName
    ) external returns (bool success);

    /**
     * @dev discontinueBucket defines a method for discontinue a bucket.
     */
    function discontinueBucket(
        string memory bucketName,
        string memory reason
    ) external returns (bool success);

    /**
     * @dev completeMigrateBucket defines a method for complete migrate a bucket.
     */
    function completeMigrateBucket(
        string memory bucketName,
        uint32 gvgFamilyId,
        GVGMapping[] memory gvgMapping
    ) external returns (bool success);

    /**
     * @dev rejectMigrateBucket defines a method for reject migrate a bucket.
     */
    function rejectMigrateBucket(
        string memory bucketName
    ) external returns (bool success);

    /**
     * @dev createObject defines a method for create a object.
     */
    function createObject(
        string memory bucketName,
        string memory objectName,
        uint64 payloadSize,
        VisibilityType visibility,
        string memory contentType,
        Approval memory primarySpApproval,
        string[] memory expectChecksums,
        RedundancyType redundancyType
    ) external returns (bool success);

    /**
     * @dev listBuckets queries all the buckets.
     */
    function listBuckets(
        PageRequest calldata pagination
    )
        external
        view
        returns (
            BucketInfo[] memory bucketInfos,
            PageResponse calldata pageResponse
        );

    /**
     * @dev listObjects queries all the objects.
     */
    function listObjects(
        PageRequest calldata pagination,
        string memory bucketName
    )
        external
        view
        returns (
            ObjectInfo[] memory objectInfos,
            PageResponse calldata pageResponse
        );

    /**
     * @dev listGroups queries all the groups.
     */
    function listGroups(
        PageRequest calldata pagination,
        address groupOwner
    )
        external
        view
        returns (
            GroupInfo[] memory groupInfos,
            PageResponse calldata pageResponse
        );

    /**
     * @dev sealObject defines a method for seal a object.
     */
    function sealObject(
        address sealAddress,
        string memory bucketName,
        string memory objectName,
        uint32 globalVirtualGroupId,
        string memory secondarySpBlsAggSignatures
    ) external returns (bool success);

    /**
     * @dev sealObjectV2 defines a method for seal a object with IsAgentUpload.
     */
    function sealObjectV2(
        address sealAddress,
        string memory bucketName,
        string memory objectName,
        uint32 globalVirtualGroupId,
        string memory secondarySpBlsAggSignatures,
        string[] memory expectChecksums
    ) external returns (bool success);

    /**
     * @dev rejectSealObject defines a method for reject seal a object.
     */
    function rejectSealObject(
        string memory bucketName,
        string memory objectName
    ) external returns (bool success);

    /**
     * @dev delegateCreateObject defines a method for delegate create a object.
     */
    function delegateCreateObject(
        string memory creator,
        string memory bucketName,
        string memory objectName,
        uint64 payloadSize,
        string memory contentType,
        VisibilityType visibility,
        string[] memory expectChecksums,
        RedundancyType redundancyType
    ) external returns (bool success);

    /**
     * @dev updateObjectInfo defines a method for update object visibility.
     */
    function updateObjectInfo(
        string memory bucketName,
        string memory objectName,
        VisibilityType visibility
    ) external returns (bool success);

    /**
     * @dev delegateUpdateObjectContent defines a method for delegate update a object content.
     */
    function delegateUpdateObjectContent(
        string memory updater,
        string memory bucketName,
        string memory objectName,
        uint64 payloadSize,
        string memory contentType,
        string[] memory expectChecksums
    ) external returns (bool success);

    /**
     * @dev createGroup defines a method for create a group.
     */
    function createGroup(
        string memory groupName,
        string memory extra
    ) external returns (bool success);

    /**
     * @dev headBucket queries the bucket's info.
     */
    function headBucket(
        string memory bucketName
    )
        external
        view
        returns (
            BucketInfo memory bucketInfo,
            BucketExtraInfo memory bucketExtraInfo
        );

    /**
     * @dev headGroup queries the group's info.
     */
    function headGroup(
        address groupOwner,
        string memory groupName
    ) external view returns (GroupInfo memory groupInfo);

    /**
     * @dev updateGroup defines a method for update a group's member.
     */
    function updateGroup(
        address groupOwner,
        string memory groupName,
        address[] memory membersToAdd,
        int64[] memory expirationTime,
        address[] memory membersToDelete
    ) external returns (bool success);

    /**
     * @dev headGroupMember queries the group member's info.
     */
    function headGroupMember(
        address member,
        address groupOwner,
        string memory groupName
    ) external view returns (GroupMember memory groupMember);

    /**
     * @dev deleteGroup defines a method for delete a group.
     */
    function deleteGroup(
        string memory groupName
    ) external returns (bool success);

    /**
     * @dev renewGroupMember defines a method for update the expire time of group member.
     */
    function renewGroupMember(
        address groupOwner,
        string memory groupName,
        address[] memory members,
        int64[] memory expirationTime
    ) external returns (bool success);

    /**
     * @dev setTagForGroup defines a method for set tags for the given group.
     */
    function setTagForGroup(
        string memory groupName,
        Tag[] memory tags
    ) external returns (bool success);

    /**
     * @dev headObject queries the object's info.
     */
    function headObject(
        string memory bucketName,
        string memory objectName
    )
        external
        view
        returns (
            ObjectInfo memory objectInfo,
            GlobalVirtualGroup memory globalVirtualGroup
        );

    /**
     * @dev headObjectById queries the object's info.
     */
    function headObjectById(
        string memory objectId
    )
        external
        view
        returns (
            ObjectInfo memory objectInfo,
            GlobalVirtualGroup memory globalVirtualGroup
        );

    /**
     * @dev CreateBucket defines an Event emitted when a user create a bucket
     */
    event CreateBucket(
        address indexed creator,
        address indexed paymentAddress,
        address indexed primarySpAddress,
        uint256 id
    );

    /**
     * @dev UpdateBucketInfo defines an Event emitted when a user update a bucket
     */
    event UpdateBucketInfo(
        address indexed operator,
        bytes32 indexed bucketName,
        address indexed paymentAddress,
        uint8 visibility
    );

    /**
     * @dev DeleteBucket defines an Event emitted when a user delete a bucket
     */
    event DeleteBucket(address indexed creator);

    /**
     * @dev DiscontinueBucket defines an Event emitted when a user discontinue a bucket
     */
    event DiscontinueBucket(address indexed creator);

    /**
     * @dev CompleteMigrateBucket defines an Event emitted when a user complete migrate a bucket
     */
    event CompleteMigrateBucket(address indexed creator);

    /**
     * @dev RejectMigrateBucket defines an Event emitted when a user reject migrate a bucket
     */
    event RejectMigrateBucket(address indexed operator);

    /**
     * @dev CreateObject defines an Event emitted when a user create a object
     */
    event CreateObject(address indexed creator, uint256 id);

    /**
     * @dev Transfer defines an Event emitted when a transfer a object nft
     */
    event Transfer(
        address indexed from,
        address indexed to,
        uint256 indexed tokenId
    );

    /**
     * @dev SealObject defines an Event emitted when a user seal a object
     */
    event SealObject(address indexed creator, address indexed sealAddress);

    /**
     * @dev SealObjectV2 defines an Event emitted when a user seal a object with IsAgentUpload
     */
    event SealObjectV2(address indexed creator, address indexed sealAddress);

    /**
     * @dev RejectSealObject defines an Event emitted when a sp reject seal a object
     */
    event RejectSealObject(address indexed creator, string indexed objectName);

    /**
     * @dev DelegateCreateObject defines an Event emitted when a user delegate create a object
     */
    event DelegateCreateObject(
        address indexed creator,
        string indexed objectName
    );

    /**
     * @dev UpdateObjectInfo defines an Event emitted when a user update object visibility
     */
    event UpdateObjectInfo(address indexed creator);

    /**
     * @dev DelegateUpdateObjectContent defines an Event emitted when a user delegate update a object content
     */
    event DelegateUpdateObjectContent(
        address indexed operator,
        string indexed objectName
    );

    /**
     * @dev CreateGroup defines an Event emitted when a user create a group
     */
    event CreateGroup(address indexed creator, uint256 id);

    /**
     * @dev UpdateGroup defines an Event emitted when a user update a group's member
     */
    event UpdateGroup(address indexed creator);

    /**
     * @dev DeleteGroup defines an Event emitted when a user delete a group
     */
    event DeleteGroup(address indexed creator);

    /**
     * @dev RenewGroupMember defines an Event emitted when a user renew group member
     */
    event RenewGroupMember(address indexed creator);

    /**
     * @dev SetTagForGroup defines an Event emitted when a user set tags for the given group
     */
    event SetTagForGroup(address indexed creator);
}
