// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

interface IJsonUtil {

  /////////////////////////////////////// READ METHODS (on jsonBlob) //////////////////////////////////////////

  function get(string memory jsonBlob, string memory path) view external returns (string memory);

  function getInt(string memory jsonBlob, string memory path) view external returns (int256);

  function getUint(string memory jsonBlob, string memory path) view external returns (uint256);

  function getBool(string memory jsonBlob, string memory path) view external returns (bool);

  function dataURI(string memory jsonBlob) view external returns (string memory);

  function exists(string memory jsonBlob, string memory path) view external returns (bool);

  ////////////////////////////////////// MODIFICATION VIEWS (on jsonBlob) //////////////////////////////////////////

  function set(string memory jsonBlob, string memory path, string memory value) view external returns (string memory);

  function set(string memory jsonBlob, string[] memory paths, string[] memory values) view external returns (string memory);

  function setRaw(string memory jsonBlob, string memory path, string memory rawBlob) view external returns (string memory);

  function setRaw(string memory jsonBlob, string[] memory paths, string[] memory rawBlobs) view external returns (string memory);

  function setInt(string memory jsonBlob, string memory path, int256 value) view external returns (string memory);

  function setInt(string memory jsonBlob, string[] memory paths, int256[] memory values) view external returns (string memory);

  function setUint(string memory jsonBlob, string memory path, uint256 value) view external returns (string memory);

  function setUint(string memory jsonBlob, string[] memory paths, uint256[] memory values) view external returns (string memory);

  function setBool(string memory jsonBlob, string memory path, bool value) view external returns (string memory);

  function setBool(string memory jsonBlob, string[] memory paths, bool[] memory values) view external returns (string memory);

  function subReplace(string memory jsonBlob, string memory searchPath, string memory replacePath, string memory value) view external returns (string memory);

  function subReplace(string memory jsonBlob, string memory searchPath, string[] memory replacePaths, string[] memory values) view external returns (string memory);

  function subReplaceInt(string memory jsonBlob, string memory searchPath, string memory replacePath, int256 value) view external returns (string memory);

  function subReplaceInt(string memory jsonBlob, string memory searchPath, string[] memory replacePaths, int256[] memory values) view external returns (string memory);

  function subReplaceUint(string memory jsonBlob, string memory searchPath, string memory replacePath, uint256 value) view external returns (string memory);

  function subReplaceUint(string memory jsonBlob, string memory searchPath, string[] memory replacePaths, uint256[] memory values) view external returns (string memory);

  function subReplaceBool(string memory jsonBlob, string memory searchPath, string memory replacePath, bool value) view external returns (string memory);

  function subReplaceBool(string memory jsonBlob, string memory searchPath, string[] memory replacePaths, bool[] memory values) view external returns (string memory);

  function remove(string memory jsonBlob, string memory path) view external returns (string memory);

}
