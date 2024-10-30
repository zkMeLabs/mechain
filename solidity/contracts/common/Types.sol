// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

/**
 * @dev Coin defines a token with a denomination and an amount.
 */
struct Coin {
    string denom;
    uint256 amount;
}

/**
 * @dev DecCoin defines a token with a denomination and a decimal amount.
 */
struct DecCoin {
    string denom;
    uint256 amount;
    uint8 precision;
}

/**
 * @dev Dec defines decimal with a precision. for example, amount:16800, precision:3 means is 16.800
 */
struct Dec {
    uint256 amount;
    uint8 precision;
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

/**
 * @dev PageRequest is to be embedded in request messages for efficient pagination
 */
struct PageRequest {
    // key is a value returned in PageResponse.next_key to begin
    // querying the next page most efficiently. Only one of offset or key
    // should be set.
    bytes key;
    // offset is a numeric offset that can be used when key is unavailable.
    // It is less efficient than using key. Only one of offset or key should
    // be set.
    uint64 offset;
    // limit is the total number of results to be returned in the result page.
    // If left empty it will default to a value to be set by each app.
    uint64 limit;
    // countTotal is set to true to indicate that the result set should include
    // a count of the total number of items available for pagination in UIs.
    // count_total is only respected when offset is used. It is ignored when key
    // is set.
    bool countTotal;
    // reverse is set to true if results are to be returned in the descending order.
    bool reverse;
}

/**
 * @dev PageResponse is to be embedded in gRPC response messages where the
 * corresponding request message has used PageRequest.
 */
struct PageResponse {
    // nextKey is the key to be passed to PageRequest.key to
    // query the next page most efficiently. It will be empty if
    // there are no more results.
    bytes nextKey;
    // total is total number of results available if PageRequest.count_total
    // was set, its value is undefined otherwise
    uint64 total;
}
