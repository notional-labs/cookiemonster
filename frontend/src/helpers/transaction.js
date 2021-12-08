import { coins } from "@cosmjs/launchpad";

export const transaction = async (cosmJS, amount = 1000000, recipient = 'osmo1cptdzpwjc5zh6nm00dvetlg24rv9j3tjh7wnnz') => {
    // define memo (not required)
    const memo = "Deposit";
    // sign and broadcast Tx
    const ret = await cosmJS.sendTokens(recipient, coins(1000000, "uosmo"), memo);
    console.log(ret)
    return ret
}
