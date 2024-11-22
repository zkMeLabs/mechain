// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

/**
 * @dev ProposalStatus enumerates the valid statuses of a proposal.
 */
enum ProposalStatus {
    // Unspecified defines the default proposal status.
    Unspecified,
    // DepositPeriod defines a proposal status during the deposit.
    DepositPeriod,
    // VotingPeriod defines a proposal status during the voting.
    VotingPeriod,
    // Passed defines a proposal status of a proposal that has passed.
    Passed,
    // Rejected defines a proposal status of a proposal that has been rejected.
    Rejected,
    // Failed defines a proposal status of a proposal that has failed.
    Failed
}

/**
 * @dev VoteOption enumerates the valid vote options for a given governance proposal.
 */
enum VoteOption {
    // Unspecified defines a no-op vote option.
    Unspecified,
    // Yes defines a yes vote option.
    Yes,
    // Abstain defines an abstain vote option.
    Abstain,
    // No defines a no vote option.
    No,
    // NoWithWeto defines a no with veto vote option.
    NoWithWeto
}

/**
 * @dev TallyResult defines a standard tally for a governance proposal.
 */
struct TallyResult {
    string yesCount;
    string abstainCount;
    string noCount;
    string noWithVetoCount;
}

/**
 * @dev Proposal defines the core field members of a governance proposal.
 */
struct Proposal {
    uint64 id;
    string[] messages;
    ProposalStatus status;
    // final_tally_result is the final tally result of the proposal. When
    // querying a proposal, this field is not populated until the
    // proposal's voting period has ended.
    TallyResult finalTallyResult;
    int64 submitTime;
    int64 depositEndTime;
    Coin[] totalDeposit;
    int64 votingStartTime;
    int64 votingEndTime;
    string metadata;
    string title;
    string summary;
    address proposer;
    string failedReason;
}

/**
 * @dev WeightedVoteOption defines a unit of vote for vote split.
 */
struct WeightedVoteOption {
    VoteOption option;
    string weight;
}

/**
 * @dev VoteData defines a vote on a governance proposal.
 * A VoteData consists of a proposal ID, the voter, and the vote option.
 */
struct VoteData {
    uint64 proposalId;
    address voter;
    WeightedVoteOption[] options;
    string metadata;
}

/**
 * @dev DepositData defines an amount deposited by an account address to an active
 * proposal.
 */
struct DepositData {
    uint64 proposalId;
    address depositor;
    Coin[] amount;
}

/**
 * @dev Params defines the parameters for the gov module.
 */
struct Params {
    Coin[] minDeposit;
    int64 maxDepositPeriod;
    int64 votingPeriod;
    string quorum;
    string threshold;
    string vetoThreshold;
    string minInitialDepositRatio;
    bool burnVoteQuorum;
    bool burnProposalDepositPrevote;
    bool burnVoteVeto;
}

interface IGov {
    /**
     * @dev legacySubmitProposal defines a method to create new proposal given a content for v1beat1.
     */
    function legacySubmitProposal(
        string memory title,
        string memory description,
        Coin[] memory initialDeposit
    ) external returns (uint64 proposalId);

    /**
     * @dev submitProposal defines a method to create new proposal given a content for v1.
     */
    function submitProposal(
        string memory messages,
        Coin[] memory initialDeposit,
        string memory metadata,
        string memory title,
        string memory summary
    ) external returns (uint64 proposalId);

    /**
     * @dev vote defines a method to add a vote on a specific proposal.
     */
    function vote(
        uint64 proposalId,
        VoteOption option,
        string memory metadata
    ) external returns (bool success);

    /**
     * @dev voteWeighted defines a method to add a weighted vote on a specific proposal.
     */
    function voteWeighted(
        uint64 proposalId,
        WeightedVoteOption[] memory options,
        string memory metadata
    ) external returns (bool success);

    /**
     * @dev deposit defines a method to add deposit on a specific proposal.
     */
    function deposit(
        uint64 proposalId,
        uint256 amount
    ) external returns (bool success);

    /**
     * @dev proposal queries proposal details based on ProposalID.
     */
    function proposal(
        uint64 proposalId
    ) external view returns (Proposal calldata proposal);

    /**
     * @dev proposals queries all proposals based on given status.
     */
    function proposals(
        ProposalStatus status,
        address voter,
        address depositor,
        PageRequest calldata pagination
    ) external view returns (Proposal[] calldata proposals, PageResponse calldata pageResponse);

    /**
     * @dev vote queries voted information based on proposalID, voterAddr.
     */
    function vote(
        uint64 proposalId,
        address voter
    ) external view returns (VoteData calldata vote);

    /**
     * @dev votes queries votes of a given proposal.
     */
    function votes(
        uint64 proposalId,
        PageRequest calldata pagination
    ) external view returns (VoteData[] calldata votes, PageResponse calldata pageResponse);

    /**
     * @dev deposit queries single deposit information based proposalID, depositAddr.
     */
    function deposit(
        uint64 proposalId,
        address depositor
    ) external view returns (DepositData calldata deposit);

    /**
     * @dev deposits queries all deposits of a single proposal.
     */
    function deposits(
        uint64 proposalId,
        PageRequest calldata pagination
    ) external view returns (DepositData[] calldata deposits, PageResponse calldata pageResponse);

    /**
     * @dev params queries the gov params.
     */
    function params() external view returns (Params calldata params);

    /**
     * @dev tallyResult queries the tally of a proposal vote.
     */
    function tallyResult(
        uint64 proposalId
    ) external view returns (TallyResult calldata tallyResult);

    /**
     * @dev LegacySubmitProposal defines an Event emitted when a legacy proposal submited.
     */
    event LegacySubmitProposal(
        address indexed proposer,
        uint64 proposalId
    );

    /**
     * @dev SubmitProposal defines an Event emitted when a proposal submited.
     */
    event SubmitProposal(
        address indexed proposer,
        uint64 proposalId
    );

    /**
     * @dev Vote defines an Event emitted when a proposal voted.
     */
    event Vote(
        address indexed voter,
        uint64 proposalId,
        uint8 option
    );

    /**
     * @dev Vote defines an Event emitted when a proposal vote weighted.
     */
    event VoteWeighted(
        address indexed voter,
        uint64 proposalId
    );

    /**
     * @dev Vote defines an Event emitted when a proposal deposited by a depositor.
     */
    event Deposit(
        address indexed depositor,
        uint64 proposalId
    );
}
