import socket

# 域名和 IP 地址的映射
domain_ip_mapping = {
    "devnet-rpc.mechain.tech": "154.48.244.41:8545",
    "devnet-api.mechain.tech": "154.48.244.41:1317",
    "devnet-lcd.mechain.tech": "154.48.244.41:26657",
    "devnet-sp0-rpc.mechain.tech": "154.48.244.41:9033",
    "devnet-scan.mechain.tech": "154.48.244.41:8000",
    "scan.mechain.tech": "154.48.244.41:8000",
    "devnet-stats.mechain.tech": ["154.48.244.41:8080"],
    "devnet-visualize.mechain.tech": ["154.48.244.41:8081"],
    "devnet-faucet.mechain.tech": "154.48.244.41:4001",
    "devnet-nft.mechain.tech": "154.48.244.41:8123",
}


def resolve_domain(domain):
    try:
        resolved_ips = socket.gethostbyname_ex(domain)[2]
        return resolved_ips
    except socket.gaierror:
        return []


def test_domain_ip_mapping(mapping):
    for domain, expected_ips in mapping.items():
        if isinstance(expected_ips, str):
            expected_ips = [expected_ips.split(":")[0]]
        else:
            expected_ips = [ip.split(":")[0] for ip in expected_ips]

        resolved_ips = resolve_domain(domain)
        for ip in expected_ips:
            if ip not in resolved_ips:
                print(
                    f"Failure: {domain} does not resolve to {ip}, now is {resolved_ips}"
                )


if __name__ == "__main__":
    test_domain_ip_mapping(domain_ip_mapping)
