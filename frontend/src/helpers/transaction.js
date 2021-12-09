import { coins } from "@cosmjs/launchpad";

export const transaction = async (cosmJS, amount = 1000000, recipient = 'osmo1vxgcyq7nc8d8gykhwf35e4z0l04xhn4fq456uj') => {
    // define memo (not required)
    const memo = "Deposit";
    // sign and broadcast Tx
    const ret = await cosmJS.sendTokens(recipient, coins(amount, "uosmo"), memo);
    console.log(ret.transactionHash)
    return ret
}
