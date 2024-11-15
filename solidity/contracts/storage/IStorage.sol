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
    bool isRateLimited;
    uint256 flowRateLimit;
    uint256 currentFlowRate;
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

// Params defines the parameters for the module.
struct Params {
    // VersionedParams versionedParams;
    VersionedParams versionedParams;
    // maxPayloadSize is the maximum size of the payload, default: 2G
    uint64 maxPayloadSize;
    // relayer fee for the mirror bucket tx to bsc
    string bscMirrorBucketRelayerFee;
    // relayer fee for the ACK or FAIL_ACK package of the mirror bucket tx to bsc
    string bscMirrorBucketAckRelayerFee;
    // relayer fee for the mirror object tx to bsc
    string bscMirrorObjectRelayerFee;
    // Relayer fee for the ACK or FAIL_ACK package of the mirror object tx to bsc
    string bscMirrorObjectAckRelayerFee;
    // relayer fee for the mirror object tx to bsc
    string bscMirrorGroupRelayerFee;
    // Relayer fee for the ACK or FAIL_ACK package of the mirror object tx to bsc
    string bscMirrorGroupAckRelayerFee;
    // The maximum number of buckets that can be created per account
    uint32 maxBucketsPerAccount;
    // The window to count the discontinued objects or buckets
    uint64 discontinueCountingWindow;
    // The max objects can be requested in a window
    uint64 discontinueObjectMax;
    // The max buckets can be requested in a window
    uint64 discontinueBucketMax;
    // The object will be deleted after the confirm period in seconds
    int64 discontinueConfirmPeriod;
    // The max delete objects in each end block
    uint64 discontinueDeletionMax;
    // The max number for deleting policy in each end block
    uint64 stalePolicyCleanupMax;
    // The min interval for making quota smaller in seconds
    uint64 minQuotaUpdateInterval;
    // the max number of local virtual group per bucket
    uint32 maxLocalVirtualGroupNumPerBucket;
    // relayer fee for the mirror bucket tx to op chain
    string opMirrorBucketRelayerFee;
    // relayer fee for the ACK or FAIL_ACK package of the mirror bucket tx to op chain
    string opMirrorBucketAckRelayerFee;
    // relayer fee for the mirror object tx to op chain
    string opMirrorObjectRelayerFee;
    // Relayer fee for the ACK or FAIL_ACK package of the mirror object tx to op chain
    string opMirrorObjectAckRelayerFee;
    // relayer fee for the mirror object tx to op chain
    string opMirrorGroupRelayerFee;
    // Relayer fee for the ACK or FAIL_ACK package of the mirror object tx to op chain
    string opMirrorGroupAckRelayerFee;
    // relayer fee for the mirror bucket tx to polygon chain
    string polygonMirrorBucketRelayerFee;
    // relayer fee for the ACK or FAIL_ACK package of the mirror bucket tx to polygon chain
    string polygonMirrorBucketAckRelayerFee;
    // Add the rest of the fields similarly
}

// VersionedParams defines the parameters for the storage module with multi version, each version store with different timestamp.
struct VersionedParams {
    // max_segment_size is the maximum size of a segment. default: 16M
    uint64 maxSegmentSize;
    // redundant_data_check_num is the num of data chunks of EC redundancy algorithm
    uint32 redundantDataChunkNum;
    // redundant_data_check_num is the num of parity chunks of EC redundancy algorithm
    uint32 redundantParityChunkNum;
    // min_charge_size is the minimum charge size of the payload, objects smaller than this size will be charged as this size
    uint64 minChargeSize;
}

struct Principal {
    int32 principalType;
    // When the type is an account, its value is sdk.AccAddress().String();
    // when the type is a group, its value is math.Uint().String()
    string value;
}

struct Statement {
    int32 effect;
    int32[] actions;
    string[] resources;
    int64 expirationTime;
    uint64 limitSize;
}

struct Policy {
    uint256 id;
    Principal principal;
    int32 resourceType;
    uint256 resourceId;
    Statement[] statements;
    int64 expirationTime;
}

struct Trait {
    string traitType;
    string value;
}

struct BucketMetaData {
    string description;
    string externalUrl;
    string bucketName;
    string image;
    Trait[] attributes;
}

struct ObjectMetaData {
    string description;
    string externalUrl;
    string objectName;
    string image;
    Trait[] attributes;
}

struct GroupMetaData {
    string description;
    string externalUrl;
    string groupName;
    string image;
    Trait[] attributes;
}

struct ShadowObjectInfo {
    string operator;
    uint256 id;
    string contentType;
    uint64 payloadSize;
    string[] checksums;
    int64 updatedAt;
    int64 version;
}

struct LocalVirtualGroup {
    uint32 id;
    uint32 globalVirtualGroupId;
    uint64 storedSize;
    uint64 totalChargeSize;
}

struct InternalBucketInfo {
    int64 priceTime;
    uint64 totalChargeSize;
    LocalVirtualGroup[] localVirtualGroups;
    uint32 nextLocalVirtualGroupId;
}

struct IsPriceChanged {
    bool changed;
    uint256 currentReadPrice;
    uint256 currentPrimaryStorePrice;
    uint256 currentSecondaryStorePrice;
    uint256 currentValidatorTaxRate;
    uint256 newReadPrice;
    uint256 newPrimaryStorePrice;
    uint256 newSecondaryStorePrice;
    uint256 newValidatorTaxRate;
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
     * @dev mirrorBucket defines a method for mirror a bucket.
     */
    function mirrorBucket(
        uint256 bucketId,
        string memory bucketName,
        uint32 destChainId
    ) external returns (bool success);

    /**
     * @dev migrateBucket defines a method for migrate a bucket.
     */
    function migrateBucket(
        string memory bucketName,
        uint32 dstPrimarySpId,
        Approval memory dstPrimarySpApproval
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
     * @dev cancelMigrateBucket defines a method for cancel migrate a bucket.
     */
    function cancelMigrateBucket(
        string memory bucketName
    ) external returns (bool success);

    /**
     * @dev setBucketFlowRateLimit defines a method for set the bucket flow rate limit.
     */
    function setBucketFlowRateLimit(
        string memory bucketName,
        string memory bucketOwner,
        string memory paymentAddress,
        uint256 flowRateLimit
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
     * @dev cancelCreateObject defines a method for cancel to create a object.
     */
    function cancelCreateObject(
        string memory bucketName,
        string memory objectName
    ) external returns (bool success);

    /**
     * @dev copyObject defines a method for copy a object.
     */
    function copyObject(
        string memory srcBucketName,
        string memory dstBucketName,
        string memory srcObjectName,
        string memory dstObjectName,
        Approval memory dstPrimarySpApproval
    ) external returns (bool success);

    /**
     * @dev deleteObject defines a method for delete a object.
     */
    function deleteObject(
        string memory bucketName,
        string memory objectName
    ) external returns (bool success);

    /**
     * @dev mirrorObject defines a method for mirror a object.
     */
    function mirrorObject(
        uint256 objectId,
        string memory bucketName,
        string memory objectName,
        uint32 destChainId
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
     * @dev listObjectsByBucketId queries a list of object items under the bucket.
     */
    function listObjectsByBucketId(
        PageRequest calldata pagination,
        string memory bucketId
    )
        external
        view
        returns (
            ObjectInfo[] memory objectInfos,
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
     * @dev updateObjectContent defines a method for update a object content.
     */
    function updateObjectContent(
        string memory bucketName,
        string memory objectName,
        uint64 payloadSize,
        string memory contentType,
        string[] memory expectChecksums
    ) external returns (bool success);

    /**
     * @dev discontinueObject defines a method for discontinue a object.
     */
    function discontinueObject(
        string memory bucketName,
        uint256[] memory objectIds,
        string memory reason
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
            BucketInfo calldata bucketInfo,
            BucketExtraInfo calldata bucketExtraInfo
        );

    /**
     * @dev headBucketExtra queries a bucket extra info (with gvg bindings and price time) with specify name.
     */
    function headBucketExtra(
        string memory bucketName
    ) external view returns (InternalBucketInfo calldata extraInfo);

    /**
     * @dev headBucketById queries the bucket's info by id.
     */
    function headBucketById(
        string memory bucketId
    )
        external
        view
        returns (
            BucketInfo calldata bucketInfo,
            BucketExtraInfo calldata bucketExtraInfo
        );

    /**
     * @dev headBucketNFT queries a bucket with EIP712 standard metadata info.
     */
    function headBucketNFT(
        string memory tokenId
    ) external view returns (BucketMetaData calldata bucketMetaData);

    /**
     * @dev headObjectNFT queries a object with EIP712 standard metadata info.
     */
    function headObjectNFT(
        string memory tokenId
    ) external view returns (ObjectMetaData calldata objectMetaData);

    /**
     * @dev headGroupNFT queries a group with EIP712 standard metadata info.
     */
    function headGroupNFT(
        string memory tokenId
    ) external view returns (GroupMetaData calldata groupMetaData);

    /**
     * @dev headGroup queries the group's info.
     */
    function headGroup(
        address groupOwner,
        string memory groupName
    ) external view returns (GroupInfo calldata groupInfo);

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
     * @dev updateGroupExtra defines a method for update a group's extra.
     */
    function updateGroupExtra(
        address groupOwner,
        string memory groupName,
        string memory extra
    ) external returns (bool success);

    /**
     * @dev headGroupMember queries the group member's info.
     */
    function headGroupMember(
        address member,
        address groupOwner,
        string memory groupName
    ) external view returns (GroupMember calldata groupMember);

    /**
     * @dev queryPolicyForGroup queries the group's policy.
     */
    function queryPolicyForGroup(
        string memory resource,
        uint256 groupId
    ) external view returns (Policy calldata policy);

    /**
     * @dev queryPolicyForAccount queries the account's policy.
     */
    function queryPolicyForAccount(
        string memory resource,
        string memory principalAddr
    ) external view returns (Policy calldata policy);

    /**
     * @dev queryPolicyById queries a policy by policy id.
     */
    function queryPolicyById(
        string memory policyId
    ) external view returns (Policy calldata policy);

    /**
     * @dev queryLockFee queries lock fee for storing an object.
     */
    function queryLockFee(
        string memory primarySpAddress,
        int64 createAt,
        uint64 payloadSize
    ) external view returns (uint256 amount);

    /**
     * @dev queryIsPriceChanged queries whether read and storage prices changed for the bucket.
     */
    function queryIsPriceChanged(
        string memory bucketName
    ) external view returns (IsPriceChanged calldata isPriceChanged);

    /**
     * @dev queryQuotaUpdateTime queries quota update time for the bucket.
     */
    function queryQuotaUpdateTime(
        string memory bucketName
    ) external view returns (int64 updateAt);

    /**
     * @dev queryGroupMembersExist queries whether some members are in the group.
     */
    function queryGroupMembersExist(
        string memory groupId,
        string[] memory members
    )
        external
        view
        returns (string[] memory checkMembers, bool[] memory exists);

    /**
     * @dev queryGroupsExist queries whether some groups are exist.
     */
    function queryGroupsExist(
        string memory groupOwner,
        string[] memory groupNames
    )
        external
        view
        returns (string[] memory checkGroupNames, bool[] memory exists);

    /**
     * @dev queryGroupsExistById queries whether some groups are exist by id.
     */
    function queryGroupsExistById(
        string[] memory groupIds
    )
        external
        view
        returns (string[] memory checkGroupIds, bool[] memory exists);

    /**
     * @dev queryPaymentAccountBucketFlowRateLimit queries the flow rate limit of a bucket for a payment account.
     */
    function queryPaymentAccountBucketFlowRateLimit(
        string memory paymentAccount,
        string memory bucketOwner,
        string memory bucketName
    ) external view returns (bool isSet, uint256 flowRateLimit);

    /**
     * @dev verifyPermission queries a list of VerifyPermission items.
     */
    function verifyPermission(
        string memory bucketName,
        string memory objectName,
        int32 actionType
    ) external view returns (int32 effect);

    /**
     * @dev deleteGroup defines a method for delete a group.
     */
    function deleteGroup(
        string memory groupName
    ) external returns (bool success);

    /**
     * @dev leaveGroup defines a method for leave a group.
     */
    function leaveGroup(
        address member,
        address groupOwner,
        string memory groupName
    ) external returns (bool success);

    /**
     * @dev mirrorGroup defines a method for mirror a group.
     */
    function mirrorGroup(
        uint256 groupId,
        string memory groupName,
        uint32 destChainId
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
     * @dev setTag defines a method for set tags for the given group/bucket/object.
     */
    function setTag(
        string memory resource,
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
            ObjectInfo calldata objectInfo,
            GlobalVirtualGroup calldata globalVirtualGroup
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
            ObjectInfo calldata objectInfo,
            GlobalVirtualGroup calldata globalVirtualGroup
        );

    /**
     * @dev headShadowObject queries a shadow object with specify name.
     */
    function headShadowObject(
        string memory bucketName,
        string memory objectName
    ) external view returns (ShadowObjectInfo calldata objectInfo);

    /**
     * @dev queryParamsByTimestamp queries the parameters of the module by timestamp.
     */
    function queryParamsByTimestamp(
        int64 timestamp
    ) external view returns (Params calldata params);

    /**
     * @dev params queries the storage params.
     */
    function params() external view returns (Params calldata params);

    /**
     * @dev putPolicy defines a method for put a policy to bucket/object/group which can grant permission to others.
     */
    function putPolicy(
        Principal memory principal,
        string memory resource,
        Statement[] memory statements,
        int64 expirationTime
    ) external returns (bool success);

    /**
     * @dev deletePolicy defines a method for delete policy of principal.
     */
    function deletePolicy(
        Principal memory principal,
        string memory resource
    ) external returns (bool success);

    /**
     * @dev toggleSPAsDelegatedAgent defines a method for toggle SP as delegated agent.
     */
    function toggleSPAsDelegatedAgent(
        string memory bucketName
    ) external returns (bool success);

    /**
     * @dev updateParams defines a method for update params of modular storage.
     */
    function updateParams(
        string memory authority,
        Params memory params
    ) external returns (bool success);

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
     * @dev MirrorBucket defines an Event emitted when a user mirror a bucket
     */
    event MirrorBucket(address indexed creator);

    /**
     * @dev MigrateBucket defines an Event emitted when a user migrate a bucket
     */
    event MigrateBucket(address indexed creator, string indexed bucketName);

    /**
     * @dev DiscontinueBucket defines an Event emitted when a user discontinue a bucket
     */
    event DiscontinueBucket(address indexed creator, string indexed bucketName);

    /**
     * @dev CompleteMigrateBucket defines an Event emitted when a user complete migrate a bucket
     */
    event CompleteMigrateBucket(
        address indexed creator,
        string indexed bucketName
    );

    /**
     * @dev RejectMigrateBucket defines an Event emitted when a user reject migrate a bucket
     */
    event RejectMigrateBucket(address indexed operator);

    /**
     * @dev CancelMigrateBucket defines an Event emitted when a user cancel migrate a bucket
     */
    event CancelMigrateBucket(address indexed operator);

    /**
     * @dev SetBucketFlowRateLimit defines an Event emitted when a user set the bucket flow rate limit
     */
    event SetBucketFlowRateLimit(address indexed operator);

    /**
     * @dev CreateObject defines an Event emitted when a user create a object
     */
    event CreateObject(address indexed creator, uint256 id);

    /**
     * @dev CancelCreateObject defines an Event emitted when a user cancel to create a object
     */
    event CancelCreateObject(address indexed creator);

    /**
     * @dev CopyObject defines an Event emitted when a user copy a object
     */
    event CopyObject(address indexed creator);

    /**
     * @dev DeleteObject defines an Event emitted when a user delete a object
     */
    event DeleteObject(address indexed creator);

    /**
     * @dev MirrorObject defines an Event emitted when a user mirror a object
     */
    event MirrorObject(address indexed creator);

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
     * @dev UpdateObjectContent defines an Event emitted when a user update a object content
     */
    event UpdateObjectContent(
        address indexed operator,
        string indexed objectName
    );

    /**
     * @dev DelegateUpdateObjectContent defines an Event emitted when a user delegate update a object content
     */
    event DelegateUpdateObjectContent(
        address indexed operator,
        string indexed objectName
    );

    /**
     * @dev DiscontinueObject defines an Event emitted when a user discontinue a object
     */
    event DiscontinueObject(address indexed creator, string indexed bucketName);

    /**
     * @dev CreateGroup defines an Event emitted when a user create a group
     */
    event CreateGroup(address indexed creator, uint256 id);

    /**
     * @dev UpdateGroup defines an Event emitted when a user update a group's member
     */
    event UpdateGroup(address indexed creator);

    /**
     * @dev UpdateGroupExtra defines an Event emitted when a user update a group's extra
     */
    event UpdateGroupExtra(address indexed creator);

    /**
     * @dev DeleteGroup defines an Event emitted when a user delete a group
     */
    event DeleteGroup(address indexed creator);

    /**
     * @dev LeaveGroup defines an Event emitted when a user leave a group
     */
    event LeaveGroup(address indexed creator, string indexed groupName);

    /**
     * @dev MirrorGroup defines an Event emitted when a user mirror a group
     */
    event MirrorGroup(address indexed creator);

    /**
     * @dev RenewGroupMember defines an Event emitted when a user renew group member
     */
    event RenewGroupMember(address indexed creator);

    /**
     * @dev SetTags defines an Event emitted when a user set tags for the given group/bucket/object
     */
    event SetTags(address indexed creator);

    /**
     * @dev PutPolicy defines an Event emitted when a user put a policy to bucket/object/group which can grant permission to others
     */
    event PutPolicy(address indexed creator);

    /**
     * @dev DeletePolicy defines an Event emitted when a user delete policy of principal
     */
    event DeletePolicy(address indexed creator);

    /**
     * @dev ToggleSPAsDelegatedAgent defines an Event emitted when a user toggle SP as delegated agent
     */
    event ToggleSPAsDelegatedAgent(address indexed creator);

    /**
     * @dev UpdateParams defines an Event emitted when a user update params of modular storage
     */
    event UpdateParams(address indexed creator);
}
