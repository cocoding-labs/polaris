// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import {TokenRouter} from "@hyperlane-xyz/core/contracts/token/libs/TokenRouter.sol";
import {INativeMinter} from "./precompile/INativeMinter.sol";

contract HypNativeMinter is TokenRouter {

    INativeMinter public immutable nativeMinter;
    uint256 public immutable scale;

    constructor(address _precompile, uint256 _scale, address _mailbox) TokenRouter(_mailbox) {
        nativeMinter = INativeMinter(_precompile);
        scale = _scale;
    }

    function transferRemote(
        uint32 _destination,
        bytes32 _recipient,
        uint256 _amount
    ) public payable virtual override returns (bytes32 messageId) {
        require(msg.value >= _amount, "HypNativeMinter: amount exceeds msg.value");
        uint256 gasPayment = msg.value - _amount;
        uint256 scaledAmount = _amount / scale;
        require(scaledAmount > 0, "HypNativeMinter: destination amount < 1");
        return _transferRemote(_destination, _recipient, scaledAmount, gasPayment);
    }

    function balanceOf(
        address _account
    ) external view override returns (uint256) {
        return _account.balance;
    }

    function _transferFromSender(
        uint256 _amount
    ) internal override returns (bytes memory) {
        uint256 scaledAmount = _amount * scale;
        nativeMinter.burn(address(this), scaledAmount);
        return bytes(""); // no metadata
    }

    function _transferTo(
        address _recipient,
        uint256 _amount,
        bytes calldata // no metadata
    ) internal virtual override {
        uint256 scaledAmount = _amount * scale;
        nativeMinter.mint(_recipient, scaledAmount);
    }

}