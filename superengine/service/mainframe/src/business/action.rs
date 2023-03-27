use std::net::{IpAddr, Ipv4Addr};
use tokio::io::Error;
use tokio::net::UdpSocket;

/// Mac address size of bytes
const MAC_ADDR_SIZE: usize = 6;
/// 魔术封包大小，大小固定的 6+(6*16)，前6字节为0xff, 然后是16次mac地址
const MAGIC_PACKET_SIZE: usize = MAC_ADDR_SIZE + MAC_ADDR_SIZE * 16;

/// 默认绑定地址: 0.0.0.0
const DEFAULT_BIND_ADDR: IpAddr = IpAddr::V4(Ipv4Addr::new(0, 0, 0, 0));
// 默认绑定地址, 0代表系统分配
const DEFAULT_BIND_PORT: u16 = 0;

/// 默认广播地址: 255.255.255.255
const DEFAULT_BROADCAST_ADDR: IpAddr = IpAddr::V4(Ipv4Addr::new(255, 255, 255, 255));
/// 默认广播端口, 必须是[0/7/9], 这里用个9
const DEFAULT_BROADCAST_PORT: u16 = 9;

pub struct ComputerAction {}

impl ComputerAction {
    pub fn new() -> Self {
        return Self {};
    }

    /// 远程唤醒
    pub async fn open(
        &self,
        mac_addr: &str,
        broadcast: Option<IpAddr>,
        ipaddr: Option<IpAddr>,
    ) -> Result<(), Error> {
        let mut magic_packet = vec![0; MAGIC_PACKET_SIZE];
        for index in 0..MAC_ADDR_SIZE {
            magic_packet[index] = 0xff;
            let value = u8::from_str_radix(&mac_addr[index * 2..index * 2 + 2], 16).unwrap();
            for step in (MAC_ADDR_SIZE..MAGIC_PACKET_SIZE).step_by(MAC_ADDR_SIZE) {
                magic_packet[step + index] = value;
            }
        }
        let broadcast = broadcast.unwrap_or_else(|| DEFAULT_BROADCAST_ADDR);
        let ipaddr = ipaddr.unwrap_or_else(|| DEFAULT_BIND_ADDR);
        // Bind port, 0 means assigned by os
        let socket = UdpSocket::bind((ipaddr, DEFAULT_BIND_PORT)).await?;
        socket.set_broadcast(true)?;
        //  WoL port could be 0/7/9， use 9 here
        socket
            .send_to(&magic_packet, (broadcast, DEFAULT_BROADCAST_PORT))
            .await?;
        Ok(())
    }

    /// 远程关闭
    pub async fn close(
        &self,
        username: &str,
        password: &str,
        hostname: &str,
        port: Option<u16>,
    ) -> Result<(), Error> {
        Ok(())
    }
    /// 状态检测
    pub async fn detect(&self, hostname: &str, port: Option<u16>) -> Result<(), Error> {
        Ok(())
    }
}

#[tokio::test]
async fn test_wakeup() {
    ComputerAction::new()
        .open("A1A2A3A4A5A6", None, None)
        .await
        .unwrap();
}
