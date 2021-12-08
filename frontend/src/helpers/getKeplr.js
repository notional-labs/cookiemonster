import { SigningCosmosClient, } from "@cosmjs/launchpad";


export const getKeplr = async (chain_id = "osmosis-1") => {
    if (!window.getOfflineSigner || !window.keplr) {
        alert("Keplr Wallet not detected, please install extension");
    } else {
        await window.keplr.enable(chain_id);
        const offlineSigner = window.keplr.getOfflineSigner(chain_id);
        const accounts = await offlineSigner.getAccounts();
        return {
            accounts,
            offlineSigner,
        };
    }
}

export const getCosmosClient = (accounts, offlineSigner) => {
    const URL = "https://lcd-osmosis.keplr.app"
    const cosmJS = new SigningCosmosClient(
        URL,
        accounts[0].address,
        offlineSigner,
    );
    return cosmJS;
}

// export const getBalance = async (address) => {
//     const apiUrl = "https://lcd-osmosis.keplr.app"
//     const client = LcdClient.withExtensions(
//         apiUrl,
//         setupBankExtension,
//     )
//     const balance = await client.bank.balances(address);
//     return balance
// }