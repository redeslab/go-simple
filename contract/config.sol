//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.1;

contract owned {
    address public owner;

    constructor(){
        owner = msg.sender;
    }

    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

    function transferOwnership(address newOwner) onlyOwner public {
        owner = newOwner;
    }
}

contract Config is owned{

    struct ServerItem{
        string addr;
        string host;
    }
    address public AdvertiseAddr;
    ServerItem[] public servers;
    mapping(string=>uint) public sIdx;
    mapping(address=>bool)public Administrators;
    event ServerChanged(string, string, uint8);
    constructor() {
        Administrators[msg.sender]=true;
        servers.push(ServerItem("", ""));
    }
    modifier onlyAdmin{
        require(Administrators[msg.sender],"invalid operator");
        _;
    }
    function addAdmin(address admin) public onlyOwner{
        Administrators[admin] = true;
    }
    function removeAdmin(address admin) public onlyOwner{
        delete  Administrators[admin];
    }

    function setAdvertisAddr(address addr) public onlyAdmin{
        AdvertiseAddr = addr;
    }

    function addServer(string memory addr, string memory host)public onlyAdmin{
        require(sIdx[addr] == 0, "dulicated server");

        sIdx[addr]=servers.length;
        servers.push(ServerItem(addr, host));

        emit ServerChanged(addr, host, 0);
    }

    function removeServer(string memory addr) public onlyAdmin{
        uint idx = sIdx[addr];
        require(idx > 0, "no such server");

        if (idx == servers.length - 1){
            servers.pop();
            sIdx[addr]=0;
            return ;
        }

        ServerItem memory lastItem = servers[servers.length - 1];
        servers[idx] = lastItem;
        servers.pop();
        sIdx[addr]=0;
        sIdx[lastItem.addr] = idx;
        emit ServerChanged(addr, "remove", 1);
    }

    function changeServer(string memory addr, string memory host)public onlyAdmin{
        uint idx = sIdx[addr];
        require(idx > 0, "no such server");

        servers[idx].host = host;
        emit ServerChanged(addr, host, 2);
    }

    function ServerList() public view returns(ServerItem[] memory){
        return servers;
    }

    function QueryByOne(string memory addr) public view returns (string memory){
        uint idx = sIdx[addr];
        if (idx == 0){
            return "";
        }

        return servers[idx].host;
    }
}