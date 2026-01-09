from typing import Any
import requests


def get_rpc_urls_for_chain(chain_id: int) -> list[str]:
    url = "https://chainlist.org/rpcs.json"
    response = requests.get(url)
    response.raise_for_status()  # Raise an error if request failed

    chains: list[dict[str, Any]] = response.json()  # pyright: ignore[reportExplicitAny, reportAny]

    target_chain = None
    for chain in chains:
        if chain.get("chainId") == chain_id:
            target_chain = chain
            break

    if target_chain is None:
        print(f"No chain found with chainId {chain_id}")
        return []

    rpc_list: list[dict[str, Any]] = target_chain.get("rpc", [])  # pyright: ignore[reportAny, reportExplicitAny]
    urls = [rpc_entry["url"] for rpc_entry in rpc_list if "url" in rpc_entry]

    return urls


rpc_urls = get_rpc_urls_for_chain(42161)
for url in rpc_urls:
    print(url)
