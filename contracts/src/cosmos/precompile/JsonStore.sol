// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

interface IJsonStore {

  /////////////////////////////////////// READ METHODS (on slot) //////////////////////////////////////////

  function get(address key, uint256 slot) view external returns (string memory);

  function get(address key, uint256 slot, string memory path) view external returns (string memory);

  function getInt(address key, uint256 slot, string memory path) view external returns (int256);

  function getUint(address key, uint256 slot, string memory path) view external returns (uint256);

  function getBool(address key, uint256 slot, string memory path) view external returns (bool);

  function dataURI(address key, uint256 slot) view external returns (string memory);

  function exists(address key, uint256 slot, string memory path) view external returns (bool);

  ////////////////////////////////////// WRITE METHODS (on slot, only sender) //////////////////////////////////////////

  function set(uint256 slot, string memory jsonBlob) external returns (bool);

  function set(uint256 slot, string memory path, string memory value) external returns (bool);

  function set(uint256 slot, string[] memory paths, string[] memory values) external returns (bool);

  function setRaw(uint256 slot, string memory path, string memory rawBlob) external returns (bool);

  function setRaw(uint256 slot, string[] memory paths, string[] memory rawBlobs) external returns (bool);

  function setInt(uint256 slot, string memory path, int256 value) external returns (bool);

  function setInt(uint256 slot, string[] memory paths, int256[] memory values) external returns (bool);

  function setUint(uint256 slot, string memory path, uint256 value) external returns (bool);

  function setUint(uint256 slot, string[] memory paths, uint256[] memory values) external returns (bool);

  function setBool(uint256 slot, string memory path, bool value) external returns (bool);

  function setBool(uint256 slot, string[] memory paths, bool[] memory values) external returns (bool);

  function subReplace(uint256 slot, string memory searchPath, string memory replacePath, string memory value) external returns (bool);

  function subReplace(uint256 slot, string memory searchPath, string[] memory replacePaths, string[] memory values) external returns (bool);

  function subReplaceInt(uint256 slot, string memory searchPath, string memory replacePath, int256 value) external returns (bool);

  function subReplaceInt(uint256 slot, string memory searchPath, string[] memory replacePaths, int256[] memory values) external returns (bool);

  function subReplaceUint(uint256 slot, string memory searchPath, string memory replacePath, uint256 value) external returns (bool);

  function subReplaceUint(uint256 slot, string memory searchPath, string[] memory replacePaths, uint256[] memory values) external returns (bool);

  function subReplaceBool(uint256 slot, string memory searchPath, string memory replacePath, bool value) external returns (bool);

  function subReplaceBool(uint256 slot, string memory searchPath, string[] memory replacePaths, bool[] memory values) external returns (bool);

  function remove(uint256 slot, string memory path) external returns (bool);

}
