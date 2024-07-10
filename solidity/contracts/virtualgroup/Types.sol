// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

// A global virtual group consists of one primary SP (SP) and multiple secondary SP.
// Every global virtual group must belong to a GVG family, and the objects of each
// bucket must be stored in a GVG within a group family.
struct GlobalVirtualGroup {
  // ID represents the unique identifier of the global virtual group.
  uint32 id;
  // Family ID represents the identifier of the GVG family that the group belongs to.
  uint32 familyId;
  // Primary SP ID represents the unique identifier of the primary storage provider in the group.
  uint32 primarySpId;
  // Secondary SP IDs represents the list of unique identifiers of the secondary storage providers in the group.
  uint32[] secondarySpIds;
  // Stored size represents the size of the stored objects within the group.
  uint64 storedSize;
  // Virtual payment address represents the payment address associated with the group.
  address virtualPaymentAddress;
  // Total deposit represents the number of tokens deposited by this storage provider for staking.
  string totalDeposit;
}