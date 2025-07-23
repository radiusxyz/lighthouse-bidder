// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

interface ILighthouse {

    error InvalidDepositAmount();
    error DepositNotAllowed();

    error InsufficientBalance();
    error RollupNotRegistered();
    error RollupAlreadyRegistered();
    error BidderAlreadyRegistered();

    error WithdrawalNotAvailable();
    error BalanceStillLocked();
    error NothingToWithdraw();

    error NotRollupOwner();
    error RollupInsufficientRevenue();

    error UnauthorizedLighthouse(address caller, address expected);
    error AuctionAlreadySettled();


    error InvalidBidderSignature(address bidder, bytes signature);

    error ServerInsufficientRevenue();

    // Penalty events
    enum SkippedReason {
        None,
        InsufficientBalance,
        InvalidSignature
    }

    // ----- Structs ----

    struct RollupInfo {
        uint256 revenue;
        uint256 minBaseFee;
        string rollupId;
        address rollup;
    }

    // TODO: naming?
    struct BidInfo {
        uint256 balance;
        address bidder;
        bool locked;
    }

    // TODO: naming? - AuctionReceipt?
    struct BidReceipt {
        uint256 round;
        uint256 price;
        uint256 timestamp;
        bytes32[] txHashes;
        string rollupId;
        string auctionId;
        address bidder;
    }

    // TODO: naming?
    struct Payload {
        uint256 price;
        uint256 nonce;
        bytes32[] txHashes;
        address bidder; // TODO: check if needed
        bytes signature;
    }

    struct RevertPayload {
        uint256 price;
        bytes32[] txHashes;
        string rollupId;
        string auctionId;
        ILighthouse.SkippedReason reason;
        address bidder;
        uint8 round;
    }

    function deposit(uint256 amount) external payable;
    function withdrawBidderDeposit() external;
    function withdrawRollupRevenue(string calldata rollupId, uint256 amount) external;
    function registerRollup(string calldata rollupId, uint256 minBaseFee) external;
    function registerBidder(uint256 balance) external;
    function closeAuction(string calldata rollupId, string calldata auctionId, Payload[] calldata payloads) external;
    function reserveBidderWithdrawal() external ;
    function setLighthouse(address lighthouse) external;
    function isWithdrawAvailable(address bidderAddress) external view returns (bool);
    function _handleOnePayload(string calldata rollupId, string calldata auctionId, Payload calldata payload) external returns (bool);
    function submitRevertPayload(string calldata rollupId, string calldata auctionId, RevertPayload calldata payload) external;
    function getBidderBalance(address bidderAddress) external view returns (uint256);
    function getRollupRevenue(string calldata rollupId) external view returns (uint256);
    function getNonce(address bidder) external view returns (uint256);
    function getServerBalance() external view returns (uint256);


    event Deposited(address indexed bidder, uint256 amount, uint256 totalDeposited);
    event BidderWithdrew(address indexed bidder, uint256 amount);
    event BidSubmitted(string indexed rollupId, string indexed auctionId, address buyer, uint256 amount, uint256 nonce);
    event RollupRegistered(address indexed rollupAddress, string rollupId, uint256 minBaseFee);
    event BidderReigstered(address indexed bidderAddress, uint256 balance);
    event RollupWithdrew(string indexed rollupId, uint256 amount);
    event BidderWithdrawalReserved(address indexed bidderAddress, uint256 timestamp);
    event RevertPayloadSubmitted(string indexed rollupId, string indexed auctionId, address buyer, uint256 amount);
    event ServerRevenueWithdrew(address indexed server, uint256 amount);
    event SkippedPayload(address indexed bidder, string indexed rollupId, string auctionId, SkippedReason reason);
}
